//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_backup/allKVsBackup/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class AllKVsBackup_Input extends $pb.GeneratedMessage {
  factory AllKVsBackup_Input() => create();

  AllKVsBackup_Input._() : super();

  factory AllKVsBackup_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory AllKVsBackup_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'AllKVsBackup.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'allKVsBackup'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  AllKVsBackup_Input clone() => AllKVsBackup_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  AllKVsBackup_Input copyWith(void Function(AllKVsBackup_Input) updates) =>
      super.copyWith((message) => updates(message as AllKVsBackup_Input))
          as AllKVsBackup_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsBackup_Input create() => AllKVsBackup_Input._();

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
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    if (file != null) {
      $result.file = file;
    }
    return $result;
  }

  AllKVsBackup_Output._() : super();

  factory AllKVsBackup_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory AllKVsBackup_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'AllKVsBackup.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'allKVsBackup'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..a<$core.List<$core.int>>(
        2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY,
        protoName: 'File')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  AllKVsBackup_Output clone() => AllKVsBackup_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  AllKVsBackup_Output copyWith(void Function(AllKVsBackup_Output) updates) =>
      super.copyWith((message) => updates(message as AllKVsBackup_Output))
          as AllKVsBackup_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsBackup_Output create() => AllKVsBackup_Output._();

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
  set name($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);

  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get file => $_getN(1);

  @$pb.TagNumber(2)
  set file($core.List<$core.int> v) {
    $_setBytes(1, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(1);

  @$pb.TagNumber(2)
  void clearFile() => clearField(2);
}

/// all kvs backup
class AllKVsBackup extends $pb.GeneratedMessage {
  factory AllKVsBackup() => create();

  AllKVsBackup._() : super();

  factory AllKVsBackup.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory AllKVsBackup.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'AllKVsBackup',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'allKVsBackup'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  AllKVsBackup clone() => AllKVsBackup()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  AllKVsBackup copyWith(void Function(AllKVsBackup) updates) =>
      super.copyWith((message) => updates(message as AllKVsBackup))
          as AllKVsBackup;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsBackup create() => AllKVsBackup._();

  AllKVsBackup createEmptyInstance() => create();

  static $pb.PbList<AllKVsBackup> createRepeated() =>
      $pb.PbList<AllKVsBackup>();

  @$core.pragma('dart2js:noInline')
  static AllKVsBackup getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<AllKVsBackup>(create);
  static AllKVsBackup? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
