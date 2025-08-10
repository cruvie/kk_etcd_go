// This is a generated file - do not edit.
//
// Generated from internal/api_hub/backup/api_def/SnapshotRestore.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class SnapshotRestore_Input extends $pb.GeneratedMessage {
  factory SnapshotRestore_Input() => create();

  SnapshotRestore_Input._();

  factory SnapshotRestore_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory SnapshotRestore_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'SnapshotRestore.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotRestore_Input clone() =>
      SnapshotRestore_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotRestore_Input copyWith(
          void Function(SnapshotRestore_Input) updates) =>
      super.copyWith((message) => updates(message as SnapshotRestore_Input))
          as SnapshotRestore_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotRestore_Input create() => SnapshotRestore_Input._();
  @$core.override
  SnapshotRestore_Input createEmptyInstance() => create();
  static $pb.PbList<SnapshotRestore_Input> createRepeated() =>
      $pb.PbList<SnapshotRestore_Input>();
  @$core.pragma('dart2js:noInline')
  static SnapshotRestore_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<SnapshotRestore_Input>(create);
  static SnapshotRestore_Input? _defaultInstance;
}

class SnapshotRestore_Output extends $pb.GeneratedMessage {
  factory SnapshotRestore_Output({
    $core.String? cmdStr,
  }) {
    final result = create();
    if (cmdStr != null) result.cmdStr = cmdStr;
    return result;
  }

  SnapshotRestore_Output._();

  factory SnapshotRestore_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory SnapshotRestore_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'SnapshotRestore.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'CmdStr', protoName: 'CmdStr')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotRestore_Output clone() =>
      SnapshotRestore_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotRestore_Output copyWith(
          void Function(SnapshotRestore_Output) updates) =>
      super.copyWith((message) => updates(message as SnapshotRestore_Output))
          as SnapshotRestore_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotRestore_Output create() => SnapshotRestore_Output._();
  @$core.override
  SnapshotRestore_Output createEmptyInstance() => create();
  static $pb.PbList<SnapshotRestore_Output> createRepeated() =>
      $pb.PbList<SnapshotRestore_Output>();
  @$core.pragma('dart2js:noInline')
  static SnapshotRestore_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<SnapshotRestore_Output>(create);
  static SnapshotRestore_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get cmdStr => $_getSZ(0);
  @$pb.TagNumber(1)
  set cmdStr($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasCmdStr() => $_has(0);
  @$pb.TagNumber(1)
  void clearCmdStr() => $_clearField(1);
}

/// snapshot restore
class SnapshotRestore extends $pb.GeneratedMessage {
  factory SnapshotRestore() => create();

  SnapshotRestore._();

  factory SnapshotRestore.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory SnapshotRestore.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'SnapshotRestore',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotRestore clone() => SnapshotRestore()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  SnapshotRestore copyWith(void Function(SnapshotRestore) updates) =>
      super.copyWith((message) => updates(message as SnapshotRestore))
          as SnapshotRestore;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotRestore create() => SnapshotRestore._();
  @$core.override
  SnapshotRestore createEmptyInstance() => create();
  static $pb.PbList<SnapshotRestore> createRepeated() =>
      $pb.PbList<SnapshotRestore>();
  @$core.pragma('dart2js:noInline')
  static SnapshotRestore getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<SnapshotRestore>(create);
  static SnapshotRestore? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
