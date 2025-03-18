import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import 'api.pb.dart';

/// snapshot info
Future<void> apiSnapshotInfo(
  SnapshotInfo_Input input,
  SnapshotInfo_Output output,
  Future<PBResponse> Function(String path, GeneratedMessage requestData)
      requestFunc, {
  Function()? okFunc,
  Function()? errorFunc,
}) async {
  await kkBaseRequest(
    "/backup/snapshotInfo",
    input,
    output,
    requestFunc,
    okFunc: okFunc,
    errorFunc: errorFunc,
  );
}
