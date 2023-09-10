///
//  Generated code. Do not modify.
//  source: pb_service.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PBService extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'PBService', package: const $pb.PackageName(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'models'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names')
        ? ''
        : 'ServiceName', protoName: 'ServiceName')..aOS(2,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'ServiceAddr', protoName: 'ServiceAddr')
    ..hasRequiredFields = false
  ;

  PBService._() : super();

  factory PBService({
    $core.String? serviceName,
    $core.String? serviceAddr,
  }) {
    final _result = create();
    if (serviceName != null) {
      _result.serviceName = serviceName;
    }
    if (serviceAddr != null) {
      _result.serviceAddr = serviceAddr;
    }
    return _result;
  }

  factory PBService.fromBuffer($core.List<$core.int> i,
      [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()
        ..mergeFromBuffer(i, r);

  factory PBService.fromJson($core.String i,
      [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()
        ..mergeFromJson(i, r);

  @$core.Deprecated(
      'Using this can add significant overhead to your binary. '
          'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
          'Will be removed in next major version')
  PBService clone() =>
      PBService()
        ..mergeFromMessage(this);

  @$core.Deprecated(
      'Using this can add significant overhead to your binary. '
          'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
          'Will be removed in next major version')
  PBService copyWith(void Function(PBService) updates) =>
      super.copyWith((message) =>
          updates(
              message as PBService)) as PBService; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBService create() => PBService._();

  PBService createEmptyInstance() => create();

  static $pb.PbList<PBService> createRepeated() => $pb.PbList<PBService>();

  @$core.pragma('dart2js:noInline')
  static PBService getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBService>(create);
  static PBService? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get serviceName => $_getSZ(0);

  @$pb.TagNumber(1)
  set serviceName($core.String v) {
    $_setString(0, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasServiceName() => $_has(0);

  @$pb.TagNumber(1)
  void clearServiceName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get serviceAddr => $_getSZ(1);

  @$pb.TagNumber(2)
  set serviceAddr($core.String v) {
    $_setString(1, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasServiceAddr() => $_has(1);

  @$pb.TagNumber(2)
  void clearServiceAddr() => clearField(2);
}

class PBListService extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'PBListService', package: const $pb.PackageName(
      const $core.bool.fromEnvironment('protobuf.omit_message_names')
          ? ''
          : 'models'), createEmptyInstance: create)
    ..pc<PBService>(1,
        const $core.bool.fromEnvironment('protobuf.omit_field_names')
            ? ''
            : 'ListService', $pb.PbFieldType.PM, protoName: 'ListService',
        subBuilder: PBService.create)
    ..hasRequiredFields = false
  ;

  PBListService._() : super();

  factory PBListService({
    $core.Iterable<PBService>? listService,
  }) {
    final _result = create();
    if (listService != null) {
      _result.listService.addAll(listService);
    }
    return _result;
  }

  factory PBListService.fromBuffer($core.List<$core.int> i,
      [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()
        ..mergeFromBuffer(i, r);

  factory PBListService.fromJson($core.String i,
      [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()
        ..mergeFromJson(i, r);

  @$core.Deprecated(
      'Using this can add significant overhead to your binary. '
          'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
          'Will be removed in next major version')
  PBListService clone() =>
      PBListService()
        ..mergeFromMessage(this);

  @$core.Deprecated(
      'Using this can add significant overhead to your binary. '
          'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
          'Will be removed in next major version')
  PBListService copyWith(void Function(PBListService) updates) =>
      super.copyWith((message) =>
          updates(
              message as PBListService)) as PBListService; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListService create() => PBListService._();

  PBListService createEmptyInstance() => create();

  static $pb.PbList<PBListService> createRepeated() =>
      $pb.PbList<PBListService>();

  @$core.pragma('dart2js:noInline')
  static PBListService getDefault() =>
      _defaultInstance ??=
          $pb.GeneratedMessage.$_defaultFor<PBListService>(create);
  static PBListService? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<PBService> get listService => $_getList(0);
}

