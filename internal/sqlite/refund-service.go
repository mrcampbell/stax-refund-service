package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mrcampbell/stax-refund-service/app"
	"github.com/mrcampbell/stax-refund-service/internal/sqlc"
)

// RefundService is a service that handles refunding payments
// ensure that the service implements the RefundService interface
var _ app.RefundService = &RefundService{}

type RefundService struct {
	queries *sqlc.Queries
}

func NewRefundService(queries *sqlc.Queries) *RefundService {
	return &RefundService{queries: queries}
}

func (r *RefundService) RefundPayment(ctx context.Context, userID uuid.UUID, paymentID uuid.UUID, description string) (app.RefundStatus, error) {
	// first, verify no refund has been triggered for this payment
	existingRefund, err := r.queries.GetRefundByPaymentID(ctx, sqlc.GetRefundByPaymentIDParams{
		PaymentID: paymentID.String(),
		UserID:    userID.String(),
	})

	// a success case is when there is an error, and that error is specifically that no rows were found
	// this means that no refund has been triggered for this payment. Any other error is one to be concerned about
	if err != nil && err != sql.ErrNoRows {
		return app.RefundStatusUnknown, fmt.Errorf("unknown error occurred: %v", err)
	}

	if existingStatus := app.RefundStatus(existingRefund.Status); existingStatus != app.RefundStatusUnknown {
		return existingStatus, app.ErrorAlreadyExists
	}

	// verify payment exists
	_, err = r.queries.GetPaymentByID(ctx, sqlc.GetPaymentByIDParams{
		ID:     paymentID.String(),
		UserID: userID.String(),
	})
	if err != nil && err == sql.ErrNoRows {
		return app.RefundStatusUnknown, app.ErrorNotFound
	}

	// todo: there may be other determinating factors for refund status, but we'll keep it simple for now
	refundStatus := app.RefundStatusPending

	// trigger refund request
	_, err = r.queries.CreateRefund(ctx, sqlc.CreateRefundParams{
		ID:          uuid.New().String(),
		PaymentID:   paymentID.String(),
		UserID:      userID.String(),
		Description: description,
		Status:      int64(refundStatus),
		RequestedAt: time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return app.RefundStatusUnknown, fmt.Errorf("error creating refund: %v", err)
	}

	// return refund status
	return refundStatus, nil
}

func (r *RefundService) ListAll(ctx context.Context, userID uuid.UUID) ([]app.Refund, error) {
	rows, err := r.queries.ListRefunds(ctx, userID.String())
	if err != nil {
		return nil, err
	}

	refunds := make([]app.Refund, 0, len(rows))
	for _, row := range rows {
		refund, err := mapRowToRefund(row)
		if err != nil {
			return nil, err
		}
		refunds = append(refunds, refund)
	}
	return refunds, nil
}

func mapRowToRefund(row sqlc.Refund) (app.Refund, error) {
	uid, err := uuid.Parse(row.ID)
	if err != nil {
		return app.Refund{}, err
	}

	paymentID, err := uuid.Parse(row.PaymentID)
	if err != nil {
		return app.Refund{}, err
	}

	return app.Refund{
		ID:          uid,
		PaymentID:   paymentID,
		Status:      app.RefundStatus(row.Status),
		UserID:      uuid.MustParse(row.UserID),
		Description: row.Description,
		RequestedAt: row.RequestedAt,
		UpdatedAt:   row.UpdatedAt,
	}, nil
}
