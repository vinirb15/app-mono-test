package helpers

import (
	"errors"
	"log"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/post-service/models"
)

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
