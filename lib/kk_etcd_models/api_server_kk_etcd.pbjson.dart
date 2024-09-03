//
//  Generated code. Do not modify.
//  source: api_server_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use serverListParamDescriptor instead')
const ServerListParam$json = {
  '1': 'ServerListParam',
  '2': [
    {'1': 'ServerType', '3': 1, '4': 1, '5': 9, '10': 'ServerType'},
    {'1': 'ServerName', '3': 2, '4': 1, '5': 9, '10': 'ServerName'},
  ],
};

/// Descriptor for `ServerListParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List serverListParamDescriptor = $convert.base64Decode(
    'Cg9TZXJ2ZXJMaXN0UGFyYW0SHgoKU2VydmVyVHlwZRgBIAEoCVIKU2VydmVyVHlwZRIeCgpTZX'
    'J2ZXJOYW1lGAIgASgJUgpTZXJ2ZXJOYW1l');

@$core.Deprecated('Use serverListResponseDescriptor instead')
const ServerListResponse$json = {
  '1': 'ServerListResponse',
  '2': [
    {'1': 'ServerList', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBListServer', '10': 'ServerList'},
  ],
};

/// Descriptor for `ServerListResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List serverListResponseDescriptor = $convert.base64Decode(
    'ChJTZXJ2ZXJMaXN0UmVzcG9uc2USPAoKU2VydmVyTGlzdBgBIAEoCzIcLmtrX2V0Y2RfbW9kZW'
    'xzLlBCTGlzdFNlcnZlclIKU2VydmVyTGlzdA==');

@$core.Deprecated('Use deregisterServerParamDescriptor instead')
const DeregisterServerParam$json = {
  '1': 'DeregisterServerParam',
  '2': [
    {'1': 'Server', '3': 1, '4': 1, '5': 11, '6': '.kk_etcd_models.PBServer', '10': 'Server'},
  ],
};

/// Descriptor for `DeregisterServerParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deregisterServerParamDescriptor = $convert.base64Decode(
    'ChVEZXJlZ2lzdGVyU2VydmVyUGFyYW0SMAoGU2VydmVyGAEgASgLMhgua2tfZXRjZF9tb2RlbH'
    'MuUEJTZXJ2ZXJSBlNlcnZlcg==');

@$core.Deprecated('Use deregisterServerResponseDescriptor instead')
const DeregisterServerResponse$json = {
  '1': 'DeregisterServerResponse',
};

/// Descriptor for `DeregisterServerResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deregisterServerResponseDescriptor = $convert.base64Decode(
    'ChhEZXJlZ2lzdGVyU2VydmVyUmVzcG9uc2U=');

