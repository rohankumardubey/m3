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

package bytes

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

// SliceArrayPool provides a pool for []byte slices.
type SliceArrayPool interface {
	// Init initializes the array pool, it needs to be called
	// before Get/Put use.
	Init()

	// Get returns the a slice from the pool.
	Get() [][]byte

	// Put returns the provided slice to the pool.
	Put(elems [][]byte)
}

type SliceFinalizeFn func([][]byte) [][]byte

type SliceArrayPoolOpts struct {
	Options     pool.ObjectPoolOptions
	Capacity    int
	MaxCapacity int
	FinalizeFn  SliceFinalizeFn
}

type SliceArrPool struct {
	opts SliceArrayPoolOpts
	pool pool.ObjectPool
}

func NewSliceArrayPool(opts SliceArrayPoolOpts) SliceArrayPool {
	if opts.FinalizeFn == nil {
		opts.FinalizeFn = defaultSliceFinalizerFn
	}
	p := pool.NewObjectPool(opts.Options)
	return &SliceArrPool{opts, p}
}

func (p *SliceArrPool) Init() {
	p.pool.Init(func() interface{} {
		return make([][]byte, 0, p.opts.Capacity)
	})
}

func (p *SliceArrPool) Get() [][]byte {
	return p.pool.Get().([][]byte)
}

func (p *SliceArrPool) Put(arr [][]byte) {
	arr = p.opts.FinalizeFn(arr)
	if max := p.opts.MaxCapacity; max > 0 && cap(arr) > max {
		return
	}
	p.pool.Put(arr)
}

func defaultSliceFinalizerFn(elems [][]byte) [][]byte {
	var empty []byte
	for i := range elems {
		elems[i] = empty
	}
	elems = elems[:0]
	return elems
}

type SliceArr [][]byte

func (elems SliceArr) grow(n int) [][]byte {
	if cap(elems) < n {
		elems = make([][]byte, n)
	}
	elems = elems[:n]
	// following compiler optimized memcpy impl
	// https://github.com/golang/go/wiki/CompilerOptimizations#optimized-memclr
	var empty []byte
	for i := range elems {
		elems[i] = empty
	}
	return elems
}
