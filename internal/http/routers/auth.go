package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) applyAuthRoutes(r *gin.Engine) {
	r.POST("/api/users/login", s.login)
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) login(c *gin.Context) {
	var loginRequest loginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := s.resources.AuthService.Login(c.Request.Context(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// todo: implement a real token generation
	// todo: handle bad credentials
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
