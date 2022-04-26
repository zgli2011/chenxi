package initialize

import (
	"log"

	"github.com/elastic/go-elasticsearch"
)

var ESClient *elasticsearch.Client

func NewESClient() *elasticsearch.Client {
	ESClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: Config.Elasticsearch,
	})
	if err != nil {
		log.Panicf(err.Error())
	}
	return ESClient
}
