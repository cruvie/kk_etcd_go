

clear
(
cd "$ProjectPath" || exit
pwd
# gen go
protoc  --proto_path=. \
        --go_out=. \
        --go_opt=paths=source_relative \
       kk_etcd_api_hub/server/deregisterServer/*.proto \
kk_etcd_api_hub/server/serverList/*.proto \
kk_etcd_models/*proto


# gen dart
protoc --proto_path=. \
       --dart_out=grpc:$DartAPILibPath \
kk_etcd_api_hub/server/deregisterServer/*.proto \
kk_etcd_api_hub/server/serverList/*.proto \
kk_etcd_models/*proto \



# gen typescript
#npx buf generate --template
)
