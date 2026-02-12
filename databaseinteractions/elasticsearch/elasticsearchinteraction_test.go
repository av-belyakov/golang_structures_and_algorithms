package elasticsearchexamples_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	elasticsearchexamples "github.com/av-belyakov/golang_structures_and_algorithms/databaseinteractions/elasticsearch"
	"github.com/stretchr/testify/assert"
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
		es  *elasticsearchexamples.EsClient
		err error

		testIndexName     string = "index_test_documents"
		fullTestIndexName string = fmt.Sprintf("%s_%v", testIndexName, time.Now().UnixMicro())
	)

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

		t.Run("Тест 2.1. Добавляем новые документы в заданный индекс", func(t *testing.T) {
			for _, v := range GetExampleDocuments() {
				_, err := es.InsertDocument(t.Context(), fullTestIndexName, v)
				assert.NoError(t, err)
			}

			time.Sleep(1 * time.Second)
		})

		t.Run("Тест 2.2. Проверяем наличие заданного индекса", func(t *testing.T) {
			indexes, err := es.GetExistingIndexes(t.Context(), testIndexName)
			assert.NoError(t, err)
			assert.NotEmpty(t, indexes)
		})

		t.Run("Тест 2.3. Получаем настройки выбранного индекса", func(t *testing.T) {
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
			dslQuery := strings.NewReader(`{
				"index": {
					"mapping": {
						"total_fields": {
							"limit": "1999"
							}
						}
					}
				}`)

			_, err := es.SetIndexSetting(t.Context(), []string{fullTestIndexName}, dslQuery)
			assert.NoError(t, err)
		})

		t.Run("Тест 2.5. Ищем документ в определенном заданном индексе", func(t *testing.T) {
			source := "source_test_3"

			// формируем Query DSL запрос
			dslQuery := strings.NewReader(fmt.Sprintf(
				`{"query": 
		  			{"bool": 
		    			{"must": [
			  				{"match": {"source": "%s"}}, 
			  				{"match": {"event.rootId": "~91686760480"}}
						]}}}`, source))

			res, err := es.GetDocument(t.Context(), []string{fullTestIndexName}, dslQuery)
			assert.NoError(t, err)

			fmt.Println("get document response:", string(res))
		})

		t.Run(fmt.Sprintf("2.6. Добавляем новый документ в индекс '%s'", fullTestIndexName), func(t *testing.T) {
			rootId := "~9111111111"

			_, err := es.InsertDocument(t.Context(), fullTestIndexName, fmt.Appendf(nil, `{
					"source": "source_test_1",
    				"event": {
        				"operation": "update",
        				"rootId": "%s",
        				"startDate": 1745568135456
					}
				}`, rootId))
			assert.NoError(t, err)

			time.Sleep(1 * time.Second)

			// проверяем наличие добавленного документа
			dslQuery := strings.NewReader(fmt.Sprintf(
				`{"query": 
		  			{"bool": 
		    			{"must": [
			  				{"match": {"rootId": "%s"}}
						]}}}`, rootId))
			res, err := es.GetDocument(t.Context(), []string{fullTestIndexName}, dslQuery)
			assert.NoError(t, err)

			fmt.Println("get NEW document:", string(res))

			/*
				что то не находит опять
			*/

		})

		/*t.Run("Тест 2.5. Удаляем все индексы подходящие под определённый шаблон", func(t *testing.T) {
			indexes, err := es.GetExistingIndexes(t.Context(), testIndexName)
			assert.NoError(t, err)
			assert.NotEmpty(t, indexes)

			assert.NoError(t, es.DelIndexSetting(t.Context(), indexes))
		})*/
	})

	t.Cleanup(func() {
		es.ConnectionClose(t.Context())
	})
}

func GetExampleDocuments() [][]byte {
	return [][]byte{
		[]byte(`{
					"source": "source_test_1",
    				"event": {
        				"operation": "update",
        				"rootId": "~91686760480",
        				"startDate": 1745568135456,
						"name": "Elemnt_1"
					}}`),
		[]byte(`{
					"source": "source_test_2",
    				"event": {
        				"operation": "insert",
        				"rootId": "~11186760480",
        				"startDate": 1743368135445,
						"name": "Prizma_2"
					}}`),
		[]byte(`{
					"source": "source_test_1",
    				"event": {
        				"operation": "insert",
        				"rootId": "~91686760111",
        				"startDate": 1745568135655,
						"name": "Elemnt_2"
					}}`),
		[]byte(`{
					"source": "source_test_3",
    				"event": {
        				"operation": "insert",
        				"rootId": "~91686760480",
        				"startDate": 1745568135456,
						"name": "Catra"
					},
					"object": {
						"objectId": "~91686760480",
        				"organisation": "LV231",
        				"organisationId": "~4192",
        				"requestId": "50ae5f0dc59a6842:-3d107b9c:1963e52bc3d:-8000:1468002",
        				"objectType": "message"
					}}`),
		[]byte(`{
					"source": "source_test_3",
    				"event": {
        				"operation": "update",
        				"rootId": "~91686763728",
        				"startDate": 17454445434,
						"name": "Catra"
					},
					"object": {
						"objectId": "~91686760480",
        				"organisation": "LV672",
        				"organisationId": "~44552",
        				"requestId": "50ae5f0dc59a6842:-3d107b9c:1963e52bc3d:-8000:1468002",
        				"objectType": "message"
					}}`),
		[]byte(`{
					"source": "source_test_3",
    				"event": {
        				"operation": "insert",
        				"rootId": "~91686760480",
        				"startDate": 17343545332345,
						"name": "Catra"
					},
					"object": {
						"objectId": "~91686760480",
        				"organisation": "LV819",
        				"organisationId": "~433",
        				"requestId": "50ae5f0dc59a6842:-3d107b9c:1963e52bc3d:-8000:1468002",
        				"objectType": "message"
					}}`),
	}
}
