package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/stax-refund-service/config"
	"github.com/mrcampbell/stax-refund-service/internal/http/middleware"
)

func (s *Server) allRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ApplyCors())

	s.applyHelloRoutes(r)
	s.applyAuthRoutes(r)
	s.applyPaymentRoutes(r)
	s.applyRefundRoutes(r)

	if config.IsDev() {
		s.applyDocsRoute(r)
	}

	return r
}
