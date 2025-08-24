// This is a generated file - do not edit.
//
// Generated from internal/service_hub/user/api_def/MyInfo.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class MyInfo_Input extends $pb.GeneratedMessage {
  factory MyInfo_Input() => create();

  MyInfo_Input._();

  factory MyInfo_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory MyInfo_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'MyInfo.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  MyInfo_Input clone() => MyInfo_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  MyInfo_Input copyWith(void Function(MyInfo_Input) updates) =>
      super.copyWith((message) => updates(message as MyInfo_Input))
          as MyInfo_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static MyInfo_Input create() => MyInfo_Input._();
  @$core.override
  MyInfo_Input createEmptyInstance() => create();
  static $pb.PbList<MyInfo_Input> createRepeated() =>
      $pb.PbList<MyInfo_Input>();
  @$core.pragma('dart2js:noInline')
  static MyInfo_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<MyInfo_Input>(create);
  static MyInfo_Input? _defaultInstance;
}

class MyInfo_Output extends $pb.GeneratedMessage {
  factory MyInfo_Output({
    $core.String? userName,
    $core.Iterable<$core.String>? roles,
  }) {
    final result = create();
    if (userName != null) result.userName = userName;
    if (roles != null) result.roles.addAll(roles);
    return result;
  }

  MyInfo_Output._();

  factory MyInfo_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory MyInfo_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'MyInfo.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..aOS(2, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..pPS(4, _omitFieldNames ? '' : 'Roles', protoName: 'Roles')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  MyInfo_Output clone() => MyInfo_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  MyInfo_Output copyWith(void Function(MyInfo_Output) updates) =>
      super.copyWith((message) => updates(message as MyInfo_Output))
          as MyInfo_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static MyInfo_Output create() => MyInfo_Output._();
  @$core.override
  MyInfo_Output createEmptyInstance() => create();
  static $pb.PbList<MyInfo_Output> createRepeated() =>
      $pb.PbList<MyInfo_Output>();
  @$core.pragma('dart2js:noInline')
  static MyInfo_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<MyInfo_Output>(create);
  static MyInfo_Output? _defaultInstance;

  @$pb.TagNumber(2)
  $core.String get userName => $_getSZ(0);
  @$pb.TagNumber(2)
  set userName($core.String value) => $_setString(0, value);
  @$pb.TagNumber(2)
  $core.bool hasUserName() => $_has(0);
  @$pb.TagNumber(2)
  void clearUserName() => $_clearField(2);

  @$pb.TagNumber(4)
  $pb.PbList<$core.String> get roles => $_getList(1);
}

/// get my info
class MyInfo extends $pb.GeneratedMessage {
  factory MyInfo() => create();

  MyInfo._();

  factory MyInfo.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory MyInfo.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'MyInfo',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  MyInfo clone() => MyInfo()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  MyInfo copyWith(void Function(MyInfo) updates) =>
      super.copyWith((message) => updates(message as MyInfo)) as MyInfo;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static MyInfo create() => MyInfo._();
  @$core.override
  MyInfo createEmptyInstance() => create();
  static $pb.PbList<MyInfo> createRepeated() => $pb.PbList<MyInfo>();
  @$core.pragma('dart2js:noInline')
  static MyInfo getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MyInfo>(create);
  static MyInfo? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
