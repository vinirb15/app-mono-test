package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/profile-service/models"
)

func FindUserByEmail(ctx context.Context, db *sql.DB, email string) (models.User, error) {
	var u models.User
	err := db.QueryRowContext(ctx, `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`, email).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	return u, err
}

func FollowUser(ctx context.Context, db *sql.DB, userID, followingID uuid.UUID) (models.Follower, error) {
	follower := models.Follower{
		ID:          uuid.New(),
		UserID:      userID,
		FollowingID: followingID,
		CreatedAt:   time.Now().UTC(),
	}

	_, err := db.ExecContext(ctx, `
		INSERT INTO followers (id, user_id, following_id)
		VALUES ($1, $2, $3)
	`, follower.ID, follower.UserID, follower.FollowingID)
	if err != nil {
		return models.Follower{}, err
	}

	return follower, nil
}

func GetFollowers(ctx context.Context, db *sql.DB, userID uuid.UUID) ([]models.Follower, error) {
	rows, err := db.QueryContext(ctx, `
        SELECT id, user_id, following_id
        FROM followers
        WHERE user_id = $1
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.Follower
	for rows.Next() {
		var f models.Follower
		if err := rows.Scan(&f.ID, &f.UserID, &f.FollowingID); err != nil {
			return nil, err
		}
		followers = append(followers, f)
	}
	return followers, nil
}
