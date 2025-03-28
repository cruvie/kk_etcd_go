

clear
(
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
       --dart_out=grpc:$DartAPILibPath \
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
