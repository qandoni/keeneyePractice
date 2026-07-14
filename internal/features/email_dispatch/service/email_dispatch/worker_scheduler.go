package email_dispatch_service

import "context"

type WorkerRetryScheduler struct {
}

func NewWorkerRetryScheduler() *WorkerRetryScheduler {
	return &WorkerRetryScheduler{}
}

func (w *WorkerRetryScheduler) ScheduleRetry(
	ctx context.Context,
	requestID int,
	retry int,
) error {

	return nil
}

func (w *WorkerRetryScheduler) ScheduleRegistrationEmail(
	ctx context.Context,
	requestID int,
) error {
	return nil
}
