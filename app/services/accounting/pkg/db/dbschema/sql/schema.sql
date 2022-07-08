-- Version: 1.1
-- Description: Create table accounts,transactions

create table IF NOT EXISTS account
(
    id         serial PRIMARY KEY,
    user_id    int       NOT NULL,
    user_type  varchar,
    balance    numeric   NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

create table IF NOT EXISTS transactions
(
    id              serial primary key,
    from_account_id int       NOT NULL DEFAULT 0,
    to_account_id   int       NOT NULL DEFAULT 0,
    amount          numeric   NOT NULL,
    created_at      timestamp NOT NULL,
    updated_at      timestamp NOT NULL,
    valid           bool      NOT NULL DEFAULT false
);
