\c order_db

DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'order_db') THEN

      RAISE NOTICE 'Role "order_db" already exists. Skipping.';
ELSE
CREATE ROLE order_db LOGIN PASSWORD 'order_db_pass';
END IF;
END
$do$;

alter user order_db with superuser;

grant all privileges on database order_db to order_db;

create table orders
(
    id serial not null primary key,
    customer_id integer not null,
    restaurant_id integer not null,
    order_items integer array not null,
    status varchar not null
);
