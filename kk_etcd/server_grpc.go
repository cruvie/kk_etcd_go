package kk_etcd

import (
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
	etcdResolver, err := resolver.NewBuilder(GetClient())
	if err != nil {
		return nil, client, err
	}

	options := []grpc.DialOption{
		grpc.WithResolvers(etcdResolver),
		//bug https://github.com/etcd-io/etcd/issues/19700#issuecomment-2779093288
		//grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"pick_first"}`),
	}
	options = append(options, opts...)

	conn, err = grpc.NewClient("etcd:///"+kk_etcd_models.Grpc.String()+"/"+serverName,
		options...,
	)
	client = clientBuilder(conn)
	return conn, client, err
}
