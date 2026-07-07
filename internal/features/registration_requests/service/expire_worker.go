package registration_service

import (
	"context"
	"log"
	"time"
)

type ExpireWorker struct {
	repo RegistrationRequestRepository
}

func NewExpireWorker(
	repo RegistrationRequestRepository,
) *ExpireWorker {
	return &ExpireWorker{
		repo: repo,
	}
}

func (w *ExpireWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := w.repo.ExpireRequests(ctx); err != nil {
				log.Printf("expire requests: %v", err)
			}
		}
	}
}
