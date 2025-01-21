package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApplyCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:8080"}
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	return cors.New(config)
}
