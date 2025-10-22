package natsgo_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/stretchr/testify/assert"
)

func TestNatsJetStream(t *testing.T) {
	var (
		conn      *nats.Conn
		subscribe *nats.Subscription
		ctx       context.Context
		ctxCancel context.CancelFunc
		stream    jetstream.Stream
		cons      jetstream.Consumer
		err       error

		host string = "localhost"
		port int    = 4222

		testMessage []string = []string{
			"remaining server",
			"balances scalability with a tolerance for failure",
			"even recommended in some cases",
			"others optimized solely",
			"enabled servers in the list of clusters",
			"you normally would by specifying",
			"the configuration",
			"should be configured",
			"following configuration uses a bcrypted password",
			"provides methods to manage",
			"perform stream-specific operations",
			"messages from stream",
			"main goal of this package",
			"a simple and clear way to interact",
			"key differences between jetstream and nats packages",
			"to continuously receive incoming messages",
			"predictable approach to consuming messages",
			"simpler interfaces to manage streams and consumers",
		}
	)

	chDone := make(chan struct{})

	// Загружаем переменные окружения
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	host = os.Getenv("GO_TEST_NATS_HOST")
	if host == "" {
		log.Fatalln("environment 'GO_TEST_NATS_HOST' is not defined")
	}

	port, err = strconv.Atoi(os.Getenv("GO_TEST_NATS_PORT"))
	if err != nil {
		log.Fatalln("environment 'GO_TEST_NATS_PORT' is not defined")
	}
	if port == 0 {
		log.Fatalln("the port cannot be equal to 0")
	}

	t.Run("Тест 0. Соединение с NATS", func(t *testing.T) {
		conn, err = nats.Connect(
			fmt.Sprintf("%s:%d", host, port),
			nats.Name("client for testing"),
			nats.MaxReconnects(-1),
			nats.ReconnectWait(3*time.Second),
			nats.Timeout(3*time.Second))
		assert.NoError(t, err)
		assert.NotNil(t, conn)
	})

	t.Run("Тест 1. Передача сообщений с использованием JetStream", func(t *testing.T) {
		// Подробнее о потоках https://github.com/nats-io/nats.go/blob/main/jetstream/README.md

		ctx, ctxCancel = context.WithTimeout(context.Background(), 30*time.Second)

		// Создание JetStream интерфейса управления
		js, err := jetstream.New(conn)
		assert.NoError(t, err)

		// Дополнительная проверка - получаем информацию о JetStream
		info, err := js.AccountInfo(ctx)
		assert.NoError(t, err)

		fmt.Printf("JetStream доступен. Лимиты: %+v\n", info)

		// Попытка получить не существующий дескриптор потока
		_, err = js.Stream(ctx, "foo")
		assert.ErrorIs(t, err, jetstream.ErrStreamNotFound)

		// Создание нового дискриптора
		stream, err = js.CreateStream(ctx, jetstream.StreamConfig{
			Name:        "new_orders",
			Description: "Stream for test",
			Subjects:    []string{"new_orders.*"},
		})
		assert.NoError(t, err)

		// Публикация некоторых сообщений
		for range testMessage {
			/*pubAck*/ _, err := js.Publish(
				ctx,
				"new_orders.small",
				fmt.Appendf(nil, "datetime:'%s', message:'%s', jetStream:'true'", time.Now().Format(time.RFC3339), testMessage[rand.Intn(len(testMessage))]))
			assert.NoError(t, err)
		}

		// Создание надежного потребителя
		cons, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
			Durable:   "durable-consumer",
			AckPolicy: jetstream.AckExplicitPolicy,
		})
		assert.NoError(t, err)

		// Получаем некоторое количество сообщений для потребителя
		msgs, err := cons.Fetch(len(testMessage))
		assert.NoError(t, err)
		assert.NoError(t, msgs.Error())

		var messageCounter int

		// Читаем полученые сообщения
		for msg := range msgs.Messages() {
			// Подтверждение сообщения. Метод сообщает серверу, что сообщение было
			//успешно обработано, и сервер может перейти к передачи следующего сообщения
			msg.Ack()

			fmt.Printf("Received a JetStream message via fetch: %s\n", string(msg.Data()))

			messageCounter++
		}
		assert.Equal(t, messageCounter, len(testMessage))

		messageCounter = 0

		// Непрерывно получать сообщения в режиме обратного вызова
		consCallback, err := cons.Consume(func(msg jetstream.Msg) {
			msg.Ack()
			fmt.Printf("Received a JetStream message via callback: %s\n", string(msg.Data()))
			messageCounter++
		})
		assert.NoError(t, err)

		if messageCounter == len(testMessage) {
			consCallback.Stop()
			assert.Equal(t, messageCounter, len(testMessage))
		}
	})

	/*
	   t.Run("", func(t *testing.T) {

	   })
	*/

	t.Cleanup(func() {
		ctxCancel()

		subscribe.Unsubscribe()
		conn.Close()
		close(chDone)

		os.Setenv("GO_TEST_NATS_HOST", "")
		os.Setenv("GO_TEST_NATS_PORT", "")
		os.Setenv("GO_TEST_NATS_CLIENT_PASSWD", "")
		os.Setenv("GO_TEST_NATS_SERVICE_PASSWD", "")
	})
}
