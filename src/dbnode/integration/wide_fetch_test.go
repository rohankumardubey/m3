// +build integration
//
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

package integration

import (
	"fmt"
	"io/ioutil"
	"sort"
	"sync/atomic"
	"testing"
	"time"

	"github.com/m3db/m3/src/dbnode/namespace"
	"github.com/m3db/m3/src/dbnode/persist/fs"
	"github.com/m3db/m3/src/dbnode/sharding"
	"github.com/m3db/m3/src/dbnode/storage"
	"github.com/m3db/m3/src/dbnode/storage/index"
	"github.com/m3db/m3/src/m3ninx/idx"
	xclock "github.com/m3db/m3/src/x/clock"
	"github.com/m3db/m3/src/x/context"
	"github.com/m3db/m3/src/x/ident"
	xtest "github.com/m3db/m3/src/x/test"
	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type shardedIndexChecksum struct {
	shard     uint32
	checksums []ident.IndexChecksum
}

func buildExpectedChecksumsByShard(
	ids []string,
	allowedShards []uint32,
	shardSet sharding.ShardSet,
) []ident.IndexChecksum {
	shardedChecksums := make([]shardedIndexChecksum, 0, len(ids))
	count := 0
	for i, id := range ids {
		checksum := ident.IndexChecksum{ID: []byte(id), Checksum: int64(i)}
		shard := shardSet.Lookup(ident.BytesID([]byte(id)))

		if len(allowedShards) > 0 {
			shardInUse := false
			for _, allowed := range allowedShards {
				if allowed == shard {
					shardInUse = true
					break
				}
			}

			if !shardInUse {
				continue
			}
		}

		found := false
		for idx, sharded := range shardedChecksums {
			if shard != sharded.shard {
				continue
			}

			found = true
			shardedChecksums[idx].checksums = append(sharded.checksums, checksum)
			break
		}

		if found {
			continue
		}

		count++
		shardedChecksums = append(shardedChecksums, shardedIndexChecksum{
			shard:     shard,
			checksums: []ident.IndexChecksum{checksum},
		})
	}

	sort.Slice(shardedChecksums, func(i, j int) bool {
		return shardedChecksums[i].shard < shardedChecksums[j].shard
	})

	checksums := make([]ident.IndexChecksum, 0, count)
	for _, sharded := range shardedChecksums {
		checksums = append(checksums, sharded.checksums...)
	}

	return checksums
}

func TestWideFetch(t *testing.T) {
	if testing.Short() {
		t.SkipNow() // Just skip if we're doing a short run
	}

	var (
		batchSize     = 15
		seriesCount   = 15
		blockSize     = time.Hour * 2
		verifyTimeout = 2 * time.Minute
	)

	// Test setup
	idxOpts := namespace.NewIndexOptions().
		SetEnabled(true).
		SetBlockSize(time.Hour * 2)
	nsOpts := namespace.NewOptions().
		SetIndexOptions(idxOpts).
		SetRepairEnabled(false).
		SetRetentionOptions(defaultIntegrationTestRetentionOpts)

	nsID := testNamespaces[0]
	nsMetadata, err := namespace.NewMetadata(nsID, nsOpts)
	require.NoError(t, err)

	// Set up file path prefix
	postfix := atomic.AddUint64(&created, 1) - 1
	filePathPrefix, err := ioutil.TempDir("", fmt.Sprintf("integration-test-%d", postfix))
	require.NoError(t, err)

	testOpts := NewTestOptions(t).
		SetTickMinimumInterval(time.Second).
		SetNamespaces([]namespace.Metadata{nsMetadata}).
		SetFilePathPrefix(filePathPrefix)

	require.NoError(t, err)
	fsOpts := fs.NewOptions().SetFilePathPrefix(filePathPrefix)
	decOpts := fsOpts.DecodingOptions().SetIndexEntryHasher(xtest.NewParsedIndexHasher(t))
	fsOpts = fsOpts.SetDecodingOptions(decOpts)

	testSetup, err := NewTestSetup(t, testOpts, fsOpts,
		func(opt storage.Options) storage.Options {
			return opt.SetWideBatchSize(batchSize)
		})
	require.NoError(t, err)
	defer testSetup.Close()

	// Start the server with filesystem bootstrapper
	log := testSetup.StorageOpts().InstrumentOptions().Logger()
	log.Debug("wide fetch test")
	require.NoError(t, testSetup.StartServer())
	log.Debug("server is now up")

	// Stop the server
	defer func() {
		require.NoError(t, testSetup.StopServer())
		log.Debug("server is now down")
	}()

	// Setup test data
	now := testSetup.NowFn()()
	indexWrites := make(TestIndexWrites, 0, seriesCount)
	ids := make([]string, 0, seriesCount)
	for i := 0; i < seriesCount; i++ {
		// Keep in lex order.
		padCount := i / 10
		pad := ""
		for i := 0; i < padCount; i++ {
			pad = fmt.Sprintf("%so", pad)
		}

		id := fmt.Sprintf("foo%s-%d", pad, i)
		ids = append(ids, id)
		indexWrites = append(indexWrites, testIndexWrite{
			id:    ident.StringID(id),
			tags:  ident.MustNewTagStringsIterator("abc", fmt.Sprintf("def%d", i)),
			ts:    now,
			value: float64(i),
		})
	}

	log.Debug("write test data")
	client := testSetup.M3DBClient()
	session, err := client.DefaultSession()
	require.NoError(t, err)

	start := time.Now()
	indexWrites.Write(t, nsID, session)
	log.Info("test data written", zap.Duration("took", time.Since(start)))

	log.Info("waiting until data is indexed")
	indexed := xclock.WaitUntil(func() bool {
		numIndexed := indexWrites.NumIndexed(t, nsID, session)
		return numIndexed == len(indexWrites)
	}, verifyTimeout)
	require.True(t, indexed)
	log.Info("verified data is indexed", zap.Duration("took", time.Since(start)))

	// Advance time to make sure all data are flushed. Because data
	// are flushed to disk asynchronously, need to poll to check
	// when data are written.
	testSetup.SetNowFn(testSetup.NowFn()().Add(blockSize * 2))
	log.Info("waiting until filesets found on disk")
	found := xclock.WaitUntil(func() bool {
		at := now.Truncate(blockSize).Add(-1 * blockSize)
		filesets, err := fs.IndexFileSetsAt(testSetup.FilePathPrefix(), nsID, at)
		require.NoError(t, err)
		return len(filesets) == 1
	}, verifyTimeout)
	require.True(t, found)
	log.Info("filesets found on disk")

	var (
		ctx      = context.NewContext()
		query    = index.Query{Query: idx.MustCreateRegexpQuery([]byte("abc"), []byte("def.*"))}
		iterOpts = index.IterationOptions{}
	)

	// Verify data.
	chk, err := testSetup.DB().WideQuery(ctx, nsMetadata.ID(), query, now, nil, iterOpts)
	require.NoError(t, err)

	expected := buildExpectedChecksumsByShard(ids, nil, testSetup.ShardSet())
	require.Equal(t, len(expected), len(chk))
	for i, c := range chk {
		assert.Equal(t, expected[i].Checksum, c.Checksum)
	}
}
