//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/role/api_def/RoleAdd.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class RoleAdd_Input extends $pb.GeneratedMessage {
  factory RoleAdd_Input({
    $core.String? name,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  RoleAdd_Input._() : super();
  factory RoleAdd_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleAdd_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleAdd.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleAdd_Input clone() => RoleAdd_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleAdd_Input copyWith(void Function(RoleAdd_Input) updates) => super.copyWith((message) => updates(message as RoleAdd_Input)) as RoleAdd_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleAdd_Input create() => RoleAdd_Input._();
  RoleAdd_Input createEmptyInstance() => create();
  static $pb.PbList<RoleAdd_Input> createRepeated() => $pb.PbList<RoleAdd_Input>();
  @$core.pragma('dart2js:noInline')
  static RoleAdd_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleAdd_Input>(create);
  static RoleAdd_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);
}

class RoleAdd_Output extends $pb.GeneratedMessage {
  factory RoleAdd_Output() => create();
  RoleAdd_Output._() : super();
  factory RoleAdd_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleAdd_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleAdd.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleAdd_Output clone() => RoleAdd_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleAdd_Output copyWith(void Function(RoleAdd_Output) updates) => super.copyWith((message) => updates(message as RoleAdd_Output)) as RoleAdd_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleAdd_Output create() => RoleAdd_Output._();
  RoleAdd_Output createEmptyInstance() => create();
  static $pb.PbList<RoleAdd_Output> createRepeated() => $pb.PbList<RoleAdd_Output>();
  @$core.pragma('dart2js:noInline')
  static RoleAdd_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleAdd_Output>(create);
  static RoleAdd_Output? _defaultInstance;
}

/// add role
class RoleAdd extends $pb.GeneratedMessage {
  factory RoleAdd() => create();
  RoleAdd._() : super();
  factory RoleAdd.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleAdd.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleAdd', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleAdd clone() => RoleAdd()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleAdd copyWith(void Function(RoleAdd) updates) => super.copyWith((message) => updates(message as RoleAdd)) as RoleAdd;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleAdd create() => RoleAdd._();
  RoleAdd createEmptyInstance() => create();
  static $pb.PbList<RoleAdd> createRepeated() => $pb.PbList<RoleAdd>();
  @$core.pragma('dart2js:noInline')
  static RoleAdd getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleAdd>(create);
  static RoleAdd? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
