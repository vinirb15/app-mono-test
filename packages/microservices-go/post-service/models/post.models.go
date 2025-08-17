package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type FeedPost struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Content      string         `json:"content"`
	Caption      sql.NullString `json:"caption,omitempty"`
	CreatedAt    string         `json:"created_at"`
	LikeCount    int            `json:"like_count"`
	CommentCount int            `json:"comment_count"`
}

type Post struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	Caption   string    `json:"caption,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type Like struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	PostID  uuid.UUID `json:"post_id"`
	LikedAt time.Time `json:"liked_at"`
}

type Comment struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	PostID      uuid.UUID `json:"post_id"`
	Text        string    `json:"text"`
	CommentedAt time.Time `json:"commented_at"`
}

type CommentResp struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Text        string    `json:"text"`
	CommentedAt time.Time `json:"commented_at"`
}

type LikeResp struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}

type PostDetail struct {
	ID        uuid.UUID      `json:"id"`
	UserID    uuid.UUID      `json:"user_id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Content   string         `json:"content"`
	Caption   sql.NullString `json:"caption,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	Likes     []LikeResp     `json:"likes"`
	Comments  []CommentResp  `json:"comments"`
}
