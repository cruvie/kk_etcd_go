// This is a generated file - do not edit.
//
// Generated from kk_etcd_api_hub/role/api_def/RoleDelete.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class RoleDelete_Input extends $pb.GeneratedMessage {
  factory RoleDelete_Input({
    $core.String? name,
  }) {
    final result = create();
    if (name != null) result.name = name;
    return result;
  }

  RoleDelete_Input._();

  factory RoleDelete_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleDelete_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleDelete.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleDelete_Input clone() => RoleDelete_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleDelete_Input copyWith(void Function(RoleDelete_Input) updates) =>
      super.copyWith((message) => updates(message as RoleDelete_Input))
          as RoleDelete_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleDelete_Input create() => RoleDelete_Input._();
  @$core.override
  RoleDelete_Input createEmptyInstance() => create();
  static $pb.PbList<RoleDelete_Input> createRepeated() =>
      $pb.PbList<RoleDelete_Input>();
  @$core.pragma('dart2js:noInline')
  static RoleDelete_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleDelete_Input>(create);
  static RoleDelete_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);
}

class RoleDelete_Output extends $pb.GeneratedMessage {
  factory RoleDelete_Output() => create();

  RoleDelete_Output._();

  factory RoleDelete_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleDelete_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleDelete.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleDelete_Output clone() => RoleDelete_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleDelete_Output copyWith(void Function(RoleDelete_Output) updates) =>
      super.copyWith((message) => updates(message as RoleDelete_Output))
          as RoleDelete_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleDelete_Output create() => RoleDelete_Output._();
  @$core.override
  RoleDelete_Output createEmptyInstance() => create();
  static $pb.PbList<RoleDelete_Output> createRepeated() =>
      $pb.PbList<RoleDelete_Output>();
  @$core.pragma('dart2js:noInline')
  static RoleDelete_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleDelete_Output>(create);
  static RoleDelete_Output? _defaultInstance;
}

/// delete role
class RoleDelete extends $pb.GeneratedMessage {
  factory RoleDelete() => create();

  RoleDelete._();

  factory RoleDelete.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory RoleDelete.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'RoleDelete',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleDelete clone() => RoleDelete()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  RoleDelete copyWith(void Function(RoleDelete) updates) =>
      super.copyWith((message) => updates(message as RoleDelete)) as RoleDelete;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleDelete create() => RoleDelete._();
  @$core.override
  RoleDelete createEmptyInstance() => create();
  static $pb.PbList<RoleDelete> createRepeated() => $pb.PbList<RoleDelete>();
  @$core.pragma('dart2js:noInline')
  static RoleDelete getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<RoleDelete>(create);
  static RoleDelete? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
