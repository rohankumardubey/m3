// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/metrics/generated/proto/metricpb/metadata.proto

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

package metricpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import aggregationpb "github.com/m3db/m3/src/metrics/generated/proto/aggregationpb"
import policypb "github.com/m3db/m3/src/metrics/generated/proto/policypb"
import pipelinepb "github.com/m3db/m3/src/metrics/generated/proto/pipelinepb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PipelineMetadata struct {
	AggregationId   aggregationpb.AggregationID `protobuf:"bytes,1,opt,name=aggregation_id,json=aggregationId" json:"aggregation_id"`
	StoragePolicies []policypb.StoragePolicy    `protobuf:"bytes,2,rep,name=storage_policies,json=storagePolicies" json:"storage_policies"`
	Pipeline        pipelinepb.AppliedPipeline  `protobuf:"bytes,3,opt,name=pipeline" json:"pipeline"`
	DropPolicy      policypb.DropPolicy         `protobuf:"varint,4,opt,name=drop_policy,json=dropPolicy,proto3,enum=policypb.DropPolicy" json:"drop_policy,omitempty"`
	ResendEnabled   bool                        `protobuf:"varint,5,opt,name=resend_enabled,json=resendEnabled,proto3" json:"resend_enabled,omitempty"`
}

func (m *PipelineMetadata) Reset()                    { *m = PipelineMetadata{} }
func (m *PipelineMetadata) String() string            { return proto.CompactTextString(m) }
func (*PipelineMetadata) ProtoMessage()               {}
func (*PipelineMetadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{0} }

func (m *PipelineMetadata) GetAggregationId() aggregationpb.AggregationID {
	if m != nil {
		return m.AggregationId
	}
	return aggregationpb.AggregationID{}
}

func (m *PipelineMetadata) GetStoragePolicies() []policypb.StoragePolicy {
	if m != nil {
		return m.StoragePolicies
	}
	return nil
}

func (m *PipelineMetadata) GetPipeline() pipelinepb.AppliedPipeline {
	if m != nil {
		return m.Pipeline
	}
	return pipelinepb.AppliedPipeline{}
}

func (m *PipelineMetadata) GetDropPolicy() policypb.DropPolicy {
	if m != nil {
		return m.DropPolicy
	}
	return policypb.DropPolicy_NONE
}

func (m *PipelineMetadata) GetResendEnabled() bool {
	if m != nil {
		return m.ResendEnabled
	}
	return false
}

type Metadata struct {
	Pipelines []PipelineMetadata `protobuf:"bytes,1,rep,name=pipelines" json:"pipelines"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{1} }

func (m *Metadata) GetPipelines() []PipelineMetadata {
	if m != nil {
		return m.Pipelines
	}
	return nil
}

type StagedMetadata struct {
	CutoverNanos int64    `protobuf:"varint,1,opt,name=cutover_nanos,json=cutoverNanos,proto3" json:"cutover_nanos,omitempty"`
	Tombstoned   bool     `protobuf:"varint,2,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	Metadata     Metadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata"`
}

func (m *StagedMetadata) Reset()                    { *m = StagedMetadata{} }
func (m *StagedMetadata) String() string            { return proto.CompactTextString(m) }
func (*StagedMetadata) ProtoMessage()               {}
func (*StagedMetadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{2} }

func (m *StagedMetadata) GetCutoverNanos() int64 {
	if m != nil {
		return m.CutoverNanos
	}
	return 0
}

func (m *StagedMetadata) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *StagedMetadata) GetMetadata() Metadata {
	if m != nil {
		return m.Metadata
	}
	return Metadata{}
}

type StagedMetadatas struct {
	Metadatas []StagedMetadata `protobuf:"bytes,1,rep,name=metadatas" json:"metadatas"`
}

func (m *StagedMetadatas) Reset()                    { *m = StagedMetadatas{} }
func (m *StagedMetadatas) String() string            { return proto.CompactTextString(m) }
func (*StagedMetadatas) ProtoMessage()               {}
func (*StagedMetadatas) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{3} }

func (m *StagedMetadatas) GetMetadatas() []StagedMetadata {
	if m != nil {
		return m.Metadatas
	}
	return nil
}

type ForwardMetadata struct {
	AggregationId     aggregationpb.AggregationID `protobuf:"bytes,1,opt,name=aggregation_id,json=aggregationId" json:"aggregation_id"`
	StoragePolicy     policypb.StoragePolicy      `protobuf:"bytes,2,opt,name=storage_policy,json=storagePolicy" json:"storage_policy"`
	Pipeline          pipelinepb.AppliedPipeline  `protobuf:"bytes,3,opt,name=pipeline" json:"pipeline"`
	SourceId          uint32                      `protobuf:"varint,4,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	NumForwardedTimes int32                       `protobuf:"varint,5,opt,name=num_forwarded_times,json=numForwardedTimes,proto3" json:"num_forwarded_times,omitempty"`
	ResendEnabled     bool                        `protobuf:"varint,6,opt,name=resend_enabled,json=resendEnabled,proto3" json:"resend_enabled,omitempty"`
}

func (m *ForwardMetadata) Reset()                    { *m = ForwardMetadata{} }
func (m *ForwardMetadata) String() string            { return proto.CompactTextString(m) }
func (*ForwardMetadata) ProtoMessage()               {}
func (*ForwardMetadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{4} }

func (m *ForwardMetadata) GetAggregationId() aggregationpb.AggregationID {
	if m != nil {
		return m.AggregationId
	}
	return aggregationpb.AggregationID{}
}

func (m *ForwardMetadata) GetStoragePolicy() policypb.StoragePolicy {
	if m != nil {
		return m.StoragePolicy
	}
	return policypb.StoragePolicy{}
}

func (m *ForwardMetadata) GetPipeline() pipelinepb.AppliedPipeline {
	if m != nil {
		return m.Pipeline
	}
	return pipelinepb.AppliedPipeline{}
}

func (m *ForwardMetadata) GetSourceId() uint32 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *ForwardMetadata) GetNumForwardedTimes() int32 {
	if m != nil {
		return m.NumForwardedTimes
	}
	return 0
}

func (m *ForwardMetadata) GetResendEnabled() bool {
	if m != nil {
		return m.ResendEnabled
	}
	return false
}

type TimedMetadata struct {
	AggregationId aggregationpb.AggregationID `protobuf:"bytes,1,opt,name=aggregation_id,json=aggregationId" json:"aggregation_id"`
	StoragePolicy policypb.StoragePolicy      `protobuf:"bytes,2,opt,name=storage_policy,json=storagePolicy" json:"storage_policy"`
}

func (m *TimedMetadata) Reset()                    { *m = TimedMetadata{} }
func (m *TimedMetadata) String() string            { return proto.CompactTextString(m) }
func (*TimedMetadata) ProtoMessage()               {}
func (*TimedMetadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{5} }

func (m *TimedMetadata) GetAggregationId() aggregationpb.AggregationID {
	if m != nil {
		return m.AggregationId
	}
	return aggregationpb.AggregationID{}
}

func (m *TimedMetadata) GetStoragePolicy() policypb.StoragePolicy {
	if m != nil {
		return m.StoragePolicy
	}
	return policypb.StoragePolicy{}
}

func init() {
	proto.RegisterType((*PipelineMetadata)(nil), "metricpb.PipelineMetadata")
	proto.RegisterType((*Metadata)(nil), "metricpb.Metadata")
	proto.RegisterType((*StagedMetadata)(nil), "metricpb.StagedMetadata")
	proto.RegisterType((*StagedMetadatas)(nil), "metricpb.StagedMetadatas")
	proto.RegisterType((*ForwardMetadata)(nil), "metricpb.ForwardMetadata")
	proto.RegisterType((*TimedMetadata)(nil), "metricpb.TimedMetadata")
}
func (m *PipelineMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PipelineMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.AggregationId.Size()))
	n1, err := m.AggregationId.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.StoragePolicies) > 0 {
		for _, msg := range m.StoragePolicies {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMetadata(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.Pipeline.Size()))
	n2, err := m.Pipeline.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.DropPolicy != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.DropPolicy))
	}
	if m.ResendEnabled {
		dAtA[i] = 0x28
		i++
		if m.ResendEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Pipelines) > 0 {
		for _, msg := range m.Pipelines {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMetadata(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StagedMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StagedMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CutoverNanos != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.CutoverNanos))
	}
	if m.Tombstoned {
		dAtA[i] = 0x10
		i++
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.Metadata.Size()))
	n3, err := m.Metadata.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *StagedMetadatas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StagedMetadatas) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Metadatas) > 0 {
		for _, msg := range m.Metadatas {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMetadata(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ForwardMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.AggregationId.Size()))
	n4, err := m.AggregationId.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x12
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.StoragePolicy.Size()))
	n5, err := m.StoragePolicy.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x1a
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.Pipeline.Size()))
	n6, err := m.Pipeline.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.SourceId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.SourceId))
	}
	if m.NumForwardedTimes != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.NumForwardedTimes))
	}
	if m.ResendEnabled {
		dAtA[i] = 0x30
		i++
		if m.ResendEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *TimedMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimedMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.AggregationId.Size()))
	n7, err := m.AggregationId.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x12
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.StoragePolicy.Size()))
	n8, err := m.StoragePolicy.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PipelineMetadata) Size() (n int) {
	var l int
	_ = l
	l = m.AggregationId.Size()
	n += 1 + l + sovMetadata(uint64(l))
	if len(m.StoragePolicies) > 0 {
		for _, e := range m.StoragePolicies {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	l = m.Pipeline.Size()
	n += 1 + l + sovMetadata(uint64(l))
	if m.DropPolicy != 0 {
		n += 1 + sovMetadata(uint64(m.DropPolicy))
	}
	if m.ResendEnabled {
		n += 2
	}
	return n
}

func (m *Metadata) Size() (n int) {
	var l int
	_ = l
	if len(m.Pipelines) > 0 {
		for _, e := range m.Pipelines {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	return n
}

func (m *StagedMetadata) Size() (n int) {
	var l int
	_ = l
	if m.CutoverNanos != 0 {
		n += 1 + sovMetadata(uint64(m.CutoverNanos))
	}
	if m.Tombstoned {
		n += 2
	}
	l = m.Metadata.Size()
	n += 1 + l + sovMetadata(uint64(l))
	return n
}

func (m *StagedMetadatas) Size() (n int) {
	var l int
	_ = l
	if len(m.Metadatas) > 0 {
		for _, e := range m.Metadatas {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	return n
}

func (m *ForwardMetadata) Size() (n int) {
	var l int
	_ = l
	l = m.AggregationId.Size()
	n += 1 + l + sovMetadata(uint64(l))
	l = m.StoragePolicy.Size()
	n += 1 + l + sovMetadata(uint64(l))
	l = m.Pipeline.Size()
	n += 1 + l + sovMetadata(uint64(l))
	if m.SourceId != 0 {
		n += 1 + sovMetadata(uint64(m.SourceId))
	}
	if m.NumForwardedTimes != 0 {
		n += 1 + sovMetadata(uint64(m.NumForwardedTimes))
	}
	if m.ResendEnabled {
		n += 2
	}
	return n
}

func (m *TimedMetadata) Size() (n int) {
	var l int
	_ = l
	l = m.AggregationId.Size()
	n += 1 + l + sovMetadata(uint64(l))
	l = m.StoragePolicy.Size()
	n += 1 + l + sovMetadata(uint64(l))
	return n
}

func sovMetadata(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PipelineMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: PipelineMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PipelineMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AggregationId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoragePolicies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoragePolicies = append(m.StoragePolicies, policypb.StoragePolicy{})
			if err := m.StoragePolicies[len(m.StoragePolicies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pipeline.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DropPolicy", wireType)
			}
			m.DropPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DropPolicy |= (policypb.DropPolicy(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResendEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ResendEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipelines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pipelines = append(m.Pipelines, PipelineMetadata{})
			if err := m.Pipelines[len(m.Pipelines)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
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
func (m *StagedMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: StagedMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StagedMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CutoverNanos", wireType)
			}
			m.CutoverNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CutoverNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
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
func (m *StagedMetadatas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: StagedMetadatas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StagedMetadatas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadatas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadatas = append(m.Metadatas, StagedMetadata{})
			if err := m.Metadatas[len(m.Metadatas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
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
func (m *ForwardMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: ForwardMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AggregationId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoragePolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StoragePolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pipeline.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceId", wireType)
			}
			m.SourceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SourceId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumForwardedTimes", wireType)
			}
			m.NumForwardedTimes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumForwardedTimes |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResendEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ResendEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
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
func (m *TimedMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: TimedMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimedMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AggregationId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoragePolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StoragePolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetadata
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
				next, err := skipMetadata(dAtA[start:])
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
	ErrInvalidLengthMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/metrics/generated/proto/metricpb/metadata.proto", fileDescriptorMetadata)
}

var fileDescriptorMetadata = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xde, 0x74, 0x7f, 0x90, 0x7d, 0x35, 0xed, 0x1a, 0x05, 0x43, 0x57, 0x6a, 0xa9, 0x08, 0xbd,
	0x98, 0x40, 0xab, 0x78, 0x51, 0x61, 0x97, 0x5a, 0xb6, 0x82, 0xeb, 0x92, 0xf5, 0xe4, 0x25, 0x24,
	0x99, 0xd9, 0x18, 0x68, 0x32, 0x61, 0x66, 0xa2, 0xf4, 0x6f, 0xf0, 0xe2, 0x9f, 0xe0, 0x1f, 0xe3,
	0x61, 0x8f, 0x5e, 0xbc, 0x8a, 0xd4, 0x7f, 0x44, 0x92, 0x99, 0x49, 0xd2, 0x65, 0x41, 0xaa, 0x08,
	0xde, 0xde, 0xfb, 0xe6, 0xbd, 0xaf, 0xdf, 0xf7, 0xfa, 0xb5, 0x30, 0x8b, 0x62, 0xfe, 0x2e, 0x0f,
	0xec, 0x90, 0x24, 0x4e, 0x32, 0x41, 0x81, 0x93, 0x4c, 0x1c, 0x46, 0x43, 0x27, 0xc1, 0x9c, 0xc6,
	0x21, 0x73, 0x22, 0x9c, 0x62, 0xea, 0x73, 0x8c, 0x9c, 0x8c, 0x12, 0x4e, 0x24, 0x9e, 0x05, 0x45,
	0xe1, 0x23, 0x9f, 0xfb, 0x76, 0x89, 0x9b, 0xba, 0x7a, 0xe8, 0x3d, 0x6c, 0x30, 0x46, 0x24, 0x22,
	0x62, 0x31, 0xc8, 0x2f, 0xca, 0x4e, 0xb0, 0x14, 0x95, 0x58, 0xec, 0x9d, 0x6e, 0x28, 0xc0, 0x8f,
	0x22, 0x8a, 0x23, 0x9f, 0xc7, 0x24, 0xcd, 0x82, 0x66, 0x27, 0xf9, 0xa6, 0x1b, 0xf2, 0x65, 0x64,
	0x11, 0x87, 0xcb, 0x2c, 0x90, 0x85, 0x64, 0x39, 0xd9, 0x94, 0x25, 0xce, 0xf0, 0x22, 0x4e, 0x71,
	0xc1, 0x23, 0x4b, 0xc1, 0x34, 0xfc, 0xd2, 0x82, 0x83, 0x33, 0x09, 0xbd, 0x92, 0x37, 0x33, 0xe7,
	0xd0, 0x69, 0x28, 0xf7, 0x62, 0x64, 0x69, 0x03, 0x6d, 0xd4, 0x1e, 0xdf, 0xb5, 0xd7, 0xec, 0xd9,
	0x47, 0x75, 0x37, 0x9f, 0x1e, 0xef, 0x5c, 0x7e, 0xbf, 0xb7, 0xe5, 0x1a, 0x8d, 0x91, 0x39, 0x32,
	0x4f, 0xe0, 0x80, 0x71, 0x42, 0xfd, 0x08, 0x7b, 0xa5, 0x83, 0x18, 0x33, 0xab, 0x35, 0xd8, 0x1e,
	0xb5, 0xc7, 0x77, 0x6c, 0xe5, 0xcd, 0x3e, 0x17, 0x13, 0x67, 0x65, 0x2f, 0x79, 0xba, 0xac, 0x01,
	0xc6, 0x98, 0x99, 0xcf, 0x40, 0x57, 0xda, 0xad, 0xed, 0x52, 0xce, 0xa1, 0x5d, 0xfb, 0xb2, 0x8f,
	0xb2, 0x6c, 0x11, 0x63, 0xa4, 0xbc, 0x48, 0x96, 0x6a, 0xc5, 0x7c, 0x0c, 0x6d, 0x44, 0x49, 0x26,
	0x54, 0x2c, 0xad, 0x9d, 0x81, 0x36, 0xea, 0x8c, 0x6f, 0xd7, 0x1a, 0xa6, 0x94, 0x64, 0x42, 0x80,
	0x0b, 0xa8, 0xaa, 0xcd, 0x07, 0xd0, 0xa1, 0x98, 0xe1, 0x14, 0x79, 0x38, 0xf5, 0x83, 0x05, 0x46,
	0xd6, 0xee, 0x40, 0x1b, 0xe9, 0xae, 0x21, 0xd0, 0x17, 0x02, 0x1c, 0xbe, 0x04, 0xbd, 0xba, 0xde,
	0x73, 0xd8, 0x57, 0x9f, 0xca, 0x2c, 0xad, 0xf4, 0xda, 0xb3, 0x55, 0xfe, 0xec, 0xab, 0xc7, 0x96,
	0x42, 0xeb, 0x95, 0xe1, 0x47, 0x0d, 0x3a, 0xe7, 0xdc, 0x8f, 0x30, 0xaa, 0x28, 0xef, 0x83, 0x11,
	0xe6, 0x9c, 0xbc, 0xc7, 0xd4, 0x4b, 0xfd, 0x94, 0xb0, 0xf2, 0xfb, 0xd8, 0x76, 0x6f, 0x48, 0xf0,
	0xb4, 0xc0, 0xcc, 0x3e, 0x00, 0x27, 0x49, 0xc0, 0x38, 0x49, 0x31, 0xb2, 0x5a, 0xa5, 0xcc, 0x06,
	0x62, 0x3e, 0x02, 0x5d, 0xfd, 0x2a, 0xe4, 0x01, 0xcd, 0x5a, 0xd6, 0x15, 0x39, 0xd5, 0xe4, 0xf0,
	0x35, 0x74, 0xd7, 0xc5, 0x30, 0xf3, 0x29, 0xec, 0xab, 0x67, 0x65, 0xd0, 0xaa, 0x99, 0xd6, 0xa7,
	0x95, 0xbd, 0x6a, 0x61, 0xf8, 0xad, 0x05, 0xdd, 0x19, 0xa1, 0x1f, 0x7c, 0x8a, 0xfe, 0x45, 0xe0,
	0xa6, 0xd0, 0x59, 0x0b, 0xdc, 0xb2, 0xbc, 0xc4, 0x6f, 0xe3, 0x66, 0x34, 0xe3, 0xb6, 0xfc, 0xdb,
	0xb0, 0x1d, 0xc2, 0x3e, 0x23, 0x39, 0x0d, 0x71, 0x61, 0xa5, 0x88, 0x9a, 0xe1, 0xea, 0x02, 0x98,
	0x23, 0xd3, 0x86, 0x5b, 0x69, 0x9e, 0x78, 0x17, 0xe2, 0x06, 0x18, 0x79, 0x3c, 0x4e, 0x30, 0x2b,
	0x73, 0xb5, 0xeb, 0xde, 0x4c, 0xf3, 0x64, 0xa6, 0x5e, 0xde, 0x14, 0x0f, 0xd7, 0x44, 0x70, 0xef,
	0xba, 0x08, 0x7e, 0xd6, 0xc0, 0x28, 0x16, 0xfe, 0xdf, 0xab, 0x1e, 0xcf, 0x2f, 0x57, 0x7d, 0xed,
	0xeb, 0xaa, 0xaf, 0xfd, 0x58, 0xf5, 0xb5, 0x4f, 0x3f, 0xfb, 0x5b, 0x6f, 0x9f, 0xfc, 0xe1, 0xff,
	0x7b, 0xb0, 0x57, 0xf6, 0x93, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0x47, 0xc8, 0xa8, 0x21,
	0x06, 0x00, 0x00,
}
