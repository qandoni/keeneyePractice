package registration_contracts

import (
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/enum"
)

type UpdateEmailAttemptInput struct {
	ID      int
	Version int

	TokenHash string

	ExpiresAt time.Time

	EmailStatus enum.EmailStatus

	EmailRetryCount int

	LastEmailAttempt *time.Time

	EmailSentAt *time.Time
}
