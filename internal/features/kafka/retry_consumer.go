package kafka

import (
	"context"
	"encoding/json"
	"math/rand"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	core_logger "github.com/qandoni/keeneyePractice/internal/core/logger"
	"go.uber.org/zap"
)

type RetryConsumer struct {
	consumer *kafka.Consumer
	service  EmailDispatchService
	producer *Producer
	logger   core_logger.Logger
}

func NewRetryConsumer(
	address []string,
	topic string,
	group string,
	producer *Producer,
	emailDispatchService EmailDispatchService,
	logger *core_logger.Logger,
) (*RetryConsumer, error) {

	cfg := &kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(address, ","),
		"group.id":                 group,
		"enable.auto.commit":       false,
		"auto.offset.reset":        "earliest",
		"allow.auto.create.topics": true,
	}

	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	if err := c.Subscribe(topic, nil); err != nil {
		c.Close()
		return nil, err
	}

	return &RetryConsumer{
		consumer: c,
		producer: producer,
		service:  emailDispatchService,
		logger:   *logger,
	}, nil
}

func (c *RetryConsumer) Run(ctx context.Context) error {
	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			c.logger.Error("read retry message", zap.Error(err))
			continue
		}
		var event RetryMessage
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			c.logger.Error("unmarshal", zap.Error(err))
			continue
		}
		if rand.Intn(100) < 30 {
			c.logger.Error(
				"retry processing failed",
				zap.Int("retry", event.Retry),
			)
			if event.Retry >= 10 {
				if err := c.producer.PublishDeadRegistrationRequest(
					ctx,
					event.RequestID,
				); err != nil {
					c.logger.Error("publish dead", zap.Error(err))
					continue
				}

			} else {

				if err := c.producer.PublishRetryRegistrationRequest(
					ctx,
					event.RequestID,
					event.Retry+1,
				); err != nil {
					c.logger.Error("publish retry", zap.Error(err))
					continue
				}
			}
			_, _ = c.consumer.CommitMessage(msg)
			continue
		}

		if err := c.service.SendRegistrationEmail(
			ctx,
			event.RequestID,
		); err != nil {
			if event.Retry >= 10 {
				if err := c.producer.PublishDeadRegistrationRequest(
					ctx,
					event.RequestID,
				); err != nil {
					continue
				}
			} else {
				if err := c.producer.PublishRetryRegistrationRequest(
					ctx,
					event.RequestID,
					event.Retry+1,
				); err != nil {
					continue
				}
			}
			_, _ = c.consumer.CommitMessage(msg)
			continue
		}

		_, _ = c.consumer.CommitMessage(msg)
	}

}
