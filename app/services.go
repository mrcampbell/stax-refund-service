package app

import (
	"context"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(ctx context.Context, username string, password string) (string, error)
	VerifyToken(ctx context.Context, token string) error
	UserIDFromToken(ctx context.Context, token string) (uuid.UUID, error)
}

type PaymentClient interface {
	ListAll(ctx context.Context, userID uuid.UUID) ([]Payment, error)
	GetPaymentByID(ctx context.Context, userID uuid.UUID, paymentID uuid.UUID) (Payment, error)
}

type RefundService interface {
	RefundPayment(ctx context.Context, userID uuid.UUID, paymentID uuid.UUID, description string) (RefundStatus, error)
	ListAll(ctx context.Context, userID uuid.UUID) ([]Refund, error)
}
