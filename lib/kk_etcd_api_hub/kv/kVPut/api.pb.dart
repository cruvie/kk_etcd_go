//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/kv/kVPut/api.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class KVPut_Input extends $pb.GeneratedMessage {
  factory KVPut_Input({
    $core.String? key,
    $core.String? value,
  }) {
    final $result = create();
    if (key != null) {
      $result.key = key;
    }
    if (value != null) {
      $result.value = value;
    }
    return $result;
  }
  KVPut_Input._() : super();
  factory KVPut_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVPut_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVPut.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'kVPut'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..aOS(2, _omitFieldNames ? '' : 'Value', protoName: 'Value')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVPut_Input clone() => KVPut_Input()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVPut_Input copyWith(void Function(KVPut_Input) updates) => super.copyWith((message) => updates(message as KVPut_Input)) as KVPut_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVPut_Input create() => KVPut_Input._();
  KVPut_Input createEmptyInstance() => create();
  static $pb.PbList<KVPut_Input> createRepeated() => $pb.PbList<KVPut_Input>();
  @$core.pragma('dart2js:noInline')
  static KVPut_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVPut_Input>(create);
  static KVPut_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get value => $_getSZ(1);
  @$pb.TagNumber(2)
  set value($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasValue() => $_has(1);
  @$pb.TagNumber(2)
  void clearValue() => $_clearField(2);
}

class KVPut_Output extends $pb.GeneratedMessage {
  factory KVPut_Output() => create();
  KVPut_Output._() : super();
  factory KVPut_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVPut_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVPut.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'kVPut'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVPut_Output clone() => KVPut_Output()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVPut_Output copyWith(void Function(KVPut_Output) updates) => super.copyWith((message) => updates(message as KVPut_Output)) as KVPut_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVPut_Output create() => KVPut_Output._();
  KVPut_Output createEmptyInstance() => create();
  static $pb.PbList<KVPut_Output> createRepeated() => $pb.PbList<KVPut_Output>();
  @$core.pragma('dart2js:noInline')
  static KVPut_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVPut_Output>(create);
  static KVPut_Output? _defaultInstance;
}

/// put kv
class KVPut extends $pb.GeneratedMessage {
  factory KVPut() => create();
  KVPut._() : super();
  factory KVPut.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVPut.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVPut', package: const $pb.PackageName(_omitMessageNames ? '' : 'kVPut'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  KVPut clone() => KVPut()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  KVPut copyWith(void Function(KVPut) updates) => super.copyWith((message) => updates(message as KVPut)) as KVPut;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVPut create() => KVPut._();
  KVPut createEmptyInstance() => create();
  static $pb.PbList<KVPut> createRepeated() => $pb.PbList<KVPut>();
  @$core.pragma('dart2js:noInline')
  static KVPut getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVPut>(create);
  static KVPut? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
