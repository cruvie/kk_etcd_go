#!/bin/bash
#go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
#dart pub global activate protoc_plugin
# 生成go文件命令
protoc -I=. --go_out=./ *.proto

# 生成dart文件命令
protoc -I=. --dart_out=../lib/kk_etcd_models *.proto google/protobuf/timestamp.proto
#google/protobuf/any.proto
