package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/server_hub"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"google.golang.org/grpc"
)

var serServer server_hub.SerServer

// RegisterService register service to etcd
func RegisterService(registration *kk_etcd_models.PBServerRegistration) error {
	err := serServer.RegisterService(registration)
	return err
}

// GetGrpcClient get grpc client for serverName
//
// don't forget to close conn conn.Close()
func GetGrpcClient[T any](
	serverName string,
	clientBuilder func(grpcConn grpc.ClientConnInterface) (client T),
	opts ...grpc.DialOption,
) (conn *grpc.ClientConn, client T, err error) {

	addr, err := serServer.GetConns(kk_etcd_models.PBServerType_Grpc, serverName)
	if err != nil {
		return nil, client, err
	}

	conn, err = grpc.NewClient(addr,
		opts...,
	)
	client = clientBuilder(conn)
	return conn, client, err
}
