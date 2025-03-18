//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_role/roleList/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use roleListDescriptor instead')
const RoleList$json = {
  '1': 'RoleList',
  '3': [RoleList_Input$json, RoleList_Output$json],
};

@$core.Deprecated('Use roleListDescriptor instead')
const RoleList_Input$json = {
  '1': 'Input',
};

@$core.Deprecated('Use roleListDescriptor instead')
const RoleList_Output$json = {
  '1': 'Output',
  '2': [
    {
      '1': 'ListRole',
      '3': 1,
      '4': 1,
      '5': 11,
      '6': '.kk_etcd_models.PBListRole',
      '10': 'ListRole'
    },
  ],
};

/// Descriptor for `RoleList`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List roleListDescriptor = $convert.base64Decode(
    'CghSb2xlTGlzdBoHCgVJbnB1dBpACgZPdXRwdXQSNgoITGlzdFJvbGUYASABKAsyGi5ra19ldG'
    'NkX21vZGVscy5QQkxpc3RSb2xlUghMaXN0Um9sZQ==');
