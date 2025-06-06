//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/kv/api_def/KVDel.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class KVDel_Input extends $pb.GeneratedMessage {
  factory KVDel_Input({
    $core.String? key,
  }) {
    final $result = create();
    if (key != null) {
      $result.key = key;
    }
    return $result;
  }
  KVDel_Input._() : super();
  factory KVDel_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVDel_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVDel.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Key', protoName: 'Key')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVDel_Input clone() => KVDel_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVDel_Input copyWith(void Function(KVDel_Input) updates) => super.copyWith((message) => updates(message as KVDel_Input)) as KVDel_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVDel_Input create() => KVDel_Input._();
  KVDel_Input createEmptyInstance() => create();
  static $pb.PbList<KVDel_Input> createRepeated() => $pb.PbList<KVDel_Input>();
  @$core.pragma('dart2js:noInline')
  static KVDel_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVDel_Input>(create);
  static KVDel_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get key => $_getSZ(0);
  @$pb.TagNumber(1)
  set key($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasKey() => $_has(0);
  @$pb.TagNumber(1)
  void clearKey() => $_clearField(1);
}

class KVDel_Output extends $pb.GeneratedMessage {
  factory KVDel_Output() => create();
  KVDel_Output._() : super();
  factory KVDel_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVDel_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVDel.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVDel_Output clone() => KVDel_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVDel_Output copyWith(void Function(KVDel_Output) updates) => super.copyWith((message) => updates(message as KVDel_Output)) as KVDel_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVDel_Output create() => KVDel_Output._();
  KVDel_Output createEmptyInstance() => create();
  static $pb.PbList<KVDel_Output> createRepeated() => $pb.PbList<KVDel_Output>();
  @$core.pragma('dart2js:noInline')
  static KVDel_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVDel_Output>(create);
  static KVDel_Output? _defaultInstance;
}

/// del kv
class KVDel extends $pb.GeneratedMessage {
  factory KVDel() => create();
  KVDel._() : super();
  factory KVDel.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory KVDel.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'KVDel', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVDel clone() => KVDel()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  KVDel copyWith(void Function(KVDel) updates) => super.copyWith((message) => updates(message as KVDel)) as KVDel;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static KVDel create() => KVDel._();
  KVDel createEmptyInstance() => create();
  static $pb.PbList<KVDel> createRepeated() => $pb.PbList<KVDel>();
  @$core.pragma('dart2js:noInline')
  static KVDel getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<KVDel>(create);
  static KVDel? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
