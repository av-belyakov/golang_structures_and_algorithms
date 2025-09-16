package confluentkafkagopackage

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// KafkaApiModule настройки для API Kafka
type KafkaApiModule struct {
	counter      Counter
	logger       Logger
	consumer     *kafka.Consumer
	topics       map[string]string
	settings     kafkaApiSettings
	chFromModule chan ChOutputSettings
	chToModule   chan ChInputSettings
}

// kafkaApiSettings настройки
type kafkaApiSettings struct {
	nameRegionalObject string
	host               string
	cachettl           int
	port               int
}

// KafkaApiOptions функциональные опции
type KafkaApiOptions func(*KafkaApiModule) error

// ChOutputSettings канал вывода данных из модуля
type ChOutputSettings struct {
	TopicType string
	Data      any
}

// ChInputSettings канал для приема данных в модуль
type ChInputSettings struct {
	TopicType string
	Data      any
}
