package kafkagopackage_test

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func TestKafkaConnectionSSL(t *testing.T) {
	publicCert, err := os.ReadFile("./kafkaimage/certs/ca.crt")
	if err != nil {
		log.Fatalln(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(publicCert)

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS: &tls.Config{
			RootCAs: caCertPool,
		},
	}

	//conn, err := dialer.DialContext(t.Context(), "tcp", "192.168.13.3:9093")
	conn, err := dialer.DialContext(t.Context(), "tcp", "127.0.0.1:9093")
	assert.NoError(t, err)

	/*
		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		 Topic: topic-A1 пока нужно создавать руками
		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	*/

	w := kafka.Writer{
		//Addr:     kafka.TCP("192.168.13.3:9093"),
		Addr:     kafka.TCP("127.0.0.1:9093"),
		Topic:    "topic-A1",
		Balancer: &kafka.Hash{},
		Transport: &kafka.Transport{
			TLS: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	err = w.WriteMessages(t.Context(), kafka.Message{
		//Topic: "topic-A1",
		Value: []byte("This is test message across Kafak broker."),
	})
	assert.NoError(t, err)

	r := kafka.NewReader(kafka.ReaderConfig{
		//Brokers: []string{"192.168.13.3:9093"},
		Brokers:     []string{"127.0.0.1:9093"},
		GroupID:     "consumer-group-id",
		GroupTopics: []string{"topic-A1"},
		Dialer:      dialer,
	})
	msg, err := r.ReadMessage(t.Context())
	assert.NoError(t, err)

	//fmt.Printf("Message:'%#v'\n", msg)
	fmt.Printf("Reseived message from topic:'%s'\n", msg.Topic)
	fmt.Printf("Data:'%s'\n", string(msg.Value))

	t.Cleanup(func() {
		w.Close()
		r.Close()
		conn.Close()
	})
}
