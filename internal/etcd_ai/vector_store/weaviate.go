package vector_store

// InitWeaviate initializes a weaviate client for our application.
//func InitWeaviate(weaviateIpPort string) (initialized bool, err error) {
//	client, err := weaviate.NewClient(weaviate.Config{
//		Host:   weaviateIpPort,
//		Scheme: "http",
//	})
//	if err != nil {
//		return false, fmt.Errorf("initializing weaviate: %w", err)
//	}
//
//	// Create a new class (collection) in weaviate if it doesn't exist yet.
//	cls := &models.Class{
//		Class:      "DocEtcdAI",
//		Vectorizer: "none",
//	}
//	exists, err := client.Schema().ClassExistenceChecker().WithClassName(cls.Class).Do(context.Background())
//	if err != nil {
//		return false, fmt.Errorf("weaviate error: %w", err)
//	}
//	if exists && config.Config.AI.ReInitWeaviate {
//		err := client.Schema().ClassDeleter().WithClassName(cls.Class).Do(context.Background())
//		if err != nil {
//			return false, fmt.Errorf("delete doc class error: %w", err)
//		}
//		exists = false
//	}
//	if !exists {
//		err = client.Schema().ClassCreator().WithClass(cls).Do(context.Background())
//		if err != nil {
//			return false, fmt.Errorf("weaviate error: %w", err)
//		}
//	} else {
//		slog.Info("Weaviate class already exists skip init", "class", cls.Class)
//	}
//	return exists, nil
//}

//func AddEtcdDocuments() error {
//	var wvDocs []schema.Document
//	//read all etcd docs
//	filePaths, err := kk_file.ListFiles(config.Config.AI.EtcdDocPath, ".md")
//	if err != nil {
//		return err
//	}
//	for _, path := range filePaths {
//		file, err := os.ReadFile(path)
//		if err != nil {
//			return err
//		}
//		wvDocs = append(wvDocs, schema.Document{PageContent: string(file)})
//	}
//	//add to weaviate
//	slog.Info("start add etcd docs to weaviate...", "doc count", len(wvDocs))
//	_, err = rag_server.RagServerClient.VectorStore.AddDocuments(context.Background(), wvDocs)
//	if err != nil {
//		return err
//	}
//	slog.Info("add etcd docs to weaviate success")
//	return nil
//}
//
//func SearchDocuments(text string) (docsContents []string, err error) {
//	// Find the most similar documents.
//	docs, err := rag_server.RagServerClient.VectorStore.SimilaritySearch(
//		context.Background(),
//		text,
//		1,
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, doc := range docs {
//		docsContents = append(docsContents, doc.PageContent)
//	}
//	return docsContents, nil
//}
