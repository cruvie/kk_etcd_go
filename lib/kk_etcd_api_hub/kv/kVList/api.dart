import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import 'api.pb.dart';

/// list kv
Future<void> apiKVList(
  KVList_Input input,
  KVList_Output output,
  Future<PBResponse> Function(String path, GeneratedMessage requestData)
      requestFunc, {
  Function()? okFunc,
  Function()? errorFunc,
}) async {
  await kkBaseRequest(
    "/kv/kVList",
    input,
    output,
    requestFunc,
    okFunc: okFunc,
    errorFunc: errorFunc,
  );
}