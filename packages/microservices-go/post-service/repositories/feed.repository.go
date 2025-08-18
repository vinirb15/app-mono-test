package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/post-service/models"
)

func FindFeedByUserID(ctx context.Context, db *sql.DB, userID uuid.UUID) ([]models.FeedPost, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT 
			p.id,
			p.user_id,
			u.username,
			u.email,
			p.content,
			p.caption,
			p.created_at,
			COALESCE(l.like_count, 0) AS like_count,
			COALESCE(c.comment_count, 0) AS comment_count
		FROM posts p
		JOIN users u ON u.id = p.user_id
		LEFT JOIN (
			SELECT post_id, COUNT(*) AS like_count
			FROM likes
			GROUP BY post_id
		) l ON l.post_id = p.id
		LEFT JOIN (
			SELECT post_id, COUNT(*) AS comment_count
			FROM comments
			GROUP BY post_id
		) c ON c.post_id = p.id
		WHERE p.user_id = $1
		   OR p.user_id IN (
		       SELECT following_id
		       FROM followers
		       WHERE user_id = $1
		   )
		ORDER BY p.created_at DESC
	`, userID)
	if err != nil {
		log.Println("Error querying feed:", err)
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	var feed []models.FeedPost
	for rows.Next() {
		var f models.FeedPost
		err := rows.Scan(
			&f.ID,
			&f.UserID,
			&f.Username,
			&f.Email,
			&f.Content,
			&f.Caption,
			&f.CreatedAt,
			&f.LikeCount,
			&f.CommentCount,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		feed = append(feed, f)
	}

	return feed, rows.Err()
}
