// Copyright 2024 Google LLC
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
// source: google/spanner/v1/query_plan.proto

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

// The kind of [PlanNode][google.spanner.v1.PlanNode]. Distinguishes between the two different kinds of
// nodes that can appear in a query plan.
type PlanNode_Kind int32

const (
	// Not specified.
	PlanNode_KIND_UNSPECIFIED PlanNode_Kind = 0
	// Denotes a Relational operator node in the expression tree. Relational
	// operators represent iterative processing of rows during query execution.
	// For example, a `TableScan` operation that reads rows from a table.
	PlanNode_RELATIONAL PlanNode_Kind = 1
	// Denotes a Scalar node in the expression tree. Scalar nodes represent
	// non-iterable entities in the query plan. For example, constants or
	// arithmetic operators appearing inside predicate expressions or references
	// to column names.
	PlanNode_SCALAR PlanNode_Kind = 2
)

// Enum value maps for PlanNode_Kind.
var (
	PlanNode_Kind_name = map[int32]string{
		0: "KIND_UNSPECIFIED",
		1: "RELATIONAL",
		2: "SCALAR",
	}
	PlanNode_Kind_value = map[string]int32{
		"KIND_UNSPECIFIED": 0,
		"RELATIONAL":       1,
		"SCALAR":           2,
	}
)

func (x PlanNode_Kind) Enum() *PlanNode_Kind {
	p := new(PlanNode_Kind)
	*p = x
	return p
}

func (x PlanNode_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanNode_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_google_spanner_v1_query_plan_proto_enumTypes[0].Descriptor()
}

func (PlanNode_Kind) Type() protoreflect.EnumType {
	return &file_google_spanner_v1_query_plan_proto_enumTypes[0]
}

func (x PlanNode_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanNode_Kind.Descriptor instead.
func (PlanNode_Kind) EnumDescriptor() ([]byte, []int) {
	return file_google_spanner_v1_query_plan_proto_rawDescGZIP(), []int{0, 0}
}

// Node information for nodes appearing in a [QueryPlan.plan_nodes][google.spanner.v1.QueryPlan.plan_nodes].
type PlanNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The `PlanNode`'s index in [node list][google.spanner.v1.QueryPlan.plan_nodes].
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// Used to determine the type of node. May be needed for visualizing
	// different kinds of nodes differently. For example, If the node is a
	// [SCALAR][google.spanner.v1.PlanNode.Kind.SCALAR] node, it will have a condensed representation
	// which can be used to directly embed a description of the node in its
	// parent.
	Kind PlanNode_Kind `protobuf:"varint,2,opt,name=kind,proto3,enum=google.spanner.v1.PlanNode_Kind" json:"kind,omitempty"`
	// The display name for the node.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// List of child node `index`es and their relationship to this parent.
	ChildLinks []*PlanNode_ChildLink `protobuf:"bytes,4,rep,name=child_links,json=childLinks,proto3" json:"child_links,omitempty"`
	// Condensed representation for [SCALAR][google.spanner.v1.PlanNode.Kind.SCALAR] nodes.
	ShortRepresentation *PlanNode_ShortRepresentation `protobuf:"bytes,5,opt,name=short_representation,json=shortRepresentation,proto3" json:"short_representation,omitempty"`
	// Attributes relevant to the node contained in a group of key-value pairs.
	// For example, a Parameter Reference node could have the following
	// information in its metadata:
	//
	//	{
	//	  "parameter_reference": "param1",
	//	  "parameter_type": "array"
	//	}
	Metadata *structpb.Struct `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The execution statistics associated with the node, contained in a group of
	// key-value pairs. Only present if the plan was returned as a result of a
	// profile query. For example, number of executions, number of rows/time per
	// execution etc.
	ExecutionStats *structpb.Struct `protobuf:"bytes,7,opt,name=execution_stats,json=executionStats,proto3" json:"execution_stats,omitempty"`
}

func (x *PlanNode) Reset() {
	*x = PlanNode{}
	mi := &file_google_spanner_v1_query_plan_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanNode) ProtoMessage() {}

func (x *PlanNode) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_query_plan_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanNode.ProtoReflect.Descriptor instead.
func (*PlanNode) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_query_plan_proto_rawDescGZIP(), []int{0}
}

func (x *PlanNode) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *PlanNode) GetKind() PlanNode_Kind {
	if x != nil {
		return x.Kind
	}
	return PlanNode_KIND_UNSPECIFIED
}

func (x *PlanNode) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *PlanNode) GetChildLinks() []*PlanNode_ChildLink {
	if x != nil {
		return x.ChildLinks
	}
	return nil
}

func (x *PlanNode) GetShortRepresentation() *PlanNode_ShortRepresentation {
	if x != nil {
		return x.ShortRepresentation
	}
	return nil
}

func (x *PlanNode) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PlanNode) GetExecutionStats() *structpb.Struct {
	if x != nil {
		return x.ExecutionStats
	}
	return nil
}

// Contains an ordered list of nodes appearing in the query plan.
type QueryPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The nodes in the query plan. Plan nodes are returned in pre-order starting
	// with the plan root. Each [PlanNode][google.spanner.v1.PlanNode]'s `id` corresponds to its index in
	// `plan_nodes`.
	PlanNodes []*PlanNode `protobuf:"bytes,1,rep,name=plan_nodes,json=planNodes,proto3" json:"plan_nodes,omitempty"`
}

func (x *QueryPlan) Reset() {
	*x = QueryPlan{}
	mi := &file_google_spanner_v1_query_plan_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPlan) ProtoMessage() {}

func (x *QueryPlan) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_query_plan_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPlan.ProtoReflect.Descriptor instead.
func (*QueryPlan) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_query_plan_proto_rawDescGZIP(), []int{1}
}

func (x *QueryPlan) GetPlanNodes() []*PlanNode {
	if x != nil {
		return x.PlanNodes
	}
	return nil
}

// Metadata associated with a parent-child relationship appearing in a
// [PlanNode][google.spanner.v1.PlanNode].
type PlanNode_ChildLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node to which the link points.
	ChildIndex int32 `protobuf:"varint,1,opt,name=child_index,json=childIndex,proto3" json:"child_index,omitempty"`
	// The type of the link. For example, in Hash Joins this could be used to
	// distinguish between the build child and the probe child, or in the case
	// of the child being an output variable, to represent the tag associated
	// with the output variable.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Only present if the child node is [SCALAR][google.spanner.v1.PlanNode.Kind.SCALAR] and corresponds
	// to an output variable of the parent node. The field carries the name of
	// the output variable.
	// For example, a `TableScan` operator that reads rows from a table will
	// have child links to the `SCALAR` nodes representing the output variables
	// created for each column that is read by the operator. The corresponding
	// `variable` fields will be set to the variable names assigned to the
	// columns.
	Variable string `protobuf:"bytes,3,opt,name=variable,proto3" json:"variable,omitempty"`
}

func (x *PlanNode_ChildLink) Reset() {
	*x = PlanNode_ChildLink{}
	mi := &file_google_spanner_v1_query_plan_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanNode_ChildLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanNode_ChildLink) ProtoMessage() {}

func (x *PlanNode_ChildLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_query_plan_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanNode_ChildLink.ProtoReflect.Descriptor instead.
func (*PlanNode_ChildLink) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_query_plan_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PlanNode_ChildLink) GetChildIndex() int32 {
	if x != nil {
		return x.ChildIndex
	}
	return 0
}

func (x *PlanNode_ChildLink) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PlanNode_ChildLink) GetVariable() string {
	if x != nil {
		return x.Variable
	}
	return ""
}

// Condensed representation of a node and its subtree. Only present for
// `SCALAR` [PlanNode(s)][google.spanner.v1.PlanNode].
type PlanNode_ShortRepresentation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A string representation of the expression subtree rooted at this node.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// A mapping of (subquery variable name) -> (subquery node id) for cases
	// where the `description` string of this node references a `SCALAR`
	// subquery contained in the expression subtree rooted at this node. The
	// referenced `SCALAR` subquery may not necessarily be a direct child of
	// this node.
	Subqueries map[string]int32 `protobuf:"bytes,2,rep,name=subqueries,proto3" json:"subqueries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *PlanNode_ShortRepresentation) Reset() {
	*x = PlanNode_ShortRepresentation{}
	mi := &file_google_spanner_v1_query_plan_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanNode_ShortRepresentation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanNode_ShortRepresentation) ProtoMessage() {}

func (x *PlanNode_ShortRepresentation) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_v1_query_plan_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanNode_ShortRepresentation.ProtoReflect.Descriptor instead.
func (*PlanNode_ShortRepresentation) Descriptor() ([]byte, []int) {
	return file_google_spanner_v1_query_plan_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PlanNode_ShortRepresentation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PlanNode_ShortRepresentation) GetSubqueries() map[string]int32 {
	if x != nil {
		return x.Subqueries
	}
	return nil
}

var File_google_spanner_v1_query_plan_proto protoreflect.FileDescriptor

var file_google_spanner_v1_query_plan_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x06, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x34, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x4e,
	0x6f, 0x64, 0x65, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x46, 0x0a, 0x0b, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x4e,
	0x6f, 0x64, 0x65, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0a, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x62, 0x0a, 0x14, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x40, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x1a, 0x5c, 0x0a, 0x09, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x1a, 0xd7, 0x01, 0x0a, 0x13, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x0a,
	0x73, 0x75, 0x62, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x75, 0x62, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x3d, 0x0a,
	0x0f, 0x53, 0x75, 0x62, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x04,
	0x4b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45,
	0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43,
	0x41, 0x4c, 0x41, 0x52, 0x10, 0x02, 0x22, 0x47, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6c, 0x61, 0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x42,
	0xb1, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73,
	0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x6c, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73,
	0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x70, 0x62, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x70, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_spanner_v1_query_plan_proto_rawDescOnce sync.Once
	file_google_spanner_v1_query_plan_proto_rawDescData = file_google_spanner_v1_query_plan_proto_rawDesc
)

func file_google_spanner_v1_query_plan_proto_rawDescGZIP() []byte {
	file_google_spanner_v1_query_plan_proto_rawDescOnce.Do(func() {
		file_google_spanner_v1_query_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_spanner_v1_query_plan_proto_rawDescData)
	})
	return file_google_spanner_v1_query_plan_proto_rawDescData
}

var file_google_spanner_v1_query_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_spanner_v1_query_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_spanner_v1_query_plan_proto_goTypes = []any{
	(PlanNode_Kind)(0),                   // 0: google.spanner.v1.PlanNode.Kind
	(*PlanNode)(nil),                     // 1: google.spanner.v1.PlanNode
	(*QueryPlan)(nil),                    // 2: google.spanner.v1.QueryPlan
	(*PlanNode_ChildLink)(nil),           // 3: google.spanner.v1.PlanNode.ChildLink
	(*PlanNode_ShortRepresentation)(nil), // 4: google.spanner.v1.PlanNode.ShortRepresentation
	nil,                                  // 5: google.spanner.v1.PlanNode.ShortRepresentation.SubqueriesEntry
	(*structpb.Struct)(nil),              // 6: google.protobuf.Struct
}
var file_google_spanner_v1_query_plan_proto_depIdxs = []int32{
	0, // 0: google.spanner.v1.PlanNode.kind:type_name -> google.spanner.v1.PlanNode.Kind
	3, // 1: google.spanner.v1.PlanNode.child_links:type_name -> google.spanner.v1.PlanNode.ChildLink
	4, // 2: google.spanner.v1.PlanNode.short_representation:type_name -> google.spanner.v1.PlanNode.ShortRepresentation
	6, // 3: google.spanner.v1.PlanNode.metadata:type_name -> google.protobuf.Struct
	6, // 4: google.spanner.v1.PlanNode.execution_stats:type_name -> google.protobuf.Struct
	1, // 5: google.spanner.v1.QueryPlan.plan_nodes:type_name -> google.spanner.v1.PlanNode
	5, // 6: google.spanner.v1.PlanNode.ShortRepresentation.subqueries:type_name -> google.spanner.v1.PlanNode.ShortRepresentation.SubqueriesEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_spanner_v1_query_plan_proto_init() }
func file_google_spanner_v1_query_plan_proto_init() {
	if File_google_spanner_v1_query_plan_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_spanner_v1_query_plan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_spanner_v1_query_plan_proto_goTypes,
		DependencyIndexes: file_google_spanner_v1_query_plan_proto_depIdxs,
		EnumInfos:         file_google_spanner_v1_query_plan_proto_enumTypes,
		MessageInfos:      file_google_spanner_v1_query_plan_proto_msgTypes,
	}.Build()
	File_google_spanner_v1_query_plan_proto = out.File
	file_google_spanner_v1_query_plan_proto_rawDesc = nil
	file_google_spanner_v1_query_plan_proto_goTypes = nil
	file_google_spanner_v1_query_plan_proto_depIdxs = nil
}
