package confluentkafkagopackage

import (
	"context"
	"fmt"
	"maps"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// NewWithSSL настраивает модуль взаимодействия с API Kafka с использование SSL
func NewWithSSL(counter Counter, logger Logger, opts ...KafkaApiOptions) (*KafkaApiModule, error) {
	api := &KafkaApiModule{
		counter: counter,
		logger:  logger,
		settings: kafkaApiSettings{
			cachettl: 15,
		},
		chFromModule: make(chan ChOutputSettings),
		chToModule:   make(chan ChInputSettings),
	}

	for _, opt := range opts {
		if err := opt(api); err != nil {
			return api, err
		}
	}

	return api, nil
}

// StartWithSSL инициализирует новый модуль взаимодействия с API Kafka по SSL протоколу,
// принимает функцию-обработчик входящих, через Kafka, событий. При инициализации
// возращается канал для взаимодействия с модулем, все запросы к модулю выполняются
// через данный канал.
func (api *KafkaApiModule) StartWithSSL(ctx context.Context, handlerFunc func() func(context.Context, *KafkaApiModule)) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	/*
		kafpkg.WithSSLKeyPassword(os.Getenv("KAFKA_SSL_KEY_PASSWORD")),
		kafpkg.WithSSLKeyStorePassword(os.Getenv("KAFKA_SSL_KEYSTORE_PASSWORD")),
		kafpkg.WithSSLTruststorePassword(os.Getenv("KAFKA_SSL_TRUSTSTORE_PASSWORD")),
		kafpkg.WithSSLKeyStoreLocation("./kafkaimage/keys/kafka.keystore.jks"),
		kafpkg.WithSSLTruststoreLocation("./kafkaimage/keys/kafka.truststore.jks"),
	*/

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%d", api.settings.host, api.settings.port),
		"security.protocol": "ssl",
		// SSL сертификаты
		"ssl.ca.location":          "./kafkaimage/certs/ca.crt ",
		"ssl.certificate.location": "./kafkaimage/certs/client.crt",
		"ssl.key.location":         "./kafkaimage/certs/client.key",
		"ssl.key.password":         api.settings.sslKeyPassword,
		// Настройки потребителя
		"group.id":          fmt.Sprintf("%s-group", api.settings.nameRegionalObject), // Идентификатор группы
		"auto.offset.reset": "earliest",                                               // Читать с начала
		"client.id":         "go-kafka-ssl-consumer",
	})
	if err != nil {
		return err
	}
	api.consumer = consumer
	context.AfterFunc(ctx, func() {
		consumer.Close()

		close(api.chToModule)
		close(api.chFromModule)
	})

	var topics []string
	mapTopics := maps.Values(api.topics)
	for topic := range mapTopics {
		topics = append(topics, topic)
	}

	// подписка на топик
	err = api.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	//обработчик подписок
	hf := handlerFunc()
	go hf(ctx, api)
	// в данном случае обрабочик добавляется путём передачи вспомогательной функции
	// но можно сделать обработчик одним из методов KafkaApiModule

	return nil
}
