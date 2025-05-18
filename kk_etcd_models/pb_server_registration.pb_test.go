package kk_etcd_models

import "testing"

func TestName(t *testing.T) {
	t.Log(PBServerType_Http.Type())
	t.Log(PBServerType_Http.Descriptor())
	t.Log(PBServerType_Http.Enum())
	t.Log(PBServerType_Http.Number())
	t.Log(PBServerType_Http.String())
}
