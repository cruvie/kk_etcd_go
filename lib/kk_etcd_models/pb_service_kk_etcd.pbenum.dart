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

class PBService_ServiceStatus extends $pb.ProtobufEnum {
  static const PBService_ServiceStatus UnKnown =
      PBService_ServiceStatus._(0, _omitEnumNames ? '' : 'UnKnown');
  static const PBService_ServiceStatus Running =
      PBService_ServiceStatus._(1, _omitEnumNames ? '' : 'Running');
  static const PBService_ServiceStatus Stop =
      PBService_ServiceStatus._(2, _omitEnumNames ? '' : 'Stop');

  static const $core.List<PBService_ServiceStatus> values =
      <PBService_ServiceStatus>[
    UnKnown,
    Running,
    Stop,
  ];

  static final $core.List<PBService_ServiceStatus?> _byValue =
      $pb.ProtobufEnum.$_initByValueList(values, 2);
  static PBService_ServiceStatus? valueOf($core.int value) =>
      value < 0 || value >= _byValue.length ? null : _byValue[value];

  const PBService_ServiceStatus._(super.value, super.name);
}

const $core.bool _omitEnumNames =
    $core.bool.fromEnvironment('protobuf.omit_enum_names');
