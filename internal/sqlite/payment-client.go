package sqlite

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrcampbell/stax-refund-service/app"
	"github.com/mrcampbell/stax-refund-service/internal/sqlc"
)

// Ensure PaymentClient implements the app.PaymentClient interface
var _ app.PaymentClient = &PaymentClient{}

type PaymentClient struct {
	queries *sqlc.Queries
}

func NewPaymentClient(queries *sqlc.Queries) *PaymentClient {
	return &PaymentClient{queries: queries}
}

func (p *PaymentClient) GetPaymentByID(ctx context.Context, userID uuid.UUID, paymentID uuid.UUID) (app.Payment, error) {
	row, err := p.queries.GetPaymentByID(ctx, sqlc.GetPaymentByIDParams{
		ID:     paymentID.String(),
		UserID: userID.String(),
	})
	// todo: check for 404
	if err != nil {
		return app.Payment{}, err
	}
	payment, err := toPayment(row)
	if err != nil {
		return app.Payment{}, err
	}

	return payment, nil
}

func (p *PaymentClient) ListAll(ctx context.Context, userID uuid.UUID) ([]app.Payment, error) {
	rows, err := p.queries.ListPayments(ctx, userID.String())
	if err != nil {
		return nil, err
	}
	payments := make([]app.Payment, 0, len(rows))
	for _, row := range rows {
		payment, err := toPayment(row)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, err
}

func toPayment(row sqlc.Payment) (app.Payment, error) {
	uid, err := uuid.Parse(row.ID)
	if err != nil {
		return app.Payment{}, err
	}

	userID, err := uuid.Parse(row.UserID)
	if err != nil {
		return app.Payment{}, err
	}
	return app.Payment{
		ID:            uid,
		UserID:        userID,
		AmountInCents: int(row.AmountInCents),
		Description:   row.Description,
		Timestamp:     row.CreatedAt,
	}, nil
}
