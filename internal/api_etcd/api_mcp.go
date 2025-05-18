package api_etcd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/cruvie/kk_go_kit/kk_log"
	"gitee.com/cruvie/kk_go_kit/kk_server"
	"gitee.com/cruvie/kk_go_kit/kk_swagger"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/swagger"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type MCPConfig struct {
	OriginalHttpServerPort        int
	OriginalHttpServerSwaggerJson string
	MCPAddr                       string
}

func ApiMCP() *kk_server.KKRunServer {
	if !config.Config.MCPServer.Enable {
		return nil
	}

	mcpServer := server.NewMCPServer(
		"kk_etcd",
		"1.0.0",
		server.WithLogging(),
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithToolCapabilities(true),
	)

	{
		mcpServer.AddTools(genTools()...)
	}

	sseServer := server.NewSSEServer(
		mcpServer,
		server.WithStaticBasePath("/kk_etcd"),
	)

	run := func() {
		slog.Info(fmt.Sprintf("Starting SSE server http://127.0.0.1:%d/kk_etcd/sse", config.Config.MCPServer.Port))
		err := sseServer.Start(fmt.Sprintf(":%d", config.Config.MCPServer.Port))
		if err != nil {
			panic(err)
		}
	}

	done := func(quitCh <-chan struct{}) {
		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		<-quitCh
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := sseServer.Shutdown(ctx); err != nil {
			slog.Error("ApiMCP shutdown error", kk_log.NewLog(nil).Error(err).Args()...)
		}
	}

	return &kk_server.KKRunServer{
		Run:  run,
		Done: done,
	}
}

func genTools() (tools []server.ServerTool) {
	apiInfos := kk_swagger.ParseSwaggerJson(swagger.SwaggerJson)

	for _, info := range apiInfos {
		tools = append(tools, buildTool(&info))
	}
	return tools
}

func buildTool(info *kk_swagger.ApiInfo) server.ServerTool {
	return server.ServerTool{
		Tool: mcp.NewTool(info.OperationId,
			mcp.WithDescription(info.Description),
			mcp.WithString("token",
				mcp.Description("auth token")),
			mcp.WithString("UserName"),
			mcp.WithString("Password"),
			mcp.WithObject("input",
				mcp.Required(),
				mcp.Description(info.InputJson),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			arguments := request.Params.Arguments

			userName, ok := arguments["UserName"].(string)
			if !ok {
				return nil, fmt.Errorf("invalid UserName argument")
			}
			password := arguments["Password"].(string)
			if password == "" {
				return nil, fmt.Errorf("invalid password argument")
			}

			input := bytes.NewReader([]byte("{}"))
			if info.InputJson != "" {
				in := arguments["input"]
				marshal, err := json.Marshal(in)
				if err != nil {
					return nil, err
				}
				input = bytes.NewReader(marshal)
			}

			token, ok := arguments["token"].(string)
			if !ok {
				return nil, fmt.Errorf("invalid token argument")
			}

			//转发请求
			req, err := http.NewRequestWithContext(
				ctx,
				info.Method, fmt.Sprintf(
					"http://127.0.0.1:%d%s",
					config.Config.Port, info.Path,
				),
				input,
			)
			if err != nil {
				return nil, err
			}
			req.Header.Set("Authorization", token)
			req.Header.Set("UserName", userName)
			req.Header.Set("Password", password)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")

			// 发送请求
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return nil, err
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					slog.Error("Error executing request", "err", err)
				}
			}(resp.Body)

			// 读取响应
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			// 返回响应
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: string(body),
					},
				},
			}, nil
		},
	}
}
