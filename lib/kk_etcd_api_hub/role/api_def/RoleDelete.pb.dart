//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/role/api_def/RoleDelete.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class RoleDelete_Input extends $pb.GeneratedMessage {
  factory RoleDelete_Input({
    $core.String? name,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  RoleDelete_Input._() : super();
  factory RoleDelete_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleDelete_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleDelete.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'Name', protoName: 'Name')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleDelete_Input clone() => RoleDelete_Input()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleDelete_Input copyWith(void Function(RoleDelete_Input) updates) => super.copyWith((message) => updates(message as RoleDelete_Input)) as RoleDelete_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleDelete_Input create() => RoleDelete_Input._();
  RoleDelete_Input createEmptyInstance() => create();
  static $pb.PbList<RoleDelete_Input> createRepeated() => $pb.PbList<RoleDelete_Input>();
  @$core.pragma('dart2js:noInline')
  static RoleDelete_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleDelete_Input>(create);
  static RoleDelete_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);
}

class RoleDelete_Output extends $pb.GeneratedMessage {
  factory RoleDelete_Output() => create();
  RoleDelete_Output._() : super();
  factory RoleDelete_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleDelete_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleDelete.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleDelete_Output clone() => RoleDelete_Output()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleDelete_Output copyWith(void Function(RoleDelete_Output) updates) => super.copyWith((message) => updates(message as RoleDelete_Output)) as RoleDelete_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleDelete_Output create() => RoleDelete_Output._();
  RoleDelete_Output createEmptyInstance() => create();
  static $pb.PbList<RoleDelete_Output> createRepeated() => $pb.PbList<RoleDelete_Output>();
  @$core.pragma('dart2js:noInline')
  static RoleDelete_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleDelete_Output>(create);
  static RoleDelete_Output? _defaultInstance;
}

/// delete role
class RoleDelete extends $pb.GeneratedMessage {
  factory RoleDelete() => create();
  RoleDelete._() : super();
  factory RoleDelete.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RoleDelete.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'RoleDelete', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  RoleDelete clone() => RoleDelete()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  RoleDelete copyWith(void Function(RoleDelete) updates) => super.copyWith((message) => updates(message as RoleDelete)) as RoleDelete;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static RoleDelete create() => RoleDelete._();
  RoleDelete createEmptyInstance() => create();
  static $pb.PbList<RoleDelete> createRepeated() => $pb.PbList<RoleDelete>();
  @$core.pragma('dart2js:noInline')
  static RoleDelete getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RoleDelete>(create);
  static RoleDelete? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
