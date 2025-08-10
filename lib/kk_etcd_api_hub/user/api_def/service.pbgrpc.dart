// This is a generated file - do not edit.
//
// Generated from kk_etcd_api_hub/user/api_def/service.proto.

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

import 'GetUser.pb.dart' as $0;
import 'Login.pb.dart' as $1;
import 'Logout.pb.dart' as $2;
import 'MyInfo.pb.dart' as $3;
import 'UserAdd.pb.dart' as $4;
import 'UserDelete.pb.dart' as $5;
import 'UserGrantRole.pb.dart' as $6;
import 'UserList.pb.dart' as $7;

export 'service.pb.dart';

@$pb.GrpcServiceName('api_def.User')
class UserClient extends $grpc.Client {
  /// The hostname for this service.
  static const $core.String defaultHost = '';

  /// OAuth scopes needed for the client.
  static const $core.List<$core.String> oauthScopes = [
    '',
  ];

  UserClient(super.channel, {super.options, super.interceptors});

  $grpc.ResponseFuture<$0.GetUser_Output> getUser(
    $0.GetUser_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$getUser, request, options: options);
  }

  $grpc.ResponseFuture<$1.Login_Output> login(
    $1.Login_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$login, request, options: options);
  }

  $grpc.ResponseFuture<$2.Logout_Output> logout(
    $2.Logout_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$logout, request, options: options);
  }

  $grpc.ResponseFuture<$3.MyInfo_Output> myInfo(
    $3.MyInfo_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$myInfo, request, options: options);
  }

  $grpc.ResponseFuture<$4.UserAdd_Output> userAdd(
    $4.UserAdd_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$userAdd, request, options: options);
  }

  $grpc.ResponseFuture<$5.UserDelete_Output> userDelete(
    $5.UserDelete_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$userDelete, request, options: options);
  }

  $grpc.ResponseFuture<$6.UserGrantRole_Output> userGrantRole(
    $6.UserGrantRole_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$userGrantRole, request, options: options);
  }

  $grpc.ResponseFuture<$7.UserList_Output> userList(
    $7.UserList_Input request, {
    $grpc.CallOptions? options,
  }) {
    return $createUnaryCall(_$userList, request, options: options);
  }

  // method descriptors

  static final _$getUser =
      $grpc.ClientMethod<$0.GetUser_Input, $0.GetUser_Output>(
          '/api_def.User/GetUser',
          ($0.GetUser_Input value) => value.writeToBuffer(),
          $0.GetUser_Output.fromBuffer);
  static final _$login = $grpc.ClientMethod<$1.Login_Input, $1.Login_Output>(
      '/api_def.User/Login',
      ($1.Login_Input value) => value.writeToBuffer(),
      $1.Login_Output.fromBuffer);
  static final _$logout = $grpc.ClientMethod<$2.Logout_Input, $2.Logout_Output>(
      '/api_def.User/Logout',
      ($2.Logout_Input value) => value.writeToBuffer(),
      $2.Logout_Output.fromBuffer);
  static final _$myInfo = $grpc.ClientMethod<$3.MyInfo_Input, $3.MyInfo_Output>(
      '/api_def.User/MyInfo',
      ($3.MyInfo_Input value) => value.writeToBuffer(),
      $3.MyInfo_Output.fromBuffer);
  static final _$userAdd =
      $grpc.ClientMethod<$4.UserAdd_Input, $4.UserAdd_Output>(
          '/api_def.User/UserAdd',
          ($4.UserAdd_Input value) => value.writeToBuffer(),
          $4.UserAdd_Output.fromBuffer);
  static final _$userDelete =
      $grpc.ClientMethod<$5.UserDelete_Input, $5.UserDelete_Output>(
          '/api_def.User/UserDelete',
          ($5.UserDelete_Input value) => value.writeToBuffer(),
          $5.UserDelete_Output.fromBuffer);
  static final _$userGrantRole =
      $grpc.ClientMethod<$6.UserGrantRole_Input, $6.UserGrantRole_Output>(
          '/api_def.User/UserGrantRole',
          ($6.UserGrantRole_Input value) => value.writeToBuffer(),
          $6.UserGrantRole_Output.fromBuffer);
  static final _$userList =
      $grpc.ClientMethod<$7.UserList_Input, $7.UserList_Output>(
          '/api_def.User/UserList',
          ($7.UserList_Input value) => value.writeToBuffer(),
          $7.UserList_Output.fromBuffer);
}

@$pb.GrpcServiceName('api_def.User')
abstract class UserServiceBase extends $grpc.Service {
  $core.String get $name => 'api_def.User';

  UserServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.GetUser_Input, $0.GetUser_Output>(
        'GetUser',
        getUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetUser_Input.fromBuffer(value),
        ($0.GetUser_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.Login_Input, $1.Login_Output>(
        'Login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Login_Input.fromBuffer(value),
        ($1.Login_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.Logout_Input, $2.Logout_Output>(
        'Logout',
        logout_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.Logout_Input.fromBuffer(value),
        ($2.Logout_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$3.MyInfo_Input, $3.MyInfo_Output>(
        'MyInfo',
        myInfo_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.MyInfo_Input.fromBuffer(value),
        ($3.MyInfo_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.UserAdd_Input, $4.UserAdd_Output>(
        'UserAdd',
        userAdd_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.UserAdd_Input.fromBuffer(value),
        ($4.UserAdd_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$5.UserDelete_Input, $5.UserDelete_Output>(
        'UserDelete',
        userDelete_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $5.UserDelete_Input.fromBuffer(value),
        ($5.UserDelete_Output value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$6.UserGrantRole_Input, $6.UserGrantRole_Output>(
            'UserGrantRole',
            userGrantRole_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $6.UserGrantRole_Input.fromBuffer(value),
            ($6.UserGrantRole_Output value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$7.UserList_Input, $7.UserList_Output>(
        'UserList',
        userList_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $7.UserList_Input.fromBuffer(value),
        ($7.UserList_Output value) => value.writeToBuffer()));
  }

  $async.Future<$0.GetUser_Output> getUser_Pre(
      $grpc.ServiceCall $call, $async.Future<$0.GetUser_Input> $request) async {
    return getUser($call, await $request);
  }

  $async.Future<$0.GetUser_Output> getUser(
      $grpc.ServiceCall call, $0.GetUser_Input request);

  $async.Future<$1.Login_Output> login_Pre(
      $grpc.ServiceCall $call, $async.Future<$1.Login_Input> $request) async {
    return login($call, await $request);
  }

  $async.Future<$1.Login_Output> login(
      $grpc.ServiceCall call, $1.Login_Input request);

  $async.Future<$2.Logout_Output> logout_Pre(
      $grpc.ServiceCall $call, $async.Future<$2.Logout_Input> $request) async {
    return logout($call, await $request);
  }

  $async.Future<$2.Logout_Output> logout(
      $grpc.ServiceCall call, $2.Logout_Input request);

  $async.Future<$3.MyInfo_Output> myInfo_Pre(
      $grpc.ServiceCall $call, $async.Future<$3.MyInfo_Input> $request) async {
    return myInfo($call, await $request);
  }

  $async.Future<$3.MyInfo_Output> myInfo(
      $grpc.ServiceCall call, $3.MyInfo_Input request);

  $async.Future<$4.UserAdd_Output> userAdd_Pre(
      $grpc.ServiceCall $call, $async.Future<$4.UserAdd_Input> $request) async {
    return userAdd($call, await $request);
  }

  $async.Future<$4.UserAdd_Output> userAdd(
      $grpc.ServiceCall call, $4.UserAdd_Input request);

  $async.Future<$5.UserDelete_Output> userDelete_Pre($grpc.ServiceCall $call,
      $async.Future<$5.UserDelete_Input> $request) async {
    return userDelete($call, await $request);
  }

  $async.Future<$5.UserDelete_Output> userDelete(
      $grpc.ServiceCall call, $5.UserDelete_Input request);

  $async.Future<$6.UserGrantRole_Output> userGrantRole_Pre(
      $grpc.ServiceCall $call,
      $async.Future<$6.UserGrantRole_Input> $request) async {
    return userGrantRole($call, await $request);
  }

  $async.Future<$6.UserGrantRole_Output> userGrantRole(
      $grpc.ServiceCall call, $6.UserGrantRole_Input request);

  $async.Future<$7.UserList_Output> userList_Pre($grpc.ServiceCall $call,
      $async.Future<$7.UserList_Input> $request) async {
    return userList($call, await $request);
  }

  $async.Future<$7.UserList_Output> userList(
      $grpc.ServiceCall call, $7.UserList_Input request);
}
