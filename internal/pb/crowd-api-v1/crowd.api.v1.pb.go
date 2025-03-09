// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: crowd-api-v1/crowd.api.v1.proto

package crowd_api_v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_crowd_api_v1_crowd_api_v1_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_crowd_api_v1_crowd_api_v1_proto_rawDescGZIP(), []int{1}
}

type ResolveTasksByProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	UserId    int32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ResolveTasksByProjectRequest) Reset() {
	*x = ResolveTasksByProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveTasksByProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveTasksByProjectRequest) ProtoMessage() {}

func (x *ResolveTasksByProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveTasksByProjectRequest.ProtoReflect.Descriptor instead.
func (*ResolveTasksByProjectRequest) Descriptor() ([]byte, []int) {
	return file_crowd_api_v1_crowd_api_v1_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveTasksByProjectRequest) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *ResolveTasksByProjectRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ResolveTasksByProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*UserTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *ResolveTasksByProjectResponse) Reset() {
	*x = ResolveTasksByProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveTasksByProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveTasksByProjectResponse) ProtoMessage() {}

func (x *ResolveTasksByProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveTasksByProjectResponse.ProtoReflect.Descriptor instead.
func (*ResolveTasksByProjectResponse) Descriptor() ([]byte, []int) {
	return file_crowd_api_v1_crowd_api_v1_proto_rawDescGZIP(), []int{3}
}

func (x *ResolveTasksByProjectResponse) GetTasks() []*UserTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type UserTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InputData string                 `protobuf:"bytes,2,opt,name=input_data,json=inputData,proto3" json:"input_data,omitempty"`
	Deadline  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *UserTask) Reset() {
	*x = UserTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTask) ProtoMessage() {}

func (x *UserTask) ProtoReflect() protoreflect.Message {
	mi := &file_crowd_api_v1_crowd_api_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTask.ProtoReflect.Descriptor instead.
func (*UserTask) Descriptor() ([]byte, []int) {
	return file_crowd_api_v1_crowd_api_v1_proto_rawDescGZIP(), []int{4}
}

func (x *UserTask) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserTask) GetInputData() string {
	if x != nil {
		return x.InputData
	}
	return ""
}

func (x *UserTask) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

var File_crowd_api_v1_crowd_api_v1_proto protoreflect.FileDescriptor

var file_crowd_api_v1_crowd_api_v1_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x6f, 0x77, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d,
	0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a,
	0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x0a,
	0x1c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x79, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x1d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x22, 0x71, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x36, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x32, 0x8b, 0x02, 0x0a, 0x0a, 0x43, 0x72, 0x6f, 0x77,
	0x64, 0x41, 0x50, 0x49, 0x56, 0x31, 0x12, 0x53, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x19,
	0x2e, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x72, 0x6f, 0x77,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x12, 0xa7, 0x01, 0x0a, 0x15,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x79, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2a, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x79, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x62, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x6c, 0x61, 0x6b, 0x33, 0x34, 0x2f, 0x63, 0x72, 0x6f, 0x77, 0x64,
	0x2d, 0x61, 0x70, 0x69, 0x3b, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crowd_api_v1_crowd_api_v1_proto_rawDescOnce sync.Once
	file_crowd_api_v1_crowd_api_v1_proto_rawDescData = file_crowd_api_v1_crowd_api_v1_proto_rawDesc
)

func file_crowd_api_v1_crowd_api_v1_proto_rawDescGZIP() []byte {
	file_crowd_api_v1_crowd_api_v1_proto_rawDescOnce.Do(func() {
		file_crowd_api_v1_crowd_api_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_crowd_api_v1_crowd_api_v1_proto_rawDescData)
	})
	return file_crowd_api_v1_crowd_api_v1_proto_rawDescData
}

var file_crowd_api_v1_crowd_api_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_crowd_api_v1_crowd_api_v1_proto_goTypes = []interface{}{
	(*PingRequest)(nil),                   // 0: crowd.api.v1.PingRequest
	(*PingResponse)(nil),                  // 1: crowd.api.v1.PingResponse
	(*ResolveTasksByProjectRequest)(nil),  // 2: crowd.api.v1.ResolveTasksByProjectRequest
	(*ResolveTasksByProjectResponse)(nil), // 3: crowd.api.v1.ResolveTasksByProjectResponse
	(*UserTask)(nil),                      // 4: crowd.api.v1.UserTask
	(*timestamppb.Timestamp)(nil),         // 5: google.protobuf.Timestamp
}
var file_crowd_api_v1_crowd_api_v1_proto_depIdxs = []int32{
	4, // 0: crowd.api.v1.ResolveTasksByProjectResponse.tasks:type_name -> crowd.api.v1.UserTask
	5, // 1: crowd.api.v1.UserTask.deadline:type_name -> google.protobuf.Timestamp
	0, // 2: crowd.api.v1.CrowdAPIV1.Ping:input_type -> crowd.api.v1.PingRequest
	2, // 3: crowd.api.v1.CrowdAPIV1.ResolveTasksByProject:input_type -> crowd.api.v1.ResolveTasksByProjectRequest
	1, // 4: crowd.api.v1.CrowdAPIV1.Ping:output_type -> crowd.api.v1.PingResponse
	3, // 5: crowd.api.v1.CrowdAPIV1.ResolveTasksByProject:output_type -> crowd.api.v1.ResolveTasksByProjectResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_crowd_api_v1_crowd_api_v1_proto_init() }
func file_crowd_api_v1_crowd_api_v1_proto_init() {
	if File_crowd_api_v1_crowd_api_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crowd_api_v1_crowd_api_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_crowd_api_v1_crowd_api_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_crowd_api_v1_crowd_api_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveTasksByProjectRequest); i {
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
		file_crowd_api_v1_crowd_api_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveTasksByProjectResponse); i {
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
		file_crowd_api_v1_crowd_api_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTask); i {
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
			RawDescriptor: file_crowd_api_v1_crowd_api_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crowd_api_v1_crowd_api_v1_proto_goTypes,
		DependencyIndexes: file_crowd_api_v1_crowd_api_v1_proto_depIdxs,
		MessageInfos:      file_crowd_api_v1_crowd_api_v1_proto_msgTypes,
	}.Build()
	File_crowd_api_v1_crowd_api_v1_proto = out.File
	file_crowd_api_v1_crowd_api_v1_proto_rawDesc = nil
	file_crowd_api_v1_crowd_api_v1_proto_goTypes = nil
	file_crowd_api_v1_crowd_api_v1_proto_depIdxs = nil
}
