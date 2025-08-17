package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/models"
)

func FindUserByEmail(ctx context.Context, db *sql.DB, email string) (model.User, error) {
	var u model.User
	err := db.QueryRowContext(ctx, `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`, email).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	return u, err
}

func FindUserByID(ctx context.Context, db *sql.DB, id uuid.UUID) (model.User, error) {
	var u model.User
	err := db.QueryRowContext(ctx, `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE id = $1
	`, id).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	return u, err
}
