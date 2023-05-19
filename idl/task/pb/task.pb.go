// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: idl/task/task.proto

package pb

import (
	"reflect"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"favorite_id"
	TaskID int64 `protobuf:"varint,1,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	// @inject_tag: json:"user_id"
	UserID int64 `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	// @inject_tag: json:"status"
	Status int64 `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	// @inject_tag: json:"content"
	Content string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	// @inject_tag: json:"start_time"
	StartTime int64 `protobuf:"varint,6,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	// @inject_tag: json:"end_time"
	EndTime int64 `protobuf:"varint,7,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *TaskModel) Reset() {
	*x = TaskModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskModel) ProtoMessage() {}

func (x *TaskModel) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskModel.ProtoReflect.Descriptor instead.
func (*TaskModel) Descriptor() ([]byte, []int) {
	return file_idl_task_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskModel) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *TaskModel) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TaskModel) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskModel) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TaskModel) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TaskModel) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID    int64  `protobuf:"varint,1,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	UserID    int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Status    int64  `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Title     string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	Content   string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	StartTime int64  `protobuf:"varint,6,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime   int64  `protobuf:"varint,7,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_idl_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskRequest) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *TaskRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TaskRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TaskRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TaskRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type TasksDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskDetail []*TaskModel `protobuf:"bytes,1,rep,name=TaskDetail,proto3" json:"TaskDetail,omitempty"`
	Code       int64        `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *TasksDetailResponse) Reset() {
	*x = TasksDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksDetailResponse) ProtoMessage() {}

func (x *TasksDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksDetailResponse.ProtoReflect.Descriptor instead.
func (*TasksDetailResponse) Descriptor() ([]byte, []int) {
	return file_idl_task_task_proto_rawDescGZIP(), []int{2}
}

func (x *TasksDetailResponse) GetTaskDetail() []*TaskModel {
	if x != nil {
		return x.TaskDetail
	}
	return nil
}

func (x *TasksDetailResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_idl_task_task_proto_rawDescGZIP(), []int{3}
}

func (x *CommonResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CommonResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_idl_task_task_proto protoreflect.FileDescriptor

var file_idl_task_task_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x64, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xbb, 0x01, 0x0a, 0x09, 0x54, 0x61,
	0x73, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x13, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0xdc, 0x01,
	0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a,
	0x0a, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f,
	0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x68, 0x6f, 0x77, 0x12,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x54, 0x61, 0x73,
	0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_task_task_proto_rawDescOnce sync.Once
	file_idl_task_task_proto_rawDescData = file_idl_task_task_proto_rawDesc
)

func file_idl_task_task_proto_rawDescGZIP() []byte {
	file_idl_task_task_proto_rawDescOnce.Do(func() {
		file_idl_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_task_task_proto_rawDescData)
	})
	return file_idl_task_task_proto_rawDescData
}

var file_idl_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idl_task_task_proto_goTypes = []interface{}{
	(*TaskModel)(nil),           // 0: pb.TaskModel
	(*TaskRequest)(nil),         // 1: pb.TaskRequest
	(*TasksDetailResponse)(nil), // 2: pb.TasksDetailResponse
	(*CommonResponse)(nil),      // 3: pb.CommonResponse
}
var file_idl_task_task_proto_depIdxs = []int32{
	0, // 0: pb.TasksDetailResponse.TaskDetail:type_name -> pb.TaskModel
	1, // 1: pb.TaskService.TaskCreate:input_type -> pb.TaskRequest
	1, // 2: pb.TaskService.TaskUpdate:input_type -> pb.TaskRequest
	1, // 3: pb.TaskService.TaskShow:input_type -> pb.TaskRequest
	1, // 4: pb.TaskService.TaskDelete:input_type -> pb.TaskRequest
	3, // 5: pb.TaskService.TaskCreate:output_type -> pb.CommonResponse
	3, // 6: pb.TaskService.TaskUpdate:output_type -> pb.CommonResponse
	2, // 7: pb.TaskService.TaskShow:output_type -> pb.TasksDetailResponse
	3, // 8: pb.TaskService.TaskDelete:output_type -> pb.CommonResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_idl_task_task_proto_init() }
func file_idl_task_task_proto_init() {
	if File_idl_task_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_task_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskModel); i {
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
		file_idl_task_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_idl_task_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksDetailResponse); i {
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
		file_idl_task_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
			RawDescriptor: file_idl_task_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_task_task_proto_goTypes,
		DependencyIndexes: file_idl_task_task_proto_depIdxs,
		MessageInfos:      file_idl_task_task_proto_msgTypes,
	}.Build()
	File_idl_task_task_proto = out.File
	file_idl_task_task_proto_rawDesc = nil
	file_idl_task_task_proto_goTypes = nil
	file_idl_task_task_proto_depIdxs = nil
}