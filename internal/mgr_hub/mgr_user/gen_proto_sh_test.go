package mgr_user

import (
	"gitee.com/cruvie/kk_go_kit/kk_protobuf"
	"testing"
)

// 此文件必须在HandlerPrePath目录中
func TestGenProtoSh(t *testing.T) {
	in := kk_protobuf.GenProtoShInput{
		HandlerPrePath:    "internal/mgr_hub/mgr_user/",
		HandlerFolderList: []string{"/Users/cruvie/cruvie-space/code-hub/my-project/kk_etcd/kk_etcd_go/internal/mgr_hub/mgr_user"},
		ProjectPath:       "../../../",
		InternalProto:     []string{
			//"google/protobuf/timestamp.proto",
		},
		TsBufYamlPath: "",
		AdditionalProto: []string{
			"kk_etcd_models/*proto",
		},
		DartAPILibPath: "/Users/cruvie/cruvie-space/code-hub/my-project/kk_etcd/kk_etcd_go/lib/",
	}
	kk_protobuf.GenProtoSh(in)
}
