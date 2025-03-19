
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
        --dart_out=/Users/cruvie/cruvie-space/code-hub/my-project/kk_etcd/kk_etcd_go/lib/ \
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
