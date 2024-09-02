package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)

// GetGrpcClient get grpc client for serverName
//
// don't forget to close conn conn.Close()
func GetGrpcClient[T any](serverName string,
	clientBuilder func(grpcConn grpc.ClientConnInterface) (client T),
	opts ...grpc.DialOption) (conn *grpc.ClientConn, client T, err error) {

	//refers to https://etcd.io/docs/v3.5/dev-guide/grpc_naming/
	etcdResolver, err := resolver.NewBuilder(global_model.GetClient(internal_client.GlobalStage))
	if err != nil {
		return nil, client, err
	}

	options := []grpc.DialOption{
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	}
	options = append(options, opts...)

	conn, err = grpc.NewClient("etcd:///"+kk_etcd_models.Grpc.String()+"/"+serverName,
		options...,
	)
	client = clientBuilder(conn)
	return conn, client, err
}
