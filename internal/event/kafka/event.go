package kafka

import (
	"batchImageProcessing/config"
	"batchImageProcessing/internal/event"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type kafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(cfg config.Config) event.Publisher {
	return &kafkaPublisher{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(cfg.KafkaBroker),
			Topic: "book_events",
		},
	}
}

func (p *kafkaPublisher) Publish(eventName string, payload interface{}) error {
	msgBytes, err := json.Marshal(map[string]interface{}{
		"event":   eventName,
		"payload": payload,
	})
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(eventName),
		Value: msgBytes,
	}
	return p.writer.WriteMessages(context.Background(), msg)
}
