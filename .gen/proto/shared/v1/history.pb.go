// Copyright (c) 2020 Uber Technologies, Inc.
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
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: uber/cadence/shared/v1/history.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	v1 "github.com/uber/cadence/.gen/proto/api/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TransientDecisionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduledEvent *v1.HistoryEvent `protobuf:"bytes,1,opt,name=scheduled_event,json=scheduledEvent,proto3" json:"scheduled_event,omitempty"`
	StartedEvent   *v1.HistoryEvent `protobuf:"bytes,2,opt,name=started_event,json=startedEvent,proto3" json:"started_event,omitempty"`
}

func (x *TransientDecisionInfo) Reset() {
	*x = TransientDecisionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_shared_v1_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransientDecisionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransientDecisionInfo) ProtoMessage() {}

func (x *TransientDecisionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_shared_v1_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransientDecisionInfo.ProtoReflect.Descriptor instead.
func (*TransientDecisionInfo) Descriptor() ([]byte, []int) {
	return file_uber_cadence_shared_v1_history_proto_rawDescGZIP(), []int{0}
}

func (x *TransientDecisionInfo) GetScheduledEvent() *v1.HistoryEvent {
	if x != nil {
		return x.ScheduledEvent
	}
	return nil
}

func (x *TransientDecisionInfo) GetStartedEvent() *v1.HistoryEvent {
	if x != nil {
		return x.StartedEvent
	}
	return nil
}

// VersionHistoryItem contains signal eventId and the corresponding version.
type VersionHistoryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId int64 `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Version int64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionHistoryItem) Reset() {
	*x = VersionHistoryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_shared_v1_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionHistoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionHistoryItem) ProtoMessage() {}

func (x *VersionHistoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_shared_v1_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionHistoryItem.ProtoReflect.Descriptor instead.
func (*VersionHistoryItem) Descriptor() ([]byte, []int) {
	return file_uber_cadence_shared_v1_history_proto_rawDescGZIP(), []int{1}
}

func (x *VersionHistoryItem) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *VersionHistoryItem) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

// VersionHistory contains the version history of a branch.
type VersionHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchToken []byte                `protobuf:"bytes,1,opt,name=branch_token,json=branchToken,proto3" json:"branch_token,omitempty"`
	Items       []*VersionHistoryItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *VersionHistory) Reset() {
	*x = VersionHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_shared_v1_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionHistory) ProtoMessage() {}

func (x *VersionHistory) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_shared_v1_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionHistory.ProtoReflect.Descriptor instead.
func (*VersionHistory) Descriptor() ([]byte, []int) {
	return file_uber_cadence_shared_v1_history_proto_rawDescGZIP(), []int{2}
}

func (x *VersionHistory) GetBranchToken() []byte {
	if x != nil {
		return x.BranchToken
	}
	return nil
}

func (x *VersionHistory) GetItems() []*VersionHistoryItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// VersionHistories contains all version histories from all branches.
type VersionHistories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentVersionHistoryIndex int32             `protobuf:"varint,1,opt,name=current_version_history_index,json=currentVersionHistoryIndex,proto3" json:"current_version_history_index,omitempty"`
	Histories                  []*VersionHistory `protobuf:"bytes,2,rep,name=histories,proto3" json:"histories,omitempty"`
}

func (x *VersionHistories) Reset() {
	*x = VersionHistories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_shared_v1_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionHistories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionHistories) ProtoMessage() {}

func (x *VersionHistories) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_shared_v1_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionHistories.ProtoReflect.Descriptor instead.
func (*VersionHistories) Descriptor() ([]byte, []int) {
	return file_uber_cadence_shared_v1_history_proto_rawDescGZIP(), []int{3}
}

func (x *VersionHistories) GetCurrentVersionHistoryIndex() int32 {
	if x != nil {
		return x.CurrentVersionHistoryIndex
	}
	return 0
}

func (x *VersionHistories) GetHistories() []*VersionHistory {
	if x != nil {
		return x.Histories
	}
	return nil
}

var File_uber_cadence_shared_v1_history_proto protoreflect.FileDescriptor

var file_uber_cadence_shared_v1_history_proto_rawDesc = []byte{
	0x0a, 0x24, 0x75, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x21,
	0x75, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x01, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x0f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0x49, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x0e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x40, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x9b, 0x01, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1a, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x44, 0x0a, 0x09, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x75,
	0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x62,
	0x65, 0x72, 0x2f, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x2e, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uber_cadence_shared_v1_history_proto_rawDescOnce sync.Once
	file_uber_cadence_shared_v1_history_proto_rawDescData = file_uber_cadence_shared_v1_history_proto_rawDesc
)

func file_uber_cadence_shared_v1_history_proto_rawDescGZIP() []byte {
	file_uber_cadence_shared_v1_history_proto_rawDescOnce.Do(func() {
		file_uber_cadence_shared_v1_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_uber_cadence_shared_v1_history_proto_rawDescData)
	})
	return file_uber_cadence_shared_v1_history_proto_rawDescData
}

var file_uber_cadence_shared_v1_history_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_uber_cadence_shared_v1_history_proto_goTypes = []interface{}{
	(*TransientDecisionInfo)(nil), // 0: uber.cadence.shared.v1.TransientDecisionInfo
	(*VersionHistoryItem)(nil),    // 1: uber.cadence.shared.v1.VersionHistoryItem
	(*VersionHistory)(nil),        // 2: uber.cadence.shared.v1.VersionHistory
	(*VersionHistories)(nil),      // 3: uber.cadence.shared.v1.VersionHistories
	(*v1.HistoryEvent)(nil),       // 4: uber.cadence.api.v1.HistoryEvent
}
var file_uber_cadence_shared_v1_history_proto_depIdxs = []int32{
	4, // 0: uber.cadence.shared.v1.TransientDecisionInfo.scheduled_event:type_name -> uber.cadence.api.v1.HistoryEvent
	4, // 1: uber.cadence.shared.v1.TransientDecisionInfo.started_event:type_name -> uber.cadence.api.v1.HistoryEvent
	1, // 2: uber.cadence.shared.v1.VersionHistory.items:type_name -> uber.cadence.shared.v1.VersionHistoryItem
	2, // 3: uber.cadence.shared.v1.VersionHistories.histories:type_name -> uber.cadence.shared.v1.VersionHistory
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_uber_cadence_shared_v1_history_proto_init() }
func file_uber_cadence_shared_v1_history_proto_init() {
	if File_uber_cadence_shared_v1_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uber_cadence_shared_v1_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransientDecisionInfo); i {
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
		file_uber_cadence_shared_v1_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionHistoryItem); i {
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
		file_uber_cadence_shared_v1_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionHistory); i {
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
		file_uber_cadence_shared_v1_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionHistories); i {
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
			RawDescriptor: file_uber_cadence_shared_v1_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uber_cadence_shared_v1_history_proto_goTypes,
		DependencyIndexes: file_uber_cadence_shared_v1_history_proto_depIdxs,
		MessageInfos:      file_uber_cadence_shared_v1_history_proto_msgTypes,
	}.Build()
	File_uber_cadence_shared_v1_history_proto = out.File
	file_uber_cadence_shared_v1_history_proto_rawDesc = nil
	file_uber_cadence_shared_v1_history_proto_goTypes = nil
	file_uber_cadence_shared_v1_history_proto_depIdxs = nil
}
