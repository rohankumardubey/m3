// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package client

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	murmur3 "github.com/m3db/stackmurmur3/v2"
	"github.com/uber-go/tally"
	"github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/thrift"
	"go.uber.org/zap"

	"github.com/m3db/m3/src/dbnode/generated/thrift/rpc"
	"github.com/m3db/m3/src/dbnode/network/server/tchannelthrift/node/channel"
	"github.com/m3db/m3/src/dbnode/topology"
)

const (
	channelName = "Client"
)

var (
	errConnectionPoolClosed           = errors.New("connection pool closed")
	errConnectionPoolHasNoConnections = newHostNotAvailableError(errors.New("connection pool has no connections"))
	errNodeNotBootstrapped            = errors.New("node not bootstrapped")
)

type connPool struct {
	sync.RWMutex

	opts             Options
	host             topology.Host
	pool             []conn
	poolLen          int64
	used             int64
	connectRand      rand.Source
	healthCheckRand  rand.Source
	sleepConnect     sleepFn
	sleepHealth      sleepFn
	sleepHealthRetry sleepFn
	status           status
	healthStatus     tally.Gauge
}

type conn struct {
	channel Channel
	client  rpc.TChanNode
}

// NewConnectionFn is a function that creates a connection.
type NewConnectionFn func(channelName string, opts Options) (Channel, error)

// NewClientFn constructs client given a channel.
type NewClientFn func(c Channel, address string) (rpc.TChanNode, error)

// HealthCheckFn is a function that checks if connection is still healthy and should be kept in the pool.
type HealthCheckFn func(client rpc.TChanNode, opts Options, checkBootstrapped bool) error

type sleepFn func(t time.Duration)

func newConnectionPool(host topology.Host, opts Options) connectionPool {
	seed := int64(murmur3.StringSum32(host.Address()))

	scope := opts.InstrumentOptions().
		MetricsScope().
		Tagged(map[string]string{
			"hostID": host.ID(),
		})

	p := &connPool{
		opts:             opts,
		host:             host,
		pool:             make([]conn, 0, opts.MaxConnectionCount()),
		poolLen:          0,
		connectRand:      rand.NewSource(seed),
		healthCheckRand:  rand.NewSource(seed + 1),
		sleepConnect:     time.Sleep,
		sleepHealth:      time.Sleep,
		sleepHealthRetry: time.Sleep,
		healthStatus:     scope.Gauge("health-status"),
	}

	return p
}

func (p *connPool) Open() {
	p.Lock()
	defer p.Unlock()

	if p.status != statusNotOpen {
		return
	}

	p.status = statusOpen

	connectEvery := p.opts.BackgroundConnectInterval()
	connectStutter := p.opts.BackgroundConnectStutter()
	go p.connectEvery(connectEvery, connectStutter)

	healthCheckEvery := p.opts.BackgroundHealthCheckInterval()
	healthCheckStutter := p.opts.BackgroundHealthCheckStutter()
	go p.healthCheckEvery(healthCheckEvery, healthCheckStutter)
}

func (p *connPool) ConnectionCount() int {
	p.RLock()
	poolLen := p.poolLen
	p.RUnlock()
	return int(poolLen)
}

func (p *connPool) NextClient() (rpc.TChanNode, Channel, error) {
	p.RLock()
	if p.status != statusOpen {
		p.RUnlock()
		return nil, nil, errConnectionPoolClosed
	}
	if p.poolLen < 1 {
		p.RUnlock()
		return nil, nil, errConnectionPoolHasNoConnections
	}
	n := atomic.AddInt64(&p.used, 1)
	conn := p.pool[n%p.poolLen]
	p.RUnlock()
	return conn.client, conn.channel, nil
}

func (p *connPool) Close() {
	p.Lock()
	if p.status != statusOpen {
		p.Unlock()
		return
	}
	p.status = statusClosed
	p.Unlock()

	for i := range p.pool {
		p.pool[i].channel.Close()
	}
}

func (p *connPool) connectEvery(interval time.Duration, stutter time.Duration) {
	log := p.opts.InstrumentOptions().Logger()
	target := p.opts.MaxConnectionCount()

	for {
		p.RLock()
		state := p.status
		poolLen := int(p.poolLen)
		p.RUnlock()
		if state != statusOpen {
			return
		}

		address := p.host.Address()

		var wg sync.WaitGroup
		for i := 0; i < target-poolLen; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				ch, client, healthCheckErr, err := establishNewConnection(address, false, p.opts)
				if err != nil {
					if healthCheckErr {
						p.maybeEmitHealthStatus(healthStatusCheckFailed)
					}
					log.Warn("could not connect", zap.String("host", address), zap.Error(err))
					return
				}

				p.maybeEmitHealthStatus(healthStatusOK)
				p.Lock()
				if p.status == statusOpen {
					p.pool = append(p.pool, conn{ch, client})
					p.poolLen = int64(len(p.pool))
				} else {
					// NB(antanas): just being defensive.
					// It's likely a corner case and happens only during server shutdown.
					ch.Close()
				}
				p.Unlock()
			}()
		}

		wg.Wait()

		p.sleepConnect(interval + randStutter(p.connectRand, stutter))
	}
}

func (p *connPool) maybeEmitHealthStatus(hs healthStatus) {
	if p.opts.HostQueueEmitsHealthStatus() {
		p.healthStatus.Update(float64(hs))
	}
}

func (p *connPool) healthCheckEvery(interval time.Duration, stutter time.Duration) {
	log := p.opts.InstrumentOptions().Logger()
	nowFn := p.opts.ClockOptions().NowFn()
	healthCheckFn := p.opts.HealthCheckFn()

	for {
		p.RLock()
		state := p.status
		p.RUnlock()
		if state != statusOpen {
			return
		}

		var (
			wg       sync.WaitGroup
			start    = nowFn()
			deadline = start.Add(interval + randStutter(p.healthCheckRand, stutter))
		)

		p.RLock()
		for i := int64(0); i < p.poolLen; i++ {
			wg.Add(1)
			go func(client rpc.TChanNode) {
				defer wg.Done()

				var (
					attempts = p.opts.BackgroundHealthCheckFailLimit()
					failed   = 0
					checkErr error
				)
				for j := 0; j < attempts; j++ {
					if err := healthCheckFn(client, p.opts, false); err != nil {
						checkErr = err
						failed++
						throttleDuration := time.Duration(math.Max(
							float64(time.Second),
							p.opts.BackgroundHealthCheckFailThrottleFactor()*
								float64(p.opts.HostConnectTimeout())))
						p.sleepHealthRetry(throttleDuration)
						continue
					}
					// Healthy
					break
				}

				healthy := failed < attempts
				if !healthy {
					// Log health check error
					log.Debug("health check failed", zap.String("host", p.host.Address()), zap.Error(checkErr))

					// Swap with tail and decrement pool size
					p.Lock()
					if p.status != statusOpen {
						p.Unlock()
						return
					}
					var c conn
					for j := int64(0); j < p.poolLen; j++ {
						if client == p.pool[j].client {
							c = p.pool[j]
							p.pool[j] = p.pool[p.poolLen-1]
							p.pool = p.pool[:p.poolLen-1]
							p.poolLen = int64(len(p.pool))
							break
						}
					}
					p.Unlock()

					// Close the client's channel
					c.channel.Close()
				}
			}(p.pool[i].client)
		}
		p.RUnlock()

		wg.Wait()

		now := nowFn()
		if !now.Before(deadline) {
			// Exceeded deadline, start next health check loop
			p.sleepHealth(0) // Call sleep 0 for tests to intercept this loop continuation
			continue
		}

		p.sleepHealth(deadline.Sub(now))
	}
}

func defaultHealthCheck(client rpc.TChanNode, opts Options, checkBootstrapped bool) error {
	tctx, _ := thrift.NewContext(opts.HostConnectTimeout())
	result, err := client.Health(tctx)
	if err != nil {
		return err
	}
	if !result.Ok {
		return fmt.Errorf("status not ok: %s", result.Status)
	}
	if checkBootstrapped && !result.Bootstrapped {
		return errNodeNotBootstrapped
	}
	return nil
}

func defaultNewConnectionFn(
	channelName string, clientOpts Options,
) (Channel, error) {
	// NB(r): Keep ref to a local channel options since it's actually modified
	// by TChannel itself to set defaults.
	var opts *tchannel.ChannelOptions
	if chanOpts := clientOpts.ChannelOptions(); chanOpts != nil {
		immutableOpts := *chanOpts
		opts = &immutableOpts
	}
	return tchannel.NewChannel(channelName, opts)
}

func defaultNewClientFn(c Channel, address string) (rpc.TChanNode, error) {
	tc, ok := c.(*tchannel.Channel)
	if !ok {
		return nil, errors.New("can't create new client: not a *tchannel.Channel")
	}
	endpoint := &thrift.ClientOptions{HostPort: address}
	thriftClient := thrift.NewClient(tc, channel.ChannelName, endpoint)
	client := rpc.NewTChanNodeClient(thriftClient)
	return client, nil
}

func randStutter(source rand.Source, t time.Duration) time.Duration {
	amount := float64(source.Int63()) / float64(math.MaxInt64)
	return time.Duration(float64(t) * amount)
}

func establishNewConnection(
	addr string,
	bootstrappedNodesOnly bool,
	opts Options,
) (Channel, rpc.TChanNode, bool, error) {
	var (
		newConnFn            = opts.NewConnectionFn()
		newClientFn          = opts.NewClientFn()
		healthCheckNewConnFn = opts.HealthCheckNewConnFn()
	)

	ch, err := newConnFn(channelName, opts)
	if err != nil {
		return nil, nil, false, err
	}

	cl, err := newClientFn(ch, addr)
	if err != nil {
		ch.Close()
		return nil, nil, false, err
	}

	if err := healthCheckNewConnFn(cl, opts, bootstrappedNodesOnly); err != nil {
		ch.Close()
		return nil, nil, true, err
	}
	return ch, cl, false, nil
}
