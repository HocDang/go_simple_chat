package bootstrap

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func InitElasticsearch(host string, port string) *elasticsearch.Client {
	// Cấu hình client
	cfg := elasticsearch.Config{
		Addresses: []string{host + ":" + port},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal("❌ Elasticsearch connection failed :", err)
		panic(err)
	}

	res, err := client.Info()
	if err != nil {
		log.Fatal("❌ Unable to connect to Elasticsearch:", err)
		panic(err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatal("❌ Elasticsearch returned an error:", res.String())
		panic("Elasticsearch error response")
	}

	log.Println("✅ Connected to Elasticsearch")
	return client
}
