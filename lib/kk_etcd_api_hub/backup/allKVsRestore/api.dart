import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import 'api.pb.dart';

/// all kvs restore
Future<void> apiAllKVsRestore(
  AllKVsRestore_Input input,
  AllKVsRestore_Output output,
  Future<PBResponse> Function(String path, GeneratedMessage requestData)
      requestFunc, {
  Function()? okFunc,
  Function()? errorFunc,
}) async {
  await kkBaseRequest(
    "/backup/allKVsRestore",
    input,
    output,
    requestFunc,
    okFunc: okFunc,
    errorFunc: errorFunc,
  );
}