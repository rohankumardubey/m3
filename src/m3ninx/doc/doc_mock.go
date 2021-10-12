// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/m3ninx/doc/types.go

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

// Package doc is a generated GoMock package.
package doc

import (
	"reflect"

	"github.com/m3db/m3/src/x/time"

	"github.com/golang/mock/gomock"
)

// MockMetadataIterator is a mock of MetadataIterator interface.
type MockMetadataIterator struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataIteratorMockRecorder
}

// MockMetadataIteratorMockRecorder is the mock recorder for MockMetadataIterator.
type MockMetadataIteratorMockRecorder struct {
	mock *MockMetadataIterator
}

// NewMockMetadataIterator creates a new mock instance.
func NewMockMetadataIterator(ctrl *gomock.Controller) *MockMetadataIterator {
	mock := &MockMetadataIterator{ctrl: ctrl}
	mock.recorder = &MockMetadataIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataIterator) EXPECT() *MockMetadataIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMetadataIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockMetadataIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMetadataIterator)(nil).Close))
}

// Current mocks base method.
func (m *MockMetadataIterator) Current() Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(Metadata)
	return ret0
}

// Current indicates an expected call of Current.
func (mr *MockMetadataIteratorMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockMetadataIterator)(nil).Current))
}

// Err mocks base method.
func (m *MockMetadataIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockMetadataIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockMetadataIterator)(nil).Err))
}

// Next mocks base method.
func (m *MockMetadataIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockMetadataIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockMetadataIterator)(nil).Next))
}

// MockIterator is a mock of Iterator interface.
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIteratorMockRecorder
}

// MockIteratorMockRecorder is the mock recorder for MockIterator.
type MockIteratorMockRecorder struct {
	mock *MockIterator
}

// NewMockIterator creates a new mock instance.
func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &MockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIterator) EXPECT() *MockIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIterator)(nil).Close))
}

// Current mocks base method.
func (m *MockIterator) Current() Document {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(Document)
	return ret0
}

// Current indicates an expected call of Current.
func (mr *MockIteratorMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockIterator)(nil).Current))
}

// Err mocks base method.
func (m *MockIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockIterator)(nil).Err))
}

// Next mocks base method.
func (m *MockIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIterator)(nil).Next))
}

// MockQueryDocIterator is a mock of QueryDocIterator interface.
type MockQueryDocIterator struct {
	ctrl     *gomock.Controller
	recorder *MockQueryDocIteratorMockRecorder
}

// MockQueryDocIteratorMockRecorder is the mock recorder for MockQueryDocIterator.
type MockQueryDocIteratorMockRecorder struct {
	mock *MockQueryDocIterator
}

// NewMockQueryDocIterator creates a new mock instance.
func NewMockQueryDocIterator(ctrl *gomock.Controller) *MockQueryDocIterator {
	mock := &MockQueryDocIterator{ctrl: ctrl}
	mock.recorder = &MockQueryDocIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryDocIterator) EXPECT() *MockQueryDocIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockQueryDocIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockQueryDocIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockQueryDocIterator)(nil).Close))
}

// Current mocks base method.
func (m *MockQueryDocIterator) Current() Document {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(Document)
	return ret0
}

// Current indicates an expected call of Current.
func (mr *MockQueryDocIteratorMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockQueryDocIterator)(nil).Current))
}

// Done mocks base method.
func (m *MockQueryDocIterator) Done() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockQueryDocIteratorMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockQueryDocIterator)(nil).Done))
}

// Err mocks base method.
func (m *MockQueryDocIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockQueryDocIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockQueryDocIterator)(nil).Err))
}

// Next mocks base method.
func (m *MockQueryDocIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockQueryDocIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockQueryDocIterator)(nil).Next))
}

// MockOnIndexSeries is a mock of OnIndexSeries interface.
type MockOnIndexSeries struct {
	ctrl     *gomock.Controller
	recorder *MockOnIndexSeriesMockRecorder
}

// MockOnIndexSeriesMockRecorder is the mock recorder for MockOnIndexSeries.
type MockOnIndexSeriesMockRecorder struct {
	mock *MockOnIndexSeries
}

// NewMockOnIndexSeries creates a new mock instance.
func NewMockOnIndexSeries(ctrl *gomock.Controller) *MockOnIndexSeries {
	mock := &MockOnIndexSeries{ctrl: ctrl}
	mock.recorder = &MockOnIndexSeriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOnIndexSeries) EXPECT() *MockOnIndexSeriesMockRecorder {
	return m.recorder
}

// IfAlreadyIndexedMarkIndexSuccessAndFinalize mocks base method.
func (m *MockOnIndexSeries) IfAlreadyIndexedMarkIndexSuccessAndFinalize(blockStart time.UnixNano) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IfAlreadyIndexedMarkIndexSuccessAndFinalize", blockStart)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IfAlreadyIndexedMarkIndexSuccessAndFinalize indicates an expected call of IfAlreadyIndexedMarkIndexSuccessAndFinalize.
func (mr *MockOnIndexSeriesMockRecorder) IfAlreadyIndexedMarkIndexSuccessAndFinalize(blockStart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IfAlreadyIndexedMarkIndexSuccessAndFinalize", reflect.TypeOf((*MockOnIndexSeries)(nil).IfAlreadyIndexedMarkIndexSuccessAndFinalize), blockStart)
}

// IndexedForBlockStart mocks base method.
func (m *MockOnIndexSeries) IndexedForBlockStart(blockStart time.UnixNano) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexedForBlockStart", blockStart)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IndexedForBlockStart indicates an expected call of IndexedForBlockStart.
func (mr *MockOnIndexSeriesMockRecorder) IndexedForBlockStart(blockStart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexedForBlockStart", reflect.TypeOf((*MockOnIndexSeries)(nil).IndexedForBlockStart), blockStart)
}

// IndexedRange mocks base method.
func (m *MockOnIndexSeries) IndexedRange() (time.UnixNano, time.UnixNano) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndexedRange")
	ret0, _ := ret[0].(time.UnixNano)
	ret1, _ := ret[1].(time.UnixNano)
	return ret0, ret1
}

// IndexedRange indicates an expected call of IndexedRange.
func (mr *MockOnIndexSeriesMockRecorder) IndexedRange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexedRange", reflect.TypeOf((*MockOnIndexSeries)(nil).IndexedRange))
}

// NeedsIndexGarbageCollected mocks base method.
func (m *MockOnIndexSeries) NeedsIndexGarbageCollected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsIndexGarbageCollected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NeedsIndexGarbageCollected indicates an expected call of NeedsIndexGarbageCollected.
func (mr *MockOnIndexSeriesMockRecorder) NeedsIndexGarbageCollected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsIndexGarbageCollected", reflect.TypeOf((*MockOnIndexSeries)(nil).NeedsIndexGarbageCollected))
}

// NeedsIndexUpdate mocks base method.
func (m *MockOnIndexSeries) NeedsIndexUpdate(indexBlockStartForWrite time.UnixNano) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsIndexUpdate", indexBlockStartForWrite)
	ret0, _ := ret[0].(bool)
	return ret0
}

// NeedsIndexUpdate indicates an expected call of NeedsIndexUpdate.
func (mr *MockOnIndexSeriesMockRecorder) NeedsIndexUpdate(indexBlockStartForWrite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsIndexUpdate", reflect.TypeOf((*MockOnIndexSeries)(nil).NeedsIndexUpdate), indexBlockStartForWrite)
}

// OnIndexFinalize mocks base method.
func (m *MockOnIndexSeries) OnIndexFinalize(blockStart time.UnixNano) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnIndexFinalize", blockStart)
}

// OnIndexFinalize indicates an expected call of OnIndexFinalize.
func (mr *MockOnIndexSeriesMockRecorder) OnIndexFinalize(blockStart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnIndexFinalize", reflect.TypeOf((*MockOnIndexSeries)(nil).OnIndexFinalize), blockStart)
}

// OnIndexPrepare mocks base method.
func (m *MockOnIndexSeries) OnIndexPrepare(blockStart time.UnixNano) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnIndexPrepare", blockStart)
}

// OnIndexPrepare indicates an expected call of OnIndexPrepare.
func (mr *MockOnIndexSeriesMockRecorder) OnIndexPrepare(blockStart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnIndexPrepare", reflect.TypeOf((*MockOnIndexSeries)(nil).OnIndexPrepare), blockStart)
}

// OnIndexSuccess mocks base method.
func (m *MockOnIndexSeries) OnIndexSuccess(blockStart time.UnixNano) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnIndexSuccess", blockStart)
}

// OnIndexSuccess indicates an expected call of OnIndexSuccess.
func (mr *MockOnIndexSeriesMockRecorder) OnIndexSuccess(blockStart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnIndexSuccess", reflect.TypeOf((*MockOnIndexSeries)(nil).OnIndexSuccess), blockStart)
}

// TryMarkIndexGarbageCollected mocks base method.
func (m *MockOnIndexSeries) TryMarkIndexGarbageCollected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryMarkIndexGarbageCollected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TryMarkIndexGarbageCollected indicates an expected call of TryMarkIndexGarbageCollected.
func (mr *MockOnIndexSeriesMockRecorder) TryMarkIndexGarbageCollected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryMarkIndexGarbageCollected", reflect.TypeOf((*MockOnIndexSeries)(nil).TryMarkIndexGarbageCollected))
}
