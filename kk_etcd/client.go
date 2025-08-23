package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/service_housekeeper"
	apiDefAI "github.com/cruvie/kk_etcd_go/internal/service_hub/ai/api_def"
	apiDefBackup "github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_def"
	apiDefKv "github.com/cruvie/kk_etcd_go/internal/service_hub/kv/api_def"
	apiDefRole "github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_def"
	apiDefService "github.com/cruvie/kk_etcd_go/internal/service_hub/service/api_def"
	apiDefUser "github.com/cruvie/kk_etcd_go/internal/service_hub/user/api_def"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/util_user"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type InitClientConfig struct {
	clientv3.Config
	GrpcAddr  string
	GrpcOpts  []grpc.DialOption
	DebugMode bool
}

func (x *InitClientConfig) Check() {
	if len(x.Endpoints) == 0 {
		panic("endpoints is empty")
	}
	if x.Username == "" {
		panic("userName is empty")
	}
	if x.Password == "" {
		panic("password is empty")
	}
}

type CloseFunc func() error

func InitClient(cfg *InitClientConfig) (CloseFunc, error) {
	cfg.Check()
	internal_client.InitGlobalStage()

	client, err := clientv3.New(cfg.Config)
	if err != nil {
		return nil, err
	}
	global_model.SetClient(internal_client.GlobalStage, client)

	user, err := util_user.GetUser(internal_client.GlobalStage, cfg.Username)
	if err != nil {
		return nil, err
	}

	global_model.SetLoginUser(internal_client.GlobalStage, user)
	service_housekeeper.InitKubernetesClient(cfg.Config)
	closeFunc := func() error {
		err := GetClient().Close()
		internal_client.CloseGlobalStage()
		service_housekeeper.CloseKubernetesClient()
		return err
	}
	return closeFunc, nil
}

func GetClient() *clientv3.Client {
	return global_model.GetClient(internal_client.GlobalStage)
}

func NewAIClient(cfg *InitClientConfig) (*grpc.ClientConn, apiDefAI.AIClient, error) {
	conn, err := grpc.NewClient(cfg.GrpcAddr,
		cfg.GrpcOpts...,
	)
	client := apiDefAI.NewAIClient(conn)
	return conn, client, err
}

func NewBackupClient(cfg *InitClientConfig) (*grpc.ClientConn, apiDefBackup.BackupClient, error) {
	conn, err := grpc.NewClient(cfg.GrpcAddr,
		cfg.GrpcOpts...,
	)
	client := apiDefBackup.NewBackupClient(conn)
	return conn, client, err

}

func NewKVClient(cfg *InitClientConfig) (*grpc.ClientConn, apiDefKv.KVClient, error) {
	conn, err := grpc.NewClient(cfg.GrpcAddr,
		cfg.GrpcOpts...,
	)
	client := apiDefKv.NewKVClient(conn)
	return conn, client, err

}

func NewUserClient(cfg *InitClientConfig) (*grpc.ClientConn, apiDefUser.UserClient, error) {
	conn, err := grpc.NewClient(cfg.GrpcAddr,
		cfg.GrpcOpts...,
	)
	client := apiDefUser.NewUserClient(conn)
	return conn, client, err

}

func NewRoleClient(cfg *InitClientConfig) (*grpc.ClientConn, apiDefRole.RoleClient, error) {
	conn, err := grpc.NewClient(cfg.GrpcAddr,
		cfg.GrpcOpts...,
	)
	client := apiDefRole.NewRoleClient(conn)
	return conn, client, err

}

func NewServiceClient(cfg *InitClientConfig) (*grpc.ClientConn, apiDefService.ServiceClient, error) {
	conn, err := grpc.NewClient(cfg.GrpcAddr,
		cfg.GrpcOpts...,
	)
	client := apiDefService.NewServiceClient(conn)
	return conn, client, err

}
