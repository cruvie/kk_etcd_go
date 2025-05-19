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
	ServiceKey             = "kk_service/"
	ServiceRegistrationKey = ServiceKey + "registration/"
	ServiceAliveKey        = ServiceKey + "alive/"
)

func (x *PBServiceRegistration_PBCheckConfig) Check() error {
	if x.Type != PBServiceType_Http && x.Type != PBServiceType_Grpc {
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

func (x *PBServiceRegistration) Check() error {
	if x.ServiceName == "" {
		msg := "server name cannot be empty"
		return errors.New(msg)
	}
	if x.ServiceAddr == "" {
		msg := "server address cannot be empty"
		return errors.New(msg)
	}
	return x.CheckConfig.Check()
}

func (x *PBServiceRegistration) Key() string {
	return fmt.Sprintf("%s%s/%s/%s", ServiceRegistrationKey, x.ServiceType.String(), x.ServiceName, x.ServiceAddr)
}

func (x *PBServiceRegistration) Marshal() (string, error) {
	marshal, err := protojson.Marshal(x)
	return string(marshal), err
}

func (x *PBServiceRegistration) UnMarshal(data []byte) error {
	err := protojson.Unmarshal(data, x)
	return err
}
