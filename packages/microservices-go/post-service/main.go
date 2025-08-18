package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/post-service/config"
	"github.com/leandro-andrade-candido/post-service/database"
	"github.com/leandro-andrade-candido/post-service/handlers"
	"github.com/leandro-andrade-candido/post-service/middleware"
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
	router.GET("/feed", middleware.AuthMiddleware(cfg), handlers.FeedHandler(db, cfg))
	router.GET("/posts/:postID", handlers.GetPostDetailHandler(db, cfg))
	router.POST("/posts", middleware.AuthMiddleware(cfg), handlers.CreatePostHandler(db, cfg))
	router.PUT("/posts/:postID", middleware.AuthMiddleware(cfg), handlers.UpdatePostHandler(db, cfg))
	router.POST("/likes", middleware.AuthMiddleware(cfg), handlers.CreateLikeHandler(db, cfg))
	router.POST("/comments", middleware.AuthMiddleware(cfg), handlers.CreateCommentHandler(db, cfg))

	log.Printf("listening on %s\n", cfg.ListenAddr)
	if err := router.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
