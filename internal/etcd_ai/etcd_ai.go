package etcd_ai

import (
	"gitee.com/cruvie/kk_go_kit/kk_server"
	"github.com/cruvie/kk_etcd_go/internal/config"
	"github.com/cruvie/kk_etcd_go/internal/etcd_ai/rag_server"
	"github.com/tmc/langchaingo/llms/ollama"
)

func EtcdAIService() *kk_server.KKRunServer {
	if !config.Config.AI.Enable {
		return nil
	}

	run := func() {
		opts := []ollama.Option{
			ollama.WithModel(config.Config.AI.LLMModel),
		}
		if config.Config.AI.OllamaService != "" {
			opts = append(opts, ollama.WithServerURL(config.Config.AI.OllamaService))
		}

		{
			//initialized, err := vector_store.InitWeaviate(config.Config.AI.WeaviateIpPort)
			//if err != nil {
			//	panic(err)
			//}
			rag_server.RagServerClient = func() *rag_server.RagServer {
				ollamaLLM, err := ollama.New(opts...)
				if err != nil {
					panic(err)
				}
				ragServer := &rag_server.RagServer{LLM: ollamaLLM}
				//wvStore, err := weaviate.New(
				//	weaviate.WithEmbedder(ragServer),
				//	weaviate.WithScheme("http"),
				//	weaviate.WithHost(config.Config.AI.WeaviateIpPort),
				//	weaviate.WithIndexName("DocEtcdAI"),
				//)
				//if err != nil {
				//	panic(err)
				//}
				//ragServer.VectorStore = &wvStore
				return ragServer
			}()

			//if !initialized {
			//	err := vector_store.AddEtcdDocuments()
			//	if err != nil {
			//		panic(err)
			//	}
			//}
		}
	}
	done := func(quitCh <-chan struct{}) {
		<-quitCh
	}
	return &kk_server.KKRunServer{
		Run:  run,
		Done: done,
	}
}
