package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	// "github.com/elastic/go-elasticsearch/v8/esapi"
)

func CreateDocument(es *elasticsearch.Client, indexName, docID string, doc interface{}) error {
	data, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("error marshaling document: %w", err)
	}

	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: docID,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	ctx := context.Background()

	res, err := req.Do(ctx, es)
	if err != nil {
		return fmt.Errorf("error indexing document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	return nil
}

func GetDocument(es *elasticsearch.Client, indexName, docID string) (map[string]interface{}, error) {
	req := esapi.GetRequest{
		Index:      indexName,
		DocumentID: docID,
	}

	ctx := context.Background()

	res, err := req.Do(ctx, es)
	if err != nil {
		return nil, fmt.Errorf("error getting document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	var doc map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %w", err)
	}

	return doc, nil
}

func DeleteDocument(es *elasticsearch.Client, indexName, docID string) error {
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: docID,
		Refresh:    "true",
	}

	ctx := context.Background()

	res, err := req.Do(ctx, es)
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	return nil
}

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	doc := map[string]interface{}{
		"name": "Jane Doe",
		"age":  27,
		"country": map[string]interface{}{
			"name": "United States",
			"code": "US",
		},
		"occupation": "Data Analyst",
	}

	err = CreateDocument(es, "users", "1", doc)
	if err != nil {
		log.Fatalf("Error creating document: %s", err)
	}

	log.Println("Document created successfully")

	doc, err = GetDocument(es, "users", "1")
	if err != nil {
		log.Fatalf("Error getting document: %s", err)
	}

	log.Printf("Document retrieved successfully: %+v", doc)

	err = DeleteDocument(es, "users", "1")
	if err != nil {
		log.Fatalf("Error deleting document: %s", err)
	}

	log.Println("Document deleted successfully")
}