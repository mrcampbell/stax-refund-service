-- name: CreatePayment :one
insert into
    payments (
        id,
        user_id,
        amount_in_cents,
        "description",
        created_at
    )
values
    (?, ?, ?, ?, ?) returning id;

-- name: ListPayments :many
select
    *
from
    payments
where 
    user_id = ?;

-- name: GetPaymentByID :one
select
    *
from
    payments
where
    id = ?
    and user_id = ?;

-- name: GetRefundByID :one
select
    *
from
    refunds
where
    id = ?
    and payment_id = ?
    and user_id = ?;

-- name: ListRefunds :many
select
    *
from
    refunds
where
    user_id = ?;

-- name: GetRefundByPaymentID :one
select
    *
from
    refunds
where
    payment_id = ?
    and user_id = ?;

-- name: CreateRefund :one
insert into
    refunds (
        id,
        payment_id,
        user_id,
        "description",
        requested_at,
        updated_at,
        "status"
    )
values
    (?, ?, ?, ?, ?, ?, ?) returning id;

-- name: UpdateRefund :one
update refunds
set
    "description" = ?,
    updated_at = ?,
    "status" = ?,
    user_id = ?
where
    id = ? returning *;