package elasticsearchexamples

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9"
)

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

// CommonResponse общий ответ на запрос документов
type CommonResponse struct {
	Hits struct {
		Hits []struct {
			ID     string          `json:"_id"`
			Index  string          `json:"_index"`
			Score  float64         `json:"_score"`
			Source json.RawMessage `json:"_source"`
		} `json:"hits"`
		MaxScore float64 `json:"max_score"`
		Total    struct {
			Relation string `json:"relation"`
			Value    int    `json:"value"`
		} `json:"total"`
	} `json:"hits"`
	Shards struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
}
