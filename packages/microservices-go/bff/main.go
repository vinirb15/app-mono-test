package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/bff/config"
	_ "github.com/leandro-andrade-candido/bff/docs"
	"github.com/leandro-andrade-candido/bff/handlers"
	"github.com/leandro-andrade-candido/bff/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           BFF API
// @version         1.0
// @description     API Gateway / BFF que proxy requests para Auth, Profile e Post Services.
// @host            localhost:8084
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()
	router := gin.Default()

	// Swagger UI
	router.Use(middleware.AddBearerMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check
	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"})
	})

	// ---------------------------
	// Auth Service
	// ---------------------------

	router.POST("/login", handlers.LoginHandler(cfg))
	router.POST("/logout", handlers.LogoutHandler(cfg))
	router.POST("/signup", handlers.SignupHandler(cfg))
	router.POST("/refresh", handlers.RefreshHandler(cfg))

	// ---------------------------
	// Profile Service
	// ---------------------------

	router.GET("/me", middleware.AuthMiddleware(cfg), handlers.MeHandler(cfg))
	router.GET("/followers/:user_id", middleware.AuthMiddleware(cfg), handlers.GetFollowersHandler(cfg))
	router.POST("/followers", middleware.AuthMiddleware(cfg), handlers.FollowUserHandler(cfg))

	// ---------------------------
	// Posts Service
	// ---------------------------

	router.GET("/feed", middleware.AuthMiddleware(cfg), handlers.GetFeedHandler(cfg))
	router.POST("/posts", middleware.AuthMiddleware(cfg), handlers.CreatePostHandler(cfg))
	router.GET("/posts/:post_id", middleware.AuthMiddleware(cfg), handlers.GetPostHandler(cfg))
	router.PUT("/posts/:post_id", middleware.AuthMiddleware(cfg), handlers.UpdatePostHandler(cfg))
	router.POST("/likes", middleware.AuthMiddleware(cfg), handlers.LikePostHandler(cfg))
	router.POST("/comments", middleware.AuthMiddleware(cfg), handlers.CommentPostHandler(cfg))

	log.Printf("listening on %s\n", cfg.ListenAddr)
	if err := router.Run(cfg.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
