package kk_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"log"
	"testing"
)

func TestInitEtcd(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")
	if err != nil {
		log.Println(err)
		return
	}
}
