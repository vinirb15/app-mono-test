package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leandro-andrade-candido/bff/models"
)

func ProxyHandler(cfg models.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var target string

		// Descobre qual serviço usar baseado na rota
		path := c.FullPath()
		if strings.HasPrefix(path, "/login") ||
			strings.HasPrefix(path, "/logout") ||
			strings.HasPrefix(path, "/signup") ||
			strings.HasPrefix(path, "/refresh") {
			target = cfg.AuthServiceURL
		} else if strings.HasPrefix(path, "/me") ||
			strings.HasPrefix(path, "/followers") {
			target = cfg.ProfileServiceURL
		} else if strings.HasPrefix(path, "/feed") ||
			strings.HasPrefix(path, "/posts") ||
			strings.HasPrefix(path, "/likes") ||
			strings.HasPrefix(path, "/comments") {
			target = cfg.PostServiceURL
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "route not mapped"})
			return
		}

		// Monta URL final
		url := target + c.Request.URL.Path

		// Cria request para o microserviço
		req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
			return
		}

		// Copia headers originais
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

		// Copia status e body da resposta
		body, _ := io.ReadAll(resp.Body)
		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
	}
}
