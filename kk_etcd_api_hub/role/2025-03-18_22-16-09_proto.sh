
# 工具准备
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# npm install -g @bufbuild/protobuf @bufbuild/protoc-gen-es @bufbuild/buf

clear
(
ProjectPath=../../
cd "$ProjectPath" || exit
pwd
# gen go
protoc  --proto_path=. \
        --go_out=. \
        --go_opt=paths=source_relative \
       kk_etcd_api_hub/role/roleAdd/*.proto \
kk_etcd_api_hub/role/roleDelete/*.proto \
kk_etcd_api_hub/role/roleGet/*.proto \
kk_etcd_api_hub/role/roleGrantPermission/*.proto \
kk_etcd_api_hub/role/roleList/*.proto \
kk_etcd_api_hub/role/roleRevokePermission/*.proto \
kk_etcd_models/*proto


# gen dart
protoc --proto_path=. \
        --dart_out=/Users/cruvie/cruvie-space/code-hub/my-project/kk_etcd/kk_etcd_go/lib/ \
kk_etcd_api_hub/role/roleAdd/*.proto \
kk_etcd_api_hub/role/roleDelete/*.proto \
kk_etcd_api_hub/role/roleGet/*.proto \
kk_etcd_api_hub/role/roleGrantPermission/*.proto \
kk_etcd_api_hub/role/roleList/*.proto \
kk_etcd_api_hub/role/roleRevokePermission/*.proto \
kk_etcd_models/*proto \
google/protobuf/timestamp.proto


# gen typescript
#npx buf generate --template
)
