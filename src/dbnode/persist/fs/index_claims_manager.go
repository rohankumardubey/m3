// Copyright (c) 2020 Uber Technologies, Inc.
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

package fs

import (
	"errors"
	"sync"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/m3db/m3/src/dbnode/namespace"
	"github.com/m3db/m3/src/dbnode/retention"
	"github.com/m3db/m3/src/x/clock"
	"github.com/m3db/m3/src/x/ident"
	"github.com/m3db/m3/src/x/instrument"
	xtime "github.com/m3db/m3/src/x/time"
)

var (
	// errMustUseSingleClaimsManager returned when a second claims manager
	// created, since this is a violation of expected behavior.
	errMustUseSingleClaimsManager = errors.New("not using single global claims manager")

	globalIndexClaimsManagers uint64

	globalIndexClaimsManagersCheckEnabled = true
)

// ResetIndexClaimsManagersUnsafe should only be used from tests or integration
// tests, it resets the count of index claim managers to allow new claim
// managers to be created.
// By default this is restricted to just once instantiation since otherwise
// concurrency issues can be skipped without realizing.
func ResetIndexClaimsManagersUnsafe() {
	atomic.StoreUint64(&globalIndexClaimsManagers, 0)
}

// DisableIndexClaimsManagersCheckUnsafe disables the global index claims
// managers check which is useful during tests that spin up multiple DB
// nodes pointed at different data directories within the same process.
// Outside of that specific context, this method should never be used.
func DisableIndexClaimsManagersCheckUnsafe() {
	globalIndexClaimsManagersCheckEnabled = false
}

type indexClaimsManager struct {
	sync.Mutex

	filePathPrefix                string
	nowFn                         clock.NowFn
	nextIndexFileSetVolumeIndexFn nextIndexFileSetVolumeIndexFn

	// Map of ns ID string -> blockStart -> volumeIndexClaim.
	volumeIndexClaims map[string]map[xtime.UnixNano]volumeIndexClaim
}

type volumeIndexClaim struct {
	volumeIndex int
}

// NewIndexClaimsManager returns an instance of the index claim manager. This manages
// concurrent claims for volume indices per ns and block start.
// NB(bodu): There should be only a single shared index claim manager among all threads
// writing index data filesets.
func NewIndexClaimsManager(opts Options) (IndexClaimsManager, error) {
	if globalIndexClaimsManagersCheckEnabled && atomic.AddUint64(&globalIndexClaimsManagers, 1) != 1 {
		err := errMustUseSingleClaimsManager
		instrument.EmitAndLogInvariantViolation(opts.InstrumentOptions(),
			func(l *zap.Logger) {
				l.Error(err.Error())
			})
		return nil, err
	}

	return &indexClaimsManager{
		filePathPrefix:                opts.FilePathPrefix(),
		nowFn:                         opts.ClockOptions().NowFn(),
		volumeIndexClaims:             make(map[string]map[xtime.UnixNano]volumeIndexClaim),
		nextIndexFileSetVolumeIndexFn: NextIndexFileSetVolumeIndex,
	}, nil
}

func (i *indexClaimsManager) ClaimNextIndexFileSetVolumeIndex(
	md namespace.Metadata,
	blockStart xtime.UnixNano,
) (int, error) {
	i.Lock()
	now := xtime.ToUnixNano(i.nowFn())
	earliestBlockStart := retention.FlushTimeStartForRetentionPeriod(
		md.Options().RetentionOptions().RetentionPeriod(),
		md.Options().IndexOptions().BlockSize(),
		now,
	)
	defer func() {
		i.deleteOutOfRetentionEntriesWithLock(md.ID(), earliestBlockStart)
		i.Unlock()
	}()

	// Reject out of retention claims.
	if blockStart.Before(earliestBlockStart) {
		return 0, ErrIndexOutOfRetention
	}

	volumeIndexClaimsByBlockStart, ok := i.volumeIndexClaims[md.ID().String()]
	if !ok {
		volumeIndexClaimsByBlockStart = make(map[xtime.UnixNano]volumeIndexClaim)
		i.volumeIndexClaims[md.ID().String()] = volumeIndexClaimsByBlockStart
	}

	if curr, ok := volumeIndexClaimsByBlockStart[blockStart]; ok {
		// Already had a previous claim, return the next claim.
		next := curr
		next.volumeIndex++
		volumeIndexClaimsByBlockStart[blockStart] = next
		return next.volumeIndex, nil
	}

	volumeIndex, err := i.nextIndexFileSetVolumeIndexFn(i.filePathPrefix, md.ID(),
		blockStart)
	if err != nil {
		return 0, err
	}
	volumeIndexClaimsByBlockStart[blockStart] = volumeIndexClaim{
		volumeIndex: volumeIndex,
	}
	return volumeIndex, nil
}

func (i *indexClaimsManager) deleteOutOfRetentionEntriesWithLock(
	nsID ident.ID,
	earliestBlockStart xtime.UnixNano,
) {
	// ns ID already exists at this point since the delete call is deferred.
	for blockStart := range i.volumeIndexClaims[nsID.String()] {
		if blockStart.Before(earliestBlockStart) {
			delete(i.volumeIndexClaims[nsID.String()], blockStart)
		}
	}
}

type nextIndexFileSetVolumeIndexFn func(
	filePathPrefix string,
	namespace ident.ID,
	blockStart xtime.UnixNano,
) (int, error)
