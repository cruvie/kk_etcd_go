package kk_etcd_models

import "context"

type ServiceRegistration struct {
	Context    context.Context
	ServerType string
	Address    string
	ServerName string
	Check      *ServiceCheck
}
type ServiceCheck struct {
	TTL        int64  //kk_etcd will per TTL/3 seconds to send check request
	Timeout    int64  //check timeout, default TTL / 3
	Interval   int64  //check interval, default TTL / 3
	HTTPMethod string //GET/POST, default GET
	HTTP       string //eg: http://127.0.0.1:4321/Check
	GRPC       string //eg: 127.0.0.1:4321
}
