//
//  Generated code. Do not modify.
//  source: api_manage_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use snapshotParamDescriptor instead')
const SnapshotParam$json = {
  '1': 'SnapshotParam',
};

/// Descriptor for `SnapshotParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotParamDescriptor = $convert.base64Decode(
    'Cg1TbmFwc2hvdFBhcmFt');

@$core.Deprecated('Use snapshotResponseDescriptor instead')
const SnapshotResponse$json = {
  '1': 'SnapshotResponse',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    {'1': 'File', '3': 2, '4': 1, '5': 12, '10': 'File'},
  ],
};

/// Descriptor for `SnapshotResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotResponseDescriptor = $convert.base64Decode(
    'ChBTbmFwc2hvdFJlc3BvbnNlEhIKBE5hbWUYASABKAlSBE5hbWUSEgoERmlsZRgCIAEoDFIERm'
    'lsZQ==');

@$core.Deprecated('Use snapshotRestoreParamDescriptor instead')
const SnapshotRestoreParam$json = {
  '1': 'SnapshotRestoreParam',
};

/// Descriptor for `SnapshotRestoreParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotRestoreParamDescriptor = $convert.base64Decode(
    'ChRTbmFwc2hvdFJlc3RvcmVQYXJhbQ==');

@$core.Deprecated('Use snapshotRestoreResponseDescriptor instead')
const SnapshotRestoreResponse$json = {
  '1': 'SnapshotRestoreResponse',
  '2': [
    {'1': 'CmdStr', '3': 1, '4': 1, '5': 9, '10': 'CmdStr'},
  ],
};

/// Descriptor for `SnapshotRestoreResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotRestoreResponseDescriptor = $convert.base64Decode(
    'ChdTbmFwc2hvdFJlc3RvcmVSZXNwb25zZRIWCgZDbWRTdHIYASABKAlSBkNtZFN0cg==');

@$core.Deprecated('Use snapshotInfoParamDescriptor instead')
const SnapshotInfoParam$json = {
  '1': 'SnapshotInfoParam',
  '2': [
    {'1': 'File', '3': 2, '4': 1, '5': 12, '10': 'File'},
  ],
};

/// Descriptor for `SnapshotInfoParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotInfoParamDescriptor = $convert.base64Decode(
    'ChFTbmFwc2hvdEluZm9QYXJhbRISCgRGaWxlGAIgASgMUgRGaWxl');

@$core.Deprecated('Use snapshotInfoResponseDescriptor instead')
const SnapshotInfoResponse$json = {
  '1': 'SnapshotInfoResponse',
  '2': [
    {'1': 'Info', '3': 1, '4': 1, '5': 9, '10': 'Info'},
  ],
};

/// Descriptor for `SnapshotInfoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List snapshotInfoResponseDescriptor = $convert.base64Decode(
    'ChRTbmFwc2hvdEluZm9SZXNwb25zZRISCgRJbmZvGAEgASgJUgRJbmZv');

@$core.Deprecated('Use allKVsBackupParamDescriptor instead')
const AllKVsBackupParam$json = {
  '1': 'AllKVsBackupParam',
};

/// Descriptor for `AllKVsBackupParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List allKVsBackupParamDescriptor = $convert.base64Decode(
    'ChFBbGxLVnNCYWNrdXBQYXJhbQ==');

@$core.Deprecated('Use allKVsBackupResponseDescriptor instead')
const AllKVsBackupResponse$json = {
  '1': 'AllKVsBackupResponse',
  '2': [
    {'1': 'Name', '3': 1, '4': 1, '5': 9, '10': 'Name'},
    {'1': 'File', '3': 2, '4': 1, '5': 12, '10': 'File'},
  ],
};

/// Descriptor for `AllKVsBackupResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List allKVsBackupResponseDescriptor = $convert.base64Decode(
    'ChRBbGxLVnNCYWNrdXBSZXNwb25zZRISCgROYW1lGAEgASgJUgROYW1lEhIKBEZpbGUYAiABKA'
    'xSBEZpbGU=');

@$core.Deprecated('Use allKVsRestoreParamDescriptor instead')
const AllKVsRestoreParam$json = {
  '1': 'AllKVsRestoreParam',
  '2': [
    {'1': 'File', '3': 2, '4': 1, '5': 12, '10': 'File'},
  ],
};

/// Descriptor for `AllKVsRestoreParam`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List allKVsRestoreParamDescriptor = $convert.base64Decode(
    'ChJBbGxLVnNSZXN0b3JlUGFyYW0SEgoERmlsZRgCIAEoDFIERmlsZQ==');

@$core.Deprecated('Use allKVsRestoreResponseDescriptor instead')
const AllKVsRestoreResponse$json = {
  '1': 'AllKVsRestoreResponse',
};

/// Descriptor for `AllKVsRestoreResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List allKVsRestoreResponseDescriptor = $convert.base64Decode(
    'ChVBbGxLVnNSZXN0b3JlUmVzcG9uc2U=');

