package config

import (
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/models"
)

func LoadConfig() models.Config {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret_key"
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/dialog?sslmode=disable"
	}

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	profileService := os.Getenv("PROFILE_SERVICE_URL")
	if profileService == "" {
		profileService = "http://localhost:8082"
	}

	access := 15 * time.Minute
	if v := os.Getenv("ACCESS_TTL"); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			access = d
		}
	}

	refresh := 7 * 24 * time.Hour
	if v := os.Getenv("REFRESH_TTL"); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			refresh = d
		}
	}
	return models.Config{
		JWTSecret:         []byte(secret),
		AccessTTL:         access,
		RefreshTTL:        refresh,
		DSN:               dsn,
		ListenAddr:        addr,
		ProfileServiceURL: profileService,
	}
}
