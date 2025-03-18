// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: kk_etcd_models/pb_user_kk_etcd.proto

package kk_etcd_models

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

type PBUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserName      string                 `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Roles         []string               `protobuf:"bytes,4,rep,name=Roles,proto3" json:"Roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PBUser) Reset() {
	*x = PBUser{}
	mi := &file_kk_etcd_models_pb_user_kk_etcd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PBUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBUser) ProtoMessage() {}

func (x *PBUser) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_models_pb_user_kk_etcd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBUser.ProtoReflect.Descriptor instead.
func (*PBUser) Descriptor() ([]byte, []int) {
	return file_kk_etcd_models_pb_user_kk_etcd_proto_rawDescGZIP(), []int{0}
}

func (x *PBUser) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *PBUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PBUser) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type PBListUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ListUser      []*PBUser              `protobuf:"bytes,1,rep,name=ListUser,proto3" json:"ListUser,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PBListUser) Reset() {
	*x = PBListUser{}
	mi := &file_kk_etcd_models_pb_user_kk_etcd_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PBListUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBListUser) ProtoMessage() {}

func (x *PBListUser) ProtoReflect() protoreflect.Message {
	mi := &file_kk_etcd_models_pb_user_kk_etcd_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBListUser.ProtoReflect.Descriptor instead.
func (*PBListUser) Descriptor() ([]byte, []int) {
	return file_kk_etcd_models_pb_user_kk_etcd_proto_rawDescGZIP(), []int{1}
}

func (x *PBListUser) GetListUser() []*PBUser {
	if x != nil {
		return x.ListUser
	}
	return nil
}

var File_kk_etcd_models_pb_user_kk_etcd_proto protoreflect.FileDescriptor

var file_kk_etcd_models_pb_user_kk_etcd_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x56, 0x0a, 0x06, 0x50, 0x42, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x40,
	0x0a, 0x0a, 0x50, 0x42, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x50, 0x42, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x72, 0x75, 0x76, 0x69, 0x65, 0x2f, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x67, 0x6f,
	0x2f, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_kk_etcd_models_pb_user_kk_etcd_proto_rawDescOnce sync.Once
	file_kk_etcd_models_pb_user_kk_etcd_proto_rawDescData []byte
)

func file_kk_etcd_models_pb_user_kk_etcd_proto_rawDescGZIP() []byte {
	file_kk_etcd_models_pb_user_kk_etcd_proto_rawDescOnce.Do(func() {
		file_kk_etcd_models_pb_user_kk_etcd_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kk_etcd_models_pb_user_kk_etcd_proto_rawDesc), len(file_kk_etcd_models_pb_user_kk_etcd_proto_rawDesc)))
	})
	return file_kk_etcd_models_pb_user_kk_etcd_proto_rawDescData
}

var file_kk_etcd_models_pb_user_kk_etcd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kk_etcd_models_pb_user_kk_etcd_proto_goTypes = []any{
	(*PBUser)(nil),     // 0: kk_etcd_models.PBUser
	(*PBListUser)(nil), // 1: kk_etcd_models.PBListUser
}
var file_kk_etcd_models_pb_user_kk_etcd_proto_depIdxs = []int32{
	0, // 0: kk_etcd_models.PBListUser.ListUser:type_name -> kk_etcd_models.PBUser
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kk_etcd_models_pb_user_kk_etcd_proto_init() }
func file_kk_etcd_models_pb_user_kk_etcd_proto_init() {
	if File_kk_etcd_models_pb_user_kk_etcd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kk_etcd_models_pb_user_kk_etcd_proto_rawDesc), len(file_kk_etcd_models_pb_user_kk_etcd_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kk_etcd_models_pb_user_kk_etcd_proto_goTypes,
		DependencyIndexes: file_kk_etcd_models_pb_user_kk_etcd_proto_depIdxs,
		MessageInfos:      file_kk_etcd_models_pb_user_kk_etcd_proto_msgTypes,
	}.Build()
	File_kk_etcd_models_pb_user_kk_etcd_proto = out.File
	file_kk_etcd_models_pb_user_kk_etcd_proto_goTypes = nil
	file_kk_etcd_models_pb_user_kk_etcd_proto_depIdxs = nil
}
