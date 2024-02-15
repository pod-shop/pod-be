package elasticsearch

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

// NewElasticsearchClient creates a new Elasticsearch client instance
func NewElasticsearchClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating the Elasticsearch client: %s", err)
	}

	return es, nil
}
