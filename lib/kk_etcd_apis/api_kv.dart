import 'package:kk_go_kit/kk_http/base_request.dart';
import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:protobuf/protobuf.dart';

import '../kk_etcd_models/api_kv_kk_etcd.pb.dart';

class ApiKV {
  /// put kv
  static Future<void> kVPut(
      KVPutParam param,
      KVPutResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/kv/kVPut", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// get kv
  static Future<void> kVGet(
      KVGetParam param,
      KVGetResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/kv/kVGet", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// del kv
  static Future<void> kVDel(
      KVDelParam param,
      KVDelResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/kv/kVDel", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }

  /// list kv
  static Future<void> kVList(
      KVListParam param,
      KVListResponse response,
      Future<PBResponse> Function(String path, GeneratedMessage requestData)
          requestFunc,
      {Function()? okFunc,
      Function()? errorFunc}) async {
    await kkBaseRequest("/kv/kVList", param, response, requestFunc,
        okFunc: okFunc, errorFunc: errorFunc);
  }
}
