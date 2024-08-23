package kk_etcd

import (
	"context"
	"errors"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"google.golang.org/grpc"
	"log/slog"
	"testing"
)

// your RPCClient form message proto
type RPCClient any

var NewRPCClient func(grpcConn grpc.ClientConnInterface) (client RPCClient)
var clientHub *ClientHub[RPCClient]

func InitGrpcClient(stage *kk_stage.Stage) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	clientHub = NewClientHub[RPCClient](kk_etcd_const.ServiceGrpc,
		"MyServerName",
		NewRPCClient)
	err := clientHub.ListenServerChange(ctx, stage)
	if err != nil {
		slog.Error(err.Error())
	}
}
func GetGrpcClient(stage *kk_stage.Stage) (RPCClient, error) {
	client := clientHub.GetClient(stage)
	if client == nil {
		return client, errors.New("nil grpc client")
	} else {
		return *client, nil
	}
}

func TestName(t *testing.T) {
	stage := kk_stage.NewStage(context.Background(), "", true)
	InitGrpcClient(stage)
	client, err := GetGrpcClient(stage)
	if err != nil {
		slog.Error(err.Error())
	}
	_ = client
}
