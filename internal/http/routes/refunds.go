package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrcampbell/stax-refund-service/app"
	"github.com/mrcampbell/stax-refund-service/internal/http/middleware"
)

func (s *Server) applyRefundRoutes(r *gin.Engine) {
	r.GET("/api/refunds", middleware.AuthMiddleware(), s.listRefunds)
	r.POST("/api/refunds", middleware.AuthMiddleware(), s.createRefund)
}

type createRefundRequest struct {
	PaymentID   string `json:"payment_id" binding:"required"`
	Description string `json:"description"`
}

func (s *Server) listRefunds(c *gin.Context) {
	userUID, err := middleware.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	payments, err := s.resources.RefundService.ListAll(c.Request.Context(), userUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// todo: map to dto, pagination, etc
	c.JSON(http.StatusOK, payments)
}

func (s *Server) createRefund(c *gin.Context) {
	userUID, err := middleware.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var refundRequest createRefundRequest
	if err := c.ShouldBindJSON(&refundRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentUUID, err := uuid.Parse(refundRequest.PaymentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment id"})
		return
	}

	status, err := s.resources.RefundService.RefundPayment(c.Request.Context(), userUID, paymentUUID, refundRequest.Description)
	if err != nil && err == app.ErrorAlreadyExists {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if err != nil && err == app.ErrorNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment with id not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}
