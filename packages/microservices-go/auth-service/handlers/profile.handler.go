package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/helpers"
	"github.com/leandro-andrade-candido/auth-service/models"
	"github.com/leandro-andrade-candido/auth-service/services"
)

func MeHandler(db *sql.DB, cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Println("Authorization header:", authHeader)
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

		user, err := services.FindUserByEmail(c, cfg, email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, models.User{
			ID:        user.ID,
			Email:     user.Email,
			UserName:  user.UserName,
			CreatedAt: user.CreatedAt,
		})
	}
}
