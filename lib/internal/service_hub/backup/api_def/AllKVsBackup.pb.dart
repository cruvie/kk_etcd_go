// This is a generated file - do not edit.
//
// Generated from internal/service_hub/backup/api_def/AllKVsBackup.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class AllKVsBackup_Input extends $pb.GeneratedMessage {
  factory AllKVsBackup_Input() => create();

  AllKVsBackup_Input._();

  factory AllKVsBackup_Input.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory AllKVsBackup_Input.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'AllKVsBackup.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsBackup_Input clone() => AllKVsBackup_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsBackup_Input copyWith(void Function(AllKVsBackup_Input) updates) =>
      super.copyWith((message) => updates(message as AllKVsBackup_Input))
          as AllKVsBackup_Input;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsBackup_Input create() => AllKVsBackup_Input._();
  @$core.override
  AllKVsBackup_Input createEmptyInstance() => create();
  static $pb.PbList<AllKVsBackup_Input> createRepeated() =>
      $pb.PbList<AllKVsBackup_Input>();
  @$core.pragma('dart2js:noInline')
  static AllKVsBackup_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<AllKVsBackup_Input>(create);
  static AllKVsBackup_Input? _defaultInstance;
}

class AllKVsBackup_Output extends $pb.GeneratedMessage {
  factory AllKVsBackup_Output({
    $core.String? name,
    $core.List<$core.int>? file,
  }) {
    final result = create();
    if (name != null) result.name = name;
    if (file != null) result.file = file;
    return result;
  }

  AllKVsBackup_Output._();

  factory AllKVsBackup_Output.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory AllKVsBackup_Output.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'AllKVsBackup.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..a<$core.List<$core.int>>(
        2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY,
        protoName: 'File')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsBackup_Output clone() => AllKVsBackup_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsBackup_Output copyWith(void Function(AllKVsBackup_Output) updates) =>
      super.copyWith((message) => updates(message as AllKVsBackup_Output))
          as AllKVsBackup_Output;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsBackup_Output create() => AllKVsBackup_Output._();
  @$core.override
  AllKVsBackup_Output createEmptyInstance() => create();
  static $pb.PbList<AllKVsBackup_Output> createRepeated() =>
      $pb.PbList<AllKVsBackup_Output>();
  @$core.pragma('dart2js:noInline')
  static AllKVsBackup_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<AllKVsBackup_Output>(create);
  static AllKVsBackup_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get file => $_getN(1);
  @$pb.TagNumber(2)
  set file($core.List<$core.int> value) => $_setBytes(1, value);
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(1);
  @$pb.TagNumber(2)
  void clearFile() => $_clearField(2);
}

/// all kvs backup
class AllKVsBackup extends $pb.GeneratedMessage {
  factory AllKVsBackup() => create();

  AllKVsBackup._();

  factory AllKVsBackup.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory AllKVsBackup.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'AllKVsBackup',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsBackup clone() => AllKVsBackup()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsBackup copyWith(void Function(AllKVsBackup) updates) =>
      super.copyWith((message) => updates(message as AllKVsBackup))
          as AllKVsBackup;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsBackup create() => AllKVsBackup._();
  @$core.override
  AllKVsBackup createEmptyInstance() => create();
  static $pb.PbList<AllKVsBackup> createRepeated() =>
      $pb.PbList<AllKVsBackup>();
  @$core.pragma('dart2js:noInline')
  static AllKVsBackup getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<AllKVsBackup>(create);
  static AllKVsBackup? _defaultInstance;
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
