package rag_server

import (
	"context"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"time"
)

var RagServerClient *RagServer

type RagServer struct {
	LLM *ollama.LLM
	//VectorStore *weaviate.Store
}

func (n *RagServer) EmbedDocuments(ctx context.Context, texts []string) ([][]float32, error) {
	embedding, err := n.LLM.CreateEmbedding(ctx, texts)
	return embedding, err
}

func (n *RagServer) EmbedQuery(ctx context.Context, text string) ([]float32, error) {
	embedding, err := n.LLM.CreateEmbedding(ctx, []string{text})
	return embedding[0], err
}

func Query(question string, docsContents []string) (answer string, err error) {
	//ragQuery := fmt.Sprintf(ragTemplateStr, question, strings.Join(docsContents, "\n"))

	parts := []llms.ContentPart{
		//llms.TextContent{Text: ragQuery},
		llms.TextContent{Text: question},
	}

	content := []llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: parts,
		},
	}

	timeout, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	generateContent, err := RagServerClient.LLM.GenerateContent(timeout, content)
	if err != nil {
		return "", err
	}
	return generateContent.Choices[0].Content, nil
}

// nolint
const ragTemplateStr = `
I will ask you a question and will provide some additional context information.
Assume this context information is factual and correct, as part of internal
documentation.
If the question relates to the context, answer it using the context.
If the question does not relate to the context, answer it as normal.

For example, let's say the context has nothing in it about tropical flowers;
then if I ask you about tropical flowers, just answer what you know about them
without referring to the context.

For example, if the context does mention minerology and I ask you about that,
provide information from the context along with general knowledge.

Question:
%s

Context:
%s
`
