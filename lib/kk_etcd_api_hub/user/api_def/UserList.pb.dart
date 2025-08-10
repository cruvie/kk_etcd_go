// This is a generated file - do not edit.
//
// Generated from kk_etcd_api_hub/user/api_def/UserList.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../kk_etcd_models/pb_user_kk_etcd.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class UserList_Input extends $pb.GeneratedMessage {
  factory UserList_Input() => create();

  UserList_Input._();

  factory UserList_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UserList_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserList.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserList_Input clone() => UserList_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserList_Input copyWith(void Function(UserList_Input) updates) =>
      super.copyWith((message) => updates(message as UserList_Input))
          as UserList_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserList_Input create() => UserList_Input._();
  @$core.override
  UserList_Input createEmptyInstance() => create();
  static $pb.PbList<UserList_Input> createRepeated() =>
      $pb.PbList<UserList_Input>();
  @$core.pragma('dart2js:noInline')
  static UserList_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserList_Input>(create);
  static UserList_Input? _defaultInstance;
}

class UserList_Output extends $pb.GeneratedMessage {
  factory UserList_Output({
    $0.PBListUser? listUser,
  }) {
    final result = create();
    if (listUser != null) result.listUser = listUser;
    return result;
  }

  UserList_Output._();

  factory UserList_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UserList_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserList.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOM<$0.PBListUser>(1, _omitFieldNames ? '' : 'ListUser',
        protoName: 'ListUser', subBuilder: $0.PBListUser.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserList_Output clone() => UserList_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserList_Output copyWith(void Function(UserList_Output) updates) =>
      super.copyWith((message) => updates(message as UserList_Output))
          as UserList_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserList_Output create() => UserList_Output._();
  @$core.override
  UserList_Output createEmptyInstance() => create();
  static $pb.PbList<UserList_Output> createRepeated() =>
      $pb.PbList<UserList_Output>();
  @$core.pragma('dart2js:noInline')
  static UserList_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UserList_Output>(create);
  static UserList_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBListUser get listUser => $_getN(0);
  @$pb.TagNumber(1)
  set listUser($0.PBListUser value) => $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasListUser() => $_has(0);
  @$pb.TagNumber(1)
  void clearListUser() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBListUser ensureListUser() => $_ensure(0);
}

/// list user
class UserList extends $pb.GeneratedMessage {
  factory UserList() => create();

  UserList._();

  factory UserList.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UserList.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UserList',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserList clone() => UserList()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserList copyWith(void Function(UserList) updates) =>
      super.copyWith((message) => updates(message as UserList)) as UserList;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserList create() => UserList._();
  @$core.override
  UserList createEmptyInstance() => create();
  static $pb.PbList<UserList> createRepeated() => $pb.PbList<UserList>();
  @$core.pragma('dart2js:noInline')
  static UserList getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserList>(create);
  static UserList? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
