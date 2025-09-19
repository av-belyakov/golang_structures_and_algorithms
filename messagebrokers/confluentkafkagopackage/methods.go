package confluentkafkagopackage

import (
	"errors"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

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

//******************* функции настройки опций kafkaapi ***********************

// WithHost имя или ip адрес хоста API
func WithHost(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		n.settings.host = v

		return nil
	}
}

// WithPort порт API
func WithPort(v int) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		n.settings.port = v

		return nil
	}
}

// WithCacheTTL время жизни для кэша хранящего функции-обработчики запросов к модулю
func WithCacheTTL(v int) KafkaApiOptions {
	return func(th *KafkaApiModule) error {
		if v <= 10 || v > 86400 {
			return errors.New("the lifetime of a cache entry should be between 10 and 86400 seconds")
		}

		th.settings.cachettl = v

		return nil
	}
}

// WithNameRegionalObject наименование
func WithNameRegionalObject(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.nameRegionalObject = v

		return nil
	}
}

// WithTopicsSubscription 'слушатель' разных топиков
func WithTopicsSubscription(v map[string]string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		if len(v) == 0 {
			return errors.New("the value of 'topics' cannot be empty")
		}

		n.topics = v

		return nil
	}
}

// WithSSLKeyPassword пароль для приватного ключа в keystore
func WithSSLKeyPassword(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.sslKeyPassword = v

		return nil
	}
}

// WithSSLKeyStorePassword пароль для keystore kafka.keystore.jks
func WithSSLKeyStorePassword(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.sslKeyStorePassword = v

		return nil
	}
}

// WithSSLTruststorePassword пароль для truststore kafka.truststore.jks
func WithSSLTruststorePassword(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.sslTruststorePassword = v

		return nil
	}
}

// WithSSLKeyStoreLocation местоположение файла kafka.keystore.jks
func WithSSLKeyStoreLocation(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.sslKeyStoreLocation = v

		return nil
	}
}

// WithSSLTruststoreLocation местоположение файла kafka.truststore.jks
func WithSSLTruststoreLocation(v string) KafkaApiOptions {
	return func(n *KafkaApiModule) error {
		n.settings.sslTruststoreLocation = v

		return nil
	}
}
