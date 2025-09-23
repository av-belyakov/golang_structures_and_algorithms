package confluentkafkagopackage_test

import (
	"fmt"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"
)

func TestKafkaSSLAuth(t *testing.T) {
	var (
		ac *kafka.AdminClient

		eventTopic string = "event-message.topic"
		alertTopic string = "alert-message.topic"

		err error
	)

	t.Run("Тест 1. Установка соединения для выполнения административных действий.", func(t *testing.T) {
		ac, err = kafka.NewAdminClient(&kafka.ConfigMap{
			"bootstrap.servers": "127.0.0.1:9093", //может содержать ip:port
			//"bootstrap.servers": "10.0.0.136", //"10.0.0.136:9092"
			"security.protocol": "ssl",
			// SSL сертификаты
			"ssl.ca.location":                     "../kafkaimage/certs/ca.crt",
			"ssl.certificate.location":            "../kafkaimage/certs/client-cert.pem",
			"ssl.key.location":                    "../kafkaimage/certs/client-key.pem",
			"client.id":                           "go-kafka-ssl-admin-client",
			"enable.ssl.certificate.verification": false,
		})
		assert.NoError(t, err)

		t.Run("Тест 1.1. Создание группы топиков, если их нет.", func(t *testing.T) {
			result, err := ac.CreateTopics(
				t.Context(),
				[]kafka.TopicSpecification{
					{
						Topic:         alertTopic,
						NumPartitions: 1,
						//ReplicationFactor: 1,
					},
					{
						Topic:         eventTopic,
						NumPartitions: 1,
						//ReplicationFactor: 1,
					},
				})
			assert.NoError(t, err)

			t.Log("=== Result topics created ===")
			for _, v := range result {
				fmt.Printf("%s\n", v.String())
			}
		})

		t.Run("Тест 1.2. Проверка наличия топиков.", func(t *testing.T) {
			kafMetadata, err := ac.GetMetadata(nil, true, 1_000)
			assert.NoError(t, err)
			assert.NotEqual(t, kafMetadata.Topics, 0)

			t.Log("=== All topics ===")
			for k, v := range kafMetadata.Topics {
				fmt.Printf("Topic '%s': %s\n", k, v.Error.String())
			}
		})

		t.Cleanup(func() {
			ac.Close()
		})
	})
}
