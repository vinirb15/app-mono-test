package models

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Config struct {
	JWTSecret  []byte
	AccessTTL  time.Duration
	RefreshTTL time.Duration
	DSN        string
	ListenAddr string
}

type User struct {
	ID           uuid.UUID `json:"user_id"`
	Email        string    `json:"email" binding:"required,email"`
	UserName     string    `json:"username" binding:"required"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type JWTClaims struct {
	UserID string `json:"uid"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type FeedPost struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Content      string         `json:"content"`
	Caption      sql.NullString `json:"caption,omitempty"`
	CreatedAt    string         `json:"created_at"`
	LikeCount    int            `json:"like_count"`
	CommentCount int            `json:"comment_count"`
}
