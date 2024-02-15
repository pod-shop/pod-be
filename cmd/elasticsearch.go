package main

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

// ElasticsearchService defines the service for interacting with Elasticsearch
type ElasticsearchService struct {
	Client *elasticsearch.Client
}

// NewElasticsearchService creates a new Elasticsearch service instance
func NewElasticsearchCLient() (*ElasticsearchService, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating the Elasticsearch client: %s", err)
	}

	return &ElasticsearchService{
		Client: es,
	}, nil
}
