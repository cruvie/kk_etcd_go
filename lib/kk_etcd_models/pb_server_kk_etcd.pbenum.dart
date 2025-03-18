//
//  Generated code. Do not modify.
//  source: kk_etcd_models/pb_server_kk_etcd.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class PBServer_ServerStatus extends $pb.ProtobufEnum {
  static const PBServer_ServerStatus UnKnown =
      PBServer_ServerStatus._(0, _omitEnumNames ? '' : 'UnKnown');
  static const PBServer_ServerStatus Running =
      PBServer_ServerStatus._(1, _omitEnumNames ? '' : 'Running');
  static const PBServer_ServerStatus Stop =
      PBServer_ServerStatus._(2, _omitEnumNames ? '' : 'Stop');
  static const PBServer_ServerStatus Init =
      PBServer_ServerStatus._(3, _omitEnumNames ? '' : 'Init');

  static const $core.List<PBServer_ServerStatus> values =
      <PBServer_ServerStatus>[
    UnKnown,
    Running,
    Stop,
    Init,
  ];

  static final $core.Map<$core.int, PBServer_ServerStatus> _byValue =
      $pb.ProtobufEnum.initByValue(values);

  static PBServer_ServerStatus? valueOf($core.int value) => _byValue[value];

  const PBServer_ServerStatus._($core.int v, $core.String n) : super(v, n);
}

const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
