
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
       internal/mgr_hub/mgr_kv/kVDel/*.proto \
internal/mgr_hub/mgr_kv/kVGet/*.proto \
internal/mgr_hub/mgr_kv/kVList/*.proto \
internal/mgr_hub/mgr_kv/kVPut/*.proto \
kk_etcd_models/*proto


# gen dart
protoc --proto_path=. \
        --dart_out=/Users/cruvie/cruvie-space/code-hub/my-project/kk_etcd/kk_etcd_go/lib/ \
internal/mgr_hub/mgr_kv/kVDel/*.proto \
internal/mgr_hub/mgr_kv/kVGet/*.proto \
internal/mgr_hub/mgr_kv/kVList/*.proto \
internal/mgr_hub/mgr_kv/kVPut/*.proto \
kk_etcd_models/*proto \



# gen typescript
#npx buf generate --template
)
