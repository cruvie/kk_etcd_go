package kk_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"log"
	"testing"
)

func TestSnapshot(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")
	if err != nil {
		log.Println(err)
		return
	}
	err = Snapshot("/Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/etcd_backup.db")
	log.Println(err)
}

func TestSnapshotRestore(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")
	if err != nil {
		log.Println(err)
		return
	}
	err = SnapshotRestore("/Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/etcd_backup.db")
	log.Println(err)
}

func TestSnapshotInfo(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd")
	if err != nil {
		log.Println(err)
		return
	}
	info, err := SnapshotInfo("/Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/etcd_backup.db")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(info)
}
