

clear
(
cd "$ProjectPath" || exit
pwd
# gen go
protoc  --proto_path=. \
        --go_out=. \
        --go_opt=paths=source_relative \
       kk_etcd_api_hub/kv/kVDel/*.proto \
kk_etcd_api_hub/kv/kVGet/*.proto \
kk_etcd_api_hub/kv/kVList/*.proto \
kk_etcd_api_hub/kv/kVPut/*.proto \
kk_etcd_models/*proto


# gen dart
protoc --proto_path=. \
       --dart_out=grpc:$DartAPILibPath \
kk_etcd_api_hub/kv/kVDel/*.proto \
kk_etcd_api_hub/kv/kVGet/*.proto \
kk_etcd_api_hub/kv/kVList/*.proto \
kk_etcd_api_hub/kv/kVPut/*.proto \
kk_etcd_models/*proto \



# gen typescript
#npx buf generate --template
)
