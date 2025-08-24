// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_service_kk_etcd.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'pb_service_kk_etcd.pbenum.dart';
import 'pb_service_registration.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

export 'pb_service_kk_etcd.pbenum.dart';

class PBService extends $pb.GeneratedMessage {
  factory PBService({
    $0.PBServiceRegistration? serviceRegistration,
    PBService_ServiceStatus? status,
  }) {
    final result = create();
    if (serviceRegistration != null)
      result.serviceRegistration = serviceRegistration;
    if (status != null) result.status = status;
    return result;
  }

  PBService._();

  factory PBService.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBService.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBService',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..aOM<$0.PBServiceRegistration>(
        1, _omitFieldNames ? '' : 'ServiceRegistration',
        protoName: 'ServiceRegistration',
        subBuilder: $0.PBServiceRegistration.create)
    ..e<PBService_ServiceStatus>(
        5, _omitFieldNames ? '' : 'Status', $pb.PbFieldType.OE,
        protoName: 'Status',
        defaultOrMaker: PBService_ServiceStatus.UnKnown,
        valueOf: PBService_ServiceStatus.valueOf,
        enumValues: PBService_ServiceStatus.values)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBService clone() => PBService()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBService copyWith(void Function(PBService) updates) =>
      super.copyWith((message) => updates(message as PBService)) as PBService;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBService create() => PBService._();
  @$core.override
  PBService createEmptyInstance() => create();
  static $pb.PbList<PBService> createRepeated() => $pb.PbList<PBService>();
  @$core.pragma('dart2js:noInline')
  static PBService getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PBService>(create);
  static PBService? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBServiceRegistration get serviceRegistration => $_getN(0);
  @$pb.TagNumber(1)
  set serviceRegistration($0.PBServiceRegistration value) =>
      $_setField(1, value);
  @$pb.TagNumber(1)
  $core.bool hasServiceRegistration() => $_has(0);
  @$pb.TagNumber(1)
  void clearServiceRegistration() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBServiceRegistration ensureServiceRegistration() => $_ensure(0);

  @$pb.TagNumber(5)
  PBService_ServiceStatus get status => $_getN(1);
  @$pb.TagNumber(5)
  set status(PBService_ServiceStatus value) => $_setField(5, value);
  @$pb.TagNumber(5)
  $core.bool hasStatus() => $_has(1);
  @$pb.TagNumber(5)
  void clearStatus() => $_clearField(5);
}

class PBListService extends $pb.GeneratedMessage {
  factory PBListService({
    $core.Iterable<PBService>? listService,
  }) {
    final result = create();
    if (listService != null) result.listService.addAll(listService);
    return result;
  }

  PBListService._();

  factory PBListService.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory PBListService.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'PBListService',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'kk_etcd'),
      createEmptyInstance: create)
    ..pc<PBService>(1, _omitFieldNames ? '' : 'ListService', $pb.PbFieldType.PM,
        protoName: 'ListService', subBuilder: PBService.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListService clone() => PBListService()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  PBListService copyWith(void Function(PBListService) updates) =>
      super.copyWith((message) => updates(message as PBListService))
          as PBListService;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PBListService create() => PBListService._();
  @$core.override
  PBListService createEmptyInstance() => create();
  static $pb.PbList<PBListService> createRepeated() =>
      $pb.PbList<PBListService>();
  @$core.pragma('dart2js:noInline')
  static PBListService getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<PBListService>(create);
  static PBListService? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<PBService> get listService => $_getList(0);
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
