package kk_etcd_models

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

const (
	HealthCheckPath = "/KKHealthCheck"
)

const (
	ServerKey             = "kk_server/"
	ServerRegistrationKey = ServerKey + "registration/"
	ServerAliveKey        = ServerKey + "alive/"
)

func (x *PBServerRegistration_PBCheckConfig) Check() error {
	if x.Type != PBServerType_Http && x.Type != PBServerType_Grpc {
		msg := "check type must be Grpc or Http"
		return errors.New(msg)
	}
	if x.Timeout.AsDuration() < 1*time.Second {
		msg := "check timeout must grater than 1 second"
		return errors.New(msg)
	}
	if x.Interval.AsDuration() < x.Timeout.AsDuration() {
		msg := "check interval must grater than Timeout"
		return errors.New(msg)
	}
	if x.Addr == "" {
		msg := "check addr cannot be empty"
		return errors.New(msg)
	}
	return nil
}

func (x *PBServerRegistration) Check() error {
	if x.ServerName == "" {
		msg := "server name cannot be empty"
		return errors.New(msg)
	}
	if x.ServerAddr == "" {
		msg := "server address cannot be empty"
		return errors.New(msg)
	}
	return x.CheckConfig.Check()
}

func (x *PBServerRegistration) Key() string {
	return fmt.Sprintf("%s%s/%s/%s", ServerRegistrationKey, x.ServerType.String(), x.ServerName, x.ServerAddr)
}

func (x *PBServerRegistration) Marshal() (string, error) {
	marshal, err := protojson.Marshal(x)
	return string(marshal), err
}

func (x *PBServerRegistration) UnMarshal(data []byte) error {
	err := protojson.Unmarshal(data, x)
	return err
}
