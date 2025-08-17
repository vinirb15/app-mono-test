package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/helpers"
	"github.com/leandro-andrade-candido/auth-service/models"
	"github.com/leandro-andrade-candido/auth-service/repositories"
)

func RefreshHandler(db *sql.DB, cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			RefreshToken string `json:"refresh_token" binding:"required"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
			return
		}

		rt, err := repositories.GetRefreshByPlain(c, db, body.RefreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh inválido"})
			return
		}
		if rt.Revoked {
			_ = repositories.RevokeTokenFamilyOnReuse(c, db, rt.ID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh revogado"})
			return
		}
		if rt.UsedAt.Valid {
			_ = repositories.RevokeTokenFamilyOnReuse(c, db, rt.ID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh já utilizado"})
			return
		}
		if time.Now().UTC().After(rt.ExpiresAt) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh expirado"})
			return
		}

		u, err := repositories.FindUserByID(c, db, rt.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "usuário inválido"})
			return
		}

		access, accessExp, err := helpers.GenerateAccessToken(cfg, u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar access token"})
			return
		}

		newPlain, newRT, err := repositories.CreateRefreshToken(c, db, u.ID, cfg.RefreshTTL, &rt.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar novo refresh"})
			return
		}

		if err := repositories.MarkRefreshUsed(c, db, rt.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar refresh antigo"})
			return
		}

		c.JSON(http.StatusOK, models.TokenResp{
			AccessToken:           access,
			AccessTokenExpiresAt:  accessExp,
			RefreshToken:          newPlain,
			RefreshTokenExpiresAt: newRT.ExpiresAt,
			TokenType:             "Bearer",
		})
	}
}
