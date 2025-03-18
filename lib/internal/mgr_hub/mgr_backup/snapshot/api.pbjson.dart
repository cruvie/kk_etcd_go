//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_backup/snapshot/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use snapshotDescriptor instead')
const Snapshot$json = {
  '1': 'Snapshot',
  '3': [Snapshot_Input$json, Snapshot_Output$json],
};

@$core.Deprecated('Use snapshotDescriptor instead')
const Snapshot_Input$json = {
  '1': 'Input',
};

@$core.Deprecated('Use snapshotDescriptor instead')
const Snapshot_Output$json = {
  '1': 'Output',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    {'1': 'File', '3': 2, '4': 1, '5': 12, '10': 'File'},
  ],
};

/// Descriptor for `Snapshot`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotDescriptor = $convert.base64Decode(
    'CghTbmFwc2hvdBoHCgVJbnB1dBowCgZPdXRwdXQSEgoETmFtZRgBIAEoCVIETmFtZRISCgRGaW'
    'xlGAIgASgMUgRGaWxl');
