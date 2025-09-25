package natsgo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestConnectNats(t *testing.T) {
	var (
		conn      *nats.Conn
		subscribe *nats.Subscription
		err       error

		host string = "localhost"
		port int    = 4222

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

	t.Run("Тест 1. Получение сообщения", func(t *testing.T) {
		var count int

		subscribe, err = conn.Subscribe("my-test-subject", func(msg *nats.Msg) {
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

			err := conn.Publish(
				"my-test-subject",
				fmt.Appendf(nil, "Datetime:'%s', message:'%s''", time.Now().Format(time.RFC3339), answers[rand.Intn(len(answers))]))
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
		conn.Close()
		close(chDone)
	})
}
