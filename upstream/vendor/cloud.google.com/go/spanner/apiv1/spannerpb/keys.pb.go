// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: google/spanner/v1/keys.proto

package spannerpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// KeyRange represents a range of rows in a table or index.
//
// A range has a start key and an end key. These keys can be open or
// closed, indicating if the range includes rows with that key.
//
// Keys are represented by lists, where the ith value in the list
// corresponds to the ith component of the table or index primary key.
// Individual values are encoded as described
// [here][google.spanner.v1.TypeCode].
//
// For example, consider the following table definition:
//
//	CREATE TABLE UserEvents (
//	  UserName STRING(MAX),
//	  EventDate STRING(10)
//	) PRIMARY KEY(UserName, EventDate);
//
// The following keys name rows in this table:
//
//	["Bob", "2014-09-23"]
//	["Alfred", "2015-06-12"]
//
// Since the `UserEvents` table's `PRIMARY KEY` clause names two
// columns, each `UserEvents` key has two elements; the first is the
// `UserName`, and the second is the `EventDate`.
//
// Key ranges with multiple components are interpreted
// lexicographically by component using the table or index key's declared
// sort order. For example, the following range returns all events for
// user `"Bob"` that occurred in the year 2015:
//
//	"start_closed": ["Bob", "2015-01-01"]
//	"end_closed": ["Bob", "2015-12-31"]
//
// Start and end keys can omit trailing key components. This affects the
// inclusion and exclusion of rows that exactly match the provided key
// components: if the key is closed, then rows that exactly match the
// provided components are included; if the key is open, then rows
// that exactly match are not included.
//
// For example, the following range includes all events for `"Bob"` that
// occurred during and after the year 2000:
//
//	"start_closed": ["Bob", "2000-01-01"]
//	"end_closed": ["Bob"]
//
// The next example retrieves all events for `"Bob"`:
//
//	"start_closed": ["Bob"]
//	"end_closed": ["Bob"]
//
// To retrieve events before the year 2000:
//
//	"start_closed": ["Bob"]
//	"end_open": ["Bob", "2000-01-01"]
//
// The following range includes all rows in the table:
//
//	"start_closed": []
//	"end_closed": []
//
// This range returns all users whose `UserName` begins with any
// character from A to C:
//
//	"start_closed": ["A"]
//	"end_open": ["D"]
//
// This range returns all users whose `UserName` begins with B:
//
//	"start_closed": ["B"]
//	"end_open": ["C"]
//
// Key ranges honor column sort order. For example, suppose a table is
// defined as follows:
//
//	CREATE TABLE DescendingSortedTable {
//	  Key INT64,
//	  ...
//	) PRIMARY KEY(Key DESC);
//
// The following range retrieves all rows with key values between 1
// and 100 inclusive:
//
//	"start_closed": ["100"]
//	"end_closed": ["1"]
//
// Note that 100 is passed as the start, and 1 is passed as the end,
// because `Key` is a descending column in the schema.
type KeyRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The start key must be provided. It can be either closed or open.
	//
	// Types that are assignable to StartKeyType:
	//
	//	*KeyRange_StartClosed
	//	*KeyRange_StartOpen
	StartKeyType isKeyRange_StartKeyType `protobuf_oneof:"start_key_type"`
	// The end key must be provided. It can be either closed or open.
	//
	// Types that are assignable to EndKeyType:
	//
	//	*KeyRange_EndClosed
	//	*KeyRange_EndOpen
	EndKeyType isKeyRange_EndKeyType `protobuf_oneof:"end_key_type"`
}

func (x *KeyRange) Reset() {
	*x = KeyRange{}
	mi := &file_google_spanner_v1_keys_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRange) ProtoMessage() {}

func (x *KeyRange) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_keys_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRange.ProtoReflect.Descriptor instead.
func (*KeyRange) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_keys_proto_rawDescGZIP(), []int{0}
}

func (m *KeyRange) GetStartKeyType() isKeyRange_StartKeyType {
	if m != nil {
		return m.StartKeyType
	}
	return nil
}

func (x *KeyRange) GetStartClosed() *structpb.ListValue {
	if x, ok := x.GetStartKeyType().(*KeyRange_StartClosed); ok {
		return x.StartClosed
	}
	return nil
}

func (x *KeyRange) GetStartOpen() *structpb.ListValue {
	if x, ok := x.GetStartKeyType().(*KeyRange_StartOpen); ok {
		return x.StartOpen
	}
	return nil
}

func (m *KeyRange) GetEndKeyType() isKeyRange_EndKeyType {
	if m != nil {
		return m.EndKeyType
	}
	return nil
}

func (x *KeyRange) GetEndClosed() *structpb.ListValue {
	if x, ok := x.GetEndKeyType().(*KeyRange_EndClosed); ok {
		return x.EndClosed
	}
	return nil
}

func (x *KeyRange) GetEndOpen() *structpb.ListValue {
	if x, ok := x.GetEndKeyType().(*KeyRange_EndOpen); ok {
		return x.EndOpen
	}
	return nil
}

type isKeyRange_StartKeyType interface {
	isKeyRange_StartKeyType()
}

type KeyRange_StartClosed struct {
	// If the start is closed, then the range includes all rows whose
	// first `len(start_closed)` key columns exactly match `start_closed`.
	StartClosed *structpb.ListValue `protobuf:"bytes,1,opt,name=start_closed,json=startClosed,proto3,oneof"`
}

type KeyRange_StartOpen struct {
	// If the start is open, then the range excludes rows whose first
	// `len(start_open)` key columns exactly match `start_open`.
	StartOpen *structpb.ListValue `protobuf:"bytes,2,opt,name=start_open,json=startOpen,proto3,oneof"`
}

func (*KeyRange_StartClosed) isKeyRange_StartKeyType() {}

func (*KeyRange_StartOpen) isKeyRange_StartKeyType() {}

type isKeyRange_EndKeyType interface {
	isKeyRange_EndKeyType()
}

type KeyRange_EndClosed struct {
	// If the end is closed, then the range includes all rows whose
	// first `len(end_closed)` key columns exactly match `end_closed`.
	EndClosed *structpb.ListValue `protobuf:"bytes,3,opt,name=end_closed,json=endClosed,proto3,oneof"`
}

type KeyRange_EndOpen struct {
	// If the end is open, then the range excludes rows whose first
	// `len(end_open)` key columns exactly match `end_open`.
	EndOpen *structpb.ListValue `protobuf:"bytes,4,opt,name=end_open,json=endOpen,proto3,oneof"`
}

func (*KeyRange_EndClosed) isKeyRange_EndKeyType() {}

func (*KeyRange_EndOpen) isKeyRange_EndKeyType() {}

// `KeySet` defines a collection of Cloud Spanner keys and/or key ranges. All
// the keys are expected to be in the same table or index. The keys need
// not be sorted in any particular way.
//
// If the same key is specified multiple times in the set (for example
// if two ranges, two keys, or a key and a range overlap), Cloud Spanner
// behaves as if the key were only specified once.
type KeySet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of specific keys. Entries in `keys` should have exactly as
	// many elements as there are columns in the primary or index key
	// with which this `KeySet` is used.  Individual key values are
	// encoded as described [here][google.spanner.v1.TypeCode].
	Keys []*structpb.ListValue `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	// A list of key ranges. See [KeyRange][google.spanner.v1.KeyRange] for more information about
	// key range specifications.
	Ranges []*KeyRange `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
	// For convenience `all` can be set to `true` to indicate that this
	// `KeySet` matches all keys in the table or index. Note that any keys
	// specified in `keys` or `ranges` are only yielded once.
	All bool `protobuf:"varint,3,opt,name=all,proto3" json:"all,omitempty"`
}

func (x *KeySet) Reset() {
	*x = KeySet{}
	mi := &file_google_spanner_v1_keys_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeySet) ProtoMessage() {}

func (x *KeySet) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_keys_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeySet.ProtoReflect.Descriptor instead.
func (*KeySet) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_keys_proto_rawDescGZIP(), []int{1}
}

func (x *KeySet) GetKeys() []*structpb.ListValue {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *KeySet) GetRanges() []*KeyRange {
	if x != nil {
		return x.Ranges
	}
	return nil
}

func (x *KeySet) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

var File_google_spanner_v1_keys_proto protoreflect.FileDescriptor

var file_google_spanner_v1_keys_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa0, 0x02, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x3b, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x0a, 0x65, 0x6e,
	0x64, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x01, 0x52, 0x09, 0x65, 0x6e,
	0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x6e,
	0x42, 0x10, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x7f, 0x0a, 0x06, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x04,
	0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x33, 0x0a, 0x06,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03,
	0x61, 0x6c, 0x6c, 0x42, 0xac, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4b,
	0x65, 0x79, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x70,
	0x62, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x70, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_spanner_v1_keys_proto_rawDescOnce sync.Once
	file_google_spanner_v1_keys_proto_rawDescData = file_google_spanner_v1_keys_proto_rawDesc
)

func file_google_spanner_v1_keys_proto_rawDescGZIP() []byte {
	file_google_spanner_v1_keys_proto_rawDescOnce.Do(func() {
		file_google_spanner_v1_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_spanner_v1_keys_proto_rawDescData)
	})
	return file_google_spanner_v1_keys_proto_rawDescData
}

var file_google_spanner_v1_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_spanner_v1_keys_proto_goTypes = []any{
	(*KeyRange)(nil),           // 0: google.spanner.v1.KeyRange
	(*KeySet)(nil),             // 1: google.spanner.v1.KeySet
	(*structpb.ListValue)(nil), // 2: google.protobuf.ListValue
}
var file_google_spanner_v1_keys_proto_depIdxs = []int32{
	2, // 0: google.spanner.v1.KeyRange.start_closed:type_name -> google.protobuf.ListValue
	2, // 1: google.spanner.v1.KeyRange.start_open:type_name -> google.protobuf.ListValue
	2, // 2: google.spanner.v1.KeyRange.end_closed:type_name -> google.protobuf.ListValue
	2, // 3: google.spanner.v1.KeyRange.end_open:type_name -> google.protobuf.ListValue
	2, // 4: google.spanner.v1.KeySet.keys:type_name -> google.protobuf.ListValue
	0, // 5: google.spanner.v1.KeySet.ranges:type_name -> google.spanner.v1.KeyRange
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_spanner_v1_keys_proto_init() }
func file_google_spanner_v1_keys_proto_init() {
	if File_google_spanner_v1_keys_proto != nil {
		return
	}
	file_google_spanner_v1_keys_proto_msgTypes[0].OneofWrappers = []any{
		(*KeyRange_StartClosed)(nil),
		(*KeyRange_StartOpen)(nil),
		(*KeyRange_EndClosed)(nil),
		(*KeyRange_EndOpen)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_spanner_v1_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_spanner_v1_keys_proto_goTypes,
		DependencyIndexes: file_google_spanner_v1_keys_proto_depIdxs,
		MessageInfos:      file_google_spanner_v1_keys_proto_msgTypes,
	}.Build()
	File_google_spanner_v1_keys_proto = out.File
	file_google_spanner_v1_keys_proto_rawDesc = nil
	file_google_spanner_v1_keys_proto_goTypes = nil
	file_google_spanner_v1_keys_proto_depIdxs = nil
}
