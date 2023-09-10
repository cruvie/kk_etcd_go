package kk_etcd

import (
	"context"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
	"github.com/cruvie/kk_etcd_go/kk_etcd_client"
	"github.com/cruvie/kk_etcd_go/models"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"gopkg.in/yaml.v3"
	"log/slog"
	"strings"
)

// GetConfig get config from etcd and unmarshal to configStruct
// eg: GetConfig("go_rec_dev", &config.Config)
func GetConfig(configKey string, configStruct any) {
	getResponse, err := kk_etcd_client.EtcdClient.Get(context.Background(), key_prefix.Config+configKey)
	if err != nil {
		slog.Error("failed to get kv:", configKey, err)
	}
	err = yaml.Unmarshal(getResponse.Kvs[0].Value, configStruct)
	if err != nil {
		slog.Error("映射配置信息到结构体失败=", err, "读取到的信息=", string(getResponse.Kvs[0].Value))
	}
}

// RegisterHttpService 注册服务
// addr:服务地址 192.123.32.11:8080
// serverName:服务名称
// ttl:租约时间,默认15秒,服务下线ttl秒后服务会自动从etcd中删除
func RegisterHttpService(ctx context.Context, addr, serverName string, ttl ...int64) {
	registerService(ctx, key_prefix.ServiceHttp, addr, serverName, ttl...)
}

// RegisterGrpcService 注册服务
// addr:服务地址 192.123.32.11:8080
// serverName:服务名称
// ttl:租约时间,默认15秒,服务下线ttl秒后服务会自动从etcd中删除
func RegisterGrpcService(ctx context.Context, addr, serverName string, ttl ...int64) {
	registerService(ctx, key_prefix.ServiceGrpc, addr, serverName, ttl...)
}

// registerService 注册服务
// serviceType:服务类型 key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// addr:服务地址 192.123.32.11:8080
// serverName:服务名称
// ttl:租约时间,默认15秒,服务下线ttl秒后服务会自动从etcd中删除
func registerService(ctx context.Context, serviceType, addr, serverName string, ttl ...int64) {
	key := serviceType + serverName
	if serverName == "" {
		slog.Error("serverName can not be empty")
	}
	if len(ttl) == 0 {
		ttl = append(ttl, 15)
	}
	etcdManager, err := endpoints.NewManager(kk_etcd_client.EtcdClient, key)
	if err != nil {
		slog.Error("failed to create etcd manager", "err:", err)
	}

	lease, err := kk_etcd_client.EtcdClient.Grant(ctx, ttl[0])
	if err != nil {
		slog.Error("failed to create lease", "err:", err)
	}

	// 添加注册节点到 etcd 中，并且携带上租约 id
	err = etcdManager.AddEndpoint(ctx, key+"/"+addr, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID))
	if err != nil {
		slog.Error("failed to add endpoint to etcd", "err:", err)
	}
	//每ttl/3秒续约一次
	respChan, err := kk_etcd_client.EtcdClient.KeepAlive(ctx, lease.ID)
	if err != nil {
		slog.Error("failed to set keep alive", "err:", err)
	}
	//启动协程监听续约结果直到context关闭
	go func() {
		for {
			select {
			case resp := <-respChan:
				if resp == nil {
					//slog.Error("keep alive channel closed")
					return
				} else {
					//slog.Info("keep alive", "resp:", resp)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

// GetServiceList
// serviceName:服务名称 should with prefix key_prefix.ServiceGrpc or key_prefix.ServiceHttp
// only give prefix to get all service list
func GetServiceList(serviceName string) (serviceList *models.PBListService) {
	etcdManager, err := endpoints.NewManager(kk_etcd_client.EtcdClient, serviceName)
	if err != nil {
		slog.Error("failed to create etcd manager", "err:", err)
	}
	endpointMap, err := etcdManager.List(context.Background())
	if err != nil {
		slog.Error("failed to list endpoints", err)
		return
	}
	//ListService:{Key2EndpointMap:{key:"kk_service_http/haha/128.2.2.3:8484"  value:{Addr:"128.2.2.3:8484"}}}
	var pBListService models.PBListService
	for key, endpoint := range endpointMap {
		split := strings.Split(key, "/")
		var name string
		if len(split) >= 2 {
			name = split[1]
		} else {
			name = key
		}
		pBListService.ListService = append(pBListService.ListService, &models.PBService{
			ServiceName: name,
			ServiceAddr: endpoint.Addr,
		})
	}
	return &pBListService
}
