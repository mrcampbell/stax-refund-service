package app

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrcampbell/stax-refund-service/internal/sqlc"
)

// todo: do not use as DTO, only this example for sake of time
type Payment struct {
	ID            uuid.UUID `json:"id"`
	AmountInCents int       `json:"amount_in_cents"`
	Description   string    `json:"description"`
	Timestamp     time.Time `json:"timestamp"`
	UserID        uuid.UUID `json:"user_id"`
}

type RefundStatus int

const (
	RefundStatusUnknown RefundStatus = iota
	RefundStatusPending
	RefundStatusSuccess
	RefundStatusFailure
)

// todo: do not use as DTO, only this example for sake of time
type Refund struct {
	ID          uuid.UUID    `json:"id"`
	PaymentID   uuid.UUID    `json:"payment_id"`
	UserID      uuid.UUID    `json:"user_id"`
	Status      RefundStatus `json:"status"`
	Description string       `json:"description"`
	RequestedAt time.Time    `json:"requested_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type ServerResources struct {
	Queries       *sqlc.Queries
	AuthService   AuthService
	PaymentClient PaymentClient
	RefundService RefundService
}
