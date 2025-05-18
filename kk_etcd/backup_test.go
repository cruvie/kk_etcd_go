package kk_etcd_test

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/backup/api_def"
	"log"
	"log/slog"
	"os"
	"testing"

	"gitee.com/cruvie/kk_go_kit/kk_log"
)

func TestSnapshot(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()

	mgrBackup := kk_etcd.NewMgrBackup()
	pBFile, err := mgrBackup.Snapshot()
	if err != nil {
		slog.Error("Failed to create snapshot", kk_log.NewLog(nil).Error(err).Args()...)
		return
	}
	workDir, _ := os.Getwd()
	backupFileName := workDir + "/" + pBFile.GetName()
	err = os.WriteFile(backupFileName, pBFile.GetFile(), 0644)
	if err != nil {
		slog.Error("Failed to write backup data to file", kk_log.NewLog(nil).Error(err).Args()...)
	}
}

func TestSnapshotRestore(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	mgrBackup := kk_etcd.NewMgrBackup()
	err, response := mgrBackup.SnapshotRestore()
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
}

func TestSnapshotInfo(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	mgrBackup := kk_etcd.NewMgrBackup()
	bytes, err := os.ReadFile("/Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/etcd_backup.db")
	if err != nil {
		slog.Error("Failed to read file", kk_log.NewLog(nil).Error(err).Args()...)
	}
	response, err := mgrBackup.SnapshotInfo(&api_def.SnapshotInfo_Input{File: bytes})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(response.GetInfo())
}

func TestAllKVsBackup(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	mgrBackup := kk_etcd.NewMgrBackup()
	pbFile, err := mgrBackup.AllKVsBackup()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(pbFile.GetName())
	log.Println(string(pbFile.GetFile()))
}

func TestAllKVsRestore(t *testing.T) {
	closeFunc := initTestEnv()
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	mgrBackup := kk_etcd.NewMgrBackup()
	jsonStr := `{"ListKV":[{"Key":"testkvbackup1","Value":"{\"UserName\":\"ww\",\"Password\":\"$2a$04$qM0YyMWVX0yz/rcVco8H/OVFBeYh/Wbc0drklFfS29BsDTekuK\"}","Key":"testkvbackup2","Value":"dsafsd\ndasfsdf"}]}`

	err, response := mgrBackup.AllKVsRestore(&api_def.AllKVsRestore_Input{File: []byte(jsonStr)})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(response)
}
