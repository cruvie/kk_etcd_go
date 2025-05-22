//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_service_kk_etcd.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'pb_service_kk_etcd.pbenum.dart';
import 'pb_service_registration.pb.dart' as $1;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

export 'pb_service_kk_etcd.pbenum.dart';

class PBService extends $pb.GeneratedMessage {
  factory PBService({
    $1.PBServiceRegistration? serviceRegistration,
    PBService_ServiceStatus? status,
  }) {
    final $result = create();
    if (serviceRegistration != null) {
      $result.serviceRegistration = serviceRegistration;
    }
    if (status != null) {
      $result.status = status;
    }
    return $result;
  }
  PBService._() : super();
  factory PBService.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBService.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBService', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..aOM<$1.PBServiceRegistration>(1, _omitFieldNames ? '' : 'ServiceRegistration', protoName: 'ServiceRegistration', subBuilder: $1.PBServiceRegistration.create)
    ..e<PBService_ServiceStatus>(5, _omitFieldNames ? '' : 'Status', $pb.PbFieldType.OE, protoName: 'Status', defaultOrMaker: PBService_ServiceStatus.UnKnown, valueOf: PBService_ServiceStatus.valueOf, enumValues: PBService_ServiceStatus.values)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBService clone() => PBService()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBService copyWith(void Function(PBService) updates) => super.copyWith((message) => updates(message as PBService)) as PBService;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBService create() => PBService._();
  PBService createEmptyInstance() => create();
  static $pb.PbList<PBService> createRepeated() => $pb.PbList<PBService>();
  @$core.pragma('dart2js:noInline')
  static PBService getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBService>(create);
  static PBService? _defaultInstance;

  @$pb.TagNumber(1)
  $1.PBServiceRegistration get serviceRegistration => $_getN(0);
  @$pb.TagNumber(1)
  set serviceRegistration($1.PBServiceRegistration v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServiceRegistration() => $_has(0);
  @$pb.TagNumber(1)
  void clearServiceRegistration() => $_clearField(1);
  @$pb.TagNumber(1)
  $1.PBServiceRegistration ensureServiceRegistration() => $_ensure(0);

  @$pb.TagNumber(5)
  PBService_ServiceStatus get status => $_getN(1);
  @$pb.TagNumber(5)
  set status(PBService_ServiceStatus v) { $_setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasStatus() => $_has(1);
  @$pb.TagNumber(5)
  void clearStatus() => $_clearField(5);
}

class PBListService extends $pb.GeneratedMessage {
  factory PBListService({
    $core.Iterable<PBService>? listService,
  }) {
    final $result = create();
    if (listService != null) {
      $result.listService.addAll(listService);
    }
    return $result;
  }
  PBListService._() : super();
  factory PBListService.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PBListService.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PBListService', package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd_models'), createEmptyInstance: create)
    ..pc<PBService>(1, _omitFieldNames ? '' : 'ListService', $pb.PbFieldType.PM, protoName: 'ListService', subBuilder: PBService.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListService clone() => PBListService()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListService copyWith(void Function(PBListService) updates) => super.copyWith((message) => updates(message as PBListService)) as PBListService;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListService create() => PBListService._();
  PBListService createEmptyInstance() => create();
  static $pb.PbList<PBListService> createRepeated() => $pb.PbList<PBListService>();
  @$core.pragma('dart2js:noInline')
  static PBListService getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBListService>(create);
  static PBListService? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<PBService> get listService => $_getList(0);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
