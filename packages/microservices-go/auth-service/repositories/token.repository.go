package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/helpers"
	"github.com/leandro-andrade-candido/auth-service/models"
)

func CreateRefreshToken(ctx context.Context, db *sql.DB, userID uuid.UUID, ttl time.Duration, parentID *uuid.UUID) (plain string, saved models.RefreshToken, err error) {
	plain, err = helpers.RandomToken(32)
	if err != nil {
		return "", models.RefreshToken{}, err
	}
	now := time.Now().UTC()
	id := uuid.New()
	var parent any
	if parentID != nil {
		parent = *parentID
	} else {
		parent = nil
	}
	_, err = db.ExecContext(ctx, `
		INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at, created_at, revoked, parent_id)
		VALUES ($1,$2,$3,$4,$5,false,$6)
	`, id, userID, helpers.HashToken(plain), now.Add(ttl), now, parent)
	if err != nil {
		return "", models.RefreshToken{}, err
	}
	return plain, models.RefreshToken{
		ID:        id,
		UserID:    userID,
		TokenHash: "",
		ExpiresAt: now.Add(ttl),
		CreatedAt: now,
		Revoked:   false,
	}, nil
}

func GetRefreshByPlain(ctx context.Context, db *sql.DB, plain string) (models.RefreshToken, error) {
	var rt models.RefreshToken
	h := helpers.HashToken(plain)
	err := db.QueryRowContext(ctx, `
		SELECT id, user_id, token_hash, expires_at, created_at, revoked, used_at, parent_id
		FROM refresh_tokens
		WHERE token_hash = $1
	`, h).Scan(&rt.ID, &rt.UserID, &rt.TokenHash, &rt.ExpiresAt, &rt.CreatedAt, &rt.Revoked, &rt.UsedAt, &rt.ParentID)
	return rt, err
}

func MarkRefreshUsed(ctx context.Context, db *sql.DB, id uuid.UUID) error {
	_, err := db.ExecContext(ctx, `
		UPDATE refresh_tokens
		SET used_at = NOW(), revoked = true
		WHERE id = $1 AND revoked = false AND used_at IS NULL
	`, id)
	return err
}

func RevokeTokenFamilyOnReuse(ctx context.Context, db *sql.DB, id uuid.UUID) error {
	_, err := db.ExecContext(ctx, `
		UPDATE refresh_tokens
		SET revoked = true
		WHERE id = $1 OR parent_id = $1
	`, id)
	return err
}
