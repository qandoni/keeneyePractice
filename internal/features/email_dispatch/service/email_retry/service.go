package email_retry_service

import "context"

type RetryScheduler interface {
	ScheduleRetry(
		ctx context.Context,
		requestID int,
		retry int,
	) error
	ScheduleRegistrationEmail(
		ctx context.Context,
		requestID int,
	) error
}
