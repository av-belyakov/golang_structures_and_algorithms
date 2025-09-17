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

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%d", api.settings.host, api.settings.port),
		"group.id":          fmt.Sprintf("%s-group", api.settings.nameRegionalObject), // Идентификатор группы
		"auto.offset.reset": "earliest",                                               // Читать с начала
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
