//
//  Generated code. Do not modify.
//  source: kk_etcd_api_hub/server/api_def/ServerList.proto
//
// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../../kk_etcd_models/pb_server_kk_etcd.pb.dart' as $0;

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

class ServerList_Input extends $pb.GeneratedMessage {
  factory ServerList_Input({
    $core.String? serverType,
    $core.String? serverName,
  }) {
    final $result = create();
    if (serverType != null) {
      $result.serverType = serverType;
    }
    if (serverName != null) {
      $result.serverName = serverName;
    }
    return $result;
  }
  ServerList_Input._() : super();
  factory ServerList_Input.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerList_Input.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ServerList.Input', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'ServerType', protoName: 'ServerType')
    ..aOS(2, _omitFieldNames ? '' : 'ServerName', protoName: 'ServerName')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServerList_Input clone() => ServerList_Input()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServerList_Input copyWith(void Function(ServerList_Input) updates) => super.copyWith((message) => updates(message as ServerList_Input)) as ServerList_Input;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ServerList_Input create() => ServerList_Input._();
  ServerList_Input createEmptyInstance() => create();
  static $pb.PbList<ServerList_Input> createRepeated() => $pb.PbList<ServerList_Input>();
  @$core.pragma('dart2js:noInline')
  static ServerList_Input getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerList_Input>(create);
  static ServerList_Input? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get serverType => $_getSZ(0);
  @$pb.TagNumber(1)
  set serverType($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasServerType() => $_has(0);
  @$pb.TagNumber(1)
  void clearServerType() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get serverName => $_getSZ(1);
  @$pb.TagNumber(2)
  set serverName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasServerName() => $_has(1);
  @$pb.TagNumber(2)
  void clearServerName() => $_clearField(2);
}

class ServerList_Output extends $pb.GeneratedMessage {
  factory ServerList_Output({
    $0.PBListServer? serverList,
  }) {
    final $result = create();
    if (serverList != null) {
      $result.serverList = serverList;
    }
    return $result;
  }
  ServerList_Output._() : super();
  factory ServerList_Output.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerList_Output.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ServerList.Output', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..aOM<$0.PBListServer>(1, _omitFieldNames ? '' : 'ServerList', protoName: 'ServerList', subBuilder: $0.PBListServer.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServerList_Output clone() => ServerList_Output()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServerList_Output copyWith(void Function(ServerList_Output) updates) => super.copyWith((message) => updates(message as ServerList_Output)) as ServerList_Output;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ServerList_Output create() => ServerList_Output._();
  ServerList_Output createEmptyInstance() => create();
  static $pb.PbList<ServerList_Output> createRepeated() => $pb.PbList<ServerList_Output>();
  @$core.pragma('dart2js:noInline')
  static ServerList_Output getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerList_Output>(create);
  static ServerList_Output? _defaultInstance;

  @$pb.TagNumber(1)
  $0.PBListServer get serverList => $_getN(0);
  @$pb.TagNumber(1)
  set serverList($0.PBListServer v) { $_setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasServerList() => $_has(0);
  @$pb.TagNumber(1)
  void clearServerList() => $_clearField(1);
  @$pb.TagNumber(1)
  $0.PBListServer ensureServerList() => $_ensure(0);
}

/// list server
class ServerList extends $pb.GeneratedMessage {
  factory ServerList() => create();
  ServerList._() : super();
  factory ServerList.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerList.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ServerList', package: const $pb.PackageName(_omitMessageNames ? '' : 'api_def'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ServerList clone() => ServerList()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ServerList copyWith(void Function(ServerList) updates) => super.copyWith((message) => updates(message as ServerList)) as ServerList;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ServerList create() => ServerList._();
  ServerList createEmptyInstance() => create();
  static $pb.PbList<ServerList> createRepeated() => $pb.PbList<ServerList>();
  @$core.pragma('dart2js:noInline')
  static ServerList getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerList>(create);
  static ServerList? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
