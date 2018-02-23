// Copyright 2017 Google Inc.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//      http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hmac.proto

/*
Package hmac_proto is a generated protocol buffer package.

It is generated from these files:
	proto/hmac.proto

It has these top-level messages:
	HmacParams
	HmacKey
	HmacKeyFormat
*/
package hmac_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_crypto_tink "github.com/google/tink/proto/common_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HmacParams struct {
	Hash    google_crypto_tink.HashType `protobuf:"varint,1,opt,name=hash,enum=google.crypto.tink.HashType" json:"hash,omitempty"`
	TagSize uint32                      `protobuf:"varint,2,opt,name=tag_size,json=tagSize" json:"tag_size,omitempty"`
}

func (m *HmacParams) Reset()                    { *m = HmacParams{} }
func (m *HmacParams) String() string            { return proto.CompactTextString(m) }
func (*HmacParams) ProtoMessage()               {}
func (*HmacParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HmacParams) GetHash() google_crypto_tink.HashType {
	if m != nil {
		return m.Hash
	}
	return google_crypto_tink.HashType_UNKNOWN_HASH
}

func (m *HmacParams) GetTagSize() uint32 {
	if m != nil {
		return m.TagSize
	}
	return 0
}

// key_type: type.googleapis.com/google.crypto.tink.HmacKey
type HmacKey struct {
	Version  uint32      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Params   *HmacParams `protobuf:"bytes,2,opt,name=params" json:"params,omitempty"`
	KeyValue []byte      `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
}

func (m *HmacKey) Reset()                    { *m = HmacKey{} }
func (m *HmacKey) String() string            { return proto.CompactTextString(m) }
func (*HmacKey) ProtoMessage()               {}
func (*HmacKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HmacKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HmacKey) GetParams() *HmacParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *HmacKey) GetKeyValue() []byte {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type HmacKeyFormat struct {
	Params  *HmacParams `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	KeySize uint32      `protobuf:"varint,2,opt,name=key_size,json=keySize" json:"key_size,omitempty"`
}

func (m *HmacKeyFormat) Reset()                    { *m = HmacKeyFormat{} }
func (m *HmacKeyFormat) String() string            { return proto.CompactTextString(m) }
func (*HmacKeyFormat) ProtoMessage()               {}
func (*HmacKeyFormat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HmacKeyFormat) GetParams() *HmacParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *HmacKeyFormat) GetKeySize() uint32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

func init() {
	proto.RegisterType((*HmacParams)(nil), "google.crypto.tink.HmacParams")
	proto.RegisterType((*HmacKey)(nil), "google.crypto.tink.HmacKey")
	proto.RegisterType((*HmacKeyFormat)(nil), "google.crypto.tink.HmacKeyFormat")
}

func init() { proto.RegisterFile("proto/hmac.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0xc9, 0x7e, 0x3f, 0xda, 0xf9, 0x68, 0x45, 0x72, 0xea, 0x74, 0x48, 0xe9, 0xc5, 0x9e,
	0x52, 0x99, 0xe0, 0x0b, 0xd8, 0x41, 0x26, 0x03, 0x29, 0x71, 0x08, 0x7a, 0x29, 0x69, 0x08, 0x6d,
	0xe9, 0xd2, 0x94, 0x36, 0x1b, 0x66, 0x2f, 0xc7, 0x57, 0x2a, 0x4d, 0xe7, 0x3f, 0xdc, 0xc5, 0xdb,
	0xf3, 0xa5, 0xdf, 0x7e, 0x3e, 0x79, 0x12, 0x38, 0x6b, 0x5a, 0xa5, 0x55, 0x5c, 0x48, 0xc6, 0x89,
	0x1d, 0x31, 0xce, 0x95, 0xca, 0xd7, 0x82, 0xf0, 0xd6, 0x34, 0x5a, 0x11, 0x5d, 0xd6, 0xd5, 0x39,
	0x1e, 0x5a, 0x5c, 0x49, 0xa9, 0xea, 0xa1, 0x17, 0x3e, 0x03, 0x2c, 0x24, 0xe3, 0x09, 0x6b, 0x99,
	0xec, 0xf0, 0x35, 0xfc, 0x2f, 0x58, 0x57, 0xf8, 0x28, 0x40, 0xd1, 0xe9, 0x6c, 0x4a, 0x7e, 0x43,
	0xc8, 0x82, 0x75, 0xc5, 0xca, 0x34, 0x82, 0xda, 0x26, 0x9e, 0xc0, 0x58, 0xb3, 0x3c, 0xed, 0xca,
	0x9d, 0xf0, 0x47, 0x01, 0x8a, 0x3c, 0xea, 0x6a, 0x96, 0x3f, 0x96, 0x3b, 0x11, 0xbe, 0x82, 0xdb,
	0xa3, 0x97, 0xc2, 0x60, 0x1f, 0xdc, 0xad, 0x68, 0xbb, 0x52, 0xd5, 0x16, 0xed, 0xd1, 0x8f, 0x88,
	0x6f, 0xc1, 0x69, 0xac, 0xdb, 0xfe, 0x7d, 0x3c, 0xbb, 0x3c, 0xe8, 0xfc, 0x3c, 0x21, 0xdd, 0xb7,
	0xf1, 0x05, 0x1c, 0x55, 0xc2, 0xa4, 0x5b, 0xb6, 0xde, 0x08, 0xff, 0x5f, 0x80, 0xa2, 0x13, 0x3a,
	0xae, 0x84, 0x79, 0xea, 0x73, 0x98, 0x81, 0xb7, 0x37, 0xdf, 0xa9, 0x56, 0x32, 0xfd, 0xcd, 0x82,
	0xfe, 0x64, 0x99, 0x40, 0x0f, 0xfd, 0xb1, 0x5d, 0x25, 0x4c, 0xbf, 0xdd, 0x9c, 0xc2, 0x94, 0x2b,
	0x79, 0x88, 0x63, 0x2f, 0x36, 0x41, 0x2f, 0x57, 0x79, 0xa9, 0x8b, 0x4d, 0x46, 0xb8, 0x92, 0xf1,
	0x50, 0x8b, 0xfb, 0xef, 0xf1, 0xd7, 0x5b, 0xa5, 0x76, 0x7c, 0x1b, 0x39, 0xab, 0xfb, 0x87, 0x65,
	0x32, 0xcf, 0x1c, 0x9b, 0x6f, 0xde, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x7f, 0xac, 0x9f, 0xcf,
	0x01, 0x00, 0x00,
}
