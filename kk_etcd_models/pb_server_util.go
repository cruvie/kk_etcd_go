package kk_etcd_models

import (
	"errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"time"
)

const (
	HealthCheckPath = "/KKHealthCheck"
)

type ServerType string

func (x ServerType) String() string {
	return string(x)
}

func (x ServerType) NewEndpointManager(client *clientv3.Client) (endpoints.Manager, error) {
	return endpoints.NewManager(client, x.String())
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
	return x.ServerType.String() + "/" + x.ServerName + "/" + x.ServerAddr
}
