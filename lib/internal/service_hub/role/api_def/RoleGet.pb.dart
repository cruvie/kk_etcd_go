// This is a generated file - do not edit.
//
// Generated from internal/service_hub/role/api_def/RoleGet.proto.

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

class RoleGet_Input extends $pb.GeneratedMessage {
  factory RoleGet_Input({
    $core.String? name,
  }) {
    final result = create();
    if (name != null) result.name = name;
    return result;
  }

  RoleGet_Input._();

  factory RoleGet_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleGet_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleGet.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGet_Input clone() => RoleGet_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGet_Input copyWith(void Function(RoleGet_Input) updates) =>
      super.copyWith((message) => updates(message as RoleGet_Input))
          as RoleGet_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGet_Input create() => RoleGet_Input._();
  @$core.override
  RoleGet_Input createEmptyInstance() => create();
  static $pb.PbList<RoleGet_Input> createRepeated() =>
      $pb.PbList<RoleGet_Input>();
  @$core.pragma('dart2js:noInline')
  static RoleGet_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleGet_Input>(create);
  static RoleGet_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);
}

class RoleGet_Output extends $pb.GeneratedMessage {
  factory RoleGet_Output({
    $0.PBRole? role,
  }) {
    final result = create();
    if (role != null) result.role = role;
    return result;
  }

  RoleGet_Output._();

  factory RoleGet_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleGet_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleGet.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOM<$0.PBRole>(1, _omitFieldNames ? '' : 'Role',
        protoName: 'Role', subBuilder: $0.PBRole.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGet_Output clone() => RoleGet_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGet_Output copyWith(void Function(RoleGet_Output) updates) =>
      super.copyWith((message) => updates(message as RoleGet_Output))
          as RoleGet_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGet_Output create() => RoleGet_Output._();
  @$core.override
  RoleGet_Output createEmptyInstance() => create();
  static $pb.PbList<RoleGet_Output> createRepeated() =>
      $pb.PbList<RoleGet_Output>();
  @$core.pragma('dart2js:noInline')
  static RoleGet_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleGet_Output>(create);
  static RoleGet_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBRole get role => $_getN(0);
  @$pb.TagNumber(1)
  set role($0.PBRole value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasRole() => $_has(0);
  @$pb.TagNumber(1)
  void clearRole() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBRole ensureRole() => $_ensure(0);
}

/// get role
class RoleGet extends $pb.GeneratedMessage {
  factory RoleGet() => create();

  RoleGet._();

  factory RoleGet.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleGet.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleGet',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGet clone() => RoleGet()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleGet copyWith(void Function(RoleGet) updates) =>
      super.copyWith((message) => updates(message as RoleGet)) as RoleGet;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleGet create() => RoleGet._();
  @$core.override
  RoleGet createEmptyInstance() => create();
  static $pb.PbList<RoleGet> createRepeated() => $pb.PbList<RoleGet>();
  @$core.pragma('dart2js:noInline')
  static RoleGet getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleGet>(create);
  static RoleGet? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
