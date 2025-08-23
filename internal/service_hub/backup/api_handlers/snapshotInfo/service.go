package snapshotInfo

import (
	"log/slog"
	"os"
	"os/exec"

	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

func (x *Api) Service(stage *kk_stage.Stage) (info string, err error) {
	span := stage.StartTrace("service")
	defer span.End()

	newLog := kk_stage.NewLog(stage)
	tempFile, err := os.CreateTemp("", "temp")
	if err != nil {
		return "", err
	}
	defer func() {
		err = os.Remove(tempFile.Name())
		if err != nil {
			slog.Error("delete temp file failed", newLog.Error(err).Args()...)
		}
	}()
	if _, err := tempFile.Write(x.In.GetFile()); err != nil {
		return "", err
	}
	if err := tempFile.Close(); err != nil {
		return "", err
	}

	//etcdutl --write-out=table snapshot status etcd_backup.snapshot
	cmd := exec.Command("etcdutl", "--write-out", "table", "snapshot", "status", tempFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
