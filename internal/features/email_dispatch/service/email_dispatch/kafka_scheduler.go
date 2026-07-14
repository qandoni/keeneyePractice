package email_dispatch_service

import "context"

type KafkaRetryScheduler struct {
	producer Producer
}

func NewKafkaRetryScheduler(
	producer Producer,
) *KafkaRetryScheduler {
	return &KafkaRetryScheduler{
		producer: producer,
	}
}

func (k *KafkaRetryScheduler) ScheduleRetry(
	ctx context.Context,
	requestID int,
	retry int,
) error {
	return k.producer.PublishRetryRegistrationRequest(
		ctx,
		requestID,
		retry,
	)
}

func (k *KafkaRetryScheduler) ScheduleRegistrationEmail(
	ctx context.Context,
	requestID int,
) error {
	return k.producer.PublishRegistrationRequest(ctx, requestID)
}
