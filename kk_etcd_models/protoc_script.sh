#!/bin/bash
# 生成go文件命令
protoc -I=. --go_out=./ *.proto

# 生成dart文件命令
protoc -I=. --dart_out=../lib/models *.proto google/protobuf/timestamp.proto
#google/protobuf/any.proto
