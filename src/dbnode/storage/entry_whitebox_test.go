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

package storage

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/m3db/m3/src/dbnode/storage/index"
	"github.com/m3db/m3/src/dbnode/storage/series"
	"github.com/m3db/m3/src/dbnode/ts/writes"
	"github.com/m3db/m3/src/m3ninx/doc"
	"github.com/m3db/m3/src/x/context"
	"github.com/m3db/m3/src/x/ident"
	xtest "github.com/m3db/m3/src/x/test"
	xtime "github.com/m3db/m3/src/x/time"
)

func TestEntryIndexAttemptRotatesSlice(t *testing.T) {
	ctrl := xtest.NewController(t)
	defer ctrl.Finish()

	e := NewEntry(NewEntryOptions{Series: newMockSeries(ctrl)})
	for i := 0; i < 10; i++ {
		ti := newTime(i)
		require.True(t, e.NeedsIndexUpdate(ti))
		require.Equal(t, i+1, e.IndexedBlockCount())
	}

	// ensure only the latest ones are held on to
	for i := 9; i >= 7; i-- {
		ti := newTime(i)
		require.False(t, e.NeedsIndexUpdate(ti))
	}
}

func TestEntryIndexSeriesRef(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()
	blockStart := newTime(0)
	mockIndexWriter := NewMockIndexWriter(ctrl)
	mockIndexWriter.EXPECT().BlockStartForWriteTime(blockStart).
		Return(blockStart).
		Times(2)

	mockSeries := series.NewMockDatabaseSeries(ctrl)
	mockSeries.EXPECT().ID().Return(ident.StringID("foo"))
	mockSeries.EXPECT().Metadata().Return(doc.Metadata{})
	mockSeries.EXPECT().Write(
		context.NewBackground(),
		blockStart,
		1.0,
		xtime.Second,
		nil,
		series.WriteOptions{},
	).Return(true, series.WarmWrite, nil)

	e := NewEntry(NewEntryOptions{
		Series:      mockSeries,
		IndexWriter: mockIndexWriter,
		NowFn: func() time.Time {
			return now
		},
	})

	mockIndexWriter.EXPECT().WritePending([]writes.PendingIndexInsert{
		{
			Entry: index.WriteBatchEntry{
				Timestamp:     blockStart,
				OnIndexSeries: e,
				EnqueuedAt:    now,
			},
			Document: doc.Metadata{},
		},
	}).Return(nil)

	ok, _, err := e.Write(
		context.NewBackground(),
		blockStart,
		1.0,
		xtime.Second,
		nil,
		series.WriteOptions{},
	)
	require.True(t, ok)
	require.NoError(t, err)
}
