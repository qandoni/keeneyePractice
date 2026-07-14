package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	core_logger "github.com/qandoni/keeneyePractice/internal/core/logger"
	email_retry_service "github.com/qandoni/keeneyePractice/internal/features/email_dispatch/service/email_retry"

	"go.uber.org/zap"
)

type EmailDispatchService interface {
	SendRegistrationEmail(
		ctx context.Context,
		requestID int,
	) error
}

type Consumer struct {
	consumer       *kafka.Consumer
	retryScheduler email_retry_service.RetryScheduler
	service        EmailDispatchService
	producer       *Producer
	logger         core_logger.Logger
}

func NewConsumer(
	address []string,
	topic string,
	group string,
	service EmailDispatchService,
	retryScheduler email_retry_service.RetryScheduler,
	producer *Producer,
	logger *core_logger.Logger,

) (*Consumer, error) {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":  strings.Join(address, ","),
		"group.id":           group,
		"enable.auto.commit": false,
		"auto.offset.reset":  "earliest",
	}
	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, fmt.Errorf("subscribe topics: %w", err)
	}
	return &Consumer{
		consumer:       c,
		service:        service,
		retryScheduler: retryScheduler,
		producer:       producer,
		logger:         *logger,
	}, nil
}

func (c *Consumer) Run(ctx context.Context) error {
	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			c.logger.Error("read message", zap.Error(err))
			continue
		}

		var event RegistrationEmailMessage
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			c.logger.Error("unmarshal", zap.Error(err))
			continue
		}

		if rand.Intn(100) < 50 {
			c.logger.Error("simulated error at", zap.String("request_id", string(event.RequestID)))

			if err := c.producer.PublishRetryRegistrationRequest(
				ctx,
				event.RequestID,
				1,
			); err != nil {
				c.logger.Error("publish retry", zap.Error(err))
				continue
			}

			if _, err := c.consumer.CommitMessage(msg); err != nil {
				c.logger.Error("commit", zap.Error(err))
			}

			continue
		}

		if err := c.service.SendRegistrationEmail(ctx, event.RequestID); err != nil {

			if err := c.producer.PublishRetryRegistrationRequest(
				ctx,
				event.RequestID,
				1,
			); err != nil {
				c.logger.Error("publish retry", zap.Error(err))
				continue
			}

			if _, err := c.consumer.CommitMessage(msg); err != nil {
				c.logger.Error("commit", zap.Error(err))
			}

			continue
		}

		if _, err := c.consumer.CommitMessage(msg); err != nil {
			c.logger.Error("commit", zap.Error(err))
		}
	}
}
