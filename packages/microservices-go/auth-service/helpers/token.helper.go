package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/models"
)

func GenerateAccessToken(cfg models.Config, u models.User) (string, time.Time, error) {
	now := time.Now().UTC()
	exp := now.Add(cfg.AccessTTL)
	claims := models.JWTClaims{
		UserID: u.ID.String(),
		Email:  u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "authsvc",
			Subject:   u.ID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := tk.SignedString(cfg.JWTSecret)
	return s, exp, err
}

func RandomToken(bytesLen int) (string, error) {
	b := make([]byte, bytesLen)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func HashToken(plain string) string {
	sum := sha256.Sum256([]byte(plain))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
