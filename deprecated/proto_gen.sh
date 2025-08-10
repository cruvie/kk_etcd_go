
clear
#go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

#export PUB_HOSTED_URL=https://pub.dev
#dart pub global activate protobuf
#dart pub global activate protoc_plugin

export ProjectPath=/Users/cruvie/cruvie-space/code-hub/my-project/kk_etcd/kk_etcd_go
export DartAPILibPath=/Users/cruvie/cruvie-space/code-hub/my-project/kk_etcd/kk_etcd_go/lib/

cd ../
(
echo "gen kk_etcd_models pb"

sh kk_etcd_models/proto_gen.sh

)


(
echo "gen kk_etcd_api_hub pb"

cd kk_etcd_api_hub || exit

sh ai/api_def/*_proto.sh

sh backup/api_def/*_proto.sh

sh kv/api_def/*_proto.sh

sh role/api_def/*_proto.sh

sh service/api_def/*_proto.sh

sh user/api_def/*_proto.sh

)
