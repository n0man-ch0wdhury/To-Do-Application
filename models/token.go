package models

import (
	"time"

	"github.com/google/uuid"
)

// BlacklistedToken represents a token that has been invalidated (logged out)
type BlacklistedToken struct {
	ID        uuid.UUID `json:"id"`
	Token     string    `json:"token"`
	UserID    uuid.UUID `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

// LogoutRequest represents the logout request payload
type LogoutRequest struct {
	Token string `json:"token"`
}