create table orders
(
    id serial not null primary key,
    customer_id integer not null,
    restaurant_id integer not null,
    order_items integer array not null,
    status varchar not null
);
