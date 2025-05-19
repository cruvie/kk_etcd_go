package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/service_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"google.golang.org/grpc"
)

var serService service_hub.SerService

// RegisterService register service to etcd
func RegisterService(registration *kk_etcd_models.PBServiceRegistration) error {
	err := serService.RegisterService(registration)
	return err
}

// GetGrpcClient get grpc client for serviceName
//
// don't forget to close conn conn.Close()
func GetGrpcClient[T any](
	serviceName string,
	clientBuilder func(grpcConn grpc.ClientConnInterface) (client T),
	opts ...grpc.DialOption,
) (conn *grpc.ClientConn, client T, err error) {

	addr, err := serService.GetConns(kk_etcd_models.PBServiceType_Grpc, serviceName)
	if err != nil {
		return nil, client, err
	}

	conn, err = grpc.NewClient(addr,
		opts...,
	)
	client = clientBuilder(conn)
	return conn, client, err
}
