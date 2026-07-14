package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const flushTimeout = 5000

var errUnknownType = errors.New("unknown event type")

type Producer struct {
	producer *kafka.Producer

	mainTopic  string
	retryTopic string
	deadTopic  string
}

func NewProducer(
	address []string,
	mainTopic string,
	retryTopic string,
	deadTopic string,
) *Producer {

	cfg := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(address, ","),
	}

	p, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil
	}

	return &Producer{
		producer:   p,
		mainTopic:  mainTopic,
		retryTopic: retryTopic,
		deadTopic:  deadTopic,
	}
}

func (p *Producer) PublishRegistrationRequest(
	ctx context.Context,
	requestID int,
) error {

	msg := RegistrationEmailMessage{
		RequestID: requestID,
	}

	data, _ := json.Marshal(msg)

	return p.publish(
		p.mainTopic,
		data,
	)
}

func (p *Producer) PublishRetryRegistrationRequest(
	ctx context.Context,
	requestID int,
	retry int,
) error {

	msg := RetryMessage{
		RequestID: requestID,
		Retry:     retry,
	}

	data, _ := json.Marshal(msg)

	return p.publish(
		p.retryTopic,
		data,
	)
}

func (p *Producer) PublishDeadRegistrationRequest(
	ctx context.Context,
	requestID int,
) error {

	msg := RegistrationEmailMessage{
		RequestID: requestID,
	}
	data, _ := json.Marshal(msg)
	return p.publish(
		p.deadTopic,
		data,
	)
}

func (p *Producer) publish(
	topic string,
	value []byte,
) error {

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}

	delivery := make(chan kafka.Event)

	err := p.producer.Produce(msg, delivery)

	if err != nil {
		return err
	}

	event := <-delivery

	result := event.(*kafka.Message)

	if result.TopicPartition.Error != nil {
		return result.TopicPartition.Error
	}

	return nil
}

func (p *Producer) Close() {

	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
