package db

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var ElasticClient *elasticsearch.Client

func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal("❌ Elasticsearch connection failed:", err)
	}
	ElasticClient = client
	log.Println("✅ Connected to Elasticsearch")
}
