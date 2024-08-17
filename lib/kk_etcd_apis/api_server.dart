import 'dart:typed_data';

import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:kk_go_kit/kk_util/kk_log.dart';

import '../kk_etcd_models/api_server_kk_etcd.pb.dart';

class ApiServer {
  /// list server
  static Future<void> serverList(
      ServerListParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(ServerListResponse)? okFunc,
      Function(ServerListResponse)? errorFunc}) async {
    ServerListResponse response = ServerListResponse();
    PBResponse res =
        await requestFunc("/server/serverList", param.writeToBuffer());
    try {
      if (res.code == 200) {
        res.data.unpackInto(response);
        if (okFunc != null) {
          await okFunc(response);
        }
      } else if (errorFunc != null) {
        await errorFunc(response);
      }
    } catch (e) {
      log.info(e);
    }
  }
}
