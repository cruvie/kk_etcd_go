// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: pb_kv_kk_etcd.proto

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

type PBKV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *PBKV) Reset() {
	*x = PBKV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_kv_kk_etcd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBKV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBKV) ProtoMessage() {}

func (x *PBKV) ProtoReflect() protoreflect.Message {
	mi := &file_pb_kv_kk_etcd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBKV.ProtoReflect.Descriptor instead.
func (*PBKV) Descriptor() ([]byte, []int) {
	return file_pb_kv_kk_etcd_proto_rawDescGZIP(), []int{0}
}

func (x *PBKV) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PBKV) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PBListKV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListKV []*PBKV `protobuf:"bytes,1,rep,name=ListKV,proto3" json:"ListKV,omitempty"`
}

func (x *PBListKV) Reset() {
	*x = PBListKV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_kv_kk_etcd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBListKV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBListKV) ProtoMessage() {}

func (x *PBListKV) ProtoReflect() protoreflect.Message {
	mi := &file_pb_kv_kk_etcd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBListKV.ProtoReflect.Descriptor instead.
func (*PBListKV) Descriptor() ([]byte, []int) {
	return file_pb_kv_kk_etcd_proto_rawDescGZIP(), []int{1}
}

func (x *PBListKV) GetListKV() []*PBKV {
	if x != nil {
		return x.ListKV
	}
	return nil
}

var File_pb_kv_kk_etcd_proto protoreflect.FileDescriptor

var file_pb_kv_kk_etcd_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x5f, 0x6b, 0x76, 0x5f, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x2e, 0x0a, 0x04, 0x50, 0x42, 0x4b, 0x56, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x08, 0x50, 0x42, 0x4c, 0x69, 0x73, 0x74, 0x4b,
	0x56, 0x12, 0x2c, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x56, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x50, 0x42, 0x4b, 0x56, 0x52, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x56, 0x42,
	0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x3b, 0x6b, 0x6b, 0x5f, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_kv_kk_etcd_proto_rawDescOnce sync.Once
	file_pb_kv_kk_etcd_proto_rawDescData = file_pb_kv_kk_etcd_proto_rawDesc
)

func file_pb_kv_kk_etcd_proto_rawDescGZIP() []byte {
	file_pb_kv_kk_etcd_proto_rawDescOnce.Do(func() {
		file_pb_kv_kk_etcd_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_kv_kk_etcd_proto_rawDescData)
	})
	return file_pb_kv_kk_etcd_proto_rawDescData
}

var file_pb_kv_kk_etcd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_kv_kk_etcd_proto_goTypes = []any{
	(*PBKV)(nil),     // 0: kk_etcd_models.PBKV
	(*PBListKV)(nil), // 1: kk_etcd_models.PBListKV
}
var file_pb_kv_kk_etcd_proto_depIdxs = []int32{
	0, // 0: kk_etcd_models.PBListKV.ListKV:type_name -> kk_etcd_models.PBKV
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_kv_kk_etcd_proto_init() }
func file_pb_kv_kk_etcd_proto_init() {
	if File_pb_kv_kk_etcd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_kv_kk_etcd_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PBKV); i {
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
		file_pb_kv_kk_etcd_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PBListKV); i {
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
			RawDescriptor: file_pb_kv_kk_etcd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_kv_kk_etcd_proto_goTypes,
		DependencyIndexes: file_pb_kv_kk_etcd_proto_depIdxs,
		MessageInfos:      file_pb_kv_kk_etcd_proto_msgTypes,
	}.Build()
	File_pb_kv_kk_etcd_proto = out.File
	file_pb_kv_kk_etcd_proto_rawDesc = nil
	file_pb_kv_kk_etcd_proto_goTypes = nil
	file_pb_kv_kk_etcd_proto_depIdxs = nil
}