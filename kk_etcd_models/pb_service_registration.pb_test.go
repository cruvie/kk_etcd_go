package kk_etcd_models

import "testing"

func TestName(t *testing.T) {
	t.Log(PBServiceType_Http.Type())
	t.Log(PBServiceType_Http.Descriptor())
	t.Log(PBServiceType_Http.Enum())
	t.Log(PBServiceType_Http.Number())
	t.Log(PBServiceType_Http.String())
}
