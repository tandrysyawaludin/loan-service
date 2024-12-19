package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
}

func NewConsumer(brokers []string, groupID string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0 // Specify the Kafka version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{consumerGroup: consumerGroup}, nil
}

func (c *Consumer) Consume(topics []string, handler sarama.ConsumerGroupHandler) {
	for {
		if err := c.consumerGroup.Consume(nil, topics, handler); err != nil {
			log.Printf("Error consuming messages: %v", err)
		}
	}
}

func (c *Consumer) Close() error {
	return c.consumerGroup.Close()
}
