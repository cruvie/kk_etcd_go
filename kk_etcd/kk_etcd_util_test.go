package kk_etcd

import (
	"context"
	"gitee.com/cruvie/kk_go_kit/kk_utils/kku_log"
	"github.com/cruvie/kk_etcd_go/consts/key_prefix"
	"log"
	"sync"
	"testing"
	"time"
)

func TestRegisterService1(t *testing.T) {
	var w sync.WaitGroup
	w.Add(1)
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")

	ctx, cancelFunc := context.WithCancel(context.Background())
	//RegisterService(ctx, "127.2.1.1", "ssss")
	//RegisterService(ctx, "127.2.2.3", "ssss")
	RegisterHttpService(ctx, "128.2.2.11:8484", "ssss")
	w.Wait()
	cancelFunc()
}
func TestRegisterService2(t *testing.T) {
	var w sync.WaitGroup
	w.Add(1)
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")

	ctx, cancelFunc := context.WithCancel(context.Background())
	//RegisterService(ctx, "127.2.1.1", "ssss")
	//RegisterService(ctx, "127.2.2.3", "ssss")
	RegisterHttpService(ctx, "128.2.2.3:8484", "ssss")
	w.Wait()
	cancelFunc()
}
func TestRegisterService3(t *testing.T) {
	var w sync.WaitGroup
	w.Add(1)
	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")

	ctx, cancelFunc := context.WithCancel(context.Background())
	//RegisterService(ctx, "127.2.1.1", "ssss")
	//RegisterService(ctx, "127.2.2.3", "ssss")
	RegisterHttpService(ctx, "128.2.2.3:8484", "haha")
	w.Wait()
	cancelFunc()
}

func TestGetServiceList(t *testing.T) {

	kku_log.InitSlog(true, nil, nil)
	InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 5)
		list := GetServiceList(key_prefix.ServiceHttp)
		log.Println("list=", list)
	}
}
