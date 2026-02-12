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

// ServiseOption сервисные опции
type ServiseOption struct {
	ID    string `json:"_id"`
	Index string `json:"_index"`
}

type IndexesSettings map[string]struct {
	Settings struct {
		Index struct {
			CreationDateString string `json:"creation_date_string"`
			Routing            struct {
				Allocation struct {
					Include map[string]string `json:"include"`
				} `json:"allocation"`
			} `json:"routing"`
			NumberShards   string `json:"number_of_shards"`
			ProvidedName   string `json:"provided_name"`
			CreationDate   string `json:"creation_date"`
			NumberReplicas string `json:"number_of_replicas"`
			UUID           string `json:"uuid"`
			Version        struct {
				CreatedString string `json:"created_string"`
				Created       string `json:"created"`
			} `json:"version"`
		} `json:"index"`
	} `json:"settings"`
}

/*
{
  "index_test_documents_1770886633258305" : {
    "settings" : {
      "index" : {
        "creation_date_string" : "2026-02-12T08:57:13.263Z",
        "routing" : {
          "allocation" : {
            "include" : {
              "_tier_preference" : "data_content"
            }
          }
        },
        "number_of_shards" : "1",
        "provided_name" : "index_test_documents_1770886633258305",
        "creation_date" : "1770886633263",
        "number_of_replicas" : "1",
        "uuid" : "jLNp-xC6QF-zsrfzK9juXQ",
        "version" : {
          "created_string" : "9.3.0",
          "created" : "9060000"
        }
      }
    }
  }
}
*/
