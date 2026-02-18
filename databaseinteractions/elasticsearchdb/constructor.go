package elasticsearchexamples

import (
	"errors"
)

// NewClient конструктор клиента
// opts - набор опций для настройки клиента в том числе авторизационных данных
// при инициализации преимущество отдаётся ApiKey, если это значение пустое
// будет использоватся стандартная авторизация через Basic Auth
func NewClient(opts ...Options) (*EsClient, error) {
	esClient := &EsClient{
		settings: Settings{
			storages: make(map[string]string),
		},
	}

	for _, opt := range opts {
		if err := opt(esClient); err != nil {
			return esClient, err
		}
	}

	return esClient, nil
}

// WithHost имя или ip адрес хоста API
func WithHost(v string) Options {
	return func(esClient *EsClient) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		esClient.settings.host = v

		return nil
	}
}

// WithPort порт API
func WithPort(v int) Options {
	return func(esClient *EsClient) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		esClient.settings.port = v

		return nil
	}
}

// WithApiKey api ключ
func WithApiKey(v string) Options {
	return func(esClient *EsClient) error {
		if v == "" {
			return errors.New("the value of 'ApiKey' cannot be empty")
		}

		esClient.settings.apiKey = v

		return nil
	}
}

// WithUser имя пользователя для доступа к БД
func WithUser(v string) Options {
	return func(esClient *EsClient) error {
		if v == "" {
			return errors.New("the value of 'user' cannot be empty")
		}

		esClient.settings.user = v

		return nil
	}
}

// WithPasswd пароль для доступа к БД
func WithPasswd(v string) Options {
	return func(esClient *EsClient) error {
		if v == "" {
			return errors.New("the value of 'passwd' cannot be empty")
		}

		esClient.settings.passwd = v

		return nil
	}
}

// WithStorage наименование коллекции или индекса БД
func WithStorage(v map[string]string) Options {
	return func(esClient *EsClient) error {
		esClient.settings.storages = v

		return nil
	}
}
