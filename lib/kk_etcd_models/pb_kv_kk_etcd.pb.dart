// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_kv_kk_etcd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class PBKV extends $pb.GeneratedMessage {
  factory PBKV({
    $core.String? key,
    $core.String? value,
  }) {
    final result = create();
    if (key != null) result.key = key;
    if (value != null) result.value = value;
    return result;
  }

  PBKV._();

  factory PBKV.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBKV.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBKV',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..aOS(2, _omitFieldNames ? '' : 'Value', protoName: 'Value')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBKV clone() => PBKV()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBKV copyWith(void Function(PBKV) updates) =>
      super.copyWith((message) => updates(message as PBKV)) as PBKV;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBKV create() => PBKV._();
  @$core.override
  PBKV createEmptyInstance() => create();
  static $pb.PbList<PBKV> createRepeated() => $pb.PbList<PBKV>();
  @$core.pragma('dart2js:noInline')
  static PBKV getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBKV>(create);
  static PBKV? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get value => $_getSZ(1);
  @$pb.TagNumber(2)
  set value($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => $_clearField(2);
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
