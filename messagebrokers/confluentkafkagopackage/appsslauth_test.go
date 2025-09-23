package confluentkafkagopackage_test

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/golang_structures_and_algorithms/errorspackage"
	"github.com/av-belyakov/golang_structures_and_algorithms/mapmanipulation"
	kafpkg "github.com/av-belyakov/golang_structures_and_algorithms/messagebrokers/confluentkafkagopackage"
)

func TestKafkaApiConnectionSSLAuth(t *testing.T) {
	var (
		kafkaApi   *kafpkg.KafkaApiModule
		testTopics map[string]string = map[string]string{
			"event-topic": "event-message.topic",
			"alert-topic": "alert-message.topic",
		}

		eventTopic string = "event-message.topic"
		alertTopic string = "alert-message.topic"

		nameRegionalObject string = "kafka-test-object"

		err error
	)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)
	t.Cleanup(func() {
		stop()
	})

	go func() {
		<-ctx.Done()

		fmt.Println("Останов модуля тестирования")

		stop()
	}()

	logging := kafpkg.NewLogging()
	counting := kafpkg.NewCounting()

	verifyCertificate := func(caPath string) error {
		caData, err := os.ReadFile(caPath)
		if err != nil {
			return fmt.Errorf("failed to read CA: %v", err)
		}

		block, _ := pem.Decode(caData)
		if block == nil {
			return fmt.Errorf("failed to parse CA PEM")
		}

		caCert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return fmt.Errorf("failed to parse CA certificate: %v", err)
		}

		fmt.Printf("CA Subject: %s\n", caCert.Subject)
		fmt.Printf("CA Valid from: %s to %s\n", caCert.NotBefore, caCert.NotAfter)

		return nil
	}

	//обработчик счётчика, логов и сообщений модуля
	go func(
		ctx context.Context,
		l kafpkg.Logger,
		c *kafpkg.Counting) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Останов обработчика счётчика и логов")

				return

			case count := <-c.GetChan():
				fmt.Println("count.message:", count.Message, " count.Count", count.Count)

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

				if msg.GetType() == "error" {
					log.Fatal(msg.GetMessage())
				}
			}
		}
	}(ctx, logging, counting)

	if err := godotenv.Load("../kafkaimage/.env"); err != nil {
		log.Fatalln(err)
	}

	t.Run("Тест 0. Проверка корневого сертификата", func(t *testing.T) {
		assert.NoError(t, verifyCertificate("../kafkaimage/certs/ca.crt"))
	})

	t.Run("Тест 1. Инициализация модуля Kafka API", func(t *testing.T) {
		kafkaApi, err = kafpkg.NewWithSSL(
			counting,
			logging,
			kafpkg.WithHost("localhost"),
			//kafpkg.WithHost("10.0.0.136"),
			kafpkg.WithPort(9093),
			kafpkg.WithLocationCertificateCA("../kafkaimage/certs/ca.crt"),
			kafpkg.WithLocationClientCertificate("../kafkaimage/certs/client-cert.pem"),
			kafpkg.WithLocationClientKey("../kafkaimage/certs/client-key.pem"),
			kafpkg.WithNameRegionalObject(nameRegionalObject),
			kafpkg.WithTopicsSubscription(testTopics))
		assert.NoError(t, err)
	})

	t.Run("Тест 2. Запуск модуля взаимодействия с Kafka", func(t *testing.T) {
		err = kafkaApi.StartWithSSL(ctx, func() func(context.Context, *kafpkg.KafkaApiModule) {
			return func(ctx context.Context, api *kafpkg.KafkaApiModule) {
				for {
					select {
					case <-ctx.Done():
						fmt.Println("Останов обработчика входящих сообщений")

						return

					default:
						consumer := api.GetConsumer()
						msg, err := consumer.ReadMessage(time.Second)
						if err == nil {
							topic := msg.TopicPartition.Topic

							topicType := "undefined_type"
							topicKey, ok := mapmanipulation.SearchValueMap(api.GetTopics(), *topic)
							if ok {
								topicType = topicKey
							}

							api.GetChanOutput() <- kafpkg.ChOutputSettings{
								TopicType: topicType,
								Data:      msg.Value,
							}
						} else if !err.(kafka.Error).IsTimeout() {
							api.GetLogger().Send("error", errorspackage.CustomError(err).Error())
						}
					}
				}
			}
		})
		assert.NoError(t, err)
	})

	t.Run("Тест 4. Инициализация Kafka Producer", func(t *testing.T) {
		p, err := kafkaApi.ProducerWithSSL(ctx)
		assert.NoError(t, err)

		// обработчик сообщений об успешности доставки
		go func(ctx context.Context, p *kafka.Producer) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Останов обработчика сообщений об успешности доставки")

					return

				case e, isNotClose := <-p.Events():
					if !isNotClose {
						return
					}

					switch ev := e.(type) {
					case *kafka.Message:
						// Отчет о доставке сообщения, указывающий на успех или
						// постоянный сбой после завершения повторных попыток.
						m := ev
						if m.TopicPartition.Error != nil {
							fmt.Printf("Доставка в Kafka не успешна: %v\n", m.TopicPartition.Error)
						} else {
							fmt.Printf("Успешная доставка в Kafka %s [%d] at offset %v\n",
								*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
						}
					case kafka.Error:
						// Типичные ошибки на уровне экземпляра клиента, такие как
						// сбои подключения к брокеру, проблемы с аутентификацией и т.д.
						assert.NoError(t, ev)
					default:
						fmt.Printf("Ignored event: %s\n", ev)
					}
				}
			}
		}(ctx, p)

		t.Run("Тест 4.1. Передача сообщения в топик 'event-message.topic'", func(t *testing.T) {
			err = p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &eventTopic,
					Partition: kafka.PartitionAny,
				},
				Value: fmt.Appendf(nil, "%s: Some small message to 'event-message.topic'. Test message.", time.Now().Format(time.RFC3339Nano)),
			}, nil)
			assert.NoError(t, err)
		})

		t.Run("Тест 4.2. Передача сообщения в топик 'alert-message.topic'", func(t *testing.T) {
			err = p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &alertTopic,
					Partition: kafka.PartitionAny,
				},
				Value: fmt.Appendf(nil, "%s: Some small message to 'alert-message.topic'. Test message.", time.Now().Format(time.RFC3339Nano)),
			}, nil)
			assert.NoError(t, err)
		})

		t.Cleanup(func() {
			p.Close()
		})
	})

	t.Run("Тест 5. Обработчик сообщений модуля", func(t *testing.T) {
		msgOne, isOpen := <-kafkaApi.GetChanOutput()
		assert.True(t, isOpen)
		assert.NotEmpty(t, msgOne.TopicType)
		assert.NotEmpty(t, msgOne.Data)

		data, ok := msgOne.Data.([]byte)
		assert.True(t, ok)
		fmt.Printf("Успешный приём сообщений из Kafka. TopicType: '%s', Data: '%s'\n", msgOne.TopicType, string(data))

		msgTwo, isOpen := <-kafkaApi.GetChanOutput()
		assert.True(t, isOpen)
		assert.NotEmpty(t, msgTwo.TopicType)
		assert.NotEmpty(t, msgTwo.Data)

		data, ok = msgTwo.Data.([]byte)
		assert.True(t, ok)
		fmt.Printf("Успешный приём сообщений из Kafka. TopicType: '%s', Data: '%s'\n", msgTwo.TopicType, string(data))
	})

	// t.Run("", func(t *testing.T) {})
}
