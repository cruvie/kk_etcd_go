package service_housekeeper

import (
	"bytes"
	"encoding/json"

	"gitee.com/cruvie/kk_go_kit/kk_sync"
	"github.com/cruvie/kk_etcd_go/kk_etcd_models"
)

type hubT struct {
	*kk_sync.Map[string, *kk_sync.Slice[*kk_etcd_models.PBServiceRegistration]]
}

func newHubT() *hubT {
	return &hubT{
		Map: kk_sync.NewMap[string, *kk_sync.Slice[*kk_etcd_models.PBServiceRegistration]](),
	}
}

func (x *hubT) MarshalJSON() ([]byte, error) {
	m := make(map[string][]*kk_etcd_models.PBServiceRegistration)
	data := x.Data()
	for k, v := range data {
		m[k] = v.Slice()
	}

	return json.Marshal(m)
}

func (x *hubT) UnmarshalJSON(b []byte) error {
	tmp := make(map[string][]*kk_etcd_models.PBServiceRegistration)
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&tmp)
	if err != nil {
		return err
	}

	for k, v := range tmp {
		slice := kk_sync.NewSlice[*kk_etcd_models.PBServiceRegistration]()
		slice.Append(v...)
		x.Add(k, slice)
	}

	return nil
}
