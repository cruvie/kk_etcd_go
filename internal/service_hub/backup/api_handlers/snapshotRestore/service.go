package snapshotRestore

import "gitee.com/cruvie/kk_go_kit/kk_stage"

func (x *Api) Service(stage *kk_stage.Stage) (string, error) {
	span := stage.StartTrace("service")
	defer span.End()

	//https://etcd.io/docs/v3.5/op-guide/recovery/
	// 执行etcdctl命令来恢复etcd数据
	//cmd := exec.Command("etcdctl", "snapshot", "restore", "/path/to/snapshot.db", "--name", "my-etcd")
	//err = cmd.Run()
	//if err != nil {
	//	fmt.Println("Failed to restore etcd data:", err)
	//	return
	//}
	cmdStr := `
			etcdutl snapshot restore /Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/snapshotFile.snapshot \
		  --data-dir  /Users/cruvie/KangXH/Coding/Uncomplete-Projects/kk_etcd/kk_etcd_go/backup/out \
		  --name node1 \
		  --initial-cluster node1=http://127.0.0.1:2380 \
		  --initial-advertise-peer-urls http://127.0.0.1:2380
`
	return cmdStr, nil
}
