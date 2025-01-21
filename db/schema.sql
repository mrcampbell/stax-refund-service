CREATE TABLE payments (
    id text PRIMARY KEY,
    user_id text NOT NULL, -- todo: move to a separate table, payment_users, etc. First we'd need an actual users table.
    amount_in_cents integer NOT NULL,
    "description" text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE refunds (
    id text PRIMARY KEY,
    payment_id text NOT NULL,
    user_id text NOT NULL, -- todo: move to a separate table, refund_users, etc. First we'd need an actual users table.
    "description" text NOT NULL,
    requested_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "status" integer NOT NULL DEFAULT 0,
    FOREIGN KEY (payment_id) REFERENCES payments(id) -- todo: add a bridge table, rather than reference directly
);