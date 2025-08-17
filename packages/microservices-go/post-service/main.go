package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/post-service/config"
	"github.com/leandro-andrade-candido/post-service/database"
	"github.com/leandro-andrade-candido/post-service/handlers"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connection(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})
	router.GET("/feed", handlers.FeedHandler(db, cfg))
	router.POST("/post", handlers.CreatePostHandler(db, cfg))

	log.Printf("listening on %s\n", cfg.ListenAddr)
	if err := router.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
