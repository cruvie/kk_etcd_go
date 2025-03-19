//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/server/deregisterServer/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use deregisterServerDescriptor instead')
const DeregisterServer$json = {
  '1': 'DeregisterServer',
  '3': [DeregisterServer_Input$json, DeregisterServer_Output$json],
};

@$core.Deprecated('Use deregisterServerDescriptor instead')
const DeregisterServer_Input$json = {
  '1': 'Input',
  '2': [
    {
      '1': 'Server',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBServer',
      '10': 'Server'
    },
  ],
};

@$core.Deprecated('Use deregisterServerDescriptor instead')
const DeregisterServer_Output$json = {
  '1': 'Output',
};

/// Descriptor for `DeregisterServer`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deregisterServerDescriptor = $convert.base64Decode(
    'ChBEZXJlZ2lzdGVyU2VydmVyGjkKBUlucHV0EjAKBlNlcnZlchgBIAEoCzIYLmtrX2V0Y2RfbW'
    '9kZWxzLlBCU2VydmVyUgZTZXJ2ZXIaCAoGT3V0cHV0');

