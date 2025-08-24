package kk_etcd

import (
	"github.com/cruvie/kk_etcd_go/internal/service_housekeeper"
	apiImplBackup "github.com/cruvie/kk_etcd_go/internal/service_hub/backup/api_impl"
	apiImplKv "github.com/cruvie/kk_etcd_go/internal/service_hub/kv/api_impl"
	apiImplRole "github.com/cruvie/kk_etcd_go/internal/service_hub/role/api_impl"
	apiImplService "github.com/cruvie/kk_etcd_go/internal/service_hub/service/api_impl"
	apiImplUser "github.com/cruvie/kk_etcd_go/internal/service_hub/user/api_impl"
	"github.com/cruvie/kk_etcd_go/internal/service_hub/user/util_user"
	"github.com/cruvie/kk_etcd_go/internal/utils/global_model"
	"github.com/cruvie/kk_etcd_go/internal/utils/internal_client"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type InitClientConfig struct {
	clientv3.Config
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

func NewBackupClient() *apiImplBackup.CServer {
	return &apiImplBackup.CServer{}
}

func NewKVClient() *apiImplKv.CServer {
	return &apiImplKv.CServer{}
}

func NewUserClient() *apiImplUser.CServer {
	return &apiImplUser.CServer{}
}

func NewRoleClient() *apiImplRole.CServer {
	return &apiImplRole.CServer{}
}

func NewServiceClient() *apiImplService.CServer {
	return &apiImplService.CServer{}
}
