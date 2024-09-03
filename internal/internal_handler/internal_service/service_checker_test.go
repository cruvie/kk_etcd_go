package internal_service

import (
	"context"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
	"testing"
	"time"
)

type MyMeta struct {
	Name string
}

func TestBug(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
		Username:    "kk_etcd",
		Password:    "kk_etcd",
	})
	if err != nil {
		panic(err)
	}
	endpointManager, err := kk_etcd_models.Http.NewEndpointManager(client)
	if err != nil {
		panic(err)
	}
	info := MyMeta{Name: "myname"}
	endpoint := endpoints.Endpoint{
		Addr:     "127.0.0.1:6666",
		Metadata: info,
	}
	//add endpoint to etcd with lease id.
	err = endpointManager.AddEndpoint(
		context.Background(),
		"mytarget/service",
		endpoint)
	if err != nil {
		panic(err)
	}
	endpointMap, err := endpointManager.List(context.Background())
	if err != nil {
		panic(err)
	}
	for _, v := range endpointMap {
		out, ok := v.Metadata.(MyMeta)
		log.Println(out, ok) // {} false
	}
}
func TestConvert(t *testing.T) {
	info := MyMeta{Name: "myname"}

	meta := endpoints.Endpoint{
		Addr:     "127.0.0.1:6666",
		Metadata: info,
	}
	out, ok := meta.Metadata.(MyMeta)
	log.Println(out, ok) //{myname} true
}
