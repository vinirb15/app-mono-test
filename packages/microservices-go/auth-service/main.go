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
	router.POST("/signup", handler.SignupHandler(db))
	router.POST("/login", handler.LoginHandler(db, cfg))
	router.POST("/refresh", handler.RefreshHandler(db, cfg))
	router.POST("/logout", handler.LogoutHandler(db))

	// Protected
	router.GET("/me", middleware.AuthMiddleware(cfg), handler.MeHandler())

	log.Printf("listening on %s\n", cfg.ListenAddr)
	if err := router.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
