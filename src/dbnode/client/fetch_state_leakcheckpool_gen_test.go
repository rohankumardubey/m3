// Copyright (c) 2021 Uber Technologies, Inc.
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

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package client

import (
	"fmt"
	"runtime/debug"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// Copyright (c) 2018 Uber Technologies, Inc.
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

// fetchStateEqualsFn allows users to override equality checks
// for `fetchState` instances.
type fetchStateEqualsFn func(a, b *fetchState) bool

// fetchStateGetHookFn allows users to override properties on items
// retrieved from the backing pools before returning in the Get()
// path.
type fetchStateGetHookFn func(*fetchState) *fetchState

// leakcheckFetchStatePoolOpts allows users to override default behaviour.
type leakcheckFetchStatePoolOpts struct {
	DisallowUntrackedPuts bool
	EqualsFn              fetchStateEqualsFn
	GetHookFn             fetchStateGetHookFn
}

// newLeakcheckFetchStatePool returns a new leakcheckFetchStatePool.
func newLeakcheckFetchStatePool(opts leakcheckFetchStatePoolOpts, backingPool fetchStatePool) *leakcheckFetchStatePool {
	if opts.EqualsFn == nil {
		// NB(prateek): fall-back to == in the worst case
		opts.EqualsFn = func(a, b *fetchState) bool {
			return a == b
		}
	}
	return &leakcheckFetchStatePool{opts: opts, fetchStatePool: backingPool}
}

// leakcheckFetchStatePool wraps the underlying fetchStatePool to make it easier to
// track leaks/allocs.
type leakcheckFetchStatePool struct {
	sync.Mutex
	fetchStatePool
	NumGets      int
	NumPuts      int
	PendingItems []leakcheckFetchState
	AllGetItems  []leakcheckFetchState

	opts leakcheckFetchStatePoolOpts
}

// leakcheckFetchState wraps `fetchState` instances along with their last Get() paths.
type leakcheckFetchState struct {
	Value         *fetchState
	GetStacktrace []byte // GetStacktrace is the stacktrace for the Get() of this item
}

func (p *leakcheckFetchStatePool) Init() {
	p.Lock()
	defer p.Unlock()
	p.fetchStatePool.Init()
}

func (p *leakcheckFetchStatePool) Get() *fetchState {
	p.Lock()
	defer p.Unlock()

	e := p.fetchStatePool.Get()
	if fn := p.opts.GetHookFn; fn != nil {
		e = fn(e)
	}

	p.NumGets++
	item := leakcheckFetchState{
		Value:         e,
		GetStacktrace: debug.Stack(),
	}
	p.PendingItems = append(p.PendingItems, item)
	p.AllGetItems = append(p.AllGetItems, item)

	return e
}

func (p *leakcheckFetchStatePool) Put(value *fetchState) {
	p.Lock()
	defer p.Unlock()

	idx := -1
	for i, item := range p.PendingItems {
		if p.opts.EqualsFn(item.Value, value) {
			idx = i
			break
		}
	}

	if idx == -1 && p.opts.DisallowUntrackedPuts {
		panic(fmt.Errorf("untracked object (%v) returned to pool", value))
	}

	if idx != -1 {
		// update slice
		p.PendingItems = append(p.PendingItems[:idx], p.PendingItems[idx+1:]...)
	}
	p.NumPuts++

	p.fetchStatePool.Put(value)
}

// Check ensures there are no leaks.
func (p *leakcheckFetchStatePool) Check(t *testing.T) {
	p.Lock()
	defer p.Unlock()

	require.Equal(t, p.NumGets, p.NumPuts)
	require.Empty(t, p.PendingItems)
}

type leakcheckFetchStateFn func(e leakcheckFetchState)

// CheckExtended ensures there are no leaks, and executes the specified fn
func (p *leakcheckFetchStatePool) CheckExtended(t *testing.T, fn leakcheckFetchStateFn) {
	p.Check(t)
	p.Lock()
	defer p.Unlock()
	for _, e := range p.AllGetItems {
		fn(e)
	}
}
