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

func TestKafkaApiSimpleConnection(t *testing.T) {
	var (
		kafkaApi   *kafpkg.KafkaApiModule
		p          *kafka.Producer
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

	t.Run("Тест 1. Инициализация модуля Kafka API", func(t *testing.T) {
		kafkaApi, err = kafpkg.New(
			counting,
			logging,
			kafpkg.WithHost("localhost"),
			//kafpkg.WithHost("10.0.0.136"),
			kafpkg.WithPort(9092),
			kafpkg.WithNameRegionalObject(nameRegionalObject),
			kafpkg.WithTopicsSubscription(testTopics))
		assert.NoError(t, err)
	})

	t.Run("Тест 2. Создание новых топиков если их нет", func(t *testing.T) {
		ac, err := kafka.NewAdminClient(&kafka.ConfigMap{
			"bootstrap.servers": "localhost", //может содержать ip:port
			//"bootstrap.servers": "10.0.0.136", //"10.0.0.136:9092"
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
		result, err := ac.CreateTopics(ctx, topics)
		assert.NoError(t, err)

		fmt.Println("Create topics result:")
		for k, v := range result {
			fmt.Printf("  %d.\n\tName:'%s'\n\tError:'%s'\n", k, v.Topic, v.Error.String())
		}
	})

	t.Run("Тест 3. Запуск модуля взаимодействия с Kafka", func(t *testing.T) {
		err = kafkaApi.Start(ctx, func() func(context.Context, *kafpkg.KafkaApiModule) {
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
		p, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost", //может содержать ip:port
			//"bootstrap.servers": "10.0.0.136", //"10.0.0.136:9092"
		})
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
	})

	t.Run("Тест 5.1. Передача сообщения в топик 'event-message.topic'", func(t *testing.T) {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &eventTopic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte("Some small message to 'event-message.topic'. Test message."),
		}, nil)
		assert.NoError(t, err)
	})

	t.Run("Тест 5.2. Передача сообщения в топик 'alert-message.topic'", func(t *testing.T) {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &alertTopic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte("Some small message to 'alert-message.topic'. Test message."),
		}, nil)
		assert.NoError(t, err)
	})

	t.Run("Тест 6. Обработчик сообщений модуля", func(t *testing.T) {
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
