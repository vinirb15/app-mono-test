package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/config"
	"github.com/leandro-andrade-candido/auth-service/database"
	"github.com/leandro-andrade-candido/auth-service/handlers"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connection(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}()

	router := gin.Default()

	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	router.POST("/signup", handlers.SignupHandler(db))
	router.POST("/login", handlers.LoginHandler(db, cfg))
	router.POST("/refresh", handlers.RefreshHandler(db, cfg))
	router.POST("/logout", handlers.LogoutHandler(db))

	log.Printf("listening on %s\n", cfg.ListenAddr)
	if err := router.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
