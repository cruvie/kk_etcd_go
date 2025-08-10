// This is a generated file - do not edit.
//
// Generated from kk_etcd_api_hub/backup/api_def/AllKVsBackup.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use allKVsBackupDescriptor instead')
const AllKVsBackup$json = {
  '1': 'AllKVsBackup',
  '3': [AllKVsBackup_Input$json, AllKVsBackup_Output$json],
};

@$core.Deprecated('Use allKVsBackupDescriptor instead')
const AllKVsBackup_Input$json = {
  '1': 'Input',
};

@$core.Deprecated('Use allKVsBackupDescriptor instead')
const AllKVsBackup_Output$json = {
  '1': 'Output',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    {'1': 'File', '3': 2, '4': 1, '5': 12, '10': 'File'},
  ],
};

/// Descriptor for `AllKVsBackup`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List allKVsBackupDescriptor = $convert.base64Decode(
    'CgxBbGxLVnNCYWNrdXAaBwoFSW5wdXQaMAoGT3V0cHV0EhIKBE5hbWUYASABKAlSBE5hbWUSEg'
    'oERmlsZRgCIAEoDFIERmlsZQ==');
