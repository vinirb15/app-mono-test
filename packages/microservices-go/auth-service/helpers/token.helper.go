package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/models"
)

func GenerateAccessToken(cfg models.Config, u models.User) (string, time.Time, error) {
	now := time.Now().UTC()
	exp := now.Add(cfg.AccessTTL)

	claims := models.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "authsvc",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	if u.ID != uuid.Nil {
		claims.UserID = u.ID.String()
		claims.Email = u.Email
		claims.Subject = u.ID.String()
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

func GetEmailFromToken(cfg models.Config, tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		return cfg.JWTSecret, nil
	})
	if err != nil {
		log.Printf("failed to parse token: %v", err)
		return "", err
	}

	if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
		return claims.Email, nil
	}

	log.Println("invalid token or missing email")
	return "", errors.New("invalid token or missing email")
}
