// Copyright (c) 2021  Uber Technologies, Inc.
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

package xos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestWatchForInterrupt(t *testing.T) {
	interruptCh := make(chan error, 1)
	interruptCh <- NewInterruptError("interrupt time")

	opts := NewInterruptOptions()
	opts.InterruptCh = interruptCh

	_ = WatchForInterrupt(zap.NewNop(), opts)

	select {
	case <-opts.InterruptedCh:
	case <-time.After(5 * time.Second):
		t.Fail()
	}
}

func TestWatchForInterruptCloseTwice(t *testing.T) {
	opts := NewInterruptOptions()
	closer := WatchForInterrupt(zap.NewNop(), opts)

	require.NotPanics(t, func() {
		closer()
		closer()
	})
}
