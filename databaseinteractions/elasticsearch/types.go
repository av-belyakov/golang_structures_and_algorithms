package elasticsearchexamples

import "github.com/elastic/go-elasticsearch/v9"

type EsClient struct {
	client   *elasticsearch.Client
	settings Settings
}

type Settings struct {
	storages map[string]string
	cloudId  string
	user     string
	passwd   string
	apiKey   string
	host     string
	port     int
}

type Options func(*EsClient) error
