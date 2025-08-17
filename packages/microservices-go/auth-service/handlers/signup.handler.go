package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/models"
	"github.com/leandro-andrade-candido/auth-service/repositories"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.SignupReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
			return
		}
		req.Email = strings.ToLower(strings.TrimSpace(req.Email))

		if _, err := repositories.FindUserByEmail(c, db, req.Email); err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "email já cadastrado"})
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar hash"})
			return
		}

		id := uuid.New()

		_, err = db.ExecContext(c, `
			INSERT INTO users (id, email, password_hash)
			VALUES ($1,$2,$3)
		`, id, req.Email, string(hash))
		if err != nil {
			log.Printf("Erro ao criar usuário: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar usuário"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id, "email": req.Email})
	}
}
