\c order_db

-- If your docker compose fails, try to comment/uncomment the next line.
-- create user order_db with encrypted password 'order_db_pass';

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
