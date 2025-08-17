package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/auth-service/repositories"
)

func LogoutHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
			return
		}

		rt, err := repository.GetRefreshByPlain(c, db, body.RefreshToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token não encontrado"})
			return
		}

		if err := repository.RevokeTokenFamilyOnReuse(c, db, rt.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao revogar token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "logout successful",
		})
	}
}
