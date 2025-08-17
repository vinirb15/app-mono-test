package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/leandro-andrade-candido/auth-service/config"
	"github.com/leandro-andrade-candido/auth-service/database"
	"github.com/leandro-andrade-candido/auth-service/handlers"
	"github.com/leandro-andrade-candido/auth-service/middleware"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connection(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	// Public
	router.POST("/signup", handlers.SignupHandler(db))
	router.POST("/login", handlers.LoginHandler(db, cfg))
	router.POST("/refresh", handlers.RefreshHandler(db, cfg))
	router.POST("/logout", handlers.LogoutHandler(db))
	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	// Protected
	router.GET("/me", middleware.AuthMiddleware(cfg), handlers.MeHandler(db, cfg))

	log.Printf("listening on %s\n", cfg.ListenAddr)
	if err := router.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
