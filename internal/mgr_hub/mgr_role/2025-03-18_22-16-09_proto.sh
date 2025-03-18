
# 工具准备
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# npm install -g @bufbuild/protobuf @bufbuild/protoc-gen-es @bufbuild/buf

clear
(
ProjectPath=../../../
cd "$ProjectPath" || exit
pwd
# gen go
protoc  --proto_path=. \
        --go_out=. \
        --go_opt=paths=source_relative \
       internal/mgr_hub/mgr_role/roleAdd/*.proto \
internal/mgr_hub/mgr_role/roleDelete/*.proto \
internal/mgr_hub/mgr_role/roleGet/*.proto \
internal/mgr_hub/mgr_role/roleGrantPermission/*.proto \
internal/mgr_hub/mgr_role/roleList/*.proto \
internal/mgr_hub/mgr_role/roleRevokePermission/*.proto \
kk_etcd_models/*proto


# gen dart
protoc --proto_path=. \
        --dart_out=/Users/cruvie/cruvie-space/code-hub/my-project/kk_etcd/kk_etcd_go/lib/ \
internal/mgr_hub/mgr_role/roleAdd/*.proto \
internal/mgr_hub/mgr_role/roleDelete/*.proto \
internal/mgr_hub/mgr_role/roleGet/*.proto \
internal/mgr_hub/mgr_role/roleGrantPermission/*.proto \
internal/mgr_hub/mgr_role/roleList/*.proto \
internal/mgr_hub/mgr_role/roleRevokePermission/*.proto \
kk_etcd_models/*proto \
google/protobuf/timestamp.proto


# gen typescript
#npx buf generate --template
)
