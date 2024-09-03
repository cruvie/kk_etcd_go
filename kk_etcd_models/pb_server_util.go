package kk_etcd_models

import (
	"errors"
	"time"
)

const (
	HealthCheckPath = "/KKHealthCheck"
)

type ServerType string

func (s ServerType) String() string {
	return string(s)
}
func (s ServerType) EndpointManagerTarget(serverName string) string {
	if serverName != "" {
		return s.String() + "/" + serverName
	}
	return s.String()
}

const (
	Server            = "kk_server/"
	Http   ServerType = Server + "http"
	Grpc   ServerType = Server + "grpc"
	// InternalServerStatus use for store server status as normal kv, cannot use key in endpoint, watch manager will watch any key with same target prefix
	InternalServerStatus = Server + "internal_server_status/"
)

type CheckConfig struct {
	Type     ServerType    // Http or Grpc
	Timeout  time.Duration //check timeout
	Interval time.Duration //check interval, kk_etcd will per Interval to send check request
	// Http default http://+Addr+/KKHealthCheck
	// Grpc Addr in ServerRegistration
	Addr string
}

func (x *CheckConfig) Check() error {
	if x.Type != Http && x.Type != Grpc {
		msg := "check type must be Grpc or Http"
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

type ServerRegistration struct {
	ServerType  ServerType
	ServerName  string
	ServerAddr  string
	Metadata    any
	CheckConfig CheckConfig
}

func (x *ServerRegistration) Check() error {
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

func (x *ServerRegistration) EndpointKey() string {
	return x.EndpointManagerTarget() + "/" + x.ServerAddr
}

func (x *ServerRegistration) EndpointManagerTarget() string {
	return x.ServerType.EndpointManagerTarget(x.ServerName)
}
