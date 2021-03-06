// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/m3ninx/index (interfaces: Reader,DocRetriever,MetadataRetriever)

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

// Package index is a generated GoMock package.
package index

import (
	"reflect"

	"github.com/m3db/m3/src/m3ninx/doc"
	"github.com/m3db/m3/src/m3ninx/postings"

	"github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// AllDocs mocks base method.
func (m *MockReader) AllDocs() (IDDocIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllDocs")
	ret0, _ := ret[0].(IDDocIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllDocs indicates an expected call of AllDocs.
func (mr *MockReaderMockRecorder) AllDocs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllDocs", reflect.TypeOf((*MockReader)(nil).AllDocs))
}

// Close mocks base method.
func (m *MockReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReader)(nil).Close))
}

// Doc mocks base method.
func (m *MockReader) Doc(arg0 postings.ID) (doc.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Doc", arg0)
	ret0, _ := ret[0].(doc.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Doc indicates an expected call of Doc.
func (mr *MockReaderMockRecorder) Doc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Doc", reflect.TypeOf((*MockReader)(nil).Doc), arg0)
}

// Docs mocks base method.
func (m *MockReader) Docs(arg0 postings.List) (doc.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Docs", arg0)
	ret0, _ := ret[0].(doc.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Docs indicates an expected call of Docs.
func (mr *MockReaderMockRecorder) Docs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Docs", reflect.TypeOf((*MockReader)(nil).Docs), arg0)
}

// MatchAll mocks base method.
func (m *MockReader) MatchAll() (postings.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchAll")
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchAll indicates an expected call of MatchAll.
func (mr *MockReaderMockRecorder) MatchAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchAll", reflect.TypeOf((*MockReader)(nil).MatchAll))
}

// MatchField mocks base method.
func (m *MockReader) MatchField(arg0 []byte) (postings.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchField", arg0)
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchField indicates an expected call of MatchField.
func (mr *MockReaderMockRecorder) MatchField(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchField", reflect.TypeOf((*MockReader)(nil).MatchField), arg0)
}

// MatchRegexp mocks base method.
func (m *MockReader) MatchRegexp(arg0 []byte, arg1 CompiledRegex) (postings.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchRegexp", arg0, arg1)
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchRegexp indicates an expected call of MatchRegexp.
func (mr *MockReaderMockRecorder) MatchRegexp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchRegexp", reflect.TypeOf((*MockReader)(nil).MatchRegexp), arg0, arg1)
}

// MatchTerm mocks base method.
func (m *MockReader) MatchTerm(arg0, arg1 []byte) (postings.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MatchTerm", arg0, arg1)
	ret0, _ := ret[0].(postings.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchTerm indicates an expected call of MatchTerm.
func (mr *MockReaderMockRecorder) MatchTerm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchTerm", reflect.TypeOf((*MockReader)(nil).MatchTerm), arg0, arg1)
}

// Metadata mocks base method.
func (m *MockReader) Metadata(arg0 postings.ID) (doc.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata", arg0)
	ret0, _ := ret[0].(doc.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metadata indicates an expected call of Metadata.
func (mr *MockReaderMockRecorder) Metadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockReader)(nil).Metadata), arg0)
}

// MetadataIterator mocks base method.
func (m *MockReader) MetadataIterator(arg0 postings.List) (doc.MetadataIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MetadataIterator", arg0)
	ret0, _ := ret[0].(doc.MetadataIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MetadataIterator indicates an expected call of MetadataIterator.
func (mr *MockReaderMockRecorder) MetadataIterator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MetadataIterator", reflect.TypeOf((*MockReader)(nil).MetadataIterator), arg0)
}

// MockDocRetriever is a mock of DocRetriever interface.
type MockDocRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockDocRetrieverMockRecorder
}

// MockDocRetrieverMockRecorder is the mock recorder for MockDocRetriever.
type MockDocRetrieverMockRecorder struct {
	mock *MockDocRetriever
}

// NewMockDocRetriever creates a new mock instance.
func NewMockDocRetriever(ctrl *gomock.Controller) *MockDocRetriever {
	mock := &MockDocRetriever{ctrl: ctrl}
	mock.recorder = &MockDocRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDocRetriever) EXPECT() *MockDocRetrieverMockRecorder {
	return m.recorder
}

// Doc mocks base method.
func (m *MockDocRetriever) Doc(arg0 postings.ID) (doc.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Doc", arg0)
	ret0, _ := ret[0].(doc.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Doc indicates an expected call of Doc.
func (mr *MockDocRetrieverMockRecorder) Doc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Doc", reflect.TypeOf((*MockDocRetriever)(nil).Doc), arg0)
}

// MockMetadataRetriever is a mock of MetadataRetriever interface.
type MockMetadataRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataRetrieverMockRecorder
}

// MockMetadataRetrieverMockRecorder is the mock recorder for MockMetadataRetriever.
type MockMetadataRetrieverMockRecorder struct {
	mock *MockMetadataRetriever
}

// NewMockMetadataRetriever creates a new mock instance.
func NewMockMetadataRetriever(ctrl *gomock.Controller) *MockMetadataRetriever {
	mock := &MockMetadataRetriever{ctrl: ctrl}
	mock.recorder = &MockMetadataRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataRetriever) EXPECT() *MockMetadataRetrieverMockRecorder {
	return m.recorder
}

// Metadata mocks base method.
func (m *MockMetadataRetriever) Metadata(arg0 postings.ID) (doc.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata", arg0)
	ret0, _ := ret[0].(doc.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metadata indicates an expected call of Metadata.
func (mr *MockMetadataRetrieverMockRecorder) Metadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockMetadataRetriever)(nil).Metadata), arg0)
}
