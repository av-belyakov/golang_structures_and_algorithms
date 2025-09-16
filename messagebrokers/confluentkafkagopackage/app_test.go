package confluentkafkagopackage_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/golang_structures_and_algorithms/errorspackage"
	"github.com/av-belyakov/golang_structures_and_algorithms/mapmanipulation"
	kafpkg "github.com/av-belyakov/golang_structures_and_algorithms/messagebrokers/confluentkafkagopackage"
)

func TestKafkaApi(t *testing.T) {
	var (
		kafkaApi   *kafpkg.KafkaApiModule
		testTopics map[string]string = map[string]string{
			"event-topic": "event-message.topic",
			"alert-topic": "alert-message.topic",
		}

		eventTopic string = "event-message.topic"
		alertTopic string = "alert-message.topic"

		err error
	)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)
	t.Cleanup(func() {
		stop()
	})

	go func() {
		<-ctx.Done()

		fmt.Println("placeholder_doc-basedb-bi-zone module is Stop")

		stop()
	}()

	logging := kafpkg.NewLogging()
	counting := kafpkg.NewCounting()

	kafkaApi, err = kafpkg.New(
		counting,
		logging,
		//kafpkg.WithHost("localhost"),
		kafpkg.WithHost("10.0.0.136"),
		//kafpkg.WithPort(19092),
		kafpkg.WithPort(9092),
		kafpkg.WithNameRegionalObject("kafka-test-object"),
		kafpkg.WithTopicsSubscription(testTopics))
	if err != nil {
		log.Fatalln(err)
	}

	go func(
		ctx context.Context,
		l kafpkg.Logger,
		c *kafpkg.Counting,
		kApi *kafpkg.KafkaApiModule) {
		for {
			select {
			case <-ctx.Done():
				return

			case count := <-c.GetChan():
				fmt.Println("count.message:", count.Message, " count.Count", count.Count)

			case msg := <-l.GetChan():
				fmt.Println("LOG:", msg)

				if msg.GetType() == "error" {
					log.Fatal(msg.GetMessage())
				}
			case msg := <-kafkaApi.GetChanOutput():
				fmt.Printf("Message from Kafka:'%#v'\n", msg)
			}
		}
	}(ctx, logging, counting, kafkaApi)

	t.Run("Тест 1. Запуск модуля взаимодействия с Kafka", func(t *testing.T) {
		err = kafkaApi.Start(ctx, func() func(context.Context, *kafpkg.KafkaApiModule) {
			return func(ctx context.Context, api *kafpkg.KafkaApiModule) {
				for {
					select {
					case <-ctx.Done():
						fmt.Println("закрываем обработчик входящих сообщений")

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

							api.GetChanInput() <- kafpkg.ChInputSettings{
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

	t.Run("Тест 2. Передача и подтверждение передачи сообщений в топики", func(t *testing.T) {
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			//"bootstrap.servers": "localhost",
			"bootstrap.servers": "10.0.0.136",
			"group.id":          fmt.Sprintf("%s-group", "group-test"),
		})
		assert.NoError(t, err)

		// обработчик сообщений об успешности доставки
		go func(ctx context.Context, p *kafka.Producer) {
			for {
				select {
				case <-ctx.Done():
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
							fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
						} else {
							fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
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

		t.Run("Тест 2.1. Передача сообщения в топик 'event-message.topic'", func(t *testing.T) {
			err = p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &eventTopic,
					Partition: kafka.PartitionAny,
				},
				Value: []byte("some small message to 'event-message.topic'"),
			}, nil)
			assert.NoError(t, err)
		})

		t.Run("Тест 2.2. Передача сообщения в топик 'alert-message.topic'", func(t *testing.T) {
			err = p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &alertTopic,
					Partition: kafka.PartitionAny,
				},
				Value: []byte("some small message to 'alert-message.topic'"),
			}, nil)
			assert.NoError(t, err)
		})
	})

	// t.Run("", func(t *testing.T) {
	// })
}
