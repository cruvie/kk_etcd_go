//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/user/api_def/Logout.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class Logout_Input extends $pb.GeneratedMessage {
  factory Logout_Input() => create();
  Logout_Input._() : super();
  factory Logout_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Logout_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Logout.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Logout_Input clone() => Logout_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Logout_Input copyWith(void Function(Logout_Input) updates) => super.copyWith((message) => updates(message as Logout_Input)) as Logout_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Logout_Input create() => Logout_Input._();
  Logout_Input createEmptyInstance() => create();
  static $pb.PbList<Logout_Input> createRepeated() => $pb.PbList<Logout_Input>();
  @$core.pragma('dart2js:noInline')
  static Logout_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Logout_Input>(create);
  static Logout_Input? _defaultInstance;
}

class Logout_Output extends $pb.GeneratedMessage {
  factory Logout_Output() => create();
  Logout_Output._() : super();
  factory Logout_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Logout_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Logout.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Logout_Output clone() => Logout_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Logout_Output copyWith(void Function(Logout_Output) updates) => super.copyWith((message) => updates(message as Logout_Output)) as Logout_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Logout_Output create() => Logout_Output._();
  Logout_Output createEmptyInstance() => create();
  static $pb.PbList<Logout_Output> createRepeated() => $pb.PbList<Logout_Output>();
  @$core.pragma('dart2js:noInline')
  static Logout_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Logout_Output>(create);
  static Logout_Output? _defaultInstance;
}

/// logout
class Logout extends $pb.GeneratedMessage {
  factory Logout() => create();
  Logout._() : super();
  factory Logout.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Logout.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Logout', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Logout clone() => Logout()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Logout copyWith(void Function(Logout) updates) => super.copyWith((message) => updates(message as Logout)) as Logout;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Logout create() => Logout._();
  Logout createEmptyInstance() => create();
  static $pb.PbList<Logout> createRepeated() => $pb.PbList<Logout>();
  @$core.pragma('dart2js:noInline')
  static Logout getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Logout>(create);
  static Logout? _defaultInstance;
}


const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
