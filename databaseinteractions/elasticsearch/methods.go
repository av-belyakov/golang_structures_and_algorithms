package elasticsearchexamples

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"

	"github.com/av-belyakov/golang_structures_and_algorithms/supportfunctions"
)

// Connect соединение с БД
func (es *EsClient) Connect(ctx context.Context) error {
	conf := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%d", es.settings.host, es.settings.port),
		},
		Transport: &http.Transport{
			MaxIdleConns:          10,              //число открытых TCP-соединений, которые в данный момент не используются
			IdleConnTimeout:       1 * time.Second, //время, через которое закрываются такие неактивные соединения
			MaxIdleConnsPerHost:   10,              //число неактивных TCP-соединений, которые допускается устанавливать на один хост
			ResponseHeaderTimeout: 2 * time.Second, //время в течении которого сервер ожидает получение ответа после записи заголовка запроса
			DialContext: (&net.Dialer{
				Timeout: 3 * time.Second,
				//KeepAlive: 1 * time.Second,
			}).DialContext,
		}}

	if es.settings.apiKey != "" {
		conf.APIKey = es.settings.apiKey
	} else if es.settings.user != "" && es.settings.passwd != "" {
		conf.Username = es.settings.user
		conf.Password = es.settings.passwd
	} else {
		return errors.New("the value of 'apiKey' or 'user' and 'passwd' cannot be empty")
	}

	client, err := elasticsearch.NewClient(conf)
	if err != nil {
		return err
	}

	es.client = client

	return nil
}

// ConnectionClose закрытие соединения с БД
func (es *EsClient) ConnectionClose(ctx context.Context) error {
	if es.client == nil {
		return errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly")
	}

	return es.client.Close(ctx)
}

// GetExistingIndexes проверка наличия индексов соответствующих определенному шаблону
func (es *EsClient) GetExistingIndexes(ctx context.Context, pattern string) ([]string, error) {
	if es.client == nil {
		return []string{}, errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly")
	}

	ctxTimeout, ctxCancel := context.WithTimeout(ctx, time.Second*15)
	defer ctxCancel()

	listIndexes := []string(nil)
	msg := []struct {
		Index string `json:"index"`
	}(nil)

	res, err := es.client.Cat.Indices(
		es.client.Cat.Indices.WithContext(ctxTimeout),
		es.client.Cat.Indices.WithFormat("json"),
	)
	if err != nil {
		return nil, err
	}
	defer bodyClose(res)

	if err = json.NewDecoder(res.Body).Decode(&msg); err != nil {
		return nil, err
	}

	for _, v := range msg {
		if !strings.Contains(v.Index, pattern) {
			continue
		}

		listIndexes = append(listIndexes, v.Index)
	}

	return listIndexes, err
}

// GetIndexSetting параметры и настройки выбранного индекса
// настройки только settings, настроек defaults нет
func (es *EsClient) GetIndexSetting(ctx context.Context, index string) (IndexesSettings, error) {
	var indexesSettings IndexesSettings

	req := esapi.IndicesGetSettingsRequest{
		Index:  []string{index},
		Pretty: true,
		Human:  true,
	}

	response, err := req.Do(ctx, es.client.Transport)
	if err != nil {
		return indexesSettings, err
	}
	defer bodyClose(response)

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("the server response when executing an index search query is equal to '%s'", response.Status())

		return indexesSettings, err
	}

	res, err := io.ReadAll(response.Body)
	if err != nil {
		return indexesSettings, err
	}

	err = json.Unmarshal(res, &indexesSettings)
	if err != nil {
		return indexesSettings, err
	}

	return indexesSettings, nil
}

// SetIndexSetting установить новые настройки индекса
func (es *EsClient) SetIndexSetting(ctx context.Context, indexes []string, body *strings.Reader) (bool, error) {
	indicesSettings := esapi.IndicesPutSettingsRequest{
		Index: indexes,
		Body:  body,
	}

	res, err := indicesSettings.Do(ctx, es.client.Transport)
	if err != nil {
		return false, err
	}
	defer bodyClose(res)

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return true, nil
	}

	r := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return true, err
	}

	if e, ok := r["error"]; ok {
		return true, fmt.Errorf("received from module Elsaticsearch: %s (%s)", res.Status(), e)
	}

	return false, nil
}

// DelIndexSetting удаление индекса и его настроек
func (es *EsClient) DelIndexSetting(ctx context.Context, indexes []string) error {
	req := esapi.IndicesDeleteRequest{Index: indexes}
	res, err := req.Do(ctx, es.client.Transport)
	if err != nil {
		return err
	}
	defer bodyClose(res)

	return err
}

// GetDocumentById поиск документа по его _id
func (es *EsClient) GetDocumentById(ctx context.Context, index string, id string) ([]byte, int, error) {
	if es.client == nil {
		return []byte{}, 0, errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly")
	}

	res, err := es.client.Get(index, id, es.client.Get.WithContext(ctx))
	if err != nil {
		return []byte{}, 0, err
	}
	defer bodyClose(res)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, res.StatusCode, err
	}

	return body, res.StatusCode, nil
}

// GetDocument поиск документа в определенном заданном индексе
func (es *EsClient) GetDocument(ctx context.Context, indexes []string, body *strings.Reader) ([]byte, error) {
	if es.client == nil {
		return []byte{}, errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly")
	}

	var res []byte

	ctxTimeout, ctxCancel := context.WithTimeout(ctx, time.Second*15)
	defer ctxCancel()

	response, err := es.client.Search(
		es.client.Search.WithContext(ctxTimeout),
		es.client.Search.WithIndex(indexes...),
		es.client.Search.WithBody(body),
	)
	if err != nil {
		return res, err
	}
	defer bodyClose(response)

	res, err = io.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	return res, nil
}

// InsertDocument добавить новый документ в заданный индекс
func (es *EsClient) InsertDocument(ctx context.Context, index string, b []byte) (int, error) {
	if es.client == nil {
		return 0, errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly")
	}

	buf := bytes.NewReader(b)
	res, err := es.client.Index(index, buf)
	if err != nil {
		return 0, err
	}
	defer bodyClose(res)

	bodyRes, err := io.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, err
	}

	result, err := supportfunctions.GetElementsFromJSON(ctx, bodyRes)
	if err != nil {
		return res.StatusCode, err
	}

	for k, v := range result {
		if strings.Contains(k, "error") {
			return res.StatusCode, errors.New(fmt.Sprint(v.Value))
		}
	}

	return res.StatusCode, nil
}

// UpdateDocument обновление документа
func (es *EsClient) UpdateDocument(ctx context.Context, index string, id string, body io.Reader) (int, error) {
	if es.client == nil {
		return 0, errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly")
	}

	ctx, ctxCancel := context.WithTimeout(ctx, time.Second*10)
	defer ctxCancel()

	reg, err := es.client.Update(index, id, body, es.client.Update.WithContext(ctx))
	if err != nil {
		return 0, err
	}
	defer bodyClose(reg)

	return reg.StatusCode, nil
}

// UpdateDocuments обновление списка документов которые совпадают с запросом
func (es *EsClient) UpdateDocuments(ctx context.Context, indexes []string, body io.Reader) (int, error) {
	if es.client == nil {
		return 0, errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly")
	}

	ctx, ctxCancel := context.WithTimeout(ctx, time.Second*10)
	defer ctxCancel()

	// подробнее по всем параметра м можно посмотреть здесь
	// https://pkg.go.dev/github.com/elastic/go-elasticsearch/v9/esapi#UpdateByQuery
	reg, err := es.client.UpdateByQuery(indexes,
		es.client.UpdateByQuery.WithContext(ctx),
		// разрешить подсчёт конфликтов версий вместо остановки и возврата к предыдущей версии, установив
		es.client.UpdateByQuery.WithConflicts("proceed"),
		es.client.UpdateByQuery.WithBody(body),
	)
	if err != nil {
		return 0, err
	}
	defer bodyClose(reg)

	return reg.StatusCode, nil
}

// UpdateDocumentWithReplacement поиск и обновление документов путём замены
func (es *EsClient) UpdateDocumentWithReplacement(ctx context.Context, currentIndex string, list []ServiseOption, document []byte) (statusCode, countDel int, err error) {
	for _, v := range list {
		res, errDel := es.client.Delete(v.Index, v.ID)
		if errDel != nil {
			err = fmt.Errorf("%v, %v", err, errDel)
		}
		bodyClose(res)

		countDel++
	}

	statusCode, err = es.InsertDocument(ctx, currentIndex, document)

	return statusCode, countDel, err
}

// bodyClose закрывает ответ с предварительной проверкой
func bodyClose(res *esapi.Response) {
	if res == nil || res.Body == nil {
		return
	}

	res.Body.Close()
}
