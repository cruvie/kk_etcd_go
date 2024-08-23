import 'dart:typed_data';

import 'package:kk_go_kit/kk_models/pb_response.pb.dart';
import 'package:kk_go_kit/kk_util/kk_log.dart';

import '../kk_etcd_models/api_kv_kk_etcd.pb.dart';

class ApiRole {
  /// put kv
  static Future<void> kVPut(
      KVPutParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(KVPutResponse)? okFunc,
      Function(KVPutResponse)? errorFunc}) async {
    KVPutResponse response = KVPutResponse();
    PBResponse res = await requestFunc("/kv/kVPut", param.writeToBuffer());
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

  /// get kv
  static Future<void> kVGet(
      KVGetParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(KVGetResponse)? okFunc,
      Function(KVGetResponse)? errorFunc}) async {
    KVGetResponse response = KVGetResponse();
    PBResponse res = await requestFunc("/kv/kVGet", param.writeToBuffer());
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

  /// del kv
  static Future<void> kVDel(
      KVDelParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(KVDelResponse)? okFunc,
      Function(KVDelResponse)? errorFunc}) async {
    KVDelResponse response = KVDelResponse();
    PBResponse res = await requestFunc("/kv/kVDel", param.writeToBuffer());
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

  /// list kv
  static Future<void> kVList(
      KVListParam param,
      Future<PBResponse> Function(String path, Uint8List requestData)
          requestFunc,
      {Function(KVListResponse)? okFunc,
      Function(KVListResponse)? errorFunc}) async {
    KVListResponse response = KVListResponse();
    PBResponse res = await requestFunc("/kv/kVList", param.writeToBuffer());
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
