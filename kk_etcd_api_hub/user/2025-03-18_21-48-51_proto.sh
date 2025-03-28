
clear
(
cd "$ProjectPath" || exit
pwd
# gen go
protoc  --proto_path=. \
        --go_out=. \
        --go_opt=paths=source_relative \
       kk_etcd_api_hub/user/getUser/*.proto \
kk_etcd_api_hub/user/login/*.proto \
kk_etcd_api_hub/user/logout/*.proto \
kk_etcd_api_hub/user/myInfo/*.proto \
kk_etcd_api_hub/user/userAdd/*.proto \
kk_etcd_api_hub/user/userDelete/*.proto \
kk_etcd_api_hub/user/userGrantRole/*.proto \
kk_etcd_api_hub/user/userList/*.proto \
kk_etcd_models/*proto


# gen dart
protoc --proto_path=. \
       --dart_out=grpc:$DartAPILibPath \
kk_etcd_api_hub/user/getUser/*.proto \
kk_etcd_api_hub/user/login/*.proto \
kk_etcd_api_hub/user/logout/*.proto \
kk_etcd_api_hub/user/myInfo/*.proto \
kk_etcd_api_hub/user/userAdd/*.proto \
kk_etcd_api_hub/user/userDelete/*.proto \
kk_etcd_api_hub/user/userGrantRole/*.proto \
kk_etcd_api_hub/user/userList/*.proto \
kk_etcd_models/*proto \



# gen typescript
#npx buf generate --template
)
