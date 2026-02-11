package elasticsearchexamples

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
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
