package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)

// GetGrpcEtcdResolverDialOption refers to https://etcd.io/docs/v3.5/dev-guide/grpc_naming/
func GetGrpcEtcdResolverDialOption() (grpc.DialOption, error) {
	etcdResolver, err := resolver.NewBuilder(global_model.GetClient(internal_client.GlobalStage))
	if err != nil {
		return nil, err
	}
	return grpc.WithResolvers(etcdResolver), nil
}

func GetGrpcRoundRobinDialOption() grpc.DialOption {
	return grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)
}
