package elasticprovider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"plentylog/plentylog"

	"github.com/elastic/go-elasticsearch/v8"
)

// ElasticProvider is a provider for writing logs to ElasticSearch
// It implements the Provider interface from the plentylog package
// and uses the ElasticSearch client to write logs to an ElasticSearch index.
type ElasticProvider struct {
	client *elasticsearch.Client
}

// ElasticProviderOptions are the options for the ElasticProvider
type ElasticProviderOptions struct{}

// NewElasticProvider creates a new ElasticProvider
// It takes an optional ElasticProviderOptions struct
// and returns a pointer to the ElasticProvider and an error if any.
func NewElasticProvider(options *ElasticProviderOptions) (*ElasticProvider, error) {
	cert, _ := os.ReadFile("http_ca.crt")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "KYs3Cc0aUCDzSCK9C9t=",
		CACert:   cert,
	}

	client, err := elasticsearch.NewClient(cfg)

	if err != nil {
		return nil, err
	}

	// client.Indices.Create("logs")

	ep := ElasticProvider{
		client: client,
	}

	return &ep, nil
}

// Write writes a log record to ElasticSearch
func (ep *ElasticProvider) Write(ctx context.Context, r plentylog.Record) {
	fmt.Printf("Writing log to ElasticSearch: %v\n", r)

	data, _ := json.Marshal(r)

	ep.client.Index("logs", bytes.NewReader(data))
}
