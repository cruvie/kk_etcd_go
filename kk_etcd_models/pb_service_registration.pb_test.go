package kk_etcd_models

import (
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

func TestName(t *testing.T) {
	t.Log(PBServiceType_Http.Type())
	t.Log(PBServiceType_Http.Descriptor())
	t.Log(PBServiceType_Http.Enum())
	t.Log(PBServiceType_Http.Number())
	t.Log(PBServiceType_Http.String())
}

func TestName2(t *testing.T) {
	pb := &PBServiceRegistration{
		ServiceType: PBServiceType_Grpc,
		ServiceName: "name",
		ServiceAddr: "1241241",
		CheckConfig: &PBServiceRegistration_PBCheckConfig{
			Type:     0,
			Timeout:  durationpb.New(10 * time.Second),
			Interval: durationpb.New(20 * time.Second),
			Addr:     "1241241",
		},
	}
	t.Log(pb.String())
}
