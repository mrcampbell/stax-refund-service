package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrcampbell/stax-refund-service/app"
)

// AuthService is a mock implementation of app.AuthService
// verify that it satisfies the interface at compile time
var _ app.AuthService = &AuthService{}

type AuthService struct {
	loginFn  func(ctx context.Context, username, password string) (string, error)
	verifyFn func(ctx context.Context, token string) error
	userIDFn func(ctx context.Context, token string) (uuid.UUID, error)
}

func NewAuthService(loginFn func(ctx context.Context, username, password string) (string, error)) *AuthService {
	return &AuthService{loginFn: loginFn}
}

func NewAuthServiceWithMockedMethods() *AuthService {
	return &AuthService{
		loginFn:  AuthServiceReturnStubbedToken,
		userIDFn: AuthServiceReturnStubbedUserID,
		verifyFn: AuthServiceVerifyTokenAlwaysValid,
	}
}

func (a *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	return a.loginFn(ctx, username, password)
}

func (a *AuthService) VerifyToken(ctx context.Context, token string) error {
	return a.verifyFn(ctx, token)
}

func (a *AuthService) UserIDFromToken(ctx context.Context, token string) (uuid.UUID, error) {
	return a.userIDFn(ctx, token)
}

func AuthServiceReturnStubbedToken(ctx context.Context, username, password string) (string, error) {
	return MockStubbedAuthToken(), nil
}

func AuthServiceReturnStubbedUserID(ctx context.Context, token string) (uuid.UUID, error) {
	return MockStubbedUserID(), nil
}

func AuthServiceVerifyTokenAlwaysValid(ctx context.Context, token string) error {
	return nil
}
