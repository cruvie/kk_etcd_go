// This is a generated file - do not edit.
//
// Generated from internal/api_hub/kv/api_def/KVPut.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class KVPut_Input extends $pb.GeneratedMessage {
  factory KVPut_Input({
    $core.String? key,
    $core.String? value,
  }) {
    final result = create();
    if (key != null) result.key = key;
    if (value != null) result.value = value;
    return result;
  }

  KVPut_Input._();

  factory KVPut_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory KVPut_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'KVPut.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..aOS(2, _omitFieldNames ? '' : 'Value', protoName: 'Value')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVPut_Input clone() => KVPut_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVPut_Input copyWith(void Function(KVPut_Input) updates) =>
      super.copyWith((message) => updates(message as KVPut_Input))
          as KVPut_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVPut_Input create() => KVPut_Input._();
  @$core.override
  KVPut_Input createEmptyInstance() => create();
  static $pb.PbList<KVPut_Input> createRepeated() => $pb.PbList<KVPut_Input>();
  @$core.pragma('dart2js:noInline')
  static KVPut_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<KVPut_Input>(create);
  static KVPut_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get value => $_getSZ(1);
  @$pb.TagNumber(2)
  set value($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => $_clearField(2);
}

class KVPut_Output extends $pb.GeneratedMessage {
  factory KVPut_Output() => create();

  KVPut_Output._();

  factory KVPut_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory KVPut_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'KVPut.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVPut_Output clone() => KVPut_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVPut_Output copyWith(void Function(KVPut_Output) updates) =>
      super.copyWith((message) => updates(message as KVPut_Output))
          as KVPut_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVPut_Output create() => KVPut_Output._();
  @$core.override
  KVPut_Output createEmptyInstance() => create();
  static $pb.PbList<KVPut_Output> createRepeated() =>
      $pb.PbList<KVPut_Output>();
  @$core.pragma('dart2js:noInline')
  static KVPut_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<KVPut_Output>(create);
  static KVPut_Output? _defaultInstance;
}

/// put kv
class KVPut extends $pb.GeneratedMessage {
  factory KVPut() => create();

  KVPut._();

  factory KVPut.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory KVPut.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'KVPut',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVPut clone() => KVPut()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVPut copyWith(void Function(KVPut) updates) =>
      super.copyWith((message) => updates(message as KVPut)) as KVPut;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVPut create() => KVPut._();
  @$core.override
  KVPut createEmptyInstance() => create();
  static $pb.PbList<KVPut> createRepeated() => $pb.PbList<KVPut>();
  @$core.pragma('dart2js:noInline')
  static KVPut getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVPut>(create);
  static KVPut? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
