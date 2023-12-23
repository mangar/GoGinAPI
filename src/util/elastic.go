package util

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Document struct {
	ID    string `json:"id"`
	API string `json:"API"`
	Status  string `json:"status"`
}

func Elastic(doc Document) error {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.1.110:9200", // URL do Elasticsearch
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Erro ao criar o cliente: %s", err)
	}

	// Documento a ser indexado
	// doc := Document{
	// 	ID:    "1",
	// 	Title: "Teste",
	// 	Body:  "Este é um teste de indexação no Elasticsearch.",
	// }

	// Serializar o documento
	var b strings.Builder
	json.NewEncoder(&b).Encode(doc)

	// Request para indexar o documento
	req := esapi.IndexRequest{
		Index:      "seu_indice", // Nome do índice
		DocumentID: doc.ID,
		Body:       bytes.NewReader([]byte(b.String())),
		Refresh:    "true",
	}

	// Executar a request
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Erro ao indexar o documento: %s", err)
	}
	defer res.Body.Close()

	// Checar se a indexação foi bem-sucedida
	if res.IsError() {
		log.Printf("Erro ao indexar o documento: %s", res.String())
	} else {
		log.Printf("Documento indexado com sucesso: %s", res.String())
	}

	return nil
}