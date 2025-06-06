pwd

# 生成go文件命令
protoc  --proto_path=. \
        --go_out=. \
        --go_opt=paths=source_relative \
         kk_etcd_models/*.proto

# 生成dart文件命令
protoc --proto_path=. \
        --dart_out="$DartAPILibPath" \
        kk_etcd_models/*proto \
        google/protobuf/duration.proto
#google/protobuf/timestamp.proto
#google/protobuf/any.proto
