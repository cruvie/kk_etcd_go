package kk_etcd_tool

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/kk_etcd_const"
	"google.golang.org/grpc"
	"testing"
)

// your RPCClient form message proto
type RPCClient any

var NewRPCClient func(grpcConn grpc.ClientConnInterface) (client RPCClient)
var clientHub *ClientHub[RPCClient]

func InitGrpcClient(stage *kk_stage.Stage) {
	clientHub = NewClientHub[RPCClient]()
	clientHub.ListenServerChange(
		stage, kk_etcd_const.ServiceGrpc,
		"MyServerName",
		clientHub,
		NewRPCClient)
}
func GetGrpcClient() RPCClient {
	return *clientHub.GetGrpcClient()
}

func TestName(t *testing.T) {
	stage := kk_stage.NewStage(nil, "")
	InitGrpcClient(stage)
	client := GetGrpcClient()
	_ = client
}
