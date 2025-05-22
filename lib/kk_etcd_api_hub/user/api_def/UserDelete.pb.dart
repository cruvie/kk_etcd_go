//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/user/api_def/UserDelete.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class UserDelete_Input extends $pb.GeneratedMessage {
  factory UserDelete_Input({
    $core.String? userName,
  }) {
    final $result = create();
    if (userName != null) {
      $result.userName = userName;
    }
    return $result;
  }
  UserDelete_Input._() : super();
  factory UserDelete_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserDelete_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'UserDelete.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'UserName', protoName: 'UserName')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete_Input clone() => UserDelete_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete_Input copyWith(void Function(UserDelete_Input) updates) => super.copyWith((message) => updates(message as UserDelete_Input)) as UserDelete_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserDelete_Input create() => UserDelete_Input._();
  UserDelete_Input createEmptyInstance() => create();
  static $pb.PbList<UserDelete_Input> createRepeated() => $pb.PbList<UserDelete_Input>();
  @$core.pragma('dart2js:noInline')
  static UserDelete_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserDelete_Input>(create);
  static UserDelete_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get userName => $_getSZ(0);
  @$pb.TagNumber(1)
  set userName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserName() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserName() => $_clearField(1);
}

class UserDelete_Output extends $pb.GeneratedMessage {
  factory UserDelete_Output() => create();
  UserDelete_Output._() : super();
  factory UserDelete_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserDelete_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'UserDelete.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete_Output clone() => UserDelete_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete_Output copyWith(void Function(UserDelete_Output) updates) => super.copyWith((message) => updates(message as UserDelete_Output)) as UserDelete_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserDelete_Output create() => UserDelete_Output._();
  UserDelete_Output createEmptyInstance() => create();
  static $pb.PbList<UserDelete_Output> createRepeated() => $pb.PbList<UserDelete_Output>();
  @$core.pragma('dart2js:noInline')
  static UserDelete_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserDelete_Output>(create);
  static UserDelete_Output? _defaultInstance;
}

/// delete user
class UserDelete extends $pb.GeneratedMessage {
  factory UserDelete() => create();
  UserDelete._() : super();
  factory UserDelete.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserDelete.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'UserDelete', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete clone() => UserDelete()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UserDelete copyWith(void Function(UserDelete) updates) => super.copyWith((message) => updates(message as UserDelete)) as UserDelete;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UserDelete create() => UserDelete._();
  UserDelete createEmptyInstance() => create();
  static $pb.PbList<UserDelete> createRepeated() => $pb.PbList<UserDelete>();
  @$core.pragma('dart2js:noInline')
  static UserDelete getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserDelete>(create);
  static UserDelete? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
