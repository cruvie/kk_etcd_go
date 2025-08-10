// This is a generated file - do not edit.
//
// Generated from internal/api_hub/role/api_def/RoleList.proto.

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

class RoleList_Input extends $pb.GeneratedMessage {
  factory RoleList_Input() => create();

  RoleList_Input._();

  factory RoleList_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleList_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleList.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleList_Input clone() => RoleList_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleList_Input copyWith(void Function(RoleList_Input) updates) =>
      super.copyWith((message) => updates(message as RoleList_Input))
          as RoleList_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleList_Input create() => RoleList_Input._();
  @$core.override
  RoleList_Input createEmptyInstance() => create();
  static $pb.PbList<RoleList_Input> createRepeated() =>
      $pb.PbList<RoleList_Input>();
  @$core.pragma('dart2js:noInline')
  static RoleList_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleList_Input>(create);
  static RoleList_Input? _defaultInstance;
}

class RoleList_Output extends $pb.GeneratedMessage {
  factory RoleList_Output({
    $0.PBListRole? listRole,
  }) {
    final result = create();
    if (listRole != null) result.listRole = listRole;
    return result;
  }

  RoleList_Output._();

  factory RoleList_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleList_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleList.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOM<$0.PBListRole>(1, _omitFieldNames ? '' : 'ListRole',
        protoName: 'ListRole', subBuilder: $0.PBListRole.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleList_Output clone() => RoleList_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleList_Output copyWith(void Function(RoleList_Output) updates) =>
      super.copyWith((message) => updates(message as RoleList_Output))
          as RoleList_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleList_Output create() => RoleList_Output._();
  @$core.override
  RoleList_Output createEmptyInstance() => create();
  static $pb.PbList<RoleList_Output> createRepeated() =>
      $pb.PbList<RoleList_Output>();
  @$core.pragma('dart2js:noInline')
  static RoleList_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleList_Output>(create);
  static RoleList_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBListRole get listRole => $_getN(0);
  @$pb.TagNumber(1)
  set listRole($0.PBListRole value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasListRole() => $_has(0);
  @$pb.TagNumber(1)
  void clearListRole() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBListRole ensureListRole() => $_ensure(0);
}

/// list role
class RoleList extends $pb.GeneratedMessage {
  factory RoleList() => create();

  RoleList._();

  factory RoleList.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleList.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleList',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleList clone() => RoleList()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleList copyWith(void Function(RoleList) updates) =>
      super.copyWith((message) => updates(message as RoleList)) as RoleList;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleList create() => RoleList._();
  @$core.override
  RoleList createEmptyInstance() => create();
  static $pb.PbList<RoleList> createRepeated() => $pb.PbList<RoleList>();
  @$core.pragma('dart2js:noInline')
  static RoleList getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleList>(create);
  static RoleList? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
