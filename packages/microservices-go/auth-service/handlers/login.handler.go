package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/helpers"
	"github.com/leandro-andrade-candido/auth-service/models"
	"github.com/leandro-andrade-candido/auth-service/repositories"
	"github.com/leandro-andrade-candido/auth-service/services"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(db *sql.DB, cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.LoginReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
			return
		}
		req.Email = strings.ToLower(strings.TrimSpace(req.Email))

		u, err := services.FindUserByEmail(c, cfg, req.Email)
		if err != nil || bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciais inválidas"})
			return
		}

		access, accessExp, err := helpers.GenerateAccessToken(cfg, u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar access token"})
			return
		}
		plainRefresh, savedRT, err := repositories.CreateRefreshToken(c, db, u.ID, cfg.RefreshTTL, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar refresh token"})
			return
		}

		c.JSON(http.StatusOK, models.TokenResp{
			AccessToken:           access,
			AccessTokenExpiresAt:  accessExp,
			RefreshToken:          plainRefresh,
			RefreshTokenExpiresAt: savedRT.ExpiresAt,
			TokenType:             "Bearer",
		})
	}
}
