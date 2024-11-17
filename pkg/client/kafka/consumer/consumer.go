package consumer

import (
	"context"
	"log/slog"
	"strings"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

// Consumer is a Kafka consumer
type Consumer struct {
	consumerGroup        sarama.ConsumerGroup
	consumerGroupHandler *GroupHandler
	logger               *slog.Logger
}

// NewConsumer creates a new Consumer
func NewConsumer(
	consumerGroup sarama.ConsumerGroup,
	consumerGroupHandler *GroupHandler,
	logger *slog.Logger,
) *Consumer {
	return &Consumer{
		consumerGroup:        consumerGroup,
		consumerGroupHandler: consumerGroupHandler,
		logger:               logger,
	}
}

// Consume starts consuming messages from the specified topic
func (c *Consumer) Consume(ctx context.Context, topicName string, handler Handler) error {
	c.consumerGroupHandler.msgHandler = handler

	return c.consume(ctx, topicName)
}

// Close closes the consumer
func (c *Consumer) Close() error {
	return c.consumerGroup.Close()
}

func (c *Consumer) consume(ctx context.Context, topicName string) error {
	for {
		err := c.consumerGroup.Consume(ctx, strings.Split(topicName, ","), c.consumerGroupHandler)
		if err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}

			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		c.logger.Info("rebalancing...")
	}
}
