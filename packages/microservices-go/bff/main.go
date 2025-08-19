package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/bff/config"
	"github.com/leandro-andrade-candido/bff/handlers"
	"github.com/leandro-andrade-candido/bff/middleware"
)

func main() {
	cfg := config.LoadConfig()
	router := gin.Default()

	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"})
	})

	// Auth Service
	router.POST("/login", handlers.ProxyHandler(cfg))
	router.POST("/logout", handlers.ProxyHandler(cfg))
	router.POST("/signup", handlers.ProxyHandler(cfg))
	router.POST("/refresh", handlers.ProxyHandler(cfg))

	// Profile Service
	router.GET("/me", middleware.AuthMiddleware(cfg), handlers.ProxyHandler(cfg))
	router.GET("/followers/:user_id", middleware.AuthMiddleware(cfg), handlers.ProxyHandler(cfg))
	router.POST("/followers", middleware.AuthMiddleware(cfg), handlers.ProxyHandler(cfg))

	// Posts Service
	router.GET("/feed", middleware.AuthMiddleware(cfg), handlers.ProxyHandler(cfg))
	router.POST("/posts", middleware.AuthMiddleware(cfg), handlers.ProxyHandler(cfg))
	router.GET("/posts/:post_id", middleware.AuthMiddleware(cfg), handlers.ProxyHandler(cfg))
	router.PUT("/posts/:post_id", middleware.AuthMiddleware(cfg), handlers.ProxyHandler(cfg))
	router.POST("/likes", middleware.AuthMiddleware(cfg), handlers.ProxyHandler(cfg))
	router.POST("/comments", middleware.AuthMiddleware(cfg), handlers.ProxyHandler(cfg))

	log.Printf("listening on %s\n", cfg.ListenAddr)
	if err := router.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
