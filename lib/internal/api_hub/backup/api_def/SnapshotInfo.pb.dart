// This is a generated file - do not edit.
//
// Generated from internal/api_hub/backup/api_def/SnapshotInfo.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class SnapshotInfo_Input extends $pb.GeneratedMessage {
  factory SnapshotInfo_Input({
    $core.List<$core.int>? file,
  }) {
    final result = create();
    if (file != null) result.file = file;
    return result;
  }

  SnapshotInfo_Input._();

  factory SnapshotInfo_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory SnapshotInfo_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'SnapshotInfo.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..a<$core.List<$core.int>>(
        2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY,
        protoName: 'File')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotInfo_Input clone() => SnapshotInfo_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotInfo_Input copyWith(void Function(SnapshotInfo_Input) updates) =>
      super.copyWith((message) => updates(message as SnapshotInfo_Input))
          as SnapshotInfo_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotInfo_Input create() => SnapshotInfo_Input._();
  @$core.override
  SnapshotInfo_Input createEmptyInstance() => create();
  static $pb.PbList<SnapshotInfo_Input> createRepeated() =>
      $pb.PbList<SnapshotInfo_Input>();
  @$core.pragma('dart2js:noInline')
  static SnapshotInfo_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<SnapshotInfo_Input>(create);
  static SnapshotInfo_Input? _defaultInstance;

  @$pb.TagNumber(2)
  $core.List<$core.int> get file => $_getN(0);
  @$pb.TagNumber(2)
  set file($core.List<$core.int> value) => $_setBytes(0, value);
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(0);
  @$pb.TagNumber(2)
  void clearFile() => $_clearField(2);
}

class SnapshotInfo_Output extends $pb.GeneratedMessage {
  factory SnapshotInfo_Output({
    $core.String? info,
  }) {
    final result = create();
    if (info != null) result.info = info;
    return result;
  }

  SnapshotInfo_Output._();

  factory SnapshotInfo_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory SnapshotInfo_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'SnapshotInfo.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Info', protoName: 'Info')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotInfo_Output clone() => SnapshotInfo_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotInfo_Output copyWith(void Function(SnapshotInfo_Output) updates) =>
      super.copyWith((message) => updates(message as SnapshotInfo_Output))
          as SnapshotInfo_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotInfo_Output create() => SnapshotInfo_Output._();
  @$core.override
  SnapshotInfo_Output createEmptyInstance() => create();
  static $pb.PbList<SnapshotInfo_Output> createRepeated() =>
      $pb.PbList<SnapshotInfo_Output>();
  @$core.pragma('dart2js:noInline')
  static SnapshotInfo_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<SnapshotInfo_Output>(create);
  static SnapshotInfo_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get info => $_getSZ(0);
  @$pb.TagNumber(1)
  set info($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasInfo() => $_has(0);
  @$pb.TagNumber(1)
  void clearInfo() => $_clearField(1);
}

/// snapshot info
class SnapshotInfo extends $pb.GeneratedMessage {
  factory SnapshotInfo() => create();

  SnapshotInfo._();

  factory SnapshotInfo.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory SnapshotInfo.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'SnapshotInfo',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotInfo clone() => SnapshotInfo()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotInfo copyWith(void Function(SnapshotInfo) updates) =>
      super.copyWith((message) => updates(message as SnapshotInfo))
          as SnapshotInfo;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotInfo create() => SnapshotInfo._();
  @$core.override
  SnapshotInfo createEmptyInstance() => create();
  static $pb.PbList<SnapshotInfo> createRepeated() =>
      $pb.PbList<SnapshotInfo>();
  @$core.pragma('dart2js:noInline')
  static SnapshotInfo getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<SnapshotInfo>(create);
  static SnapshotInfo? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
