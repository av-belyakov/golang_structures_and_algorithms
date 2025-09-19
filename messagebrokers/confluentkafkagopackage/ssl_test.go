package confluentkafkagopackage_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"
)

func TestKafkaSSL(t *testing.T) {
	var (
		eventTopic string = "event-message.topic"
		alertTopic string = "alert-message.topic"
	)

	ac, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": "192.168.13.3:9093", //может содержать ip:port
		//"bootstrap.servers": "10.0.0.136", //"10.0.0.136:9092"
		"security.protocol": "ssl",
		// SSL сертификаты
		"ssl.ca.location":          "./kafkaimage/certs/ca.crt",
		"ssl.certificate.location": "./kafkaimage/certs/client.crt",
		"ssl.key.location":         "./kafkaimage/certs/client.key",
		"ssl.key.password":         os.Getenv("KAFKA_SSL_KEY_PASSWORD"),
		"client.id":                "go-kafka-ssl-admin-client",
		//"enable.ssl.certificate.verification": false,
	})

	assert.NoError(t, err)

	topics := []kafka.TopicSpecification{
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
	}

	result, err := ac.CreateTopics(t.Context(), topics)
	assert.NoError(t, err)

	fmt.Printf("Result:\n%#v\n", result)

	t.Cleanup(func() {
		ac.Close()
	})
}
