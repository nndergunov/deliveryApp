\c accounting_db

DO
$do$
    BEGIN
        IF EXISTS (
                SELECT FROM pg_catalog.pg_roles
                WHERE  rolname = 'accounting_db') THEN

            RAISE NOTICE 'Role "accounting_db" already exists. Skipping.';
        ELSE
            CREATE ROLE accounting_db LOGIN PASSWORD 'accounting_db_pass';
        END IF;
    END
$do$;

alter user accounting_db with superuser;

grant all privileges on database accounting_db to accounting_db;

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
