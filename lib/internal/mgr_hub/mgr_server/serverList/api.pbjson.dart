//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_server/serverList/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use serverListDescriptor instead')
const ServerList$json = {
  '1': 'ServerList',
  '3': [ServerList_Input$json, ServerList_Output$json],
};

@$core.Deprecated('Use serverListDescriptor instead')
const ServerList_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'ServerType', '3': 1, '4': 1, '5': 9, '10': 'ServerType'},
    {'1': 'ServerName', '3': 2, '4': 1, '5': 9, '10': 'ServerName'},
  ],
};

@$core.Deprecated('Use serverListDescriptor instead')
const ServerList_Output$json = {
  '1': 'Output',
  '2': [
    {
      '1': 'ServerList',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBListServer',
      '10': 'ServerList'
    },
  ],
};

/// Descriptor for `ServerList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List serverListDescriptor = $convert.base64Decode(
    'CgpTZXJ2ZXJMaXN0GkcKBUlucHV0Eh4KClNlcnZlclR5cGUYASABKAlSClNlcnZlclR5cGUSHg'
    'oKU2VydmVyTmFtZRgCIAEoCVIKU2VydmVyTmFtZRpGCgZPdXRwdXQSPAoKU2VydmVyTGlzdBgB'
    'IAEoCzIcLmtrX2V0Y2RfbW9kZWxzLlBCTGlzdFNlcnZlclIKU2VydmVyTGlzdA==');
