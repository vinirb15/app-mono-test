package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/profile-service/helpers"
	"github.com/leandro-andrade-candido/profile-service/models"
	"github.com/leandro-andrade-candido/profile-service/repositories"
)

func FollowUserHandler(db *sql.DB, cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
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

		var req struct {
			FollowingID string `json:"following_id" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}

		followingID, err := uuid.Parse(req.FollowingID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid following_id"})
			return
		}

		follower, err := repositories.FollowUser(c, db, user.ID, followingID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to follow user"})
			return
		}

		c.JSON(http.StatusCreated, follower)
	}
}

func GetFollowersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDParam := c.Param("userID")
		userID, err := uuid.Parse(userIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
			return
		}

		followers, err := repositories.GetFollowers(c, db, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get followers"})
			return
		}

		c.JSON(http.StatusOK, followers)
	}
}
