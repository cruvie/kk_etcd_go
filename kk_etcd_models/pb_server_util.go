package kk_etcd_models

import (
	"errors"
	"time"
)

type CheckType string

const (
	CheckTypeGrpc   CheckType = "GRPC"
	CheckTypeHttp   CheckType = "HTTP"
	HealthCheckPath           = "/KKHealthCheck"
)

type ServiceType string

func (s ServiceType) String() string {
	return string(s)
}

const (
	Service                 = "kk_service/"
	ServiceHttp ServiceType = Service + "http"
	ServiceGrpc ServiceType = Service + "grpc"
)

type CheckConfig struct {
	Type     CheckType     // CheckTypeGrpc or CheckTypeHttp
	Timeout  time.Duration //check timeout
	Interval time.Duration //check interval, kk_etcd will per Interval to send check request
	// CheckTypeHttp default http://+Addr+/KKHealthCheck
	// CheckTypeGrpc Addr in ServiceRegistration
	Addr string
}

func (x *CheckConfig) Check() error {
	if x.Type != CheckTypeGrpc && x.Type != CheckTypeHttp {
		msg := "check type must be CheckTypeGrpc or CheckTypeHttp"
		return errors.New(msg)
	}
	if x.Timeout <= 0 {
		msg := "check timeout cannot be zero"
		return errors.New(msg)
	}
	if x.Interval <= 0 {
		msg := "check interval cannot be zero"
		return errors.New(msg)
	}
	if x.Addr == "" {
		msg := "check addr cannot be empty"
		return errors.New(msg)
	}
	return nil
}

type ServiceRegistration struct {
	ServerType  ServiceType
	ServerName  string
	Addr        string
	Metadata    any
	CheckConfig CheckConfig
}

func (x *ServiceRegistration) Check() error {
	if x.ServerName == "" {
		msg := "server name cannot be empty"
		return errors.New(msg)
	}
	if x.Addr == "" {
		msg := "server address cannot be empty"
		return errors.New(msg)
	}
	return x.CheckConfig.Check()
}

func (x *ServiceRegistration) EndpointKey() string {
	return x.ServerType.String() + "/" + x.ServerName
}
