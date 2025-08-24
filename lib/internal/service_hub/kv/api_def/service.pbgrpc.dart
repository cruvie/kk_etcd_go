// This is a generated file - do not edit.
//
// Generated from internal/service_hub/kv/api_def/service.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'KVDel.pb.dart' as $0;
import 'KVGet.pb.dart' as $1;
import 'KVList.pb.dart' as $2;
import 'KVPut.pb.dart' as $3;

export 'service.pb.dart';

@$pb.GrpcServiceName('kk_etcd.KV')
class KVClient extends $grpc.Client {
  /// The hostname for this service.
  static const $core.String defaultHost = '';

  /// OAuth scopes needed for the client.
  static const $core.List<$core.String> oauthScopes = [
    '',
  ];

  KVClient(super.channel, {super.options, super.interceptors});

  $grpc.ResponseFuture<$0.KVDel_Output> kVDel(
    $0.KVDel_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$kVDel, request, options: options);
  }

  $grpc.ResponseFuture<$1.KVGet_Output> kVGet(
    $1.KVGet_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$kVGet, request, options: options);
  }

  $grpc.ResponseFuture<$2.KVList_Output> kVList(
    $2.KVList_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$kVList, request, options: options);
  }

  $grpc.ResponseFuture<$3.KVPut_Output> kVPut(
    $3.KVPut_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$kVPut, request, options: options);
  }

  // method descriptors

  static final _$kVDel = $grpc.ClientMethod<$0.KVDel_Input, $0.KVDel_Output>(
      '/kk_etcd.KV/KVDel',
      ($0.KVDel_Input value) => value.writeToBuffer(),
      $0.KVDel_Output.fromBuffer);
  static final _$kVGet = $grpc.ClientMethod<$1.KVGet_Input, $1.KVGet_Output>(
      '/kk_etcd.KV/KVGet',
      ($1.KVGet_Input value) => value.writeToBuffer(),
      $1.KVGet_Output.fromBuffer);
  static final _$kVList = $grpc.ClientMethod<$2.KVList_Input, $2.KVList_Output>(
      '/kk_etcd.KV/KVList',
      ($2.KVList_Input value) => value.writeToBuffer(),
      $2.KVList_Output.fromBuffer);
  static final _$kVPut = $grpc.ClientMethod<$3.KVPut_Input, $3.KVPut_Output>(
      '/kk_etcd.KV/KVPut',
      ($3.KVPut_Input value) => value.writeToBuffer(),
      $3.KVPut_Output.fromBuffer);
}

@$pb.GrpcServiceName('kk_etcd.KV')
abstract class KVServiceBase extends $grpc.Service {
  $core.String get $name => 'kk_etcd.KV';

  KVServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.KVDel_Input, $0.KVDel_Output>(
        'KVDel',
        kVDel_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.KVDel_Input.fromBuffer(value),
        ($0.KVDel_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.KVGet_Input, $1.KVGet_Output>(
        'KVGet',
        kVGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.KVGet_Input.fromBuffer(value),
        ($1.KVGet_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.KVList_Input, $2.KVList_Output>(
        'KVList',
        kVList_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.KVList_Input.fromBuffer(value),
        ($2.KVList_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.KVPut_Input, $3.KVPut_Output>(
        'KVPut',
        kVPut_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.KVPut_Input.fromBuffer(value),
        ($3.KVPut_Output value) => value.writeToBuffer()));
  }

  $async.Future<$0.KVDel_Output> kVDel_Pre(
      $grpc.ServiceCall $call, $async.Future<$0.KVDel_Input> $request) async {
    return kVDel($call, await $request);
  }

  $async.Future<$0.KVDel_Output> kVDel(
      $grpc.ServiceCall call, $0.KVDel_Input request);

  $async.Future<$1.KVGet_Output> kVGet_Pre(
      $grpc.ServiceCall $call, $async.Future<$1.KVGet_Input> $request) async {
    return kVGet($call, await $request);
  }

  $async.Future<$1.KVGet_Output> kVGet(
      $grpc.ServiceCall call, $1.KVGet_Input request);

  $async.Future<$2.KVList_Output> kVList_Pre(
      $grpc.ServiceCall $call, $async.Future<$2.KVList_Input> $request) async {
    return kVList($call, await $request);
  }

  $async.Future<$2.KVList_Output> kVList(
      $grpc.ServiceCall call, $2.KVList_Input request);

  $async.Future<$3.KVPut_Output> kVPut_Pre(
      $grpc.ServiceCall $call, $async.Future<$3.KVPut_Input> $request) async {
    return kVPut($call, await $request);
  }

  $async.Future<$3.KVPut_Output> kVPut(
      $grpc.ServiceCall call, $3.KVPut_Input request);
}
