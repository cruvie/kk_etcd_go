import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import 'api.pb.dart';

/// grant permission
Future<void> apiRoleGrantPermission(
  RoleGrantPermission_Input input,
  RoleGrantPermission_Output output,
  Future<PBResponse> Function(String path, GeneratedMessage requestData)
      requestFunc, {
  Function()? okFunc,
  Function()? errorFunc,
}) async {
  await kkBaseRequest(
    "/role/roleGrantPermission",
    input,
    output,
    requestFunc,
    okFunc: okFunc,
    errorFunc: errorFunc,
  );
}