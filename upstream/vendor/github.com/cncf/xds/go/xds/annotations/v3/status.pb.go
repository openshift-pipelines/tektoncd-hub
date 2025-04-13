// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.1
// source: xds/annotations/v3/status.proto

package v3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PackageVersionStatus int32

const (
	PackageVersionStatus_UNKNOWN                      PackageVersionStatus = 0
	PackageVersionStatus_FROZEN                       PackageVersionStatus = 1
	PackageVersionStatus_ACTIVE                       PackageVersionStatus = 2
	PackageVersionStatus_NEXT_MAJOR_VERSION_CANDIDATE PackageVersionStatus = 3
)

// Enum value maps for PackageVersionStatus.
var (
	PackageVersionStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "FROZEN",
		2: "ACTIVE",
		3: "NEXT_MAJOR_VERSION_CANDIDATE",
	}
	PackageVersionStatus_value = map[string]int32{
		"UNKNOWN":                      0,
		"FROZEN":                       1,
		"ACTIVE":                       2,
		"NEXT_MAJOR_VERSION_CANDIDATE": 3,
	}
)

func (x PackageVersionStatus) Enum() *PackageVersionStatus {
	p := new(PackageVersionStatus)
	*p = x
	return p
}

func (x PackageVersionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackageVersionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_xds_annotations_v3_status_proto_enumTypes[0].Descriptor()
}

func (PackageVersionStatus) Type() protoreflect.EnumType {
	return &file_xds_annotations_v3_status_proto_enumTypes[0]
}

func (x PackageVersionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackageVersionStatus.Descriptor instead.
func (PackageVersionStatus) EnumDescriptor() ([]byte, []int) {
	return file_xds_annotations_v3_status_proto_rawDescGZIP(), []int{0}
}

type FileStatusAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkInProgress bool `protobuf:"varint,1,opt,name=work_in_progress,json=workInProgress,proto3" json:"work_in_progress,omitempty"`
}

func (x *FileStatusAnnotation) Reset() {
	*x = FileStatusAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_annotations_v3_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileStatusAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStatusAnnotation) ProtoMessage() {}

func (x *FileStatusAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_xds_annotations_v3_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStatusAnnotation.ProtoReflect.Descriptor instead.
func (*FileStatusAnnotation) Descriptor() ([]byte, []int) {
	return file_xds_annotations_v3_status_proto_rawDescGZIP(), []int{0}
}

func (x *FileStatusAnnotation) GetWorkInProgress() bool {
	if x != nil {
		return x.WorkInProgress
	}
	return false
}

type MessageStatusAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkInProgress bool `protobuf:"varint,1,opt,name=work_in_progress,json=workInProgress,proto3" json:"work_in_progress,omitempty"`
}

func (x *MessageStatusAnnotation) Reset() {
	*x = MessageStatusAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_annotations_v3_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageStatusAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageStatusAnnotation) ProtoMessage() {}

func (x *MessageStatusAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_xds_annotations_v3_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageStatusAnnotation.ProtoReflect.Descriptor instead.
func (*MessageStatusAnnotation) Descriptor() ([]byte, []int) {
	return file_xds_annotations_v3_status_proto_rawDescGZIP(), []int{1}
}

func (x *MessageStatusAnnotation) GetWorkInProgress() bool {
	if x != nil {
		return x.WorkInProgress
	}
	return false
}

type FieldStatusAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkInProgress bool `protobuf:"varint,1,opt,name=work_in_progress,json=workInProgress,proto3" json:"work_in_progress,omitempty"`
}

func (x *FieldStatusAnnotation) Reset() {
	*x = FieldStatusAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_annotations_v3_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldStatusAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldStatusAnnotation) ProtoMessage() {}

func (x *FieldStatusAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_xds_annotations_v3_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldStatusAnnotation.ProtoReflect.Descriptor instead.
func (*FieldStatusAnnotation) Descriptor() ([]byte, []int) {
	return file_xds_annotations_v3_status_proto_rawDescGZIP(), []int{2}
}

func (x *FieldStatusAnnotation) GetWorkInProgress() bool {
	if x != nil {
		return x.WorkInProgress
	}
	return false
}

type StatusAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkInProgress       bool                 `protobuf:"varint,1,opt,name=work_in_progress,json=workInProgress,proto3" json:"work_in_progress,omitempty"`
	PackageVersionStatus PackageVersionStatus `protobuf:"varint,2,opt,name=package_version_status,json=packageVersionStatus,proto3,enum=xds.annotations.v3.PackageVersionStatus" json:"package_version_status,omitempty"`
}

func (x *StatusAnnotation) Reset() {
	*x = StatusAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_annotations_v3_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAnnotation) ProtoMessage() {}

func (x *StatusAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_xds_annotations_v3_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAnnotation.ProtoReflect.Descriptor instead.
func (*StatusAnnotation) Descriptor() ([]byte, []int) {
	return file_xds_annotations_v3_status_proto_rawDescGZIP(), []int{3}
}

func (x *StatusAnnotation) GetWorkInProgress() bool {
	if x != nil {
		return x.WorkInProgress
	}
	return false
}

func (x *StatusAnnotation) GetPackageVersionStatus() PackageVersionStatus {
	if x != nil {
		return x.PackageVersionStatus
	}
	return PackageVersionStatus_UNKNOWN
}

var file_xds_annotations_v3_status_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileStatusAnnotation)(nil),
		Field:         226829418,
		Name:          "xds.annotations.v3.file_status",
		Tag:           "bytes,226829418,opt,name=file_status",
		Filename:      "xds/annotations/v3/status.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageStatusAnnotation)(nil),
		Field:         226829418,
		Name:          "xds.annotations.v3.message_status",
		Tag:           "bytes,226829418,opt,name=message_status",
		Filename:      "xds/annotations/v3/status.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldStatusAnnotation)(nil),
		Field:         226829418,
		Name:          "xds.annotations.v3.field_status",
		Tag:           "bytes,226829418,opt,name=field_status",
		Filename:      "xds/annotations/v3/status.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional xds.annotations.v3.FileStatusAnnotation file_status = 226829418;
	E_FileStatus = &file_xds_annotations_v3_status_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional xds.annotations.v3.MessageStatusAnnotation message_status = 226829418;
	E_MessageStatus = &file_xds_annotations_v3_status_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional xds.annotations.v3.FieldStatusAnnotation field_status = 226829418;
	E_FieldStatus = &file_xds_annotations_v3_status_proto_extTypes[2]
)

var File_xds_annotations_v3_status_proto protoreflect.FileDescriptor

var file_xds_annotations_v3_status_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x78, 0x64, 0x73, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x43, 0x0a, 0x17, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x41,
	0x0a, 0x15, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x5e, 0x0a, 0x16, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x14, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x5d, 0x0a, 0x14, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x4f, 0x5a, 0x45, 0x4e, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x20, 0x0a,
	0x1c, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x4d, 0x41, 0x4a, 0x4f, 0x52, 0x5f, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x3a,
	0x6a, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xea, 0xc8, 0x94,
	0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x76, 0x0a, 0x0e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xea,
	0xc8, 0x94, 0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x3a, 0x6e, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xea, 0xc8, 0x94, 0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x78, 0x64,
	0x73, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x78, 0x64,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_annotations_v3_status_proto_rawDescOnce sync.Once
	file_xds_annotations_v3_status_proto_rawDescData = file_xds_annotations_v3_status_proto_rawDesc
)

func file_xds_annotations_v3_status_proto_rawDescGZIP() []byte {
	file_xds_annotations_v3_status_proto_rawDescOnce.Do(func() {
		file_xds_annotations_v3_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_annotations_v3_status_proto_rawDescData)
	})
	return file_xds_annotations_v3_status_proto_rawDescData
}

var file_xds_annotations_v3_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_xds_annotations_v3_status_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_xds_annotations_v3_status_proto_goTypes = []interface{}{
	(PackageVersionStatus)(0),           // 0: xds.annotations.v3.PackageVersionStatus
	(*FileStatusAnnotation)(nil),        // 1: xds.annotations.v3.FileStatusAnnotation
	(*MessageStatusAnnotation)(nil),     // 2: xds.annotations.v3.MessageStatusAnnotation
	(*FieldStatusAnnotation)(nil),       // 3: xds.annotations.v3.FieldStatusAnnotation
	(*StatusAnnotation)(nil),            // 4: xds.annotations.v3.StatusAnnotation
	(*descriptorpb.FileOptions)(nil),    // 5: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 6: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 7: google.protobuf.FieldOptions
}
var file_xds_annotations_v3_status_proto_depIdxs = []int32{
	0, // 0: xds.annotations.v3.StatusAnnotation.package_version_status:type_name -> xds.annotations.v3.PackageVersionStatus
	5, // 1: xds.annotations.v3.file_status:extendee -> google.protobuf.FileOptions
	6, // 2: xds.annotations.v3.message_status:extendee -> google.protobuf.MessageOptions
	7, // 3: xds.annotations.v3.field_status:extendee -> google.protobuf.FieldOptions
	1, // 4: xds.annotations.v3.file_status:type_name -> xds.annotations.v3.FileStatusAnnotation
	2, // 5: xds.annotations.v3.message_status:type_name -> xds.annotations.v3.MessageStatusAnnotation
	3, // 6: xds.annotations.v3.field_status:type_name -> xds.annotations.v3.FieldStatusAnnotation
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	4, // [4:7] is the sub-list for extension type_name
	1, // [1:4] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_xds_annotations_v3_status_proto_init() }
func file_xds_annotations_v3_status_proto_init() {
	if File_xds_annotations_v3_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xds_annotations_v3_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileStatusAnnotation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_xds_annotations_v3_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageStatusAnnotation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_xds_annotations_v3_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldStatusAnnotation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_xds_annotations_v3_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusAnnotation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xds_annotations_v3_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_xds_annotations_v3_status_proto_goTypes,
		DependencyIndexes: file_xds_annotations_v3_status_proto_depIdxs,
		EnumInfos:         file_xds_annotations_v3_status_proto_enumTypes,
		MessageInfos:      file_xds_annotations_v3_status_proto_msgTypes,
		ExtensionInfos:    file_xds_annotations_v3_status_proto_extTypes,
	}.Build()
	File_xds_annotations_v3_status_proto = out.File
	file_xds_annotations_v3_status_proto_rawDesc = nil
	file_xds_annotations_v3_status_proto_goTypes = nil
	file_xds_annotations_v3_status_proto_depIdxs = nil
}
