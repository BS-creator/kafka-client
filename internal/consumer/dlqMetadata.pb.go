// Copyright (c) 2017 Uber Technologies, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dlqMetadata.proto

/*
Package consumer is a generated protocol buffer package.

It is generated from these files:
	dlqMetadata.proto

It has these top-level messages:
	DLQMetadata
*/
package consumer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// DLQMetadata contains metadata from the original kafka message.
// The metadata will be encoded and decoded when sending or receiving
// messages from the DLQ cluster in order to present the library
// user a seamless logical topic.
type DLQMetadata struct {
	// retry_count is an incrementing value denoting the number
	// of times a message has been redelivered.
	// It will be 0 on first delivery.
	RetryCount int64 `protobuf:"varint,1,opt,name=retry_count,json=retryCount" json:"retry_count,omitempty"`
	// topic is the original kafka topic the mesasge was received on.
	// This is analogous to the logical topic name.
	Topic string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	// partition is the original kafka partition the message was received on.
	Partition int32 `protobuf:"varint,3,opt,name=partition" json:"partition,omitempty"`
	// offset is the record offset of the original message in the original topic-partition.
	Offset int64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	// timestamp_ns is the original record timestamp of the original mesage.
	TimestampNs int64 `protobuf:"varint,5,opt,name=timestamp_ns,json=timestampNs" json:"timestamp_ns,omitempty"`
	// data is a byte buffer for storing arbitrary information.
	// This is useful if the Kafka Broker version used is < 0.11
	// and hence Kafka native record headers (KAFKA-4208) are unavaiable
	// so the DLQ metadata must be stored in the record Key or Value.
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DLQMetadata) Reset()                    { *m = DLQMetadata{} }
func (m *DLQMetadata) String() string            { return proto.CompactTextString(m) }
func (*DLQMetadata) ProtoMessage()               {}
func (*DLQMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DLQMetadata) GetRetryCount() int64 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *DLQMetadata) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *DLQMetadata) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *DLQMetadata) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DLQMetadata) GetTimestampNs() int64 {
	if m != nil {
		return m.TimestampNs
	}
	return 0
}

func (m *DLQMetadata) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DLQMetadata)(nil), "DLQMetadata")
}

func init() { proto.RegisterFile("dlqMetadata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xce, 0xbd, 0x0a, 0xc2, 0x30,
	0x14, 0x05, 0x60, 0x62, 0x7f, 0xa0, 0xb7, 0x5d, 0x0c, 0x22, 0x19, 0x04, 0xa3, 0x53, 0x26, 0x17,
	0x1f, 0x41, 0x47, 0x15, 0xcc, 0x0b, 0x94, 0xd8, 0xa6, 0x10, 0xb0, 0x4d, 0x4c, 0xae, 0x83, 0x8f,
	0xe5, 0x1b, 0x4a, 0xe3, 0xdf, 0x76, 0xcf, 0x77, 0xe1, 0x70, 0x60, 0xda, 0x5e, 0x6f, 0x47, 0x8d,
	0xaa, 0x55, 0xa8, 0x36, 0xce, 0x5b, 0xb4, 0xeb, 0x27, 0x81, 0x72, 0x7f, 0x38, 0x7f, 0x95, 0x2e,
	0xa1, 0xf4, 0x1a, 0xfd, 0xa3, 0x6e, 0xec, 0x7d, 0x40, 0x46, 0x38, 0x11, 0x89, 0x84, 0x48, 0xbb,
	0x51, 0xe8, 0x0c, 0x32, 0xb4, 0xce, 0x34, 0x6c, 0xc2, 0x89, 0x28, 0xe4, 0x3b, 0xd0, 0x05, 0x14,
	0x4e, 0x79, 0x34, 0x68, 0xec, 0xc0, 0x12, 0x4e, 0x44, 0x26, 0xff, 0x40, 0xe7, 0x90, 0xdb, 0xae,
	0x0b, 0x1a, 0x59, 0x1a, 0xfb, 0x3e, 0x89, 0xae, 0xa0, 0x42, 0xd3, 0xeb, 0x80, 0xaa, 0x77, 0xf5,
	0x10, 0x58, 0x16, 0xbf, 0xe5, 0xcf, 0x4e, 0x81, 0x52, 0x48, 0xc7, 0x5d, 0x2c, 0xe7, 0x44, 0x54,
	0x32, 0xde, 0x97, 0x3c, 0x4e, 0xdf, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x23, 0x67, 0x20, 0xf9,
	0xcf, 0x00, 0x00, 0x00,
}