package confluentkafkagopackage

import (
	"context"
	"fmt"
	"maps"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Start инициализирует новый модуль взаимодействия с API Kafka, принимает функцию-обработчик
// входящих, через Kafka, событий. При инициализации возращается канал для взаимодействия
// с модулем, все запросы к модулю выполняются через данный канал.
func (api *KafkaApiModule) Start(ctx context.Context, handlerFunc func() func(context.Context, *KafkaApiModule)) error {
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

	// выполняем после отмены контекста или истечении времени контекста
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
	// в данном случае обрабочик добавляется путём передачи вспомогательной функции
	// но можно сделать обработчик одним из методов KafkaApiModule
	go hf(ctx, api)

	return nil
}

// GetChanInput канал для передачи данных в модуль
func (api *KafkaApiModule) GetChanInput() chan ChInputSettings {
	return api.chToModule
}

// GetChanOutput канал для приёма данных из модуля
func (api *KafkaApiModule) GetChanOutput() chan ChOutputSettings {
	return api.chFromModule
}

// GetConsumer потребитель для kafka
func (api *KafkaApiModule) GetConsumer() *kafka.Consumer {
	return api.consumer
}

// GetTopics топики Kafka
func (api *KafkaApiModule) GetTopics() map[string]string {
	return api.topics
}

// GetLogger метод логирования (ВСПОМОГАТЕЛЬНЫЙ МЕТОД, ДЛЯ ТЕСТА)
func (api *KafkaApiModule) GetLogger() Logger {
	return api.logger
}
