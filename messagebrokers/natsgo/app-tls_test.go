package natsgo_test

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestConnectNatsTLS(t *testing.T) {
	var (
		conn      *nats.Conn
		subscribe *nats.Subscription
		cert      tls.Certificate
		err       error

		chDone chan struct{} = make(chan struct{})

		host string
		port int = 4443

		clientCert string = "../natimage/certs/client-cert.pem"
		clientKey  string = "../natimage/certs/client-key.pem"

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

	// Загружаем переменные окружения
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	host = os.Getenv("GO_TEST_NATS_HOST")
	if host == "" {
		log.Fatalln("environment 'GO_TEST_NATS_HOST' is not defined")
	}

	//получаем корневой сертификат
	publicCert, err := os.ReadFile("../natimage/certs/rootCA.pem")
	if err != nil {
		log.Fatalln(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(publicCert)

	//получаем сертификаты клиента
	//cert := nats.ClientCert(clientCert, clientKey) //более простая настройка
	cert, err = tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		log.Fatalln(err)
	}
	cfg := &tls.Config{
		//ServerName:   host,
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		//MinVersion: 	tls.VersionTLS12,
	}

	conn, err = nats.Connect(
		fmt.Sprintf("tls://%s:%d", host, port),
		nats.Name("tls client for testing"),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(3*time.Second),
		nats.Timeout(3*time.Second),
		nats.Secure(cfg))
	if err != nil {
		log.Fatalln(err)
	}

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
		conn.Close()
		close(chDone)

		os.Setenv("GO_TEST_NATS_HOST", "")
		os.Setenv("GO_TEST_NATS_PORT", "")
		os.Setenv("GO_TEST_NATS_CLIENT_PASSWD", "")
		os.Setenv("GO_TEST_NATS_SERVICE_PASSWD", "")
	})
}
