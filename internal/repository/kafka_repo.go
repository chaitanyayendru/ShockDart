package repository

import (
	"log"

	"github.com/Shopify/sarama"
)

type NotificationRepository interface {
	EnqueueNotification(req NotificationRequest) error
}

type kafkaRepository struct {
	producer sarama.SyncProducer
}

func NewKafkaRepository() NotificationRepository {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}

	return &kafkaRepository{producer: producer}
}

func (r *kafkaRepository) EnqueueNotification(req NotificationRequest) error {
	msg := &sarama.ProducerMessage{
		Topic: "notifications",
		Value: sarama.StringEncoder(req.Message),
	}
	_, _, err := r.producer.SendMessage(msg)
	return err
}
