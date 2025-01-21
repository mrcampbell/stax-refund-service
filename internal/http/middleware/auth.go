package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrcampbell/stax-refund-service/config"
	"github.com/mrcampbell/stax-refund-service/internal/mock"
)

const AuthUserIDContextKey = "authenticated-user-id"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		// todo: actually use the auth service, even just the stubbed one
		config.PanicIfNotDev() // because we have stubbed auth
		if token != fmt.Sprintf("Bearer %s", mock.MockStubbedAuthToken()) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.Set(AuthUserIDContextKey, mock.MockStubbedUserID().String())
		c.Next()
	}
}

func GetAuthenticatedUserID(c *gin.Context) (uuid.UUID, error) {
	id := c.GetString(AuthUserIDContextKey)
	return uuid.Parse(id)
}
