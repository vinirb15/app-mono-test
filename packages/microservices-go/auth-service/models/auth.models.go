package models

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type SignupReq struct {
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type TokenResp struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	TokenType             string    `json:"token_type"`
}

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

type RefreshToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TokenHash string
	ExpiresAt time.Time
	CreatedAt time.Time
	Revoked   bool
	UsedAt    sql.NullTime
	ParentID  sql.NullString
}

type JWTClaims struct {
	UserID   string `json:"uid"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}
