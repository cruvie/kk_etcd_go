// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_user_kk_etcd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class PBUser extends $pb.GeneratedMessage {
  factory PBUser({
    $core.String? userName,
    $core.String? password,
    $core.Iterable<$core.String>? roles,
  }) {
    final result = create();
    if (userName != null) result.userName = userName;
    if (password != null) result.password = password;
    if (roles != null) result.roles.addAll(roles);
    return result;
  }

  PBUser._();

  factory PBUser.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBUser.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBUser',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'),
      createEmptyInstance: create)
    ..aOS(2, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..aOS(3, _omitFieldNames ? '' : 'Password', protoName: 'Password')
    ..pPS(4, _omitFieldNames ? '' : 'Roles', protoName: 'Roles')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBUser clone() => PBUser()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBUser copyWith(void Function(PBUser) updates) =>
      super.copyWith((message) => updates(message as PBUser)) as PBUser;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBUser create() => PBUser._();
  @$core.override
  PBUser createEmptyInstance() => create();
  static $pb.PbList<PBUser> createRepeated() => $pb.PbList<PBUser>();
  @$core.pragma('dart2js:noInline')
  static PBUser getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBUser>(create);
  static PBUser? _defaultInstance;

  @$pb.TagNumber(2)
  $core.String get userName => $_getSZ(0);
  @$pb.TagNumber(2)
  set userName($core.String value) => $_setString(0, value);
  @$pb.TagNumber(2)
  $core.bool hasUserName() => $_has(0);
  @$pb.TagNumber(2)
  void clearUserName() => $_clearField(2);

  @$pb.TagNumber(3)
  $core.String get password => $_getSZ(1);
  @$pb.TagNumber(3)
  set password($core.String value) => $_setString(1, value);
  @$pb.TagNumber(3)
  $core.bool hasPassword() => $_has(1);
  @$pb.TagNumber(3)
  void clearPassword() => $_clearField(3);

  @$pb.TagNumber(4)
  $pb.PbList<$core.String> get roles => $_getList(2);
}

class PBListUser extends $pb.GeneratedMessage {
  factory PBListUser({
    $core.Iterable<PBUser>? listUser,
  }) {
    final result = create();
    if (listUser != null) result.listUser.addAll(listUser);
    return result;
  }

  PBListUser._();

  factory PBListUser.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBListUser.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBListUser',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'),
      createEmptyInstance: create)
    ..pc<PBUser>(1, _omitFieldNames ? '' : 'ListUser', $pb.PbFieldType.PM,
        protoName: 'ListUser', subBuilder: PBUser.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListUser clone() => PBListUser()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListUser copyWith(void Function(PBListUser) updates) =>
      super.copyWith((message) => updates(message as PBListUser)) as PBListUser;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListUser create() => PBListUser._();
  @$core.override
  PBListUser createEmptyInstance() => create();
  static $pb.PbList<PBListUser> createRepeated() => $pb.PbList<PBListUser>();
  @$core.pragma('dart2js:noInline')
  static PBListUser getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<PBListUser>(create);
  static PBListUser? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<PBUser> get listUser => $_getList(0);
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
