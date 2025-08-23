// This is a generated file - do not edit.
//
// Generated from internal/service_hub/backup/api_def/SnapshotInfo.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use snapshotInfoDescriptor instead')
const SnapshotInfo$json = {
  '1': 'SnapshotInfo',
  '3': [SnapshotInfo_Input$json, SnapshotInfo_Output$json],
};

@$core.Deprecated('Use snapshotInfoDescriptor instead')
const SnapshotInfo_Input$json = {
  '1': 'Input',
  '2': [
    {'1': 'File', '3': 2, '4': 1, '5': 12, '10': 'File'},
  ],
};

@$core.Deprecated('Use snapshotInfoDescriptor instead')
const SnapshotInfo_Output$json = {
  '1': 'Output',
  '2': [
    {'1': 'Info', '3': 1, '4': 1, '5': 9, '10': 'Info'},
  ],
};

/// Descriptor for `SnapshotInfo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotInfoDescriptor = $convert.base64Decode(
    'CgxTbmFwc2hvdEluZm8aGwoFSW5wdXQSEgoERmlsZRgCIAEoDFIERmlsZRocCgZPdXRwdXQSEg'
    'oESW5mbxgBIAEoCVIESW5mbw==');
