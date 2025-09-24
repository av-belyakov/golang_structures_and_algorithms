package kafkagopackage

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func TestKafkaConnection(t *testing.T) {
	var (
		conn *kafka.Conn
		r    *kafka.Reader
		err  error

		kafkaHost    string   = "10.0.0.136:9092" //"127.0.0.1:9093"
		listenTopics []string = []string{"soc-events", "arkime-json"}
		//listenTopics []string = []string{"topic-A0", "topic-A1"}
	)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT)
	t.Cleanup(func() {
		stop()
		conn.Close()
	})

	t.Run("Тест 0. Соединение с Kafka сервером.", func(t *testing.T) {
		conn, err = kafka.DialContext(ctx, "tcp", kafkaHost)
		assert.NoError(t, err)
	})

	t.Run("Тест 1. Посмотреть все брокеры.", func(t *testing.T) {
		brokers, err := conn.Brokers()
		assert.NoError(t, err)
		assert.NotEqual(t, len(brokers), 0)

		fmt.Println("Kafka brokers:")
		for _, v := range brokers {
			fmt.Printf("broker id:%d, host:%s\n", v.ID, v.Host)
		}
	})

	t.Run("Тест 2. Прочитать содержимое топиков.", func(t *testing.T) {
		dialer := &kafka.Dialer{
			ClientID:  "kafka-go-test-client",
			Timeout:   10 * time.Second,
			DualStack: true,
			//LocalAddr: kafka.TCP(kafkaHost),
		}

		r = kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{kafkaHost},
			Topic:       listenTopics[0],
			GroupTopics: listenTopics,
			Dialer:      dialer,
		})

		var num int
	DONE:
		for {
			select {
			case <-ctx.Done():
				break DONE
			default:
				if num > 3 {
					break DONE
				}

				nextMsg, err := r.ReadMessage(ctx)
				assert.NoError(t, err)

				fmt.Printf("topic:'%s', msg:'%s'\n", nextMsg.Topic, string(nextMsg.Value))
				num++
			}
		}

		assert.GreaterOrEqual(t, num, 3)
	})

	t.Cleanup(func() {
		if r != nil {
			r.Close()
		}

		if conn != nil {
			conn.Close()
		}
	})
}
