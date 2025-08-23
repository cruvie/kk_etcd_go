package api_def

import (
	"testing"

	"gitee.com/cruvie/kk_go_kit/kk_protobuf"
)

func TestGenProtoSh(t *testing.T) {
	in := kk_protobuf.GenProtoShInput{
		HandlerPrePath:    "service/",
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
