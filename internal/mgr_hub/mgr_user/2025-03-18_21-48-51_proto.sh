
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
       internal/mgr_hub/mgr_user/getUser/*.proto \
internal/mgr_hub/mgr_user/login/*.proto \
internal/mgr_hub/mgr_user/logout/*.proto \
internal/mgr_hub/mgr_user/myInfo/*.proto \
internal/mgr_hub/mgr_user/userAdd/*.proto \
internal/mgr_hub/mgr_user/userDelete/*.proto \
internal/mgr_hub/mgr_user/userGrantRole/*.proto \
internal/mgr_hub/mgr_user/userList/*.proto \
kk_etcd_models/*proto


# gen dart
protoc --proto_path=. \
        --dart_out=/Users/cruvie/cruvie-space/code-hub/my-project/kk_etcd/kk_etcd_go/lib/ \
internal/mgr_hub/mgr_user/getUser/*.proto \
internal/mgr_hub/mgr_user/login/*.proto \
internal/mgr_hub/mgr_user/logout/*.proto \
internal/mgr_hub/mgr_user/myInfo/*.proto \
internal/mgr_hub/mgr_user/userAdd/*.proto \
internal/mgr_hub/mgr_user/userDelete/*.proto \
internal/mgr_hub/mgr_user/userGrantRole/*.proto \
internal/mgr_hub/mgr_user/userList/*.proto \
kk_etcd_models/*proto \



# gen typescript
#npx buf generate --template
)
