package kk_etcd_models

import "context"

const (
	CheckTypeGrpc   = "GRPC"
	CheckTypeHttp   = "HTTP"
	HealthCheckPath = "/KKHealthCheck"
)

type ServiceRegistration struct {
	Context    context.Context
	ServerType string
	Address    string
	ServerName string
	Check      *ServiceCheck
}

type ServiceCheck struct {
	Type       string // CheckTypeGrpc or CheckTypeHttp
	TTL        int64  // default 30
	Timeout    int64  //check timeout, default TTL / 3
	Interval   int64  //check interval, default TTL / 3, kk_etcd will per TTL/3 seconds to send check request,
	HTTPMethod string //GET/POST, default GET
	HTTP       string //default http://+Address+/KKHealthCheck
	GRPC       string //default Address
}
