// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: distributedchessboardgeneration/dcg.proto

package distributedchessboardgeneration

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReturnMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip       string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *ReturnMessage) Reset() {
	*x = ReturnMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distributedchessboardgeneration_dcg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnMessage) ProtoMessage() {}

func (x *ReturnMessage) ProtoReflect() protoreflect.Message {
	mi := &file_distributedchessboardgeneration_dcg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnMessage.ProtoReflect.Descriptor instead.
func (*ReturnMessage) Descriptor() ([]byte, []int) {
	return file_distributedchessboardgeneration_dcg_proto_rawDescGZIP(), []int{0}
}

func (x *ReturnMessage) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ReturnMessage) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type Emptyy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Emptyy) Reset() {
	*x = Emptyy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distributedchessboardgeneration_dcg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emptyy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emptyy) ProtoMessage() {}

func (x *Emptyy) ProtoReflect() protoreflect.Message {
	mi := &file_distributedchessboardgeneration_dcg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emptyy.ProtoReflect.Descriptor instead.
func (*Emptyy) Descriptor() ([]byte, []int) {
	return file_distributedchessboardgeneration_dcg_proto_rawDescGZIP(), []int{1}
}

// Message can be empty as the server always sends the next
// available task.  However, the caller will send the server
// for logging purposes and future potential "heart-beat" messages.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distributedchessboardgeneration_dcg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_distributedchessboardgeneration_dcg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_distributedchessboardgeneration_dcg_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type ChessboardString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Board      string `protobuf:"bytes,1,opt,name=board,proto3" json:"board,omitempty"`
	IsFinished bool   `protobuf:"varint,2,opt,name=isFinished,proto3" json:"isFinished,omitempty"`
}

func (x *ChessboardString) Reset() {
	*x = ChessboardString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distributedchessboardgeneration_dcg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessboardString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessboardString) ProtoMessage() {}

func (x *ChessboardString) ProtoReflect() protoreflect.Message {
	mi := &file_distributedchessboardgeneration_dcg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessboardString.ProtoReflect.Descriptor instead.
func (*ChessboardString) Descriptor() ([]byte, []int) {
	return file_distributedchessboardgeneration_dcg_proto_rawDescGZIP(), []int{3}
}

func (x *ChessboardString) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

func (x *ChessboardString) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

var File_distributedchessboardgeneration_dcg_proto protoreflect.FileDescriptor

var file_distributedchessboardgeneration_dcg_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x63, 0x68, 0x65,
	0x73, 0x73, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x64, 0x63, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x63, 0x68, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x0d,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x08, 0x0a, 0x06, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x79, 0x22, 0x19, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x48,
	0x0a, 0x10, 0x43, 0x68, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x32, 0x82, 0x01, 0x0a, 0x18, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x43, 0x62, 0x12, 0x28,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x63, 0x68, 0x65, 0x73,
	0x73, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x63, 0x68, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x32, 0x86, 0x01,
	0x0a, 0x1d, 0x43, 0x68, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x52, 0x4c, 0x12,
	0x65, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x62, 0x12, 0x2e, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x63, 0x68, 0x65, 0x73, 0x73, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x27, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x63, 0x68, 0x65, 0x73, 0x73, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x79, 0x22, 0x00, 0x42, 0x62, 0x5a, 0x60, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x61,
	0x74, 0x6f, 0x6e, 0x42, 0x72, 0x69, 0x6e, 0x6b, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x75, 0x62,
	0x75, 0x74, 0x65, 0x64, 0x2d, 0x63, 0x68, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x63, 0x68, 0x65, 0x73, 0x73, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_distributedchessboardgeneration_dcg_proto_rawDescOnce sync.Once
	file_distributedchessboardgeneration_dcg_proto_rawDescData = file_distributedchessboardgeneration_dcg_proto_rawDesc
)

func file_distributedchessboardgeneration_dcg_proto_rawDescGZIP() []byte {
	file_distributedchessboardgeneration_dcg_proto_rawDescOnce.Do(func() {
		file_distributedchessboardgeneration_dcg_proto_rawDescData = protoimpl.X.CompressGZIP(file_distributedchessboardgeneration_dcg_proto_rawDescData)
	})
	return file_distributedchessboardgeneration_dcg_proto_rawDescData
}

var file_distributedchessboardgeneration_dcg_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_distributedchessboardgeneration_dcg_proto_goTypes = []interface{}{
	(*ReturnMessage)(nil),    // 0: distributedchessboardgeneration.ReturnMessage
	(*Emptyy)(nil),           // 1: distributedchessboardgeneration.Emptyy
	(*Message)(nil),          // 2: distributedchessboardgeneration.Message
	(*ChessboardString)(nil), // 3: distributedchessboardgeneration.ChessboardString
}
var file_distributedchessboardgeneration_dcg_proto_depIdxs = []int32{
	2, // 0: distributedchessboardgeneration.ChessboardTaskAssignment.GetCb:input_type -> distributedchessboardgeneration.Message
	0, // 1: distributedchessboardgeneration.ChessboardReturnAssignmentURL.ReturnCb:input_type -> distributedchessboardgeneration.ReturnMessage
	3, // 2: distributedchessboardgeneration.ChessboardTaskAssignment.GetCb:output_type -> distributedchessboardgeneration.ChessboardString
	1, // 3: distributedchessboardgeneration.ChessboardReturnAssignmentURL.ReturnCb:output_type -> distributedchessboardgeneration.Emptyy
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_distributedchessboardgeneration_dcg_proto_init() }
func file_distributedchessboardgeneration_dcg_proto_init() {
	if File_distributedchessboardgeneration_dcg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_distributedchessboardgeneration_dcg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnMessage); i {
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
		file_distributedchessboardgeneration_dcg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emptyy); i {
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
		file_distributedchessboardgeneration_dcg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_distributedchessboardgeneration_dcg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessboardString); i {
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
			RawDescriptor: file_distributedchessboardgeneration_dcg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_distributedchessboardgeneration_dcg_proto_goTypes,
		DependencyIndexes: file_distributedchessboardgeneration_dcg_proto_depIdxs,
		MessageInfos:      file_distributedchessboardgeneration_dcg_proto_msgTypes,
	}.Build()
	File_distributedchessboardgeneration_dcg_proto = out.File
	file_distributedchessboardgeneration_dcg_proto_rawDesc = nil
	file_distributedchessboardgeneration_dcg_proto_goTypes = nil
	file_distributedchessboardgeneration_dcg_proto_depIdxs = nil
}
