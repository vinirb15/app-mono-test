package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"user_id"`
	Email        string    `json:"email" binding:"required,email"`
	UserName     string    `json:"username" binding:"required"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}
