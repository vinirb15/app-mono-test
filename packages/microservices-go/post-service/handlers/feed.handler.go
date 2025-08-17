package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/post-service/helpers"
	"github.com/leandro-andrade-candido/post-service/models"
	"github.com/leandro-andrade-candido/post-service/repositories"
)

func FeedHandler(db *sql.DB, cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		email, err := helpers.GetEmailFromToken(cfg, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		user, err := repositories.FindUserByEmail(c, db, email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		feed, err := repositories.FindFeedByUserID(c, db, user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve feed"})
			return
		}

		c.JSON(http.StatusOK, feed)
	}
}
