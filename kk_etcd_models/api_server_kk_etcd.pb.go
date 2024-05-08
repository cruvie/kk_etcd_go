// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: api_server_kk_etcd.proto

package kk_etcd_models

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

type ServerListParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix string `protobuf:"bytes,1,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
}

func (x *ServerListParam) Reset() {
	*x = ServerListParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_kk_etcd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerListParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerListParam) ProtoMessage() {}

func (x *ServerListParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_kk_etcd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerListParam.ProtoReflect.Descriptor instead.
func (*ServerListParam) Descriptor() ([]byte, []int) {
	return file_api_server_kk_etcd_proto_rawDescGZIP(), []int{0}
}

func (x *ServerListParam) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type ServerListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerList *PBListServer `protobuf:"bytes,1,opt,name=ServerList,proto3" json:"ServerList,omitempty"`
}

func (x *ServerListResponse) Reset() {
	*x = ServerListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_kk_etcd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerListResponse) ProtoMessage() {}

func (x *ServerListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_kk_etcd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerListResponse.ProtoReflect.Descriptor instead.
func (*ServerListResponse) Descriptor() ([]byte, []int) {
	return file_api_server_kk_etcd_proto_rawDescGZIP(), []int{1}
}

func (x *ServerListResponse) GetServerList() *PBListServer {
	if x != nil {
		return x.ServerList
	}
	return nil
}

var File_api_server_kk_etcd_proto protoreflect.FileDescriptor

var file_api_server_kk_etcd_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6b, 0x6b, 0x5f,
	0x65, 0x74, 0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x6b, 0x5f, 0x65,
	0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x0f, 0x70, 0x62, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0f, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x52, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x50, 0x42, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x0a,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f,
	0x3b, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_server_kk_etcd_proto_rawDescOnce sync.Once
	file_api_server_kk_etcd_proto_rawDescData = file_api_server_kk_etcd_proto_rawDesc
)

func file_api_server_kk_etcd_proto_rawDescGZIP() []byte {
	file_api_server_kk_etcd_proto_rawDescOnce.Do(func() {
		file_api_server_kk_etcd_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_server_kk_etcd_proto_rawDescData)
	})
	return file_api_server_kk_etcd_proto_rawDescData
}

var file_api_server_kk_etcd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_server_kk_etcd_proto_goTypes = []interface{}{
	(*ServerListParam)(nil),    // 0: kk_etcd_models.ServerListParam
	(*ServerListResponse)(nil), // 1: kk_etcd_models.ServerListResponse
	(*PBListServer)(nil),       // 2: kk_etcd_models.PBListServer
}
var file_api_server_kk_etcd_proto_depIdxs = []int32{
	2, // 0: kk_etcd_models.ServerListResponse.ServerList:type_name -> kk_etcd_models.PBListServer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_server_kk_etcd_proto_init() }
func file_api_server_kk_etcd_proto_init() {
	if File_api_server_kk_etcd_proto != nil {
		return
	}
	file_pb_server_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_server_kk_etcd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerListParam); i {
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
		file_api_server_kk_etcd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerListResponse); i {
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
			RawDescriptor: file_api_server_kk_etcd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_server_kk_etcd_proto_goTypes,
		DependencyIndexes: file_api_server_kk_etcd_proto_depIdxs,
		MessageInfos:      file_api_server_kk_etcd_proto_msgTypes,
	}.Build()
	File_api_server_kk_etcd_proto = out.File
	file_api_server_kk_etcd_proto_rawDesc = nil
	file_api_server_kk_etcd_proto_goTypes = nil
	file_api_server_kk_etcd_proto_depIdxs = nil
}