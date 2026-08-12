package confluentkafkagopackage

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/av-belyakov/golang_structures_and_algorithms/errorspackage"
	"github.com/av-belyakov/golang_structures_and_algorithms/mapmanipulation"
)

func TopicHandler(ctx context.Context, api *KafkaApiModule) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			msg, err := api.consumer.ReadMessage(time.Second)
			if err == nil {
				topic := msg.TopicPartition.Topic

				topicType := "undefined_type"
				topicKey, ok := mapmanipulation.SearchValueMap(api.topics, *topic)
				if ok {
					topicType = topicKey
				}

				api.chFromModule <- ChOutputSettings{
					TopicType: topicType,
					Data:      msg.Value,
				}
			} else if !err.(kafka.Error).IsTimeout() {
				api.logger.Send("error", errorspackage.CustomError(err).Error())
			}
		}
	}
}

// topicsHandler обработчик топиков (подписок)
//func (api *kafkaApiModule) topicsHandler(ctx context.Context) {
//	for {
//		select {
//		case <-ctx.Done():
//			fmt.Println("func 'topicsHandler', закрываем канал на отправку из модуля")
//
//			//закрываем канал на отправку из модуля
//			close(api.chFromModule)
//
//			return
//
//		default:
//			msg, err := api.consumer.ReadMessage(time.Second) //-1)
//			if err == nil {
//				topic := msg.TopicPartition.Topic
//
//				subjectType := "undefined_type"
//				topicKey, ok := supportingfunctions.SearchValue(api.topics, *topic)
//				if ok {
//					subjectType = topicKey
//				}
//
//				api.chFromModule <- SettingsChanOutput{
//					SubjectType: subjectType,
//					Data:        msg.Value,
//				}
//			} else if !err.(kafka.Error).IsTimeout() {
//				api.logger.Send("error", supportingfunctions.CustomError(err).Error())
//			}
//		}
//	}
//}
