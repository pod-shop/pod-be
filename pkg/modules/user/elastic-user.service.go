package user

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

// ElasticsearchService defines the service for interacting with Elasticsearch
type ElasticsearchService struct {
	Client *elasticsearch.Client
}

// NewElasticsearchService creates a new Elasticsearch service instance
func NewElasticsearchService(client *elasticsearch.Client) *ElasticsearchService {
	return &ElasticsearchService{
		Client: client,
	}
}

// GetDocumentByID retrieves a document from Elasticsearch by ID
func (s *ElasticsearchService) GetDocumentByID(indexName, docID string) (map[string]interface{}, error) {
	res, err := s.Client.Get(indexName, docID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error response: %s", res.Status())
	}

	var doc map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	return doc, nil
}
