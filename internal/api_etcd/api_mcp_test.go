package api_etcd

import (
	"context"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/kv/kVList"
	"github.com/cruvie/kk_etcd_go/kk_etcd_api_hub/user/login"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func newClient(t *testing.T) (*client.SSEMCPClient, context.Context) {
	ssemcpClient, err := client.NewSSEMCPClient("http://127.0.0.1:2445/kk_etcd/sse")
	if err != nil {
		t.Fatalf("Failed to create ssemcpClient: %v", err)
	}

	ctx := context.Background()

	// Start the ssemcpClient
	if err := ssemcpClient.Start(ctx); err != nil {
		t.Fatalf("Failed to start ssemcpClient: %v", err)
	}

	// Initialize
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "kk_etcd",
		Version: "1.0.0",
	}

	result, err := ssemcpClient.Initialize(ctx, initRequest)
	if err != nil {
		t.Fatalf("Failed to initialize: %v", err)
	}

	if result.ServerInfo.Name != "kk_etcd" {
		t.Errorf(
			"Expected server name 'kk_etcd', got '%s'",
			result.ServerInfo.Name,
		)
	}

	//Test Ping
	if err := ssemcpClient.Ping(ctx); err != nil {
		t.Errorf("Ping failed: %v", err)
	}
	return ssemcpClient, ctx
}
func TestMCPServer(t *testing.T) {
	cli, _ := newClient(t)
	defer cli.Close()
	// Test ListTools
	toolsRequest := mcp.ListToolsRequest{}
	res, err := cli.ListTools(context.Background(), toolsRequest)
	if err != nil {
		t.Errorf("ListTools failed: %v", err)
	}
	t.Log(res)
}

func TestLogin(t *testing.T) {
	cli, ctx := newClient(t)
	defer cli.Close()

	request := mcp.CallToolRequest{}
	request.Params.Name = "Login"
	request.Params.Arguments = map[string]interface{}{
		"UserName": "root",
		"Password": "root",
		"input":    login.Login_Input{},
		"token":    "asfafdsgds",
	}

	result, err := cli.CallTool(ctx, request)
	if err != nil {
		t.Fatalf("CallTool failed: %v", err)
	}
	t.Log(result.Content)
}
func TestCallTool(t *testing.T) {
	cli, ctx := newClient(t)
	defer cli.Close()

	request := mcp.CallToolRequest{}
	request.Params.Name = "KVList"
	request.Params.Arguments = map[string]interface{}{
		"UserName": "root",
		"Password": "root",
		"input":    kVList.KVList_Input{},
		"token":    "asfafdsgds",
	}

	result, err := cli.CallTool(ctx, request)
	if err != nil {
		t.Fatalf("CallTool failed: %v", err)
	}
	t.Log(result.Content)
}
