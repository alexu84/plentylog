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

type ElasticProvider struct {
	client *elasticsearch.Client
}

type ElasticProviderOptions struct{}

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

func (ep *ElasticProvider) Write(ctx context.Context, r plentylog.Record) {
	fmt.Printf("Writing log to ElasticSearch: %v\n", r)

	data, _ := json.Marshal(r)

	ep.client.Index("logs", bytes.NewReader(data))
}
