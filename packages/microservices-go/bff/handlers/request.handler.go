package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/bff/models"
)

// -------------------- Proxy interno --------------------

func ProxyHandler(c *gin.Context, serviceURL string) {
	// Aqui você implementa o proxy real
	c.JSON(http.StatusOK, gin.H{"proxy_to": serviceURL})
}

// -------------------- Auth Handlers --------------------

// LoginHandler godoc
// @Summary Login user
// @Description Faz login do usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.User true "Login payload"
// @Success 200 {object} map[string]string
// @Router /login [post]
func LoginHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ProxyHandler(c, cfg.AuthServiceURL+"/login")
	}
}

// LogoutHandler godoc
// @Summary Logout user
// @Description Faz logout do usuário
// @Tags Auth
// @Security BearerAuth
// @Success 200 {string} string "ok"
// @Router /logout [post]
func LogoutHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ProxyHandler(c, cfg.AuthServiceURL+"/logout")
	}
}

// SignupHandler godoc
// @Summary Signup user
// @Description Cadastra novo usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.User true "Signup payload"
// @Success 201 {object} models.User
// @Router /signup [post]
func SignupHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ProxyHandler(c, cfg.AuthServiceURL+"/signup")
	}
}

// RefreshHandler godoc
// @Summary Refresh token
// @Description Atualiza token de autenticação
// @Tags Auth
// @Success 200 {object} map[string]string
// @Router /refresh [post]
func RefreshHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ProxyHandler(c, cfg.AuthServiceURL+"/refresh")
	}
}

// -------------------- Profile Handlers --------------------

// MeHandler godoc
// @Summary Get current user
// @Description Retorna informações do usuário logado
// @Tags Profile
// @Security BearerAuth
// @Success 200 {object} models.User
// @Router /me [get]
func MeHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ProxyHandler(c, cfg.ProfileServiceURL+"/me")
	}
}

// GetFollowersHandler godoc
// @Summary Get followers
// @Description Lista seguidores de um usuário
// @Tags Profile
// @Security BearerAuth
// @Param user_id path string true "User ID"
// @Success 200 {array} models.Follower
// @Router /followers/{user_id} [get]
func GetFollowersHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		ProxyHandler(c, cfg.ProfileServiceURL+"/followers/"+userID)
	}
}

// FollowUserHandler godoc
// @Summary Follow a user
// @Description Segue um usuário
// @Tags Profile
// @Security BearerAuth
// @Param body body models.Follower true "Follow payload"
// @Success 201 {object} models.Follower
// @Router /followers [post]
func FollowUserHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Follower
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ProxyHandler(c, cfg.ProfileServiceURL+"/followers")
	}
}

// -------------------- Posts Handlers --------------------

// GetFeedHandler godoc
// @Summary Get feed
// @Description Retorna feed do usuário
// @Tags Posts
// @Security BearerAuth
// @Success 200 {array} models.FeedPost
// @Router /feed [get]
func GetFeedHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ProxyHandler(c, cfg.PostServiceURL+"/feed")
	}
}

// CreatePostHandler godoc
// @Summary Create post
// @Description Cria um novo post
// @Tags Posts
// @Security BearerAuth
// @Param body body models.Post true "Post payload"
// @Success 201 {object} models.Post
// @Router /posts [post]
func CreatePostHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Post
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ProxyHandler(c, cfg.PostServiceURL+"/posts")
	}
}

// GetPostHandler godoc
// @Summary Get post by ID
// @Description Retorna post pelo ID
// @Tags Posts
// @Security BearerAuth
// @Param post_id path string true "Post ID"
// @Success 200 {object} models.PostDetail
// @Router /posts/{post_id} [get]
func GetPostHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")
		ProxyHandler(c, cfg.PostServiceURL+"/posts/"+postID)
	}
}

// UpdatePostHandler godoc
// @Summary Update post
// @Description Atualiza um post
// @Tags Posts
// @Security BearerAuth
// @Param post_id path string true "Post ID"
// @Param body body models.Post true "Post payload"
// @Success 200 {object} models.Post
// @Router /posts/{post_id} [put]
func UpdatePostHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")
		var req models.Post
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ProxyHandler(c, cfg.PostServiceURL+"/posts/"+postID)
	}
}

// LikePostHandler godoc
// @Summary Like a post
// @Description Marca um post como curtido
// @Tags Posts
// @Security BearerAuth
// @Param body body models.Like true "Like payload"
// @Success 201 {object} models.Like
// @Router /likes [post]
func LikePostHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Like
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ProxyHandler(c, cfg.PostServiceURL+"/likes")
	}
}

// CommentPostHandler godoc
// @Summary Comment on post
// @Description Comenta em um post
// @Tags Posts
// @Security BearerAuth
// @Param body body models.Comment true "Comment payload"
// @Success 201 {object} models.Comment
// @Router /comments [post]
func CommentPostHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Comment
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ProxyHandler(c, cfg.PostServiceURL+"/comments")
	}
}
