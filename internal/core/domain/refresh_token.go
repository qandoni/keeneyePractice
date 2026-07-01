package domain

import "time"

type RefreshToken struct {
	ID        int
	Version   int
	UserID    int
	TokenHash string
	ExpiresAt time.Time
}

func NewRefreshToken(
	id int,
	version int,
	userID int,
	tokenHash string,
	expiresAt time.Time,
) RefreshToken {
	return RefreshToken{
		ID:        id,
		Version:   version,
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: expiresAt,
	}
}

func NewRefreshTokenUninitialized(
	userID int,
	tokenHash string,
	expiresAt time.Time,
) RefreshToken {
	return NewRefreshToken(
		UninitializedID,
		UninitializedVersion,
		userID,
		tokenHash,
		expiresAt,
	)
}
