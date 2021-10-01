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

package aggregator

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/m3db/m3/src/metrics/metadata"
	"github.com/m3db/m3/src/metrics/metric/aggregated"
	"github.com/m3db/m3/src/metrics/metric/unaggregated"
	"github.com/m3db/m3/src/metrics/transformation"
	xtime "github.com/m3db/m3/src/x/time"

	"github.com/willf/bitset"
)

type lockedTimerAggregation struct {
	sync.Mutex

	dirty       bool
	flushed     bool
	closed      bool
	sourcesSeen map[uint32]*bitset.BitSet
	aggregation timerAggregation
	prevValues  []float64 // the previously emitted values (one per aggregation type).
}

type timedTimer struct {
	startAtNanos     int64 // start time of an aggregation window
	lockedAgg        *lockedTimerAggregation
	onConsumeExpired bool
	// this is needed to account for the extra data we are
	// tracking in timeAggregations instead of consumedVals map
	gcEligible bool
}

func (ta *timedTimer) Release() {
	ta.startAtNanos = 0
	ta.lockedAgg = nil
}

// TimerElem is an element storing time-bucketed aggregations.
type TimerElem struct {
	elemBase
	timerElemBase

	// use this instead of consumedValues
	//
	values []timedTimer // metric aggregations sorted by time in ascending order

	// internal consume state that does not need to be synchronized.
	toConsume      []timedTimer // small buffer to avoid memory allocations during consumption
	consumedValues valuesByTime
}

// NewTimerElem returns a new TimerElem.
func NewTimerElem(data ElemData, opts Options) (*TimerElem, error) {
	e := &TimerElem{
		elemBase: newElemBase(opts),
		values:   make([]timedTimer, 0, defaultNumAggregations), // in most cases values will have two entries
	}
	if err := e.ResetSetData(data); err != nil {
		return nil, err
	}
	return e, nil
}

// MustNewTimerElem returns a new TimerElem and panics if an error occurs.
func MustNewTimerElem(data ElemData, opts Options) *TimerElem {
	elem, err := NewTimerElem(data, opts)
	if err != nil {
		panic(fmt.Errorf("unable to create element: %v", err))
	}
	return elem
}

// ResetSetData resets the element and sets data.
func (e *TimerElem) ResetSetData(data ElemData) error {
	useDefaultAggregation := data.AggTypes.IsDefault()
	if useDefaultAggregation {
		data.AggTypes = e.DefaultAggregationTypes(e.aggTypesOpts)
	}
	if err := e.elemBase.resetSetData(data, useDefaultAggregation); err != nil {
		return err
	}
	return e.timerElemBase.ResetSetData(e.aggTypesOpts, data.AggTypes, useDefaultAggregation)
}

// ResendEnabled returns true if resends are enabled for the element.
func (e *TimerElem) ResendEnabled() bool {
	return e.resendEnabled
}

// AddUnion adds a metric value union at a given timestamp.
func (e *TimerElem) AddUnion(timestamp time.Time, mu unaggregated.MetricUnion) error {
	alignedStart := timestamp.Truncate(e.sp.Resolution().Window).UnixNano()
	lockedAgg, err := e.findOrCreate(alignedStart, createAggregationOptions{})
	if err != nil {
		return err
	}
	lockedAgg.Lock()
	if lockedAgg.closed {
		lockedAgg.Unlock()
		return errAggregationClosed
	}
	lockedAgg.aggregation.AddUnion(timestamp, mu)
	lockedAgg.dirty = true
	lockedAgg.Unlock()
	return nil
}

// AddValue adds a metric value at a given timestamp.
func (e *TimerElem) AddValue(timestamp time.Time, value float64, annotation []byte) error {
	alignedStart := timestamp.Truncate(e.sp.Resolution().Window).UnixNano()
	lockedAgg, err := e.findOrCreate(alignedStart, createAggregationOptions{})
	if err != nil {
		return err
	}
	lockedAgg.Lock()
	if lockedAgg.closed {
		lockedAgg.Unlock()
		return errAggregationClosed
	}
	lockedAgg.aggregation.Add(timestamp, value, annotation)
	lockedAgg.dirty = true
	lockedAgg.Unlock()
	return nil
}

// AddUnique adds a metric value from a given source at a given timestamp.
// If previous values from the same source have already been added to the
// same aggregation, the incoming value is discarded.
//nolint: dupl
func (e *TimerElem) AddUnique(
	timestamp time.Time,
	metric aggregated.ForwardedMetric,
	metadata metadata.ForwardMetadata,
) error {
	alignedStart := timestamp.Truncate(e.sp.Resolution().Window).UnixNano()
	lockedAgg, err := e.findOrCreate(alignedStart, createAggregationOptions{initSourceSet: true})
	if err != nil {
		return err
	}
	lockedAgg.Lock()
	if lockedAgg.closed {
		lockedAgg.Unlock()
		return errAggregationClosed
	}
	versionsSeen := lockedAgg.sourcesSeen[metadata.SourceID]
	if versionsSeen == nil {
		// N.B - these bitsets will be transitively cached through the cached sources seen.
		versionsSeen = bitset.New(defaultNumVersions)
		lockedAgg.sourcesSeen[metadata.SourceID] = versionsSeen
	}
	version := uint(metric.Version)
	if versionsSeen.Test(version) {
		lockedAgg.Unlock()
		return errDuplicateForwardingSource
	}
	versionsSeen.Set(version)

	if metric.Version > 0 {
		e.metrics.updatedValues.Inc(1)
		for i := range metric.Values {
			if err := lockedAgg.aggregation.UpdateVal(timestamp, metric.Values[i], metric.PrevValues[i]); err != nil {
				return err
			}
		}
	} else {
		for _, v := range metric.Values {
			lockedAgg.aggregation.Add(timestamp, v, metric.Annotation)
		}
	}

	lockedAgg.dirty = true
	lockedAgg.Unlock()
	return nil
}

// Consume consumes values before a given time and removes them from the element
// after they are consumed, returning whether the element can be collected after
// the consumption is completed.
// NB: Consume is not thread-safe and must be called within a single goroutine
// to avoid race conditions.
func (e *TimerElem) Consume(
	targetNanos int64,
	isEarlierThanFn isEarlierThanFn,
	timestampNanosFn timestampNanosFn,
	flushLocalFn flushLocalMetricFn,
	flushForwardedFn flushForwardedMetricFn,
	onForwardedFlushedFn onForwardingElemFlushedFn,
) bool {
	resolution := e.sp.Resolution().Window
	e.Lock()
	if e.closed {
		e.Unlock()
		return false
	}
	e.toConsume = e.toConsume[:0]
	e.consumedValues = nil

	// Evaluate and GC expired items.
	var (
		prev           *timedTimer
		firstToConsume bool
	)
	valuesForConsideration := e.values
	e.values = e.values[:0]
	for i, value := range valuesForConsideration {
		if value.gcEligible {
			continue
		}

		if !isEarlierThanFn(value.startAtNanos, resolution, targetNanos) {
			e.values = append(e.values, value)
			continue
		}

		if !firstToConsume && i > 0 {
			prev = &valuesForConsideration[i-1]
		}
		firstToConsume = true

		expired := true
		if e.resendEnabled {
			// If resend is enabled, we only expire if the value is now outside the buffer past. It is safe to expire
			// since any metrics intended for this value are rejected for being too late.
			expiredNanos := targetNanos - e.bufferForPastTimedMetricFn(resolution).Nanoseconds()
			expired = value.startAtNanos < expiredNanos
		}

		// Modify the by value copy with whether it needs time flush and accumulate.
		copiedValue := value
		copiedValue.onConsumeExpired = expired
		if !value.onConsumeExpired {
			e.toConsume = append(e.toConsume, copiedValue)
		}

		// Keep item. Expired values are GC'd below after consuming.
		e.values = append(e.values, value)
	}
	canCollect := len(e.values) == 0 && e.tombstoned
	e.Unlock()

	var (
		cascadeDirty bool
	)
	// Process the aggregations that are ready for consumption.
	for i := range e.toConsume {
		expired := e.toConsume[i].onConsumeExpired
		timeNanos := timestampNanosFn(e.toConsume[i].startAtNanos, resolution)

		var (
			prevAgg       *timedTimer
			prevLockedAgg *lockedTimerAggregation
		)
		if i > 0 {
			prevAgg = &e.toConsume[i-1]
			prevLockedAgg = prevAgg.lockedAgg
		} else if prev != nil {
			prevAgg = prev
			prevLockedAgg = prevAgg.lockedAgg
		}

		if prevLockedAgg != nil {
			prevLockedAgg.Lock()
		}
		e.toConsume[i].lockedAgg.Lock()

		// if a previous timestamps was dirty, that value might impact a future derivative calculation, so
		// cascade the dirty bit to all succeeding values. there is a check later to not resend a value if it doesn't
		// change, so it's ok to optimistically mark dirty.
		if cascadeDirty || e.toConsume[i].lockedAgg.dirty {
			cascadeDirty = e.processValueWithAggregationLock(
				timeNanos,
				e.toConsume[i].lockedAgg,
				prevLockedAgg,
				flushLocalFn,
				flushForwardedFn,
				resolution,
			)
			e.toConsume[i].lockedAgg.flushed = true
			e.toConsume[i].lockedAgg.dirty = false
		}

		// Closes the aggregation object after it's processed.
		if expired {
			// Cleanup expired item.
			e.toConsume[i].lockedAgg.closed = true
			e.toConsume[i].lockedAgg.aggregation.Close()
			if e.toConsume[i].lockedAgg.sourcesSeen != nil {
				e.cachedSourceSetsLock.Lock()
				// This is to make sure there aren't too many cached source sets taking up
				// too much space.
				if len(e.cachedSourceSets) < e.opts.MaxNumCachedSourceSets() {
					e.cachedSourceSets = append(e.cachedSourceSets, e.toConsume[i].lockedAgg.sourcesSeen)
				}
				e.cachedSourceSetsLock.Unlock()
				e.toConsume[i].lockedAgg.sourcesSeen = nil
			}
		}

		e.toConsume[i].lockedAgg.Unlock()
		if expired {
			//e.toConsume[i].Release()
			if prevAgg != nil {
				//prevAgg.gcEligible = true
			}
		}

		if prevLockedAgg != nil {
			prevLockedAgg.Unlock()
		}
	}

	if e.parsedPipeline.HasRollup {
		forwardedAggregationKey, _ := e.ForwardedAggregationKey()
		onForwardedFlushedFn(e.onForwardedAggregationWrittenFn, forwardedAggregationKey)
	}

	return canCollect
}

// Close closes the element.
func (e *TimerElem) Close() {
	e.Lock()
	if e.closed {
		e.Unlock()
		return
	}
	e.closed = true
	e.id = nil
	e.parsedPipeline = parsedPipeline{}
	e.writeForwardedMetricFn = nil
	e.onForwardedAggregationWrittenFn = nil
	for idx := range e.cachedSourceSets {
		e.cachedSourceSets[idx] = nil
	}
	e.cachedSourceSets = nil
	for idx := range e.values {
		// Close the underlying aggregation objects.
		e.values[idx].lockedAgg.sourcesSeen = nil
		e.values[idx].lockedAgg.aggregation.Close()
		e.values[idx].Release()
	}
	e.values = e.values[:0]
	e.timerElemBase.Close()
	aggTypesPool := e.aggTypesOpts.TypesPool()
	pool := e.ElemPool(e.opts)
	e.Unlock()

	// internal consumption state that doesn't need to be synchronized.
	e.toConsume = e.toConsume[:0]

	if !e.useDefaultAggregation {
		aggTypesPool.Put(e.aggTypes)
	}
	pool.Put(e)
}

// findOrCreate finds the aggregation for a given time, or creates one
// if it doesn't exist.
func (e *TimerElem) findOrCreate(
	alignedStart int64,
	createOpts createAggregationOptions,
) (*lockedTimerAggregation, error) {
	e.RLock()
	if e.closed {
		e.RUnlock()
		return nil, errElemClosed
	}
	idx, found := e.indexOfWithLock(alignedStart)
	if found {
		agg := e.values[idx].lockedAgg
		e.RUnlock()
		return agg, nil
	}
	e.RUnlock()

	e.Lock()
	if e.closed {
		e.Unlock()
		return nil, errElemClosed
	}
	idx, found = e.indexOfWithLock(alignedStart)
	if found {
		agg := e.values[idx].lockedAgg
		e.Unlock()
		return agg, nil
	}

	// If not found, create a new aggregation.
	numValues := len(e.values)
	e.values = append(e.values, timedTimer{})
	copy(e.values[idx+1:numValues+1], e.values[idx:numValues])

	var sourcesSeen map[uint32]*bitset.BitSet
	if createOpts.initSourceSet {
		e.cachedSourceSetsLock.Lock()
		if numCachedSourceSets := len(e.cachedSourceSets); numCachedSourceSets > 0 {
			sourcesSeen = e.cachedSourceSets[numCachedSourceSets-1]
			e.cachedSourceSets[numCachedSourceSets-1] = nil
			e.cachedSourceSets = e.cachedSourceSets[:numCachedSourceSets-1]
			for _, bs := range sourcesSeen {
				bs.ClearAll()
			}
		} else {
			sourcesSeen = make(map[uint32]*bitset.BitSet)
		}
		e.cachedSourceSetsLock.Unlock()
	}
	e.values[idx] = timedTimer{
		startAtNanos: alignedStart,
		lockedAgg: &lockedTimerAggregation{
			sourcesSeen: sourcesSeen,
			aggregation: e.NewAggregation(e.opts, e.aggOpts),
			prevValues:  make([]float64, len(e.aggTypes)),
		},
	}
	agg := e.values[idx].lockedAgg
	e.Unlock()
	return agg, nil
}

// indexOfWithLock finds the smallest element index whose timestamp
// is no smaller than the start time passed in, and true if it's an
// exact match, false otherwise.
func (e *TimerElem) indexOfWithLock(alignedStart int64) (int, bool) {
	numValues := len(e.values)
	// Optimize for the common case.
	if numValues > 0 && e.values[numValues-1].startAtNanos == alignedStart {
		return numValues - 1, true
	}
	// Binary search for the unusual case. We intentionally do not
	// use the sort.Search() function because it requires passing
	// in a closure.
	left, right := 0, numValues
	for left < right {
		mid := left + (right-left)/2 // avoid overflow
		if e.values[mid].startAtNanos < alignedStart {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// If the current timestamp is equal to or larger than the target time,
	// return the index as is.
	if left < numValues && e.values[left].startAtNanos == alignedStart {
		return left, true
	}
	return left, false
}

// returns true if a datapoint is emitted.
func (e *TimerElem) processValueWithAggregationLock(
	timeNanos int64,
	lockedAgg *lockedTimerAggregation,
	lockedPrevAgg *lockedTimerAggregation,
	flushLocalFn flushLocalMetricFn,
	flushForwardedFn flushForwardedMetricFn,
	resolution time.Duration) bool {
	var (
		transformations  = e.parsedPipeline.Transformations
		discardNaNValues = e.opts.DiscardNaNAggregatedValues()
		emitted          bool
	)
	for aggTypeIdx, aggType := range e.aggTypes {
		var extraDp transformation.Datapoint
		value := lockedAgg.aggregation.ValueOf(aggType)
		for _, transformOp := range transformations {
			unaryOp, isUnaryOp := transformOp.UnaryTransform()
			binaryOp, isBinaryOp := transformOp.BinaryTransform()
			unaryMultiOp, isUnaryMultiOp := transformOp.UnaryMultiOutputTransform()
			switch {
			case isUnaryOp:
				curr := transformation.Datapoint{
					TimeNanos: int64(timeNanos),
					Value:     value,
				}

				res := unaryOp.Evaluate(curr)

				value = res.Value

			case isBinaryOp:
				// lazily construct consumedValues since they are only needed by binary transforms.
				if e.consumedValues == nil {
					e.consumedValues = make(valuesByTime)
				}
				var prevDp transformation.Datapoint
				if lockedPrevAgg == nil {
					prevDp = transformation.Datapoint{
						Value: nan,
					}
				} else {
					prevDp = transformation.Datapoint{
						Value: lockedPrevAgg.prevValues[aggTypeIdx],
					}
				}

				curr := transformation.Datapoint{
					TimeNanos: int64(timeNanos),
					Value:     value,
				}
				res := binaryOp.Evaluate(prevDp, curr, transformation.FeatureFlags{})

				// NB: we only need to record the value needed for derivative transformations.
				// We currently only support first-order derivative transformations so we only
				// need to keep one value. In the future if we need to support higher-order
				// derivative transformations, we need to store an array of values here.
				if !math.IsNaN(curr.Value) {
					t := xtime.UnixNano(timeNanos)
					if e.consumedValues[t] == nil {
						e.consumedValues[t] = make([]transformation.Datapoint, len(e.aggTypes))
					}
					e.consumedValues[t][aggTypeIdx] = curr
				}

				value = res.Value
			case isUnaryMultiOp:
				curr := transformation.Datapoint{
					TimeNanos: int64(timeNanos),
					Value:     value,
				}

				var res transformation.Datapoint
				res, extraDp = unaryMultiOp.Evaluate(curr, resolution)
				value = res.Value
			}
		}

		if discardNaNValues && math.IsNaN(value) {
			continue
		}

		// It's ok to send a 0 prevValue on the first forward because it's not used in AddUnique unless it's a
		// resend (version > 0)
		prevValue := lockedAgg.prevValues[aggTypeIdx]
		lockedAgg.prevValues[aggTypeIdx] = value
		if lockedAgg.flushed {
			// no need to resend a value that hasn't changed.
			if (math.IsNaN(prevValue) && math.IsNaN(value)) || (prevValue == value) {
				continue
			}
		}
		emitted = true

		if !e.parsedPipeline.HasRollup {
			toFlush := make([]transformation.Datapoint, 0, 2)
			toFlush = append(toFlush, transformation.Datapoint{
				TimeNanos: int64(timeNanos),
				Value:     value,
			})
			if extraDp.TimeNanos != 0 {
				toFlush = append(toFlush, extraDp)
			}
			for _, point := range toFlush {
				switch e.idPrefixSuffixType {
				case NoPrefixNoSuffix:
					flushLocalFn(nil, e.id, nil, point.TimeNanos, point.Value, lockedAgg.aggregation.Annotation(), e.sp)
				case WithPrefixWithSuffix:
					flushLocalFn(e.FullPrefix(e.opts), e.id, e.TypeStringFor(e.aggTypesOpts, aggType),
						point.TimeNanos, point.Value, lockedAgg.aggregation.Annotation(), e.sp)
				}
			}
		} else {
			forwardedAggregationKey, _ := e.ForwardedAggregationKey()
			flushForwardedFn(e.writeForwardedMetricFn, forwardedAggregationKey,
				int64(timeNanos), value, prevValue, lockedAgg.aggregation.Annotation())
		}
	}
	return emitted
}
