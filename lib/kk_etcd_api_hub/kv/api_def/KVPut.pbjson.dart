//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/kv/api_def/KVPut.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use kVPutDescriptor instead')
const KVPut$json = {
  '1': 'KVPut',
  '3': [KVPut_Input$json, KVPut_Output$json],
};

@$core.Deprecated('Use kVPutDescriptor instead')
const KVPut_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'Key', '3': 1, '4': 1, '5': 9, '10': 'Key'},
    {'1': 'Value', '3': 2, '4': 1, '5': 9, '10': 'Value'},
  ],
};

@$core.Deprecated('Use kVPutDescriptor instead')
const KVPut_Output$json = {
  '1': 'Output',
};

/// Descriptor for `KVPut`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List kVPutDescriptor = $convert.base64Decode(
    'CgVLVlB1dBovCgVJbnB1dBIQCgNLZXkYASABKAlSA0tleRIUCgVWYWx1ZRgCIAEoCVIFVmFsdW'
    'UaCAoGT3V0cHV0');

