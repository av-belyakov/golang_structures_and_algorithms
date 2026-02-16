package elasticsearchexamples_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"

	elasticsearchexamples "github.com/av-belyakov/golang_structures_and_algorithms/databaseinteractions/elasticsearch"
)

const (
	Host = "localhost"
	Port = 9200
	// получаем после создания пользователя
	User   = "development"
	Passwd = "7kY-bQv-HNy-vwx"
	// получаем после создания ApiKey
	ApiKey = "VTkxclRKd0J3WGVtNzJadzlnamk6ZTBMU3BzSTdoWmc0NmpXNUkzM3Mxdw=="
)

func TestElasticsearchIntaraction(t *testing.T) {
	var (
		es             *elasticsearchexamples.EsClient
		commonResponse *elasticsearchexamples.CommonResponse
		err            error

		exampleRootId string = "~9111111111"

		testIndexName     string = "index_test_documents"
		fullTestIndexName string = fmt.Sprintf("%s_%v", testIndexName, time.Now().UnixMicro())
		documentId        string
	)

	responseDocumentId := struct {
		Source struct {
			Source    string `json:"source"`
			PartyDate string `json:"party_date"`
			Event     struct {
				Operation string `json:"operation"`
				RootId    string `json:"rootId"`
				StartDate int64  `json:"startDate"`
			} `json:"event"`
			Hobby struct {
				Name     string `json:"name"`
				Weekday  string `json:"weekday"`
				Time     string `json:"time"`
				Duration string `json:"duration"`
			} `json:"hobby"`
			AboutOue struct {
				Name    string `json:"name"`
				Address string `json:"address"`
				Age     int    `json:"age"`
			} `json:"aboutOur"`
		} `json:"_source"`
		Index       string `json:"_index"`
		Id          string `json:"_id"`
		Version     int    `json:"_version"`
		SeqNo       int    `json:"_seq_no"`
		PrimaryTerm int    `json:"_primary_term"`
		Found       bool   `json:"found"`
	}{}

	t.Run("Тест 1.1. Подключение к БД с использованием API-ключа.", func(t *testing.T) {
		es, err := elasticsearchexamples.NewClient(
			elasticsearchexamples.WithApiKey(ApiKey),
			elasticsearchexamples.WithHost(Host),
			elasticsearchexamples.WithPort(Port),
			elasticsearchexamples.WithStorage(map[string]string{
				"index-test": "index_for_test",
			}),
		)
		assert.NoError(t, err)
		assert.NoError(t, es.Connect(t.Context()))
		assert.NoError(t, es.ConnectionClose(t.Context()))
	})

	t.Run("Тест 1.2. Подключение к БД с использованием имени пользователя и пароля.", func(t *testing.T) {
		es, err := elasticsearchexamples.NewClient(
			elasticsearchexamples.WithUser(User),
			elasticsearchexamples.WithPasswd(Passwd),
			elasticsearchexamples.WithHost(Host),
			elasticsearchexamples.WithPort(Port),
			elasticsearchexamples.WithStorage(map[string]string{
				"index-test": "index_for_test",
			}),
		)
		assert.NoError(t, err)
		assert.NoError(t, es.Connect(t.Context()))
		assert.NoError(t, es.ConnectionClose(t.Context()))
	})

	t.Run("Тест 2. Выполнение различных запросов к БД.", func(t *testing.T) {
		es, err = elasticsearchexamples.NewClient(
			elasticsearchexamples.WithApiKey(ApiKey),
			elasticsearchexamples.WithHost(Host),
			elasticsearchexamples.WithPort(Port),
			elasticsearchexamples.WithStorage(map[string]string{
				"index-test": "index_for_test",
			}),
		)
		assert.NoError(t, err)
		assert.NoError(t, es.Connect(t.Context()))

		t.Run("2.1. Добавляем новые документы в заданный индекс", func(t *testing.T) {
			for _, v := range GetExampleDocuments() {
				_, err := es.InsertDocument(t.Context(), fullTestIndexName, v)
				assert.NoError(t, err)
			}

			time.Sleep(1 * time.Second)
		})

		t.Run("2.2. Проверяем наличие заданного индекса", func(t *testing.T) {
			indexes, err := es.GetExistingIndexes(t.Context(), testIndexName)
			assert.NoError(t, err)
			assert.NotEmpty(t, indexes)
		})

		t.Run("2.3. Получаем настройки выбранного индекса", func(t *testing.T) {
			indexSettings, err := es.GetIndexSetting(t.Context(), fullTestIndexName)
			assert.NoError(t, err)
			assert.NotEmpty(t, indexSettings)
		})

		t.Run(fmt.Sprintf("2.4. Устанавливаем новые настройки для индекса '%s'", fullTestIndexName), func(t *testing.T) {
			// Устанавливаем максимальный лимит полей для выбранного индекса в 1999.
			// Это позволяет убрать ошибку 'Elasticsearch типа Limit of total
			// fields [1000] has been exceeded while adding new fields' которая
			// возникает при установленном максимальном количестве полей в 1000, что
			// соответствует дефолтному значению.

			// Следует обратить внимание что при установки этого параметра он переезжает
			// из объекта 'defaults' настроек в 'settings'
			_, err := es.SetIndexSetting(
				t.Context(),
				[]string{fullTestIndexName},
				strings.NewReader(`{
				"index": {
					"mapping": {
						"total_fields": {
							"limit": "1999"
							}
						}
					}
				}`))
			assert.NoError(t, err)
		})

		t.Run("2.5. Ищем документ в определенном заданном индексе", func(t *testing.T) {
			// формируем Query DSL запрос
			res, err := es.GetDocument(
				t.Context(),
				[]string{fullTestIndexName},
				strings.NewReader(`{"query": 
		  			{"bool": 
		    			{"must": [
			  				{"match": {"source": "source_test_3"}}, 
			  				{"match": {"event.rootId": "~91686760480"}}
						]}}}`))
			assert.NoError(t, err)

			commonResponse = &elasticsearchexamples.CommonResponse{}
			assert.NoError(t, json.Unmarshal(res, commonResponse))
			assert.Equal(t, len(commonResponse.Hits.Hits), 2)
			assert.Equal(t, commonResponse.Hits.Total.Value, 2)

			//fmt.Println("2.5. Found document:")
			//godump.NewDumper().DumpJSON(commonResponse)
		})

		t.Run(fmt.Sprintf("2.6. Добавляем новый документ в индекс '%s'", fullTestIndexName), func(t *testing.T) {
			_, err := es.InsertDocument(
				t.Context(),
				fullTestIndexName,
				fmt.Appendf(nil, `{
					"source": "source_test_1",
    				"event": {
        				"operation": "create",
        				"rootId": "%s",
        				"startDate": 1745568135456
					}
				}`, exampleRootId))
			assert.NoError(t, err)

			time.Sleep(1_000 * time.Millisecond)

			// проверяем наличие добавленного документа
			res, err := es.GetDocument(
				t.Context(),
				[]string{fullTestIndexName},
				strings.NewReader(
					fmt.Sprintf(
						`{"query": 
							{"bool": 
								{"must": [
									{"match": {
										"event.rootId": "%s"}}]}}}`,
						exampleRootId)))
			assert.NoError(t, err)

			commonResponse = &elasticsearchexamples.CommonResponse{}
			assert.NoError(t, json.Unmarshal(res, commonResponse))
			assert.Equal(t, len(commonResponse.Hits.Hits), 1)
			assert.Equal(t, commonResponse.Hits.Total.Value, 1)

			documentId = commonResponse.Hits.Hits[0].ID

			//fmt.Println("2.6. Found document:")
			//godump.DumpJSON(commonResponse)
		})

		t.Run(fmt.Sprintf("2.7. Изменяем некоторые параметры документа с rootId '%s' для индекса '%s'", exampleRootId, fullTestIndexName), func(t *testing.T) {
			// обновление документа в Elasticsearch возможно с помощью запуска скрипта
			// или путём передачи частичного документа подробнее
			// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-update и
			// https://www.elastic.co/docs/reference/elasticsearch/rest-apis/update-document
			// для полной замены документа нужно использовать индексный метод API
			// то есть выполнить InsertDocument но возможно понадобятся дополнительные настройки
			// метода, подробнее по настройкам https://pkg.go.dev/github.com/elastic/go-elasticsearch/v9/esapi#Index

			t.Run("2.7.1. Изменение параметров документа с помощью локального скрипта Elasticsearch", func(t *testing.T) {
				statusCode, err := es.UpdateDocument(
					t.Context(),
					fullTestIndexName,
					documentId,
					strings.NewReader(`{
						"script" : {
							"source": "ctx._source.event.operation = params.operation",
							"lang": "painless",
							"params" : {
								"operation" : "update"
							}
						}
					}`),
				)
				assert.NoError(t, err)
				assert.Equal(t, statusCode, http.StatusOK, fmt.Sprintf(
					"Ошибка при обновлении документа с rootId '%s' в индексе '%s', код ошибки '%d'\n",
					exampleRootId,
					fullTestIndexName,
					statusCode,
				))

				// проверяем изменение документа
				res, statusCode, err := es.GetDocumentById(t.Context(), fullTestIndexName, documentId)
				assert.NoError(t, err)
				assert.Equal(t, statusCode, http.StatusOK)

				assert.NoError(t, json.Unmarshal(res, &responseDocumentId))
				assert.Equal(t, responseDocumentId.Source.Event.Operation, "update")
			})

			t.Run("2.7.2. Изменение параметров документа путём частичной замены документа", func(t *testing.T) {
				// для частичного изменения документа нужно использовать конструкцию
				// "doc": { ... }
				// подробнее https://www.elastic.co/docs/reference/elasticsearch/rest-apis/update-document
				statusCode, err := es.UpdateDocument(
					t.Context(),
					fullTestIndexName,
					documentId,
					strings.NewReader(`{
						"doc": {
							"hobby": {
								"name": "MotoCross",
								"weekday": "Sunday",
								"time": "10:05:00",
								"duration": "3h"
							}
						}
					}`),
				)
				assert.NoError(t, err, fmt.Sprintf("update error: '%+v'", err))
				assert.Equal(t, statusCode, http.StatusOK, fmt.Sprintf(
					"Ошибка при обновлении документа с rootId '%s' в индексе '%s', код ошибки '%d'\n",
					exampleRootId,
					fullTestIndexName,
					statusCode,
				))

				// выполняем обновление если значение найдено или добавляем
				// значение если оно не было найдено
				statusCode, err = es.UpdateDocument(
					t.Context(),
					fullTestIndexName,
					documentId,
					strings.NewReader(`{
						"doc": {
							"hobby": {
								"duration": "4h15m"
							}
						},
						"doc_as_upsert": true
					}`),
				)

				// проверяем изменение документа
				res, statusCode, err := es.GetDocumentById(t.Context(), fullTestIndexName, documentId)
				assert.NoError(t, err)
				assert.Equal(t, statusCode, http.StatusOK)

				//fmt.Println("Document id:", documentId)
				//fmt.Println("Response:", string(res))

				assert.NoError(t, json.Unmarshal(res, &responseDocumentId))
				assert.Equal(t, responseDocumentId.Source.Hobby.Name, "MotoCross")
				assert.Equal(t, responseDocumentId.Source.Hobby.Duration, "4h15m")
			})
		})

		t.Run(fmt.Sprintf("2.8. Изменяем некоторые параметры во всех документах в списке индексов, для теста тольоко индекс с имененм '%s'", fullTestIndexName), func(t *testing.T) {
			partyDate := "2026-02-23 15:00:00"
			source := "source_test_3"

			statusCode, err := es.UpdateDocuments(
				t.Context(),
				[]string{fullTestIndexName},
				strings.NewReader(fmt.Sprintf(`{
					"script": {
						"source": "ctx._source.party_date = params.party_date",
						"lang": "painless",
						"params" : {
							"party_date": "%s"
						}
					},
					"query": {
						"term": {
							"source": "%s"
						}
					}
				}`, partyDate, source)),
			)
			assert.NoError(t, err)
			assert.Equal(t, statusCode, http.StatusOK)

			time.Sleep(1_000 * time.Millisecond)

			res, err := es.GetDocument(
				t.Context(),
				[]string{fullTestIndexName},
				strings.NewReader(fmt.Sprintf(`
					{"query": 
						{"bool": 
							{"must": [
								{"match": {"party_date": "%s"}}]}}}
							`, partyDate)),
			)
			assert.NoError(t, err)

			fmt.Println("___ Updated documents:", string(res))

			commonResponse = &elasticsearchexamples.CommonResponse{}
			assert.NoError(t, json.Unmarshal(res, commonResponse))
			assert.Equal(t, commonResponse.Hits.Total.Value, 3)
		})
	})

	/*t.Run("Тест 3. Удаляем все индексы подходящие под определённый шаблон", func(t *testing.T) {
		indexes, err := es.GetExistingIndexes(t.Context(), testIndexName)
		assert.NoError(t, err)
		assert.NotEmpty(t, indexes)

		assert.NoError(t, es.DelIndexSetting(t.Context(), indexes))
	})

	t.Cleanup(func() {
		es.ConnectionClose(t.Context())
	})*/
}

func GetExampleDocuments() [][]byte {
	return [][]byte{
		fmt.Appendf(nil, `{
					"source": "source_test_1",
    				"event": {
        				"operation": "update",
        				"rootId": "~91686760480",
        				"startDate": %d,
						"name": "%s"
					}}`,
			gofakeit.Date().Unix(),
			gofakeit.Name()),
		fmt.Appendf(nil, `{
					"source": "source_test_2",
    				"event": {
        				"operation": "insert",
        				"rootId": "~11186760480",
        				"startDate": %d,
						"name": "%s"
					}}`,
			gofakeit.Date().Unix(),
			gofakeit.Name()),
		fmt.Appendf(nil, `{
					"source": "source_test_1",
    				"event": {
        				"operation": "insert",
        				"rootId": "~91686760111",
        				"startDate": %d,
						"name": "%s"
					}}`,
			gofakeit.Date().Unix(),
			gofakeit.Name()),
		fmt.Appendf(nil, `{
					"source": "source_test_3",
    				"event": {
        				"operation": "insert",
        				"rootId": "~91686760480",
        				"startDate": %d,
						"name": "%s"
					},
					"object": {
						"objectId": "~91686760480",
        				"organisation": "LV231",
        				"organisationId": "~4192",
        				"requestId": "50ae5f0dc59a6842:-3d107b9c:1963e52bc3d:-8000:1468002",
        				"objectType": "message"
					}}`,
			gofakeit.Date().Unix(),
			gofakeit.Name()),
		fmt.Appendf(nil, `{
					"source": "source_test_3",
    				"event": {
        				"operation": "update",
        				"rootId": "~91686763728",
        				"startDate": %d,
						"name": "%s"
					},
					"object": {
						"objectId": "~91686760480",
        				"organisation": "LV672",
        				"organisationId": "~44552",
        				"requestId": "50ae5f0dc59a6842:-3d107b9c:1963e52bc3d:-8000:1468002",
        				"objectType": "message"
					}}`,
			gofakeit.Date().Unix(),
			gofakeit.Name()),
		fmt.Appendf(nil, `{
					"source": "source_test_3",
    				"event": {
        				"operation": "insert",
        				"rootId": "~91686760480",
        				"startDate": %d,
						"name": "%s"
					},
					"object": {
						"objectId": "~91686760480",
        				"organisation": "LV819",
        				"organisationId": "~433",
        				"requestId": "50ae5f0dc59a6842:-3d107b9c:1963e52bc3d:-8000:1468002",
        				"objectType": "message"
					}}`,
			gofakeit.Date().Unix(),
			gofakeit.Name()),
	}
}
