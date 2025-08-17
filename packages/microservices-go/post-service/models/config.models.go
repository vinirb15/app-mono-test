package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	JWTSecret  []byte
	AccessTTL  time.Duration
	RefreshTTL time.Duration
	DSN        string
	ListenAddr string
}

type JWTClaims struct {
	UserID string `json:"uid"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
