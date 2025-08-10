// This is a generated file - do not edit.
//
// Generated from internal/api_hub/backup/api_def/SnapshotRestore.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use snapshotRestoreDescriptor instead')
const SnapshotRestore$json = {
  '1': 'SnapshotRestore',
  '3': [SnapshotRestore_Input$json, SnapshotRestore_Output$json],
};

@$core.Deprecated('Use snapshotRestoreDescriptor instead')
const SnapshotRestore_Input$json = {
  '1': 'Input',
};

@$core.Deprecated('Use snapshotRestoreDescriptor instead')
const SnapshotRestore_Output$json = {
  '1': 'Output',
  '2': [
    {'1': 'CmdStr', '3': 1, '4': 1, '5': 9, '10': 'CmdStr'},
  ],
};

/// Descriptor for `SnapshotRestore`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotRestoreDescriptor = $convert.base64Decode(
    'Cg9TbmFwc2hvdFJlc3RvcmUaBwoFSW5wdXQaIAoGT3V0cHV0EhYKBkNtZFN0chgBIAEoCVIGQ2'
    '1kU3Ry');
