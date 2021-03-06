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

package index

import (
	"github.com/m3db/m3/src/x/pool"
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

// AggregateResultsEntryArrayPool provides a pool for aggregateResultsEntry slices.
type AggregateResultsEntryArrayPool interface {
	// Init initializes the array pool, it needs to be called
	// before Get/Put use.
	Init()

	// Get returns the a slice from the pool.
	Get() []AggregateResultsEntry

	// Put returns the provided slice to the pool.
	Put(elems []AggregateResultsEntry)
}

type AggregateResultsEntryFinalizeFn func([]AggregateResultsEntry) []AggregateResultsEntry

type AggregateResultsEntryArrayPoolOpts struct {
	Options     pool.ObjectPoolOptions
	Capacity    int
	MaxCapacity int
	FinalizeFn  AggregateResultsEntryFinalizeFn
}

type AggregateResultsEntryArrPool struct {
	opts AggregateResultsEntryArrayPoolOpts
	pool pool.ObjectPool
}

func NewAggregateResultsEntryArrayPool(opts AggregateResultsEntryArrayPoolOpts) AggregateResultsEntryArrayPool {
	if opts.FinalizeFn == nil {
		opts.FinalizeFn = defaultAggregateResultsEntryFinalizerFn
	}
	p := pool.NewObjectPool(opts.Options)
	return &AggregateResultsEntryArrPool{opts, p}
}

func (p *AggregateResultsEntryArrPool) Init() {
	p.pool.Init(func() interface{} {
		return make([]AggregateResultsEntry, 0, p.opts.Capacity)
	})
}

func (p *AggregateResultsEntryArrPool) Get() []AggregateResultsEntry {
	return p.pool.Get().([]AggregateResultsEntry)
}

func (p *AggregateResultsEntryArrPool) Put(arr []AggregateResultsEntry) {
	arr = p.opts.FinalizeFn(arr)
	if max := p.opts.MaxCapacity; max > 0 && cap(arr) > max {
		return
	}
	p.pool.Put(arr)
}

func defaultAggregateResultsEntryFinalizerFn(elems []AggregateResultsEntry) []AggregateResultsEntry {
	var empty AggregateResultsEntry
	for i := range elems {
		elems[i] = empty
	}
	elems = elems[:0]
	return elems
}

type AggregateResultsEntryArr []AggregateResultsEntry

func (elems AggregateResultsEntryArr) grow(n int) []AggregateResultsEntry {
	if cap(elems) < n {
		elems = make([]AggregateResultsEntry, n)
	}
	elems = elems[:n]
	// following compiler optimized memcpy impl
	// https://github.com/golang/go/wiki/CompilerOptimizations#optimized-memclr
	var empty AggregateResultsEntry
	for i := range elems {
		elems[i] = empty
	}
	return elems
}
