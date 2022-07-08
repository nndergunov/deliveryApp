-- Version: 1.1
-- Description: Create table accounts,transactions


create table IF NOT EXISTS courier
(
    id         serial    not null primary key,
    username   varchar   NOT NULL UNIQUE,
    password   varchar   NOT NULL,
    firstname  varchar   NOT NULL DEFAULT '',
    lastname   varchar   NOT NULL DEFAULT '',
    email      varchar   NOT NULL DEFAULT '',
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    phone      varchar   NOT NULL DEFAULT '',
    available  boolean   NOT NULL DEFAULT TRUE
);

create table IF NOT EXISTS location
(
    user_id     int     NOT NULL,
    latitude    varchar NOT NULL DEFAULT '',
    longitude   varchar NOT NULL DEFAULT '',
    country     varchar NOT NULL DEFAULT '',
    city        varchar NOT NULL DEFAULT '',
    region      varchar NOT NULL DEFAULT '',
    street      varchar NOT NULL DEFAULT '',
    home_number varchar NOT NULL DEFAULT '',
    floor       varchar NOT NULL DEFAULT '',
    door        varchar NOT NULL DEFAULT ''
);
