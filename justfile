# brew install just

set shell := ["zsh", "-ic"]

proto-gen:
    buf generate

proto-util-install:
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    export PUB_HOSTED_URL=https://pub.dev
    dart pub global activate protobuf
    dart pub global activate protoc_plugin
    npm run buf-install
