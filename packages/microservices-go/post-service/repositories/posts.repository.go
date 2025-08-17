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

func UpdatePost(ctx context.Context, db *sql.DB, postID uuid.UUID, userID uuid.UUID, content string, caption string) (models.Post, error) {
	_, err := db.ExecContext(ctx, `
		UPDATE posts
		SET content = $1, caption = $2
		WHERE id = $3 AND user_id = $4
	`, content, caption, postID, userID)
	if err != nil {
		return models.Post{}, err
	}

	row := db.QueryRowContext(ctx, `
		SELECT id, user_id, content, caption, created_at
		FROM posts
		WHERE id = $1
	`, postID)

	var post models.Post
	var captionNull sql.NullString
	if err := row.Scan(&post.ID, &post.UserID, &post.Content, &captionNull, &post.CreatedAt); err != nil {
		return models.Post{}, err
	}

	if captionNull.Valid {
		post.Caption = captionNull.String
	}

	return post, nil
}
