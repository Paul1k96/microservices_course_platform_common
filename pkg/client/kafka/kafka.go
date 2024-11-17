package kafka

import (
	"context"

	"github.com/Paul1k96/microservices_course_platform_common/pkg/client/kafka/consumer"
)

// Consumer is a Kafka consumer
type Consumer interface {
	Consume(ctx context.Context, topicName string, handler consumer.Handler) (err error)
	Close() error
}
