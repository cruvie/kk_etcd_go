package kk_etcd

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"testing"
)

func TestGetGrpcClient(t *testing.T) {
	initTestEnv()
	conn, client, err := GetGrpcClient[grpc_health_v1.HealthClient]("haha_grpc",
		grpc_health_v1.NewHealthClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	response, err := client.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{
		Service: "empty kk",
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Println(response)
}
