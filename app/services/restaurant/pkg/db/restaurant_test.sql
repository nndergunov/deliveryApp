create table restaurants
(
    id serial not null primary key,
    name varchar not null,
    accepting_orders boolean not null,
    city varchar not null,
    address varchar not null,
    longitude float not null,
    latitude float not null
);

create table menus
(
    id serial not null primary key,
    restaurant_id integer not null unique references restaurants (id)
);

create table menu_items
(
    id serial not null primary key,
    menu_id integer not null references menus (id),
    name varchar not null,
    price float not null,
    course varchar not null
);
