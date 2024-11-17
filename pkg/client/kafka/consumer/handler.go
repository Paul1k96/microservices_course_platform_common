package consumer

import (
	"context"
	"log/slog"

	"github.com/IBM/sarama"
)

// Handler is a function that processes a message
type Handler func(ctx context.Context, msg *sarama.ConsumerMessage) error

// GroupHandler is a handler for a consumer group
type GroupHandler struct {
	msgHandler Handler
	logger     *slog.Logger
}

// NewGroupHandler creates a new GroupHandler
func NewGroupHandler(logger *slog.Logger) *GroupHandler {
	return &GroupHandler{logger: logger}
}

// Setup is called at the beginning of a new session before ConsumeClaim is called
func (c *GroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is called at the end of the session after all ConsumeClaim goroutines have finished
func (c *GroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim should start the message consumption loop with ConsumerGroupClaim().
// After closing the Messages() channel, the handler should finish processing
func (c *GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// The code below should not be moved to a goroutine, as ConsumeClaim
	// is already started in a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L869
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				c.logger.Info("message channel was closed")
				return nil
			}

			c.logger.Debug(
				"message claimed:",
				slog.String("value", string(message.Value)),
				slog.String("timestamp", message.Timestamp.String()),
				slog.String("topic", message.Topic),
			)

			err := c.msgHandler(session.Context(), message)
			if err != nil {
				c.logger.Error("error handling message:", slog.String("err: ", err.Error()))
				continue
			}

			session.MarkMessage(message, "")

		// Should return when `session.Context()` is finished.
		// Otherwise, `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout`
		// will occur during Kafka rebalancing. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			c.logger.Info("session context done")
			return nil
		}
	}
}
