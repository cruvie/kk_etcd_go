//
//  Generated code. Do not modify.
//  source: internal/mgr_hub/mgr_backup/snapshotInfo/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class SnapshotInfo_Input extends $pb.GeneratedMessage {
  factory SnapshotInfo_Input({
    $core.List<$core.int>? file,
  }) {
    final $result = create();
    if (file != null) {
      $result.file = file;
    }
    return $result;
  }

  SnapshotInfo_Input._() : super();

  factory SnapshotInfo_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory SnapshotInfo_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'SnapshotInfo.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'snapshotInfo'),
      createEmptyInstance: create)
    ..a<$core.List<$core.int>>(
        2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY,
        protoName: 'File')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  SnapshotInfo_Input clone() => SnapshotInfo_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  SnapshotInfo_Input copyWith(void Function(SnapshotInfo_Input) updates) =>
      super.copyWith((message) => updates(message as SnapshotInfo_Input))
          as SnapshotInfo_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotInfo_Input create() => SnapshotInfo_Input._();

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
  set file($core.List<$core.int> v) {
    $_setBytes(0, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(0);

  @$pb.TagNumber(2)
  void clearFile() => clearField(2);
}

class SnapshotInfo_Output extends $pb.GeneratedMessage {
  factory SnapshotInfo_Output({
    $core.String? info,
  }) {
    final $result = create();
    if (info != null) {
      $result.info = info;
    }
    return $result;
  }

  SnapshotInfo_Output._() : super();

  factory SnapshotInfo_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory SnapshotInfo_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'SnapshotInfo.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'snapshotInfo'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Info', protoName: 'Info')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  SnapshotInfo_Output clone() => SnapshotInfo_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  SnapshotInfo_Output copyWith(void Function(SnapshotInfo_Output) updates) =>
      super.copyWith((message) => updates(message as SnapshotInfo_Output))
          as SnapshotInfo_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotInfo_Output create() => SnapshotInfo_Output._();

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
  set info($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasInfo() => $_has(0);

  @$pb.TagNumber(1)
  void clearInfo() => clearField(1);
}

/// snapshot info
class SnapshotInfo extends $pb.GeneratedMessage {
  factory SnapshotInfo() => create();

  SnapshotInfo._() : super();

  factory SnapshotInfo.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory SnapshotInfo.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'SnapshotInfo',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'snapshotInfo'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  SnapshotInfo clone() => SnapshotInfo()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  SnapshotInfo copyWith(void Function(SnapshotInfo) updates) =>
      super.copyWith((message) => updates(message as SnapshotInfo))
          as SnapshotInfo;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotInfo create() => SnapshotInfo._();

  SnapshotInfo createEmptyInstance() => create();

  static $pb.PbList<SnapshotInfo> createRepeated() =>
      $pb.PbList<SnapshotInfo>();

  @$core.pragma('dart2js:noInline')
  static SnapshotInfo getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<SnapshotInfo>(create);
  static SnapshotInfo? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
