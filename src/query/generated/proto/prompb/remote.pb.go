// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/query/generated/proto/prompb/remote.proto

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

/*
	Package prompb is a generated protocol buffer package.

	It is generated from these files:
		github.com/m3db/m3/src/query/generated/proto/prompb/remote.proto
		github.com/m3db/m3/src/query/generated/proto/prompb/types.proto

	It has these top-level messages:
		WriteRequest
		ReadRequest
		ReadResponse
		Query
		QueryResult
		Sample
		TimeSeries
		Label
		Labels
		LabelMatcher
*/
package prompb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type WriteRequest struct {
	Timeseries []TimeSeries `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptorRemote, []int{0} }

func (m *WriteRequest) GetTimeseries() []TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

type ReadRequest struct {
	Queries []*Query `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptorRemote, []int{1} }

func (m *ReadRequest) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type ReadResponse struct {
	// In same order as the request's queries.
	Results []*QueryResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptorRemote, []int{2} }

func (m *ReadResponse) GetResults() []*QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type Query struct {
	StartTimestampMs int64           `protobuf:"varint,1,opt,name=start_timestamp_ms,json=startTimestampMs,proto3" json:"start_timestamp_ms,omitempty"`
	EndTimestampMs   int64           `protobuf:"varint,2,opt,name=end_timestamp_ms,json=endTimestampMs,proto3" json:"end_timestamp_ms,omitempty"`
	Matchers         []*LabelMatcher `protobuf:"bytes,3,rep,name=matchers" json:"matchers,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptorRemote, []int{3} }

func (m *Query) GetStartTimestampMs() int64 {
	if m != nil {
		return m.StartTimestampMs
	}
	return 0
}

func (m *Query) GetEndTimestampMs() int64 {
	if m != nil {
		return m.EndTimestampMs
	}
	return 0
}

func (m *Query) GetMatchers() []*LabelMatcher {
	if m != nil {
		return m.Matchers
	}
	return nil
}

type QueryResult struct {
	Timeseries []*TimeSeries `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptorRemote, []int{4} }

func (m *QueryResult) GetTimeseries() []*TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteRequest)(nil), "m3prometheus.WriteRequest")
	proto.RegisterType((*ReadRequest)(nil), "m3prometheus.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "m3prometheus.ReadResponse")
	proto.RegisterType((*Query)(nil), "m3prometheus.Query")
	proto.RegisterType((*QueryResult)(nil), "m3prometheus.QueryResult")
}
func (m *WriteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WriteRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Timeseries) > 0 {
		for _, msg := range m.Timeseries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRemote(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ReadRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, msg := range m.Queries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRemote(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ReadResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRemote(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Query) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Query) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StartTimestampMs != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRemote(dAtA, i, uint64(m.StartTimestampMs))
	}
	if m.EndTimestampMs != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRemote(dAtA, i, uint64(m.EndTimestampMs))
	}
	if len(m.Matchers) > 0 {
		for _, msg := range m.Matchers {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRemote(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *QueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Timeseries) > 0 {
		for _, msg := range m.Timeseries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRemote(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintRemote(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *WriteRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Timeseries) > 0 {
		for _, e := range m.Timeseries {
			l = e.Size()
			n += 1 + l + sovRemote(uint64(l))
		}
	}
	return n
}

func (m *ReadRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Queries) > 0 {
		for _, e := range m.Queries {
			l = e.Size()
			n += 1 + l + sovRemote(uint64(l))
		}
	}
	return n
}

func (m *ReadResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovRemote(uint64(l))
		}
	}
	return n
}

func (m *Query) Size() (n int) {
	var l int
	_ = l
	if m.StartTimestampMs != 0 {
		n += 1 + sovRemote(uint64(m.StartTimestampMs))
	}
	if m.EndTimestampMs != 0 {
		n += 1 + sovRemote(uint64(m.EndTimestampMs))
	}
	if len(m.Matchers) > 0 {
		for _, e := range m.Matchers {
			l = e.Size()
			n += 1 + l + sovRemote(uint64(l))
		}
	}
	return n
}

func (m *QueryResult) Size() (n int) {
	var l int
	_ = l
	if len(m.Timeseries) > 0 {
		for _, e := range m.Timeseries {
			l = e.Size()
			n += 1 + l + sovRemote(uint64(l))
		}
	}
	return n
}

func sovRemote(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRemote(x uint64) (n int) {
	return sovRemote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WriteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemote
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WriteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WriteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeseries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRemote
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timeseries = append(m.Timeseries, TimeSeries{})
			if err := m.Timeseries[len(m.Timeseries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRemote
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemote
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Queries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRemote
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Queries = append(m.Queries, &Query{})
			if err := m.Queries[len(m.Queries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRemote
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReadResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemote
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRemote
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &QueryResult{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRemote
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Query) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemote
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Query: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Query: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTimestampMs", wireType)
			}
			m.StartTimestampMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTimestampMs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTimestampMs", wireType)
			}
			m.EndTimestampMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTimestampMs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Matchers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRemote
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Matchers = append(m.Matchers, &LabelMatcher{})
			if err := m.Matchers[len(m.Matchers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRemote
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRemote
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeseries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRemote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRemote
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timeseries = append(m.Timeseries, &TimeSeries{})
			if err := m.Timeseries[len(m.Timeseries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRemote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRemote
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRemote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRemote
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRemote
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRemote
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRemote
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRemote
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRemote(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRemote = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRemote   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/query/generated/proto/prompb/remote.proto", fileDescriptorRemote)
}

var fileDescriptorRemote = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x9b, 0xdb, 0x7b, 0xdb, 0xcb, 0xb4, 0x48, 0x89, 0x9b, 0xd8, 0x45, 0x94, 0xac, 0xba,
	0xb0, 0x19, 0x30, 0x20, 0x2e, 0x44, 0xa5, 0x2e, 0xdc, 0x58, 0xc1, 0xb1, 0x20, 0xb8, 0x29, 0x49,
	0x73, 0x4c, 0x03, 0x9d, 0x4c, 0x3a, 0x73, 0xb2, 0xe8, 0x5b, 0xb8, 0xf3, 0x95, 0xba, 0xf4, 0x09,
	0x44, 0xea, 0x8b, 0x48, 0x66, 0x88, 0xa6, 0xe0, 0x46, 0x37, 0x21, 0x39, 0xff, 0xf7, 0xff, 0xfc,
	0x39, 0x33, 0xe4, 0x22, 0x49, 0x71, 0x5e, 0x44, 0xfe, 0x4c, 0x70, 0xca, 0x83, 0x38, 0xa2, 0x3c,
	0xa0, 0x4a, 0xce, 0xe8, 0xb2, 0x00, 0xb9, 0xa2, 0x09, 0x64, 0x20, 0x43, 0x84, 0x98, 0xe6, 0x52,
	0xa0, 0x28, 0x9f, 0x3c, 0x8f, 0xa8, 0x04, 0x2e, 0x10, 0x7c, 0x3d, 0xb3, 0xbb, 0x3c, 0x28, 0xc7,
	0x80, 0x73, 0x28, 0x54, 0xff, 0xfc, 0x37, 0x79, 0xb8, 0xca, 0x41, 0x99, 0xb8, 0xfe, 0xb0, 0x16,
	0x90, 0x88, 0x44, 0x18, 0x32, 0x2a, 0x1e, 0xf5, 0x97, 0xb1, 0x95, 0x6f, 0x06, 0xf7, 0x6e, 0x48,
	0xf7, 0x5e, 0xa6, 0x08, 0x0c, 0x96, 0x05, 0x28, 0xb4, 0xcf, 0x08, 0xc1, 0x94, 0x83, 0x02, 0x99,
	0x82, 0x72, 0xac, 0x83, 0xe6, 0xa0, 0x73, 0xe4, 0xf8, 0xf5, 0x8a, 0xfe, 0x24, 0xe5, 0x70, 0xa7,
	0xf5, 0xd1, 0xdf, 0xf5, 0xeb, 0x7e, 0x83, 0xd5, 0x1c, 0xde, 0x29, 0xe9, 0x30, 0x08, 0xe3, 0x2a,
	0x6e, 0x48, 0xda, 0x65, 0xf3, 0xaf, 0xac, 0xdd, 0xed, 0xac, 0xdb, 0xf2, 0xb7, 0x58, 0xc5, 0x78,
	0x97, 0xa4, 0x6b, 0xdc, 0x2a, 0x17, 0x99, 0x02, 0x3b, 0x20, 0x6d, 0x09, 0xaa, 0x58, 0x60, 0x65,
	0xdf, 0xfb, 0xce, 0xae, 0x09, 0x56, 0x91, 0xde, 0xb3, 0x45, 0xfe, 0x69, 0xc1, 0x3e, 0x24, 0xb6,
	0xc2, 0x50, 0xe2, 0x54, 0x17, 0xc4, 0x90, 0xe7, 0x53, 0x5e, 0x26, 0x59, 0x83, 0x26, 0xeb, 0x69,
	0x65, 0x52, 0x09, 0x63, 0x65, 0x0f, 0x48, 0x0f, 0xb2, 0x78, 0x9b, 0xfd, 0xa3, 0xd9, 0x1d, 0xc8,
	0xe2, 0x3a, 0x79, 0x4c, 0xfe, 0xf3, 0x10, 0x67, 0x73, 0x90, 0xca, 0x69, 0xea, 0x5e, 0xfd, 0xed,
	0x5e, 0xd7, 0x61, 0x04, 0x8b, 0xb1, 0x41, 0xd8, 0x27, 0xeb, 0x5d, 0x91, 0x4e, 0xad, 0xb1, 0x7d,
	0xf2, 0x93, 0x5d, 0xd7, 0xb7, 0x3c, 0x72, 0xd6, 0x1b, 0xd7, 0x7a, 0xd9, 0xb8, 0xd6, 0xdb, 0xc6,
	0xb5, 0x9e, 0xde, 0xdd, 0xc6, 0x43, 0xcb, 0x5c, 0x85, 0xa8, 0xa5, 0x8f, 0x35, 0xf8, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x24, 0xf8, 0x63, 0xea, 0x98, 0x02, 0x00, 0x00,
}
