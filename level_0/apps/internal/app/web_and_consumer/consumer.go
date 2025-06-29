package web_and_consumer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
	"time"

	"github.com/rashid567/learning-golang/level_0/apps/internal/configs"
	"github.com/rashid567/learning-golang/level_0/apps/internal/domain/order"
	"github.com/rashid567/learning-golang/level_0/apps/internal/models"
)

func runKafkaConsumer(ctx context.Context, wg *sync.WaitGroup, conf *configs.KafkaConsumerConfig, orderService order.OrderService) {
	defer wg.Done()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     conf.Hosts(),
		Topic:       conf.Topic,
		GroupID:     conf.ConsumerGroup,
		StartOffset: kafka.FirstOffset,
	})
	defer reader.Close()

	log.Println("Kafka Consumer started!")

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping Kafka Consumer...")
			return

		default:
			// Читаем сообщение с таймаутом
			msgCtx, cancel := context.WithTimeout(ctx, time.Second*10)
			msg, err := reader.FetchMessage(msgCtx)
			cancel()
			if err != nil {
				if err == context.Canceled {
					return
				} else if err == context.DeadlineExceeded {
					continue
				}
				log.Fatalf("Kafka read error: %s\n", err)
			}

			var order models.Order
			err = json.Unmarshal(msg.Value, &order)
			if err != nil {
				log.Fatalf("Failed to parse msgs: %s\n", err)
			}

			err = orderService.SaveOrder(&order)
			if err != nil {
				log.Fatalf("Failed to save order: %s\n", err)
			}
			reader.CommitMessages(ctx, msg)
		}
	}
}
