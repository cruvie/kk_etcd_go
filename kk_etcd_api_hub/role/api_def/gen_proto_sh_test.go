package api_def

import (
	"gitee.com/cruvie/kk_go_kit/kk_protobuf"
	"testing"
)

// 此文件必须在HandlerPrePath目录中
func TestGenProtoSh(t *testing.T) {
	in := kk_protobuf.GenProtoShInput{
		HandlerPrePath:    "kk_etcd_api_hub/role/",
		HandlerFolderList: []string{"."},
		InternalProto: []string{
			"google/protobuf/timestamp.proto",
		},
		TsBufYamlPath: "",
		AdditionalProto: []string{
			"kk_etcd_models/*proto",
		},
	}
	kk_protobuf.GenProtoSh(in)
}
