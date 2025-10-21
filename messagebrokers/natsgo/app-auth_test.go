package natsgo_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestNatsAuth(t *testing.T) {
	var (
		connSubs  *nats.Conn
		connProd  *nats.Conn
		subscribe *nats.Subscription
		err       error

		host string
		port int

		answers []string = []string{
			"It is certain",
			"It is decidedly so",
			"Without a doubt",
			"Yes definitely",
			"You may rely on it",
			"As I see it yes",
			"Most likely",
			"Outlook good",
			"Yes",
			"Signs point to yes",
			"Reply hazy try again",
			"Ask again later",
			"Better not tell you now",
			"Cannot predict now",
			"Concentrate and ask again",
			"Don't count on it",
			"My reply is no",
			"My sources say no",
			"Outlook not so good",
			"Very doubtful",
		}
	)

	chDone := make(chan struct{})

	// Загружаем переменные окружения
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	// Получаем имя или адрес хоста
	host = os.Getenv("GO_TEST_NATS_HOST")
	if host == "" {
		log.Fatalln("environment 'GO_TEST_NATS_HOST' is not defined")
	}

	// Получаем сетевой порт
	port, err = strconv.Atoi(os.Getenv("GO_TEST_NATS_PORT"))
	if err != nil {
		log.Fatalln("environment 'GO_TEST_NATS_PORT' is not defined")
	}
	if port == 0 {
		log.Fatalln("the port cannot be equal to 0")
	}

	// Получаем пароль потребителя (consumer)
	consPasswd := os.Getenv("GO_TEST_NATS_CLIENT_PASSWD")
	if consPasswd == "" {
		log.Fatalln("environment 'GO_TEST_NATS_CLIENT_PASSWD' is not defined")
	}
	// Получаем пароль производителя запросов (producer)
	prodPasswd := os.Getenv("GO_TEST_NATS_SERVICE_PASSWD")
	if prodPasswd == "" {
		log.Fatalln("environment 'GO_TEST_NATS_SERVICE_PASSWD' is not defined")
	}

	t.Run("Тест 0. Соединения с NATS", func(t *testing.T) {
		t.Run("Установление соединения для потребителя (consumer)", func(t *testing.T) {
			connSubs, err = nats.Connect(
				fmt.Sprintf("%s:%d", host, port),
				nats.Name("service for testing"),
				nats.MaxReconnects(-1),
				nats.ReconnectWait(3*time.Second),
				nats.Timeout(3*time.Second),
				nats.UserInfo("service", consPasswd))
			assert.NoError(t, err)
			assert.NotNil(t, connSubs)
		})
		t.Run("Установление соединения для производителя (producer)", func(t *testing.T) {
			connProd, err = nats.Connect(
				fmt.Sprintf("%s:%d", host, port),
				nats.Name("client for testing"),
				nats.MaxReconnects(-1),
				nats.ReconnectWait(3*time.Second),
				nats.Timeout(3*time.Second),
				nats.UserInfo("client", prodPasswd))
			assert.NoError(t, err)
			assert.NotNil(t, connProd)
		})
	})

	t.Run("Тест 1. Получение сообщения", func(t *testing.T) {
		var count int

		subscribe, err = connSubs.Subscribe("case-message.test", func(msg *nats.Msg) {
			fmt.Printf("Reseived message:'%s'\n", string(msg.Data))

			count++
			if count == len(answers) {
				chDone <- struct{}{}
			}
		})
		assert.NoError(t, err)
	})

	t.Run("Тест 2. Отправка сообщения", func(t *testing.T) {
		for range answers {
			time.Sleep(time.Millisecond * 300)

			err := connProd.Publish(
				"case-message.test",
				fmt.Appendf(nil, "datetime:'%s', message:'%s'", time.Now().Format(time.RFC3339), answers[rand.Intn(len(answers))]))
			assert.NoError(t, err)
		}

		<-chDone

		assert.True(t, true, true)
	})

	/*
	   t.Run("", func(t *testing.T) {

	   })
	*/

	t.Cleanup(func() {
		subscribe.Unsubscribe()
		connSubs.Close()
		connProd.Close()
		close(chDone)

		os.Setenv("GO_TEST_NATS_HOST", "")
		os.Setenv("GO_TEST_NATS_PORT", "")
		os.Setenv("GO_TEST_NATS_CLIENT_PASSWD", "")
		os.Setenv("GO_TEST_NATS_SERVICE_PASSWD", "")
	})
}
