package api_etcd

import (
	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_stage"
	"log/slog"
	"net"
	"os"
	"path/filepath"
)

func ApiSocket(stage *kk_stage.Stage) *kk_server.KKRunServer {
	var conn net.Conn
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	run := func() {
		ln, err := net.Listen("unix", filepath.Join(workDir, "kk_etcd.sock"))
		if err != nil {
			panic(err)
		}
		conn, err = ln.Accept()
		if err != nil {
			panic(err)
		}
	}

	done := func(quitCh <-chan struct{}) {
		<-quitCh
		err := conn.Close()
		if err != nil {
			slog.Error("ApiSocket conn close error", "error", err)
		}
	}

	return &kk_server.KKRunServer{
		Run:  run,
		Done: done,
	}
}
