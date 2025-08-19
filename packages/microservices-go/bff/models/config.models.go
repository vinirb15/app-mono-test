package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	JWTSecret         []byte
	ListenAddr        string
	ProfileServiceURL string
	PostServiceURL    string
	AuthServiceURL    string
}

type JWTClaims struct {
	UserID string `json:"uid"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
