package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leandro-andrade-candido/post-service/helpers"
	"github.com/leandro-andrade-candido/post-service/models"
	"github.com/leandro-andrade-candido/post-service/repositories"
)

func CreatePostHandler(db *sql.DB, cfg models.Config) gin.HandlerFunc {
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

		var req struct {
			Content string `json:"content" binding:"required"`
			Caption string `json:"caption"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
			return
		}

		post, err := repositories.CreatePost(c, db, user.ID, req.Content, req.Caption)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post"})
			return
		}

		c.JSON(http.StatusCreated, post)
	}
}

func UpdatePostHandler(db *sql.DB, cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		postIDParam := c.Param("postID")
		postID, err := uuid.Parse(postIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
			return
		}

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

		var req struct {
			Content string `json:"content" binding:"required"`
			Caption string `json:"caption"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
			return
		}

		post, err := repositories.UpdatePost(c, db, postID, user.ID, req.Content, req.Caption)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
			return
		}

		c.JSON(http.StatusOK, post)
	}
}
