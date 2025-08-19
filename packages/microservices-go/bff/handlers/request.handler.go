package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/bff/models"
)

// -------------------- Proxy genérico --------------------
func ProxyHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var target string

		path := c.FullPath()
		switch {
		case strings.HasPrefix(path, "/login"),
			strings.HasPrefix(path, "/logout"),
			strings.HasPrefix(path, "/signup"),
			strings.HasPrefix(path, "/refresh"):
			target = cfg.AuthServiceURL
		case strings.HasPrefix(path, "/me"),
			strings.HasPrefix(path, "/followers"):
			target = cfg.ProfileServiceURL
		case strings.HasPrefix(path, "/feed"),
			strings.HasPrefix(path, "/posts"),
			strings.HasPrefix(path, "/likes"),
			strings.HasPrefix(path, "/comments"):
			target = cfg.PostServiceURL
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": "route not mapped"})
			return
		}

		url := target + c.Request.URL.Path

		req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
			return
		}

		for k, v := range c.Request.Header {
			req.Header[k] = v
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "service unavailable"})
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
	}
}

// -------------------- Auth Handlers --------------------

// @Summary Login user
// @Description Faz login do usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.LoginRequest true "Login payload"
// @Success 200 {object} map[string]string
// @Router /login [post]
func LoginHandler(cfg models.Config) gin.HandlerFunc {
	return ProxyHandler(cfg)
}

// LogoutHandler godoc
// @Summary Logout user
// @Description Faz logout do usuário usando refresh token
// @Tags Auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body models.LogoutRequest true "Logout payload"
// @Success 200 {string} string "ok"
// @Router /logout [post]
func LogoutHandler(cfg models.Config) gin.HandlerFunc {
	return ProxyHandler(cfg)
}

// SignupHandler godoc
// @Summary Signup user
// @Description Cria novo usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.SignupRequest true "Signup payload"
// @Success 201 {object} map[string]string
// @Router /signup [post]
func SignupHandler(cfg models.Config) gin.HandlerFunc {
	return ProxyHandler(cfg)
}

// RefreshHandler godoc
// @Summary Refresh token
// @Description Atualiza o access token usando refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.RefreshRequest true "Refresh payload"
// @Success 200 {object} map[string]string
// @Router /refresh [post]
func RefreshHandler(cfg models.Config) gin.HandlerFunc {
	return ProxyHandler(cfg)
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
	return ProxyHandler(cfg)
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
	return ProxyHandler(cfg)
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
	return ProxyHandler(cfg)
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
	return ProxyHandler(cfg)
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
	return ProxyHandler(cfg)
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
	return ProxyHandler(cfg)
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
	return ProxyHandler(cfg)
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
	return ProxyHandler(cfg)
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
	return ProxyHandler(cfg)
}
