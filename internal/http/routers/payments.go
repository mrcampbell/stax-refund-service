package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrcampbell/stax-refund-service/internal/http/middleware"
)

func (s *Server) applyPaymentRoutes(r *gin.Engine) {
	r.GET("/api/payments", middleware.AuthMiddleware(), s.listPayments)
}

func (s *Server) listPayments(c *gin.Context) {
	userUID, err := middleware.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	payments, err := s.resources.PaymentClient.ListAll(c.Request.Context(), userUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// todo: map to dto, pagination, etc
	c.JSON(http.StatusOK, payments)
}
