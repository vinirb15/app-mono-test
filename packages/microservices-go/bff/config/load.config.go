package config

import (
	"os"

	"github.com/leandro-andrade-candido/bff/models"
)

func LoadConfig() models.Config {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret_key"
	}

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	profileService := os.Getenv("PROFILE_SERVICE_URL")
	if profileService == "" {
		profileService = "http://localhost:8082"
	}

	postService := os.Getenv("POST_SERVICE_URL")
	if postService == "" {
		postService = "http://localhost:8081"
	}

	authService := os.Getenv("AUTH_SERVICE_URL")
	if authService == "" {
		authService = "http://localhost:8083"
	}

	return models.Config{
		JWTSecret:         []byte(secret),
		ListenAddr:        addr,
		ProfileServiceURL: profileService,
		PostServiceURL:    postService,
		AuthServiceURL:    authService,
	}
}
