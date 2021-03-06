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

package protobuf

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/m3db/m3/src/metrics/encoding"
	"github.com/m3db/m3/src/metrics/generated/proto/metricpb"
	"github.com/m3db/m3/src/x/pool"
)

// UnaggregatedIterator decodes unaggregated metrics.
type UnaggregatedIterator struct {
	pb             metricpb.MetricWithMetadatas
	reader         encoding.ByteReadScanner
	bytesPool      pool.BytesPool
	err            error
	buf            []byte
	msg            encoding.UnaggregatedMessageUnion
	maxMessageSize int
	closed         bool
}

// NewUnaggregatedIterator creates a new unaggregated iterator.
func NewUnaggregatedIterator(
	reader encoding.ByteReadScanner,
	opts UnaggregatedOptions,
) *UnaggregatedIterator {
	bytesPool := opts.BytesPool()
	return &UnaggregatedIterator{
		reader:         reader,
		bytesPool:      bytesPool,
		maxMessageSize: opts.MaxMessageSize(),
		buf:            allocate(bytesPool, opts.InitBufferSize()),
	}
}

// Close closes the iterator.
func (it *UnaggregatedIterator) Close() {
	if it.closed {
		return
	}
	it.closed = true
	it.reader = nil
	it.pb.Reset()
	it.msg = encoding.UnaggregatedMessageUnion{}
	if it.bytesPool != nil && it.buf != nil {
		it.bytesPool.Put(it.buf)
	}
	it.bytesPool = nil
	it.buf = nil
	it.err = nil
}

// Err returns the error encountered during decoding, if any
func (it *UnaggregatedIterator) Err() error { return it.err }

// Current returns the current decoded value
func (it *UnaggregatedIterator) Current() *encoding.UnaggregatedMessageUnion { return &it.msg }

// Next returns true if there are more items to decode.
func (it *UnaggregatedIterator) Next() bool {
	if it.err != nil || it.closed {
		return false
	}
	size, err := it.decodeSize()
	if err != nil {
		return false
	}
	if size > it.maxMessageSize {
		it.err = fmt.Errorf("decoded message size %d is larger than supported max message size %d", size, it.maxMessageSize)
		return false
	}
	if size <= 0 {
		it.err = fmt.Errorf("decoded message size %d is zero or negative", size)
		return false
	}
	it.ensureBufferSize(size)
	if err := it.decodeMessage(size); err != nil {
		return false
	}
	return true
}

func (it *UnaggregatedIterator) decodeSize() (int, error) {
	n, err := binary.ReadVarint(it.reader)
	if err != nil {
		it.err = err
		return 0, err
	}
	return int(n), nil
}

func (it *UnaggregatedIterator) ensureBufferSize(targetSize int) {
	it.buf = ensureBufferSize(it.buf, it.bytesPool, targetSize, dontCopyData)
}

//nolint:gocyclo
func (it *UnaggregatedIterator) decodeMessage(size int) error {
	_, err := io.ReadFull(it.reader, it.buf[:size])
	if err != nil {
		it.err = err
		return err
	}
	ReuseMetricWithMetadatasProto(&it.pb)
	if err := it.pb.Unmarshal(it.buf[:size]); err != nil {
		it.err = err
		return err
	}
	switch it.pb.Type {
	case metricpb.MetricWithMetadatas_COUNTER_WITH_METADATAS:
		it.msg.Type = encoding.CounterWithMetadatasType
		it.err = it.msg.CounterWithMetadatas.FromProto(it.pb.CounterWithMetadatas)
	case metricpb.MetricWithMetadatas_BATCH_TIMER_WITH_METADATAS:
		it.msg.Type = encoding.BatchTimerWithMetadatasType
		it.err = it.msg.BatchTimerWithMetadatas.FromProto(it.pb.BatchTimerWithMetadatas)
	case metricpb.MetricWithMetadatas_GAUGE_WITH_METADATAS:
		it.msg.Type = encoding.GaugeWithMetadatasType
		it.err = it.msg.GaugeWithMetadatas.FromProto(it.pb.GaugeWithMetadatas)
	case metricpb.MetricWithMetadatas_FORWARDED_METRIC_WITH_METADATA:
		it.msg.Type = encoding.ForwardedMetricWithMetadataType
		it.err = it.msg.ForwardedMetricWithMetadata.FromProto(it.pb.ForwardedMetricWithMetadata)
	case metricpb.MetricWithMetadatas_TIMED_METRIC_WITH_METADATA:
		it.msg.Type = encoding.TimedMetricWithMetadataType
		it.err = it.msg.TimedMetricWithMetadata.FromProto(it.pb.TimedMetricWithMetadata)
	case metricpb.MetricWithMetadatas_TIMED_METRIC_WITH_METADATAS:
		it.msg.Type = encoding.TimedMetricWithMetadatasType
		it.err = it.msg.TimedMetricWithMetadatas.FromProto(it.pb.TimedMetricWithMetadatas)
	case metricpb.MetricWithMetadatas_TIMED_METRIC_WITH_STORAGE_POLICY:
		it.msg.Type = encoding.PassthroughMetricWithMetadataType
		it.err = it.msg.PassthroughMetricWithMetadata.FromProto(it.pb.TimedMetricWithStoragePolicy)
	default:
		it.err = fmt.Errorf("unrecognized message type: %v", it.pb.Type)
	}
	return it.err
}
