// This is a generated file - do not edit.
//
// Generated from internal/service_hub/role/api_def/RoleGrantPermission.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../../kk_etcd_models/pb_role_kk_etcd.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class RoleGrantPermission_Input extends $pb.GeneratedMessage {
  factory RoleGrantPermission_Input({
    $core.String? name,
    $0.PBPermission? perm,
  }) {
    final result = create();
    if (name != null) result.name = name;
    if (perm != null) result.perm = perm;
    return result;
  }

  RoleGrantPermission_Input._();

  factory RoleGrantPermission_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleGrantPermission_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleGrantPermission.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..aOM<$0.PBPermission>(2, _omitFieldNames ? '' : 'Perm',
        protoName: 'Perm', subBuilder: $0.PBPermission.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGrantPermission_Input clone() =>
      RoleGrantPermission_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGrantPermission_Input copyWith(
          void Function(RoleGrantPermission_Input) updates) =>
      super.copyWith((message) => updates(message as RoleGrantPermission_Input))
          as RoleGrantPermission_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission_Input create() => RoleGrantPermission_Input._();
  @$core.override
  RoleGrantPermission_Input createEmptyInstance() => create();
  static $pb.PbList<RoleGrantPermission_Input> createRepeated() =>
      $pb.PbList<RoleGrantPermission_Input>();
  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleGrantPermission_Input>(create);
  static RoleGrantPermission_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);

  @$pb.TagNumber(2)
  $0.PBPermission get perm => $_getN(1);
  @$pb.TagNumber(2)
  set perm($0.PBPermission value) => $_setField(2, value);
  @$pb.TagNumber(2)
  $core.bool hasPerm() => $_has(1);
  @$pb.TagNumber(2)
  void clearPerm() => $_clearField(2);
  @$pb.TagNumber(2)
  $0.PBPermission ensurePerm() => $_ensure(1);
}

class RoleGrantPermission_Output extends $pb.GeneratedMessage {
  factory RoleGrantPermission_Output() => create();

  RoleGrantPermission_Output._();

  factory RoleGrantPermission_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleGrantPermission_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleGrantPermission.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGrantPermission_Output clone() =>
      RoleGrantPermission_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGrantPermission_Output copyWith(
          void Function(RoleGrantPermission_Output) updates) =>
      super.copyWith(
              (message) => updates(message as RoleGrantPermission_Output))
          as RoleGrantPermission_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission_Output create() => RoleGrantPermission_Output._();
  @$core.override
  RoleGrantPermission_Output createEmptyInstance() => create();
  static $pb.PbList<RoleGrantPermission_Output> createRepeated() =>
      $pb.PbList<RoleGrantPermission_Output>();
  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleGrantPermission_Output>(create);
  static RoleGrantPermission_Output? _defaultInstance;
}

/// grant permission
class RoleGrantPermission extends $pb.GeneratedMessage {
  factory RoleGrantPermission() => create();

  RoleGrantPermission._();

  factory RoleGrantPermission.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleGrantPermission.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleGrantPermission',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGrantPermission clone() => RoleGrantPermission()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGrantPermission copyWith(void Function(RoleGrantPermission) updates) =>
      super.copyWith((message) => updates(message as RoleGrantPermission))
          as RoleGrantPermission;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission create() => RoleGrantPermission._();
  @$core.override
  RoleGrantPermission createEmptyInstance() => create();
  static $pb.PbList<RoleGrantPermission> createRepeated() =>
      $pb.PbList<RoleGrantPermission>();
  @$core.pragma('dart2js:noInline')
  static RoleGrantPermission getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleGrantPermission>(create);
  static RoleGrantPermission? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
