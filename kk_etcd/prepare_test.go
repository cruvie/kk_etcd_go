package kk_etcd_test

import (
	"log"
	"time"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"github.com/cruvie/kk_etcd_go/internal/utils/consts"
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/service/api_def"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	testServiceClient api_def.ServiceClient
)

func initGrpcClient() {
	serverAddress := "localhost:8081"
	conn, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	testServiceClient = api_def.NewServiceClient(conn)
}
func initTestEnv() kk_etcd.CloseFunc {
	configLog := kk_stage.ConfigLog{
		DebugMode:  true,
		Lumberjack: kk_stage.DefaultLogConfig(time.Now(), consts.ServiceName),
		StartTime:  time.Now(),
	}
	configLog.Init()
	defer configLog.Close()
	closeFunc, err := kk_etcd.InitClient(&kk_etcd.InitClientConfig{
		Config: clientv3.Config{
			Endpoints: []string{"http://127.0.0.1:2379"},
			Username:  "root",
			Password:  "root",
		},
		DebugMode: true})
	if err != nil {
		panic(err)
	}
	return closeFunc
}
