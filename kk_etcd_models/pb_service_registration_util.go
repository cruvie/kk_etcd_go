package kk_etcd_models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
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

// UniqueKey represents every service registration.
func (x *PBServiceRegistration) UniqueKey() string {
	return fmt.Sprintf("%s%s/%s/%s", ServiceRegistrationKey, x.ServiceType.String(), x.ServiceName, x.ServiceAddr)
}

func (x *PBServiceRegistration) BuildFromUniqueKey(uniqueKey string) error {
	// kk_service/registration/Http/test_http_8843/192.168.0.100:8843
	s := strings.TrimPrefix(uniqueKey, ServiceRegistrationKey)
	split := strings.Split(s, "/")
	if len(split) != 3 {
		return errors.New("invalid unique key")
	}
	x.ServiceType = PBServiceType(PBServiceType_value[split[0]])
	x.ServiceName = split[1]
	x.ServiceAddr = split[2]
	return nil
}

// HubKey represents service registrations with same ServiceName and ServiceType
func (x *PBServiceRegistration) HubKey() string {
	return fmt.Sprintf("%s-%s", x.ServiceType.String(), x.ServiceName)
}

func (x PBServiceType) HubKey(serviceName string) string {
	return fmt.Sprintf("%s-%s", x.String(), serviceName)
}

func (x *PBServiceRegistration) Marshal() (string, error) {
	marshal, err := protojson.Marshal(x)
	return string(marshal), err
}

func (x *PBServiceRegistration) UnMarshal(data []byte) error {
	err := protojson.Unmarshal(data, x)
	return err
}

func (x *PBServiceRegistration) Equal(item *PBServiceRegistration) bool {
	return x.String() == item.String()
}
