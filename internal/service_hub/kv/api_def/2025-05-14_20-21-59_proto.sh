
(
cd "$ProjectPath" || exit
pwd

common_protos="kk_etcd_api_hub/kv/api_def/*.proto \
kk_etcd_models/*proto"

echo "gen go"
eval "protoc  --proto_path=. \
              --go_out=. \
              --go_opt=paths=source_relative \
              $common_protos"

echo "gen dart"
eval "protoc --proto_path=. \
             --dart_out=grpc:$DartAPILibPath \
             $common_protos google/protobuf/timestamp.proto"

#echo "gen python"
#eval "protoc --proto_path=. \
#             --python_out=pyi_out:$PythonLibPath \
#             $common_protos google/protobuf/timestamp.proto"
#
#echo "gen ts"
#npx buf generate --template
)