// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: internal/mgr_hub/mgr_user/logout/api.proto

package logout

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// logout
type Logout struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Logout) Reset() {
	*x = Logout{}
	mi := &file_internal_mgr_hub_mgr_user_logout_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logout) ProtoMessage() {}

func (x *Logout) ProtoReflect() protoreflect.Message {
	mi := &file_internal_mgr_hub_mgr_user_logout_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logout.ProtoReflect.Descriptor instead.
func (*Logout) Descriptor() ([]byte, []int) {
	return file_internal_mgr_hub_mgr_user_logout_api_proto_rawDescGZIP(), []int{0}
}

type Logout_Input struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Logout_Input) Reset() {
	*x = Logout_Input{}
	mi := &file_internal_mgr_hub_mgr_user_logout_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logout_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logout_Input) ProtoMessage() {}

func (x *Logout_Input) ProtoReflect() protoreflect.Message {
	mi := &file_internal_mgr_hub_mgr_user_logout_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logout_Input.ProtoReflect.Descriptor instead.
func (*Logout_Input) Descriptor() ([]byte, []int) {
	return file_internal_mgr_hub_mgr_user_logout_api_proto_rawDescGZIP(), []int{0, 0}
}

type Logout_Output struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Logout_Output) Reset() {
	*x = Logout_Output{}
	mi := &file_internal_mgr_hub_mgr_user_logout_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Logout_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logout_Output) ProtoMessage() {}

func (x *Logout_Output) ProtoReflect() protoreflect.Message {
	mi := &file_internal_mgr_hub_mgr_user_logout_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logout_Output.ProtoReflect.Descriptor instead.
func (*Logout_Output) Descriptor() ([]byte, []int) {
	return file_internal_mgr_hub_mgr_user_logout_api_proto_rawDescGZIP(), []int{0, 1}
}

var File_internal_mgr_hub_mgr_user_logout_api_proto protoreflect.FileDescriptor

var file_internal_mgr_hub_mgr_user_logout_api_proto_rawDesc = string([]byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x67, 0x72, 0x5f, 0x68,
	0x75, 0x62, 0x2f, 0x6d, 0x67, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x22, 0x1b, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x1a, 0x07,
	0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x08, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_internal_mgr_hub_mgr_user_logout_api_proto_rawDescOnce sync.Once
	file_internal_mgr_hub_mgr_user_logout_api_proto_rawDescData []byte
)

func file_internal_mgr_hub_mgr_user_logout_api_proto_rawDescGZIP() []byte {
	file_internal_mgr_hub_mgr_user_logout_api_proto_rawDescOnce.Do(func() {
		file_internal_mgr_hub_mgr_user_logout_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_mgr_hub_mgr_user_logout_api_proto_rawDesc), len(file_internal_mgr_hub_mgr_user_logout_api_proto_rawDesc)))
	})
	return file_internal_mgr_hub_mgr_user_logout_api_proto_rawDescData
}

var file_internal_mgr_hub_mgr_user_logout_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_mgr_hub_mgr_user_logout_api_proto_goTypes = []any{
	(*Logout)(nil),        // 0: logout.Logout
	(*Logout_Input)(nil),  // 1: logout.Logout.Input
	(*Logout_Output)(nil), // 2: logout.Logout.Output
}
var file_internal_mgr_hub_mgr_user_logout_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_mgr_hub_mgr_user_logout_api_proto_init() }
func file_internal_mgr_hub_mgr_user_logout_api_proto_init() {
	if File_internal_mgr_hub_mgr_user_logout_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_mgr_hub_mgr_user_logout_api_proto_rawDesc), len(file_internal_mgr_hub_mgr_user_logout_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_mgr_hub_mgr_user_logout_api_proto_goTypes,
		DependencyIndexes: file_internal_mgr_hub_mgr_user_logout_api_proto_depIdxs,
		MessageInfos:      file_internal_mgr_hub_mgr_user_logout_api_proto_msgTypes,
	}.Build()
	File_internal_mgr_hub_mgr_user_logout_api_proto = out.File
	file_internal_mgr_hub_mgr_user_logout_api_proto_goTypes = nil
	file_internal_mgr_hub_mgr_user_logout_api_proto_depIdxs = nil
}
