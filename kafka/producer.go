package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	syncProducer sarama.SyncProducer
}

func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{syncProducer: producer}, nil
}

func (p *Producer) SendMessage(topic string, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := p.syncProducer.SendMessage(msg)
	if err != nil {
		log.Printf("Error sending message to topic %s: %v", topic, err)
		return err
	}
	log.Printf("Message sent to topic %s: %s", topic, message)
	return nil
}

func (p *Producer) Close() error {
	return p.syncProducer.Close()
}
