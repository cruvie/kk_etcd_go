// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_role_kk_etcd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class PBRole extends $pb.GeneratedMessage {
  factory PBRole({
    $core.String? name,
    $core.Iterable<PBPermission>? perms,
  }) {
    final result = create();
    if (name != null) result.name = name;
    if (perms != null) result.perms.addAll(perms);
    return result;
  }

  PBRole._();

  factory PBRole.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBRole.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBRole',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..pc<PBPermission>(2, _omitFieldNames ? '' : 'Perms', $pb.PbFieldType.PM,
        protoName: 'Perms', subBuilder: PBPermission.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBRole clone() => PBRole()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBRole copyWith(void Function(PBRole) updates) =>
      super.copyWith((message) => updates(message as PBRole)) as PBRole;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBRole create() => PBRole._();
  @$core.override
  PBRole createEmptyInstance() => create();
  static $pb.PbList<PBRole> createRepeated() => $pb.PbList<PBRole>();
  @$core.pragma('dart2js:noInline')
  static PBRole getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBRole>(create);
  static PBRole? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);

  @$pb.TagNumber(2)
  $pb.PbList<PBPermission> get perms => $_getList(1);
}

class PBPermission extends $pb.GeneratedMessage {
  factory PBPermission({
    $core.String? key,
    $core.String? rangeEnd,
    $core.int? permissionType,
  }) {
    final result = create();
    if (key != null) result.key = key;
    if (rangeEnd != null) result.rangeEnd = rangeEnd;
    if (permissionType != null) result.permissionType = permissionType;
    return result;
  }

  PBPermission._();

  factory PBPermission.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBPermission.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBPermission',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..aOS(2, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..aOS(3, _omitFieldNames ? '' : 'RangeEnd', protoName: 'RangeEnd')
    ..a<$core.int>(
        4, _omitFieldNames ? '' : 'PermissionType', $pb.PbFieldType.O3,
        protoName: 'PermissionType')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBPermission clone() => PBPermission()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBPermission copyWith(void Function(PBPermission) updates) =>
      super.copyWith((message) => updates(message as PBPermission))
          as PBPermission;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBPermission create() => PBPermission._();
  @$core.override
  PBPermission createEmptyInstance() => create();
  static $pb.PbList<PBPermission> createRepeated() =>
      $pb.PbList<PBPermission>();
  @$core.pragma('dart2js:noInline')
  static PBPermission getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<PBPermission>(create);
  static PBPermission? _defaultInstance;

  @$pb.TagNumber(2)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(2)
  set key($core.String value) => $_setString(0, value);
  @$pb.TagNumber(2)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(2)
  void clearKey() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get rangeEnd => $_getSZ(1);
  @$pb.TagNumber(3)
  set rangeEnd($core.String value) => $_setString(1, value);
  @$pb.TagNumber(3)
  $core.bool hasRangeEnd() => $_has(1);
  @$pb.TagNumber(3)
  void clearRangeEnd() => $_clearField(3);

  /// authpb.READ 0
  /// authpb.WRITE 1
  /// authpb.READWRITE 2
  @$pb.TagNumber(4)
  $core.int get permissionType => $_getIZ(2);
  @$pb.TagNumber(4)
  set permissionType($core.int value) => $_setSignedInt32(2, value);
  @$pb.TagNumber(4)
  $core.bool hasPermissionType() => $_has(2);
  @$pb.TagNumber(4)
  void clearPermissionType() => $_clearField(4);
}

class PBListRole extends $pb.GeneratedMessage {
  factory PBListRole({
    $core.Iterable<PBRole>? list,
  }) {
    final result = create();
    if (list != null) result.list.addAll(list);
    return result;
  }

  PBListRole._();

  factory PBListRole.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBListRole.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBListRole',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..pc<PBRole>(1, _omitFieldNames ? '' : 'List', $pb.PbFieldType.PM,
        protoName: 'List', subBuilder: PBRole.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListRole clone() => PBListRole()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListRole copyWith(void Function(PBListRole) updates) =>
      super.copyWith((message) => updates(message as PBListRole)) as PBListRole;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListRole create() => PBListRole._();
  @$core.override
  PBListRole createEmptyInstance() => create();
  static $pb.PbList<PBListRole> createRepeated() => $pb.PbList<PBListRole>();
  @$core.pragma('dart2js:noInline')
  static PBListRole getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<PBListRole>(create);
  static PBListRole? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<PBRole> get list => $_getList(0);
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
