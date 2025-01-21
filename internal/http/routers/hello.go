package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) applyHelloRoutes(r *gin.Engine) {
	r.GET("/hello", s.hello)
}

func (s *Server) hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
