// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: pb_server_kk_etcd.proto

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

type PBServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	ServiceAddr string `protobuf:"bytes,2,opt,name=ServiceAddr,proto3" json:"ServiceAddr,omitempty"`
}

func (x *PBServer) Reset() {
	*x = PBServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_server_kk_etcd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBServer) ProtoMessage() {}

func (x *PBServer) ProtoReflect() protoreflect.Message {
	mi := &file_pb_server_kk_etcd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBServer.ProtoReflect.Descriptor instead.
func (*PBServer) Descriptor() ([]byte, []int) {
	return file_pb_server_kk_etcd_proto_rawDescGZIP(), []int{0}
}

func (x *PBServer) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *PBServer) GetServiceAddr() string {
	if x != nil {
		return x.ServiceAddr
	}
	return ""
}

type PBListServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListServer []*PBServer `protobuf:"bytes,1,rep,name=ListServer,proto3" json:"ListServer,omitempty"`
}

func (x *PBListServer) Reset() {
	*x = PBListServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_server_kk_etcd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBListServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBListServer) ProtoMessage() {}

func (x *PBListServer) ProtoReflect() protoreflect.Message {
	mi := &file_pb_server_kk_etcd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBListServer.ProtoReflect.Descriptor instead.
func (*PBListServer) Descriptor() ([]byte, []int) {
	return file_pb_server_kk_etcd_proto_rawDescGZIP(), []int{1}
}

func (x *PBListServer) GetListServer() []*PBServer {
	if x != nil {
		return x.ListServer
	}
	return nil
}

var File_pb_server_kk_etcd_proto protoreflect.FileDescriptor

var file_pb_server_kk_etcd_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6b, 0x6b, 0x5f, 0x65,
	0x74, 0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x6b, 0x5f, 0x65, 0x74,
	0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x4e, 0x0a, 0x08, 0x50, 0x42, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x48, 0x0a, 0x0c, 0x50, 0x42, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50,
	0x42, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x3b, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63,
	0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_server_kk_etcd_proto_rawDescOnce sync.Once
	file_pb_server_kk_etcd_proto_rawDescData = file_pb_server_kk_etcd_proto_rawDesc
)

func file_pb_server_kk_etcd_proto_rawDescGZIP() []byte {
	file_pb_server_kk_etcd_proto_rawDescOnce.Do(func() {
		file_pb_server_kk_etcd_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_server_kk_etcd_proto_rawDescData)
	})
	return file_pb_server_kk_etcd_proto_rawDescData
}

var file_pb_server_kk_etcd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_server_kk_etcd_proto_goTypes = []interface{}{
	(*PBServer)(nil),     // 0: kk_etcd_models.PBServer
	(*PBListServer)(nil), // 1: kk_etcd_models.PBListServer
}
var file_pb_server_kk_etcd_proto_depIdxs = []int32{
	0, // 0: kk_etcd_models.PBListServer.ListServer:type_name -> kk_etcd_models.PBServer
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_server_kk_etcd_proto_init() }
func file_pb_server_kk_etcd_proto_init() {
	if File_pb_server_kk_etcd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_server_kk_etcd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBServer); i {
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
		file_pb_server_kk_etcd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBListServer); i {
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
			RawDescriptor: file_pb_server_kk_etcd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_server_kk_etcd_proto_goTypes,
		DependencyIndexes: file_pb_server_kk_etcd_proto_depIdxs,
		MessageInfos:      file_pb_server_kk_etcd_proto_msgTypes,
	}.Build()
	File_pb_server_kk_etcd_proto = out.File
	file_pb_server_kk_etcd_proto_rawDesc = nil
	file_pb_server_kk_etcd_proto_goTypes = nil
	file_pb_server_kk_etcd_proto_depIdxs = nil
}
