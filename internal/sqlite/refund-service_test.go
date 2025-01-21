package sqlite

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/mrcampbell/stax-refund-service/app"
	"github.com/mrcampbell/stax-refund-service/db"
	"github.com/mrcampbell/stax-refund-service/internal/mock"
	"github.com/mrcampbell/stax-refund-service/internal/sqlc"
)

func TestRefundService_RefundPayment(t *testing.T) {
	ctx := context.Background()
	paymentThreeRefundDescription := "refund for payment 3"
	nonExistantPaymentID, _ := uuid.Parse("395eda5a-6ae9-473f-9efb-1bc401fec1d1")

	type fields struct {
		queries *sqlc.Queries
	}
	type args struct {
		ctx         context.Context
		userID      uuid.UUID
		paymentID   uuid.UUID
		description string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    app.RefundStatus
		wantErr bool
	}{
		{
			name: "successfully request a refund for a yet to be refunded payment",
			fields: fields{
				// todo: cleaner way to do this?
				queries: sampleDBWithStubbedEntries(ctx, "../../db/files/test1.db"),
			},
			args: args{
				ctx:         ctx,
				userID:      mock.MockStubbedUserID(),
				paymentID:   mock.PaymentThreeID(), // payment 3 has not been refunded, yet
				description: paymentThreeRefundDescription,
			},
			want:    app.RefundStatusPending,
			wantErr: false,
		},
		{
			name: "fail to request a refund for a payment that has already been refunded",
			fields: fields{
				queries: sampleDBWithStubbedEntries(ctx, "../../db/files/test2.db"),
			},
			args: args{
				ctx:         ctx,
				userID:      mock.MockStubbedUserID(),
				paymentID:   mock.PaymentOneID(), // payment 1 has already been refunded
				description: "",
			},
			want:    app.RefundStatusSuccess,
			wantErr: true,
		},
		{
			name: "fail to request a refund for a payment that does not exist",
			fields: fields{
				queries: sampleDBWithStubbedEntries(ctx, "../../db/files/test2.db"),
			},
			args: args{
				ctx:         ctx,
				userID:      mock.MockStubbedUserID(),
				paymentID:   nonExistantPaymentID,
				description: "",
			},
			want:    app.RefundStatusUnknown,
			wantErr: true,
		},
		// {
		// 	name: "fail to request a refund for a payment not assigned to the user",
		// }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRefundService(tt.fields.queries)
			got, err := r.RefundPayment(tt.args.ctx, tt.args.userID, tt.args.paymentID, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefundService.RefundPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefundService.RefundPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sampleDBWithStubbedEntries(ctx context.Context, path string) *sqlc.Queries {
	conn, err := db.NewEmptySQLiteDB(ctx, path)
	if err != nil {
		panic(err)
	}

	queries := sqlc.New(conn)

	err = mock.PopulateSamplePayments(ctx, queries)
	if err != nil {
		panic(err)
	}

	err = mock.PopulateSampleRefunds(ctx, queries)
	if err != nil {
		panic(err)
	}

	return queries
}
