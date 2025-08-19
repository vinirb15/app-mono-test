package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- Auth --------------------

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// -------------------- Profile --------------------

type FollowRequest struct {
	FollowingID uuid.UUID `json:"following_id" binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type FollowerResponse struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	FollowingID uuid.UUID `json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`
}

// -------------------- Posts --------------------

type CreatePostRequest struct {
	Content string  `json:"content" binding:"required"`
	Caption *string `json:"caption,omitempty"`
}

type UpdatePostRequest struct {
	Content string  `json:"content" binding:"required"`
	Caption *string `json:"caption,omitempty"`
}

type FeedPostResponse struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Content      string    `json:"content"`
	Caption      *string   `json:"caption,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
}

type LikeRequest struct {
	PostID uuid.UUID `json:"post_id" binding:"required"`
}

type CommentRequest struct {
	PostID uuid.UUID `json:"post_id" binding:"required"`
	Text   string    `json:"text" binding:"required"`
}

type PostDetailResponse struct {
	ID        uuid.UUID        `json:"id"`
	UserID    uuid.UUID        `json:"user_id"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Content   string           `json:"content"`
	Caption   *string          `json:"caption,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	Likes     []LikeRequest    `json:"likes"`
	Comments  []CommentRequest `json:"comments"`
}
