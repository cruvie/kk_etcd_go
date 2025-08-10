// This is a generated file - do not edit.
//
// Generated from kk_etcd_models/pb_service_registration.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PBServiceType extends $pb.ProtobufEnum {
  static const PBServiceType Unknown =
      PBServiceType._(0, _omitEnumNames ? '' : 'Unknown');
  static const PBServiceType Http =
      PBServiceType._(1, _omitEnumNames ? '' : 'Http');
  static const PBServiceType Grpc =
      PBServiceType._(2, _omitEnumNames ? '' : 'Grpc');

  static const $core.List<PBServiceType> values = <PBServiceType>[
    Unknown,
    Http,
    Grpc,
  ];

  static final $core.List<PBServiceType?> _byValue =
      $pb.ProtobufEnum.$_initByValueList(values, 2);
  static PBServiceType? valueOf($core.int value) =>
      value < 0 || value >= _byValue.length ? null : _byValue[value];

  const PBServiceType._(super.value, super.name);
}

const $core.bool _omitEnumNames =
    $core.bool.fromEnvironment('protobuf.omit_enum_names');
