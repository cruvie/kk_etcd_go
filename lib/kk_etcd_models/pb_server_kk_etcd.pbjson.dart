//
//  Generated code. Do not modify.
//  source: pb_server_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use pBServerDescriptor instead')
const PBServer$json = {
  '1': 'PBServer',
  '2': [
    {'1': 'ServerName', '3': 1, '4': 1, '5': 9, '10': 'ServerName'},
    {'1': 'ServerAddr', '3': 2, '4': 1, '5': 9, '10': 'ServerAddr'},
    {'1': 'Status', '3': 3, '4': 1, '5': 14, '6': '.kk_etcd_models.PBServer.ServerStatus', '10': 'Status'},
    {'1': 'LastCheck', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'LastCheck'},
    {'1': 'Msg', '3': 5, '4': 1, '5': 9, '10': 'Msg'},
  ],
  '4': [PBServer_ServerStatus$json],
};

@$core.Deprecated('Use pBServerDescriptor instead')
const PBServer_ServerStatus$json = {
  '1': 'ServerStatus',
  '2': [
    {'1': 'UnKnown', '2': 0},
    {'1': 'Running', '2': 1},
    {'1': 'Stop', '2': 2},
    {'1': 'Init', '2': 3},
  ],
};

/// Descriptor for `PBServer`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBServerDescriptor = $convert.base64Decode(
    'CghQQlNlcnZlchIeCgpTZXJ2ZXJOYW1lGAEgASgJUgpTZXJ2ZXJOYW1lEh4KClNlcnZlckFkZH'
    'IYAiABKAlSClNlcnZlckFkZHISPQoGU3RhdHVzGAMgASgOMiUua2tfZXRjZF9tb2RlbHMuUEJT'
    'ZXJ2ZXIuU2VydmVyU3RhdHVzUgZTdGF0dXMSOAoJTGFzdENoZWNrGAQgASgLMhouZ29vZ2xlLn'
    'Byb3RvYnVmLlRpbWVzdGFtcFIJTGFzdENoZWNrEhAKA01zZxgFIAEoCVIDTXNnIjwKDFNlcnZl'
    'clN0YXR1cxILCgdVbktub3duEAASCwoHUnVubmluZxABEggKBFN0b3AQAhIICgRJbml0EAM=');

@$core.Deprecated('Use pBListServerDescriptor instead')
const PBListServer$json = {
  '1': 'PBListServer',
  '2': [
    {'1': 'ListServer', '3': 1, '4': 3, '5': 11, '6': '.kk_etcd_models.PBServer', '10': 'ListServer'},
  ],
};

/// Descriptor for `PBListServer`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBListServerDescriptor = $convert.base64Decode(
    'CgxQQkxpc3RTZXJ2ZXISOAoKTGlzdFNlcnZlchgBIAMoCzIYLmtrX2V0Y2RfbW9kZWxzLlBCU2'
    'VydmVyUgpMaXN0U2VydmVy');

