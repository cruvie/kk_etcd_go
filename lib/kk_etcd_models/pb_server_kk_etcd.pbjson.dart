//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_server_kk_etcd.proto
//
// @dart = 3.3

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
    {'1': 'ServerRegistration', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBServerRegistration', '10': 'ServerRegistration'},
    {'1': 'Status', '3': 5, '4': 1, '5': 14, '6': '.kk_etcd_models.PBServer.ServerStatus', '10': 'Status'},
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
  ],
};

/// Descriptor for `PBServer`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pBServerDescriptor = $convert.base64Decode(
    'CghQQlNlcnZlchJUChJTZXJ2ZXJSZWdpc3RyYXRpb24YASABKAsyJC5ra19ldGNkX21vZGVscy'
    '5QQlNlcnZlclJlZ2lzdHJhdGlvblISU2VydmVyUmVnaXN0cmF0aW9uEj0KBlN0YXR1cxgFIAEo'
    'DjIlLmtrX2V0Y2RfbW9kZWxzLlBCU2VydmVyLlNlcnZlclN0YXR1c1IGU3RhdHVzIjIKDFNlcn'
    'ZlclN0YXR1cxILCgdVbktub3duEAASCwoHUnVubmluZxABEggKBFN0b3AQAg==');

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

