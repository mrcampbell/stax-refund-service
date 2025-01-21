package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrcampbell/stax-refund-service/app"
)

// PaymentClient is a mock implementation of app.PaymentClient
// verify that it satisfies the interface at compile time
var _ app.PaymentClient = &PaymentClient{}

type PaymentClient struct {
	listAllFn func(ctx context.Context, userID uuid.UUID) ([]app.Payment, error)
	getByIDFn func(ctx context.Context, userID uuid.UUID, paymentID uuid.UUID) (app.Payment, error)
}

func NewPaymentClient(
	listAllFn func(ctx context.Context, userID uuid.UUID) ([]app.Payment, error),
	getByIDFn func(ctx context.Context, userID uuid.UUID, paymentID uuid.UUID) (app.Payment, error),
) *PaymentClient {
	return &PaymentClient{listAllFn: listAllFn, getByIDFn: getByIDFn}
}

func (p *PaymentClient) ListAll(ctx context.Context, userID uuid.UUID) ([]app.Payment, error) {
	return p.listAllFn(ctx, userID)
}

func (p *PaymentClient) GetPaymentByID(ctx context.Context, userID uuid.UUID, paymentID uuid.UUID) (app.Payment, error) {
	return p.getByIDFn(ctx, userID, paymentID)
}
