\c restaurant_db
DO
$do$
    BEGIN
        IF EXISTS (
                SELECT FROM pg_catalog.pg_roles
                WHERE  rolname = 'restaurant_db') THEN

            RAISE NOTICE 'Role "restaurant_db" already exists. Skipping.';
        ELSE
            CREATE ROLE restaurant_db LOGIN PASSWORD 'restaurant_db_pass';
        END IF;
    END
$do$;

alter user restaurant_db with superuser;

grant all privileges on database restaurant_db to restaurant_db;

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
