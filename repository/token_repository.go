package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/noman/todo-application/database"
	"github.com/noman/todo-application/models"
)

// TokenRepository handles database operations for tokens
type TokenRepository struct {
	db *sql.DB
}

// NewTokenRepository creates a new TokenRepository
func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		db: database.DB,
	}
}

// BlacklistToken adds a token to the blacklist
func (r *TokenRepository) BlacklistToken(token string, userID uuid.UUID, expiresAt time.Time) error {
	// Create a new blacklisted token
	blacklistedToken := &models.BlacklistedToken{
		ID:        uuid.New(),
		Token:     token,
		UserID:    userID,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}

	// Insert the token into the database
	query := `
	INSERT INTO blacklisted_tokens (id, token, user_id, expires_at, created_at)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(query, blacklistedToken.ID, blacklistedToken.Token, blacklistedToken.UserID, blacklistedToken.ExpiresAt, blacklistedToken.CreatedAt)
	return err
}

// IsTokenBlacklisted checks if a token is blacklisted
func (r *TokenRepository) IsTokenBlacklisted(token string) (bool, error) {
	query := `
	SELECT COUNT(*) FROM blacklisted_tokens
	WHERE token = $1 AND expires_at > $2
	`

	var count int
	err := r.db.QueryRow(query, token, time.Now()).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// CleanupExpiredTokens removes expired tokens from the blacklist
func (r *TokenRepository) CleanupExpiredTokens() error {
	query := `
	DELETE FROM blacklisted_tokens
	WHERE expires_at <= $1
	`

	_, err := r.db.Exec(query, time.Now())
	return err
}