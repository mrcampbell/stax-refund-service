package mock

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mrcampbell/stax-refund-service/app"
	"github.com/mrcampbell/stax-refund-service/config"
	"github.com/mrcampbell/stax-refund-service/internal/sqlc"
)

func PopulateSampleData(ctx context.Context, queries *sqlc.Queries) {
	config.PanicIfNotDev()
	err := PopulateSamplePayments(ctx, queries)
	if err != nil {
		panic(fmt.Errorf("error populating mock payments: %v", err))
	}

	err = PopulateSampleRefunds(ctx, queries)
	if err != nil {
		panic(fmt.Errorf("error populating mock refunds: %v", err))
	}
}

func PopulateSamplePayments(ctx context.Context, queries *sqlc.Queries) error {
	err := CreateSamplePaymentWithStubbedUserID(ctx, queries, PaymentOneID(), "mock payment 1", 100, time.Now())
	if err != nil {
		return fmt.Errorf("error creating mock payment: %v", err)
	}
	err = CreateSamplePaymentWithStubbedUserID(ctx, queries, PaymentTwoID(), "mock payment 2", 200, time.Now())
	if err != nil {
		return fmt.Errorf("error creating mock payment: %v", err)
	}
	err = CreateSamplePaymentWithStubbedUserID(ctx, queries, PaymentThreeID(), "mock payment 3", 300, time.Now())
	if err != nil {
		return fmt.Errorf("error creating mock payment: %v", err)
	}
	return nil
}

func PopulateSampleRefunds(ctx context.Context, queries *sqlc.Queries) error {
	err := CreateSampleRefundWithStubbedUserID(
		ctx, queries, RefundOneID(),
		PaymentOneID(),
		"mock refund 1",
		time.Now(),
		app.RefundStatusSuccess,
	)
	if err != nil {
		return fmt.Errorf("error creating mock refund: %v", err)
	}
	err = CreateSampleRefundWithStubbedUserID(
		ctx,
		queries,
		RefundTwoID(),
		PaymentTwoID(),
		"mock refund 2",
		time.Now(),
		app.RefundStatusPending,
	)
	if err != nil {
		return fmt.Errorf("error creating mock refund: %v", err)
	}
	return nil
}

func CreateSamplePaymentWithStubbedUserID(ctx context.Context, queries *sqlc.Queries, id uuid.UUID, description string, amountInCents int, createdAt time.Time) error {
	log.Println("Creating mock payment with id: ", id.String())
	_, err := queries.CreatePayment(ctx, sqlc.CreatePaymentParams{
		AmountInCents: int64(amountInCents),
		ID:            id.String(),
		UserID:        MockStubbedUserID().String(),
		Description:   description,
		CreatedAt:     createdAt,
	})
	if err != nil {
		return fmt.Errorf("error creating mock payment: %v", err)
	}
	return nil
}

func CreateSampleRefundWithStubbedUserID(ctx context.Context, queries *sqlc.Queries, id uuid.UUID, paymentID uuid.UUID, description string, requestedAt time.Time, status app.RefundStatus) error {
	log.Println("Creating mock refund with id: ", id.String())
	_, err := queries.CreateRefund(ctx, sqlc.CreateRefundParams{
		ID:          id.String(),
		PaymentID:   paymentID.String(),
		Description: description,
		UserID:      MockStubbedUserID().String(),
		RequestedAt: requestedAt,
		UpdatedAt:   requestedAt,
		Status:      int64(status),
	})
	if err != nil {
		return fmt.Errorf("error creating mock refund: %v", err)
	}
	return nil
}
