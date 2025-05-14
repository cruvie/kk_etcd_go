//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/backup/api_def/SnapshotRestore.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class SnapshotRestore_Input extends $pb.GeneratedMessage {
  factory SnapshotRestore_Input() => create();
  SnapshotRestore_Input._() : super();
  factory SnapshotRestore_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotRestore_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SnapshotRestore.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotRestore_Input clone() => SnapshotRestore_Input()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotRestore_Input copyWith(void Function(SnapshotRestore_Input) updates) => super.copyWith((message) => updates(message as SnapshotRestore_Input)) as SnapshotRestore_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotRestore_Input create() => SnapshotRestore_Input._();
  SnapshotRestore_Input createEmptyInstance() => create();
  static $pb.PbList<SnapshotRestore_Input> createRepeated() => $pb.PbList<SnapshotRestore_Input>();
  @$core.pragma('dart2js:noInline')
  static SnapshotRestore_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotRestore_Input>(create);
  static SnapshotRestore_Input? _defaultInstance;
}

class SnapshotRestore_Output extends $pb.GeneratedMessage {
  factory SnapshotRestore_Output({
    $core.String? cmdStr,
  }) {
    final $result = create();
    if (cmdStr != null) {
      $result.cmdStr = cmdStr;
    }
    return $result;
  }
  SnapshotRestore_Output._() : super();
  factory SnapshotRestore_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotRestore_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SnapshotRestore.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'CmdStr', protoName: 'CmdStr')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotRestore_Output clone() => SnapshotRestore_Output()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotRestore_Output copyWith(void Function(SnapshotRestore_Output) updates) => super.copyWith((message) => updates(message as SnapshotRestore_Output)) as SnapshotRestore_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotRestore_Output create() => SnapshotRestore_Output._();
  SnapshotRestore_Output createEmptyInstance() => create();
  static $pb.PbList<SnapshotRestore_Output> createRepeated() => $pb.PbList<SnapshotRestore_Output>();
  @$core.pragma('dart2js:noInline')
  static SnapshotRestore_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotRestore_Output>(create);
  static SnapshotRestore_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get cmdStr => $_getSZ(0);
  @$pb.TagNumber(1)
  set cmdStr($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCmdStr() => $_has(0);
  @$pb.TagNumber(1)
  void clearCmdStr() => $_clearField(1);
}

/// snapshot restore
class SnapshotRestore extends $pb.GeneratedMessage {
  factory SnapshotRestore() => create();
  SnapshotRestore._() : super();
  factory SnapshotRestore.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SnapshotRestore.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SnapshotRestore', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SnapshotRestore clone() => SnapshotRestore()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SnapshotRestore copyWith(void Function(SnapshotRestore) updates) => super.copyWith((message) => updates(message as SnapshotRestore)) as SnapshotRestore;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SnapshotRestore create() => SnapshotRestore._();
  SnapshotRestore createEmptyInstance() => create();
  static $pb.PbList<SnapshotRestore> createRepeated() => $pb.PbList<SnapshotRestore>();
  @$core.pragma('dart2js:noInline')
  static SnapshotRestore getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SnapshotRestore>(create);
  static SnapshotRestore? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
