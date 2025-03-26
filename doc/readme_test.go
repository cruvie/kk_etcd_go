package doc_test

import (
	"github.com/cruvie/kk_etcd_go/kk_etcd"
	"gopkg.in/yaml.v3"
	"log"
	"log/slog"
	"os"
	"testing"
)

type myConfig struct {
	ServerAddr string
	Postgres   struct {
		Dsn  string
		Port int
	}
	Redis struct {
		Addr     string
		Password string
	}
}

func TestPutYaml(t *testing.T) {
	//init client
	closeFunc, err := kk_etcd.InitClient(&kk_etcd.InitClientConfig{
		Endpoints: []string{"http://127.0.0.1:2379"},
		UserName:  "root",
		Password:  "root",
		DebugMode: true})
	if err != nil {
		panic(err)
	}
	defer func() {
		err := closeFunc()
		if err != nil {
			log.Println(err)
		}
	}()
	//load config
	data, err := os.ReadFile("path/to/my_config.yml")
	if err != nil {
		panic(err)
	}
	var Config myConfig
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		slog.Error("unable to unmarshal config.yaml", "err", err)
		panic(err)
	}
	//push config to etcd in yaml format
	kv := kk_etcd.NewMgrKV()
	err = kv.PutExistUpdateYaml("my_config", &Config)
	if err != nil {
		panic(err)
	}
	//get config from etcd
	var newConfig myConfig
	err = kv.GetYaml("my_config", &newConfig)
	if err != nil {
		panic(err)
	}
}
