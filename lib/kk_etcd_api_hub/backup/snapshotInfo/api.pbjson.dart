//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/backup/snapshotInfo/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

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

