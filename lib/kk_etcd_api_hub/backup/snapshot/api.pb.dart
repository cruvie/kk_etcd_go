//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/backup/snapshot/api.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Snapshot_Input extends $pb.GeneratedMessage {
  factory Snapshot_Input() => create();

  Snapshot_Input._() : super();

  factory Snapshot_Input.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory Snapshot_Input.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Snapshot.Input',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'snapshot'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  Snapshot_Input clone() => Snapshot_Input()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  Snapshot_Input copyWith(void Function(Snapshot_Input) updates) =>
      super.copyWith((message) => updates(message as Snapshot_Input))
          as Snapshot_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Snapshot_Input create() => Snapshot_Input._();

  Snapshot_Input createEmptyInstance() => create();

  static $pb.PbList<Snapshot_Input> createRepeated() =>
      $pb.PbList<Snapshot_Input>();

  @$core.pragma('dart2js:noInline')
  static Snapshot_Input getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<Snapshot_Input>(create);
  static Snapshot_Input? _defaultInstance;
}

class Snapshot_Output extends $pb.GeneratedMessage {
  factory Snapshot_Output({
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

  Snapshot_Output._() : super();

  factory Snapshot_Output.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory Snapshot_Output.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Snapshot.Output',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'snapshot'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..a<$core.List<$core.int>>(
        2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY,
        protoName: 'File')
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  Snapshot_Output clone() => Snapshot_Output()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  Snapshot_Output copyWith(void Function(Snapshot_Output) updates) =>
      super.copyWith((message) => updates(message as Snapshot_Output))
          as Snapshot_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Snapshot_Output create() => Snapshot_Output._();

  Snapshot_Output createEmptyInstance() => create();

  static $pb.PbList<Snapshot_Output> createRepeated() =>
      $pb.PbList<Snapshot_Output>();

  @$core.pragma('dart2js:noInline')
  static Snapshot_Output getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<Snapshot_Output>(create);
  static Snapshot_Output? _defaultInstance;

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

/// snapshot
class Snapshot extends $pb.GeneratedMessage {
  factory Snapshot() => create();

  Snapshot._() : super();

  factory Snapshot.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);

  factory Snapshot.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Snapshot',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'snapshot'),
      createEmptyInstance: create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  Snapshot clone() => Snapshot()..mergeFromMessage(this);

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  Snapshot copyWith(void Function(Snapshot) updates) =>
      super.copyWith((message) => updates(message as Snapshot)) as Snapshot;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Snapshot create() => Snapshot._();

  Snapshot createEmptyInstance() => create();

  static $pb.PbList<Snapshot> createRepeated() => $pb.PbList<Snapshot>();

  @$core.pragma('dart2js:noInline')
  static Snapshot getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Snapshot>(create);
  static Snapshot? _defaultInstance;
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
