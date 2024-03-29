package kk_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"log"
	"log/slog"
	"os"
	"testing"
)

func TestSnapshot(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd", true)
	if err != nil {
		log.Println(err)
		return
	}

	pBFile, err := Snapshot()
	if err != nil {
		slog.Error("Failed to create snapshot", kk_stage.NewLog(nil).Error(err).Args()...)
		return
	}
	workDir, _ := os.Getwd()
	backupFileName := workDir + "/" + pBFile.Name
	err = os.WriteFile(backupFileName, pBFile.Bytes, 0644)
	if err != nil {
		slog.Error("Failed to write backup data to file", kk_stage.NewLog(nil).Error(err).Args()...)
	}
}

func TestSnapshotRestore(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)

	cmdStr, err := SnapshotRestore()
	if err != nil {
		log.Println(err)
	}
	log.Println(cmdStr)
}

func TestSnapshotInfo(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	bytes, err := os.ReadFile("/Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/etcd_backup.db")
	if err != nil {
		slog.Error("Failed to read file", kk_stage.NewLog(nil).Error(err).Args()...)
	}
	info, err := SnapshotInfo(bytes)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(info)
}

func TestAllKVsBackup(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd", true)
	if err != nil {
		log.Println(err)
		return
	}
	pbFile, err := AllKVsBackup()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(pbFile.Name)
	log.Println(string(pbFile.Bytes))
}

func TestAllKVsRestore(t *testing.T) {
	kk_stage.InitSlog(true, nil, nil)
	err := InitEtcd([]string{"http://127.0.0.1:2379"}, "kk_etcd", "kk_etcd", true)
	if err != nil {
		log.Println(err)
		return
	}
	jsonStr := `{"ListKV":[{"Key":"testkvbackup1","Value":"{\"UserName\":\"ww\",\"Password\":\"$2a$04$qM0YyMWVX0yz/rcVco8H/OVFBeYh/Wbc0drklFfS29BsDTekuK\"}"},{"Key":"testkvbackup2","Value":"dsafsd\ndasfsdf"}]}`

	err = AllKVsRestore([]byte(jsonStr))
	if err != nil {
		log.Println(err)
		return
	}
}
