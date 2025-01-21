package routes

import (
	_ "embed"

	"github.com/gin-gonic/gin"
)

func (s *Server) applyDocsRoute(r *gin.Engine) {
	r.GET("/docs", s.swaggerFile)
}

func (s *Server) swaggerFile(c *gin.Context) {
	c.File("docs/openapi.yml")
}
