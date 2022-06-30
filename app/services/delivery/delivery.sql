\c delivery_db

DO
$do$
    BEGIN
        IF EXISTS (
                SELECT FROM pg_catalog.pg_roles
                WHERE  rolname = 'delivery_db') THEN

            RAISE NOTICE 'Role "delivery_db" already exists. Skipping.';
        ELSE
            CREATE ROLE delivery_db LOGIN PASSWORD 'delivery_db_pass';
        END IF;
    END
$do$;

alter user delivery_db with superuser;

grant all privileges on database delivery_db to delivery_db;

create table IF NOT EXISTS delivery
(
    order_id   int,
    courier_id int
);