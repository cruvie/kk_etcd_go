// This is a generated file - do not edit.
//
// Generated from internal/service_hub/role/api_def/service.proto.

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

import 'RoleAdd.pb.dart' as $0;
import 'RoleDelete.pb.dart' as $1;
import 'RoleGet.pb.dart' as $2;
import 'RoleGrantPermission.pb.dart' as $3;
import 'RoleList.pb.dart' as $4;
import 'RoleRevokePermission.pb.dart' as $5;

export 'service.pb.dart';

@$pb.GrpcServiceName('api_def.Role')
class RoleClient extends $grpc.Client {
  /// The hostname for this service.
  static const $core.String defaultHost = '';

  /// OAuth scopes needed for the client.
  static const $core.List<$core.String> oauthScopes = [
    '',
  ];

  RoleClient(super.channel, {super.options, super.interceptors});

  $grpc.ResponseFuture<$0.RoleAdd_Output> roleAdd(
    $0.RoleAdd_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$roleAdd, request, options: options);
  }

  $grpc.ResponseFuture<$1.RoleDelete_Output> roleDelete(
    $1.RoleDelete_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$roleDelete, request, options: options);
  }

  $grpc.ResponseFuture<$2.RoleGet_Output> roleGet(
    $2.RoleGet_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$roleGet, request, options: options);
  }

  $grpc.ResponseFuture<$3.RoleGrantPermission_Output> roleGrantPermission(
    $3.RoleGrantPermission_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$roleGrantPermission, request, options: options);
  }

  $grpc.ResponseFuture<$4.RoleList_Output> roleList(
    $4.RoleList_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$roleList, request, options: options);
  }

  $grpc.ResponseFuture<$5.RoleRevokePermission_Output> roleRevokePermission(
    $5.RoleRevokePermission_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$roleRevokePermission, request, options: options);
  }

  // method descriptors

  static final _$roleAdd =
      $grpc.ClientMethod<$0.RoleAdd_Input, $0.RoleAdd_Output>(
          '/api_def.Role/RoleAdd',
          ($0.RoleAdd_Input value) => value.writeToBuffer(),
          $0.RoleAdd_Output.fromBuffer);
  static final _$roleDelete =
      $grpc.ClientMethod<$1.RoleDelete_Input, $1.RoleDelete_Output>(
          '/api_def.Role/RoleDelete',
          ($1.RoleDelete_Input value) => value.writeToBuffer(),
          $1.RoleDelete_Output.fromBuffer);
  static final _$roleGet =
      $grpc.ClientMethod<$2.RoleGet_Input, $2.RoleGet_Output>(
          '/api_def.Role/RoleGet',
          ($2.RoleGet_Input value) => value.writeToBuffer(),
          $2.RoleGet_Output.fromBuffer);
  static final _$roleGrantPermission = $grpc.ClientMethod<
          $3.RoleGrantPermission_Input, $3.RoleGrantPermission_Output>(
      '/api_def.Role/RoleGrantPermission',
      ($3.RoleGrantPermission_Input value) => value.writeToBuffer(),
      $3.RoleGrantPermission_Output.fromBuffer);
  static final _$roleList =
      $grpc.ClientMethod<$4.RoleList_Input, $4.RoleList_Output>(
          '/api_def.Role/RoleList',
          ($4.RoleList_Input value) => value.writeToBuffer(),
          $4.RoleList_Output.fromBuffer);
  static final _$roleRevokePermission = $grpc.ClientMethod<
          $5.RoleRevokePermission_Input, $5.RoleRevokePermission_Output>(
      '/api_def.Role/RoleRevokePermission',
      ($5.RoleRevokePermission_Input value) => value.writeToBuffer(),
      $5.RoleRevokePermission_Output.fromBuffer);
}

@$pb.GrpcServiceName('api_def.Role')
abstract class RoleServiceBase extends $grpc.Service {
  $core.String get $name => 'api_def.Role';

  RoleServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.RoleAdd_Input, $0.RoleAdd_Output>(
        'RoleAdd',
        roleAdd_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.RoleAdd_Input.fromBuffer(value),
        ($0.RoleAdd_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.RoleDelete_Input, $1.RoleDelete_Output>(
        'RoleDelete',
        roleDelete_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.RoleDelete_Input.fromBuffer(value),
        ($1.RoleDelete_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.RoleGet_Input, $2.RoleGet_Output>(
        'RoleGet',
        roleGet_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.RoleGet_Input.fromBuffer(value),
        ($2.RoleGet_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.RoleGrantPermission_Input,
            $3.RoleGrantPermission_Output>(
        'RoleGrantPermission',
        roleGrantPermission_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $3.RoleGrantPermission_Input.fromBuffer(value),
        ($3.RoleGrantPermission_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.RoleList_Input, $4.RoleList_Output>(
        'RoleList',
        roleList_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.RoleList_Input.fromBuffer(value),
        ($4.RoleList_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.RoleRevokePermission_Input,
            $5.RoleRevokePermission_Output>(
        'RoleRevokePermission',
        roleRevokePermission_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $5.RoleRevokePermission_Input.fromBuffer(value),
        ($5.RoleRevokePermission_Output value) => value.writeToBuffer()));
  }

  $async.Future<$0.RoleAdd_Output> roleAdd_Pre(
      $grpc.ServiceCall $call, $async.Future<$0.RoleAdd_Input> $request) async {
    return roleAdd($call, await $request);
  }

  $async.Future<$0.RoleAdd_Output> roleAdd(
      $grpc.ServiceCall call, $0.RoleAdd_Input request);

  $async.Future<$1.RoleDelete_Output> roleDelete_Pre($grpc.ServiceCall $call,
      $async.Future<$1.RoleDelete_Input> $request) async {
    return roleDelete($call, await $request);
  }

  $async.Future<$1.RoleDelete_Output> roleDelete(
      $grpc.ServiceCall call, $1.RoleDelete_Input request);

  $async.Future<$2.RoleGet_Output> roleGet_Pre(
      $grpc.ServiceCall $call, $async.Future<$2.RoleGet_Input> $request) async {
    return roleGet($call, await $request);
  }

  $async.Future<$2.RoleGet_Output> roleGet(
      $grpc.ServiceCall call, $2.RoleGet_Input request);

  $async.Future<$3.RoleGrantPermission_Output> roleGrantPermission_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$3.RoleGrantPermission_Input> $request) async {
    return roleGrantPermission($call, await $request);
  }

  $async.Future<$3.RoleGrantPermission_Output> roleGrantPermission(
      $grpc.ServiceCall call, $3.RoleGrantPermission_Input request);

  $async.Future<$4.RoleList_Output> roleList_Pre($grpc.ServiceCall $call,
      $async.Future<$4.RoleList_Input> $request) async {
    return roleList($call, await $request);
  }

  $async.Future<$4.RoleList_Output> roleList(
      $grpc.ServiceCall call, $4.RoleList_Input request);

  $async.Future<$5.RoleRevokePermission_Output> roleRevokePermission_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$5.RoleRevokePermission_Input> $request) async {
    return roleRevokePermission($call, await $request);
  }

  $async.Future<$5.RoleRevokePermission_Output> roleRevokePermission(
      $grpc.ServiceCall call, $5.RoleRevokePermission_Input request);
}
