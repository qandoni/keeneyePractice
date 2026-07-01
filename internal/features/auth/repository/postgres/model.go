package auth_postgres_repository

import (
	"time"

	"github.com/qandoni/keeneyePractice/internal/core/domain"
)

type RefreshTokenModel struct {
	ID        int
	Version   int
	UserID    int
	TokenHash string
	ExpiresAt time.Time
}

func modelToDomain(
	m RefreshTokenModel,
) domain.RefreshToken {
	return domain.RefreshToken{
		ID:        m.ID,
		Version:   m.Version,
		TokenHash: m.TokenHash,
		UserID:    m.UserID,
		ExpiresAt: m.ExpiresAt,
	}
}
