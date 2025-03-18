//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_role/roleRevokePermission/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class RoleRevokePermission_Input extends $pb.GeneratedMessage {
  factory RoleRevokePermission_Input({
    $core.String? name,
    $core.String? key,
    $core.String? rangeEnd,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    if (key != null) {
      $result.key = key;
    }
    if (rangeEnd != null) {
      $result.rangeEnd = rangeEnd;
    }
    return $result;
  }

  RoleRevokePermission_Input._() : super();

  factory RoleRevokePermission_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory RoleRevokePermission_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleRevokePermission.Input',
      package: const $pb.PackageName(
          _omitMessageNames ? '' : 'roleRevokePermission'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..aOS(2, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..aOS(3, _omitFieldNames ? '' : 'RangeEnd', protoName: 'RangeEnd')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  RoleRevokePermission_Input clone() =>
      RoleRevokePermission_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  RoleRevokePermission_Input copyWith(
          void Function(RoleRevokePermission_Input) updates) =>
      super.copyWith(
              (message) => updates(message as RoleRevokePermission_Input))
          as RoleRevokePermission_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleRevokePermission_Input create() => RoleRevokePermission_Input._();

  RoleRevokePermission_Input createEmptyInstance() => create();

  static $pb.PbList<RoleRevokePermission_Input> createRepeated() =>
      $pb.PbList<RoleRevokePermission_Input>();

  @$core.pragma('dart2js:noInline')
  static RoleRevokePermission_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleRevokePermission_Input>(create);
  static RoleRevokePermission_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);

  @$pb.TagNumber(1)
  set name($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);

  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get key => $_getSZ(1);

  @$pb.TagNumber(2)
  set key($core.String v) {
    $_setString(1, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasKey() => $_has(1);

  @$pb.TagNumber(2)
  void clearKey() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get rangeEnd => $_getSZ(2);

  @$pb.TagNumber(3)
  set rangeEnd($core.String v) {
    $_setString(2, v);
  }

  @$pb.TagNumber(3)
  $core.bool hasRangeEnd() => $_has(2);

  @$pb.TagNumber(3)
  void clearRangeEnd() => clearField(3);
}

class RoleRevokePermission_Output extends $pb.GeneratedMessage {
  factory RoleRevokePermission_Output() => create();

  RoleRevokePermission_Output._() : super();

  factory RoleRevokePermission_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory RoleRevokePermission_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleRevokePermission.Output',
      package: const $pb.PackageName(
          _omitMessageNames ? '' : 'roleRevokePermission'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  RoleRevokePermission_Output clone() =>
      RoleRevokePermission_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  RoleRevokePermission_Output copyWith(
          void Function(RoleRevokePermission_Output) updates) =>
      super.copyWith(
              (message) => updates(message as RoleRevokePermission_Output))
          as RoleRevokePermission_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleRevokePermission_Output create() =>
      RoleRevokePermission_Output._();

  RoleRevokePermission_Output createEmptyInstance() => create();

  static $pb.PbList<RoleRevokePermission_Output> createRepeated() =>
      $pb.PbList<RoleRevokePermission_Output>();

  @$core.pragma('dart2js:noInline')
  static RoleRevokePermission_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleRevokePermission_Output>(create);
  static RoleRevokePermission_Output? _defaultInstance;
}

/// revoke permission
class RoleRevokePermission extends $pb.GeneratedMessage {
  factory RoleRevokePermission() => create();

  RoleRevokePermission._() : super();

  factory RoleRevokePermission.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory RoleRevokePermission.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleRevokePermission',
      package: const $pb.PackageName(
          _omitMessageNames ? '' : 'roleRevokePermission'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  RoleRevokePermission clone() =>
      RoleRevokePermission()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  RoleRevokePermission copyWith(void Function(RoleRevokePermission) updates) =>
      super.copyWith((message) => updates(message as RoleRevokePermission))
          as RoleRevokePermission;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleRevokePermission create() => RoleRevokePermission._();

  RoleRevokePermission createEmptyInstance() => create();

  static $pb.PbList<RoleRevokePermission> createRepeated() =>
      $pb.PbList<RoleRevokePermission>();

  @$core.pragma('dart2js:noInline')
  static RoleRevokePermission getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleRevokePermission>(create);
  static RoleRevokePermission? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
