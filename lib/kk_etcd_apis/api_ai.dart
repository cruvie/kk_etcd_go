import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import '../kk_etcd_models/api_ai_kk_etcd.pb.dart';

class ApiAI {
  /// query ai
  static Future<void> query(
      QueryParam param,
      QueryResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/ai/query", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }
}
