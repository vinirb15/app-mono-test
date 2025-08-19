package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/profile-service/config"
	"github.com/leandro-andrade-candido/profile-service/database"
	"github.com/leandro-andrade-candido/profile-service/handlers"
	"github.com/leandro-andrade-candido/profile-service/middleware"
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

	router.GET("/me", middleware.AuthMiddleware(cfg), handlers.MeHandler(db, cfg))
	router.GET("/user/:email", handlers.GetUserByEmailHandler(db))
	router.POST("/followers", middleware.AuthMiddleware(cfg), handlers.FollowUserHandler(db, cfg))
	router.GET("/followers/:userID", middleware.AuthMiddleware(cfg), handlers.GetFollowersHandler(db))

	log.Printf("listening on %s\n", cfg.ListenAddr)
	if err := router.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
