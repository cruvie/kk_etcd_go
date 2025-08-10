// This is a generated file - do not edit.
//
// Generated from kk_etcd_api_hub/backup/api_def/AllKVsRestore.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use allKVsRestoreDescriptor instead')
const AllKVsRestore$json = {
  '1': 'AllKVsRestore',
  '3': [AllKVsRestore_Input$json, AllKVsRestore_Output$json],
};

@$core.Deprecated('Use allKVsRestoreDescriptor instead')
const AllKVsRestore_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'File', '3': 2, '4': 1, '5': 12, '10': 'File'},
  ],
};

@$core.Deprecated('Use allKVsRestoreDescriptor instead')
const AllKVsRestore_Output$json = {
  '1': 'Output',
};

/// Descriptor for `AllKVsRestore`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List allKVsRestoreDescriptor = $convert.base64Decode(
    'Cg1BbGxLVnNSZXN0b3JlGhsKBUlucHV0EhIKBEZpbGUYAiABKAxSBEZpbGUaCAoGT3V0cHV0');
