package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type MessageHandler struct{}

func (h *MessageHandler) Setup(_ sarama.ConsumerGroupSession) error {
	// Initialize resources before consuming messages
	return nil
}

func (h *MessageHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	// Clean up resources after consuming messages
	return nil
}

func (h *MessageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message received: topic=%s partition=%d offset=%d key=%s value=%s",
			message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}
