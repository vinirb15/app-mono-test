package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/post-service/models"
)

func CreatePost(ctx context.Context, db *sql.DB, userID uuid.UUID, content string, caption string) (models.Post, error) {
	post := models.Post{
		ID:        uuid.New(),
		UserID:    userID,
		Content:   content,
		Caption:   caption,
		CreatedAt: time.Now().UTC(),
	}

	_, err := db.ExecContext(ctx, `
		INSERT INTO posts (id, user_id, content, caption, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`, post.ID, post.UserID, post.Content, post.Caption, post.CreatedAt)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}
