//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/backup/api_def/AllKVsRestore.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class AllKVsRestore_Input extends $pb.GeneratedMessage {
  factory AllKVsRestore_Input({
    $core.List<$core.int>? file,
  }) {
    final $result = create();
    if (file != null) {
      $result.file = file;
    }
    return $result;
  }
  AllKVsRestore_Input._() : super();
  factory AllKVsRestore_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AllKVsRestore_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AllKVsRestore.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(2, _omitFieldNames ? '' : 'File', $pb.PbFieldType.OY, protoName: 'File')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsRestore_Input clone() => AllKVsRestore_Input()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsRestore_Input copyWith(void Function(AllKVsRestore_Input) updates) => super.copyWith((message) => updates(message as AllKVsRestore_Input)) as AllKVsRestore_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsRestore_Input create() => AllKVsRestore_Input._();
  AllKVsRestore_Input createEmptyInstance() => create();
  static $pb.PbList<AllKVsRestore_Input> createRepeated() => $pb.PbList<AllKVsRestore_Input>();
  @$core.pragma('dart2js:noInline')
  static AllKVsRestore_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AllKVsRestore_Input>(create);
  static AllKVsRestore_Input? _defaultInstance;

  @$pb.TagNumber(2)
  $core.List<$core.int> get file => $_getN(0);
  @$pb.TagNumber(2)
  set file($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(0);
  @$pb.TagNumber(2)
  void clearFile() => $_clearField(2);
}

class AllKVsRestore_Output extends $pb.GeneratedMessage {
  factory AllKVsRestore_Output() => create();
  AllKVsRestore_Output._() : super();
  factory AllKVsRestore_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AllKVsRestore_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AllKVsRestore.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsRestore_Output clone() => AllKVsRestore_Output()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsRestore_Output copyWith(void Function(AllKVsRestore_Output) updates) => super.copyWith((message) => updates(message as AllKVsRestore_Output)) as AllKVsRestore_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsRestore_Output create() => AllKVsRestore_Output._();
  AllKVsRestore_Output createEmptyInstance() => create();
  static $pb.PbList<AllKVsRestore_Output> createRepeated() => $pb.PbList<AllKVsRestore_Output>();
  @$core.pragma('dart2js:noInline')
  static AllKVsRestore_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AllKVsRestore_Output>(create);
  static AllKVsRestore_Output? _defaultInstance;
}

/// all kvs restore
class AllKVsRestore extends $pb.GeneratedMessage {
  factory AllKVsRestore() => create();
  AllKVsRestore._() : super();
  factory AllKVsRestore.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AllKVsRestore.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AllKVsRestore', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsRestore clone() => AllKVsRestore()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  AllKVsRestore copyWith(void Function(AllKVsRestore) updates) => super.copyWith((message) => updates(message as AllKVsRestore)) as AllKVsRestore;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AllKVsRestore create() => AllKVsRestore._();
  AllKVsRestore createEmptyInstance() => create();
  static $pb.PbList<AllKVsRestore> createRepeated() => $pb.PbList<AllKVsRestore>();
  @$core.pragma('dart2js:noInline')
  static AllKVsRestore getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AllKVsRestore>(create);
  static AllKVsRestore? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
