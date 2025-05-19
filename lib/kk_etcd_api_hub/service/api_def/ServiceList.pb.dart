//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/service/api_def/ServiceList.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../kk_etcd_models/pb_service_kk_etcd.pb.dart' as $0;
import '../../../kk_etcd_models/pb_service_registration.pbenum.dart' as $1;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class ServiceList_Input extends $pb.GeneratedMessage {
  factory ServiceList_Input({
    $1.PBServiceType? serviceType,
    $core.String? serviceName,
  }) {
    final $result = create();
    if (serviceType != null) {
      $result.serviceType = serviceType;
    }
    if (serviceName != null) {
      $result.serviceName = serviceName;
    }
    return $result;
  }
  ServiceList_Input._() : super();
  factory ServiceList_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServiceList_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ServiceList.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..e<$1.PBServiceType>(1, _omitFieldNames ? '' : 'ServiceType', $pb.PbFieldType.OE, protoName: 'ServiceType', defaultOrMaker: $1.PBServiceType.Unknown, valueOf: $1.PBServiceType.valueOf, enumValues: $1.PBServiceType.values)
    ..aOS(2, _omitFieldNames ? '' : 'ServiceName', protoName: 'ServiceName')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServiceList_Input clone() => ServiceList_Input()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServiceList_Input copyWith(void Function(ServiceList_Input) updates) => super.copyWith((message) => updates(message as ServiceList_Input)) as ServiceList_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ServiceList_Input create() => ServiceList_Input._();
  ServiceList_Input createEmptyInstance() => create();
  static $pb.PbList<ServiceList_Input> createRepeated() => $pb.PbList<ServiceList_Input>();
  @$core.pragma('dart2js:noInline')
  static ServiceList_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServiceList_Input>(create);
  static ServiceList_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $1.PBServiceType get serviceType => $_getN(0);
  @$pb.TagNumber(1)
  set serviceType($1.PBServiceType v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServiceType() => $_has(0);
  @$pb.TagNumber(1)
  void clearServiceType() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get serviceName => $_getSZ(1);
  @$pb.TagNumber(2)
  set serviceName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasServiceName() => $_has(1);
  @$pb.TagNumber(2)
  void clearServiceName() => $_clearField(2);
}

class ServiceList_Output extends $pb.GeneratedMessage {
  factory ServiceList_Output({
    $0.PBListService? serviceList,
  }) {
    final $result = create();
    if (serviceList != null) {
      $result.serviceList = serviceList;
    }
    return $result;
  }
  ServiceList_Output._() : super();
  factory ServiceList_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServiceList_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ServiceList.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOM<$0.PBListService>(1, _omitFieldNames ? '' : 'ServiceList', protoName: 'ServiceList', subBuilder: $0.PBListService.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServiceList_Output clone() => ServiceList_Output()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServiceList_Output copyWith(void Function(ServiceList_Output) updates) => super.copyWith((message) => updates(message as ServiceList_Output)) as ServiceList_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ServiceList_Output create() => ServiceList_Output._();
  ServiceList_Output createEmptyInstance() => create();
  static $pb.PbList<ServiceList_Output> createRepeated() => $pb.PbList<ServiceList_Output>();
  @$core.pragma('dart2js:noInline')
  static ServiceList_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServiceList_Output>(create);
  static ServiceList_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBListService get serviceList => $_getN(0);
  @$pb.TagNumber(1)
  set serviceList($0.PBListService v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServiceList() => $_has(0);
  @$pb.TagNumber(1)
  void clearServiceList() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBListService ensureServiceList() => $_ensure(0);
}

/// list server
class ServiceList extends $pb.GeneratedMessage {
  factory ServiceList() => create();
  ServiceList._() : super();
  factory ServiceList.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServiceList.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ServiceList', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServiceList clone() => ServiceList()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServiceList copyWith(void Function(ServiceList) updates) => super.copyWith((message) => updates(message as ServiceList)) as ServiceList;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ServiceList create() => ServiceList._();
  ServiceList createEmptyInstance() => create();
  static $pb.PbList<ServiceList> createRepeated() => $pb.PbList<ServiceList>();
  @$core.pragma('dart2js:noInline')
  static ServiceList getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServiceList>(create);
  static ServiceList? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
