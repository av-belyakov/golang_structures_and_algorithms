package kafkagopackage_test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func TestKafkaConnectionSSL(t *testing.T) {
	var (
		topicsName []string = []string{"topic-A0", "topic-A1"}
		newTopics  []kafka.TopicConfig
		pathCerts  string = "../kafkaimage/certs/"

		usersList []string = []string{
			`{"name":"Некрасова Василиса Андреевна", "Age": 31, "Address": "г. Москва, Зелёный проспект, д.23, кв. 12"}`,
			`{"name":"Толмачёв Виктор Геннадьевич", "Age":35, "Address":"г. Балашиха, ул. Весенная, д.41, кв. 56"}`,
			`{"name":"Степанова Галина Сергеевна", "Age":22, "Address":"г. Москва, ул. Красного знамени, д. 8, кв. 191"}`,
			`{"name":"Тяпков Василий Петрович", "Age":12, "Address":"г. Красногорск, ул. 8-Мая, д. 24, кв. 113"}`,
		}
	)

	s := slices.All(topicsName)
	for _, v := range s {
		newTopics = append(newTopics, kafka.TopicConfig{
			Topic:             v,
			NumPartitions:     1,
			ReplicationFactor: -1,
		})
	}

	publicCert, err := os.ReadFile(path.Join(pathCerts, "ca.crt"))
	if err != nil {
		log.Fatalln(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(publicCert)

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS: &tls.Config{
			RootCAs:            caCertPool,
			ServerName:         "localhost",
			InsecureSkipVerify: false, // true,
		},
	}

	//conn, err := dialer.DialContext(t.Context(), "tcp", "192.168.13.3:9093")
	conn, err := dialer.DialContext(t.Context(), "tcp", "127.0.0.1:9093")
	assert.NoError(t, err)

	t.Run("Тест 0. Версия API, локальный, удалённый адрес Kafka", func(t *testing.T) {
		info, err := conn.ApiVersions()
		assert.NoError(t, err)

		fmt.Printf("API informations:%#v\n", info)

		ip := conn.LocalAddr()
		fmt.Println("Local IP address:", ip.String())

		ip = conn.RemoteAddr()
		fmt.Println("Remote IP address:", ip.String())
	})

	t.Run("Тест 1. Создать новый топик если он не существует.", func(t *testing.T) {
		err := conn.CreateTopics(newTopics...)
		assert.NoError(t, err)
	})

	t.Run("Тест 2. Отправить сообщение в топик.", func(t *testing.T) {
		w := kafka.Writer{
			//Addr:     kafka.TCP("192.168.13.3:9093"),
			Addr:     kafka.TCP("127.0.0.1:9093"),
			Topic:    topicsName[1],
			Balancer: &kafka.Hash{},
			Transport: &kafka.Transport{
				TLS: &tls.Config{
					RootCAs:            caCertPool,
					ServerName:         "localhost",
					InsecureSkipVerify: false,
				},
			},
		}

		for _, v := range usersList {
			err = w.WriteMessages(t.Context(), kafka.Message{
				Value: fmt.Appendf(
					nil,
					`{{"datetime":%d}, {"info": %v}}`,
					time.Now().UnixMicro(), v),
			})
			assert.NoError(t, err)
		}

		t.Cleanup(func() {
			w.Close()
		})
	})

	t.Run("Тест 3. Принять сообщение из топика.", func(t *testing.T) {
		r := kafka.NewReader(kafka.ReaderConfig{
			//Brokers: []string{"192.168.13.3:9093"},
			Brokers:     []string{"127.0.0.1:9093"},
			GroupID:     "consumer-group-id",
			GroupTopics: topicsName,
			Dialer:      dialer,
		})

		t.Run("Тест 3.1. Получить одно сообщение", func(t *testing.T) {
			// Чтение только одного сообщения
			msg, err := r.ReadMessage(t.Context())
			assert.NoError(t, err)

			fmt.Printf("Reseived message from topic:'%s'\n", msg.Topic)
			fmt.Printf("Data:'%s'\n", string(msg.Value))
		})

		t.Run("Тест 3.2. Получить все сообщения", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(t.Context(), time.Duration(time.Second*5))
			// Чтение следующего сообщения в очереди. Блокировка если сообщений
			// в очереди больше нет.
		DONE:
			for {
				select {
				case <-ctx.Done():
					break DONE

				default:
					msg, err := r.FetchMessage(ctx)
					if err != nil && err.Error() != "context deadline exceeded" {
						assert.NoError(t, err)
					}

					fmt.Printf("Topic: '%s', Next message:'%v'\n", msg.Topic, string(msg.Value))
				}
			}

			t.Cleanup(func() {
				cancel()
			})
		})

		t.Cleanup(func() {
			r.Close()
		})
	})

	t.Run("Тест 4. ", func(t *testing.T) {

	})

	//t.Run("Тест Finally. Удаляет все созданные топики.", func(t *testing.T) {
	//	err := conn.DeleteTopics(topicsName...)
	//	assert.NoError(t, err)
	//})

	t.Cleanup(func() {
		conn.Close()
	})
}
