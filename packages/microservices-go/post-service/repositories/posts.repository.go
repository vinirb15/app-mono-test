package repositories

import (
	"context"
	"database/sql"
	"log"
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

func CreateLike(ctx context.Context, db *sql.DB, userID, postID uuid.UUID) (models.Like, error) {
	like := models.Like{
		ID:      uuid.New(),
		UserID:  userID,
		PostID:  postID,
		LikedAt: time.Now().UTC(),
	}

	_, err := db.ExecContext(ctx, `
		INSERT INTO likes (id, user_id, post_id, liked_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id, post_id) DO NOTHING
	`, like.ID, like.UserID, like.PostID, like.LikedAt)
	if err != nil {
		return models.Like{}, err
	}

	return like, nil
}

func CreateComment(ctx context.Context, db *sql.DB, userID, postID uuid.UUID, text string) (models.Comment, error) {
	comment := models.Comment{
		ID:          uuid.New(),
		UserID:      userID,
		PostID:      postID,
		Text:        text,
		CommentedAt: time.Now().UTC(),
	}

	_, err := db.ExecContext(ctx, `
		INSERT INTO comments (id, user_id, post_id, text, commented_at)
		VALUES ($1, $2, $3, $4, $5)
	`, comment.ID, comment.UserID, comment.PostID, comment.Text, comment.CommentedAt)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func GetPostWithDetails(ctx context.Context, db *sql.DB, postID uuid.UUID) (models.PostDetail, error) {
	var post models.PostDetail

	// Buscar dados principais do post
	err := db.QueryRowContext(ctx, `
		SELECT p.id, p.user_id, u.username, u.email, p.content, p.caption, p.created_at
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.id = $1
	`, postID).Scan(&post.ID, &post.UserID, &post.Username, &post.Email, &post.Content, &post.Caption, &post.CreatedAt)
	if err != nil {
		return models.PostDetail{}, err
	}

	// Buscar likes
	likeRows, err := db.QueryContext(ctx, `
		SELECT id, user_id
		FROM likes
		WHERE post_id = $1
	`, postID)
	if err != nil {
		return models.PostDetail{}, err
	}
	defer func() {
		if err := likeRows.Close(); err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	for likeRows.Next() {
		var l models.LikeResp
		if err := likeRows.Scan(&l.ID, &l.UserID); err != nil {
			return models.PostDetail{}, err
		}
		post.Likes = append(post.Likes, l)
	}

	// Buscar comentários
	commentRows, err := db.QueryContext(ctx, `
		SELECT id, user_id, text, commented_at
		FROM comments
		WHERE post_id = $1
		ORDER BY commented_at ASC
	`, postID)
	if err != nil {
		return models.PostDetail{}, err
	}
	defer func() {
		if err := commentRows.Close(); err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	for commentRows.Next() {
		var c models.CommentResp
		if err := commentRows.Scan(&c.ID, &c.UserID, &c.Text, &c.CommentedAt); err != nil {
			return models.PostDetail{}, err
		}
		post.Comments = append(post.Comments, c)
	}

	return post, nil
}
