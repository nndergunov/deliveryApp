create table IF NOT EXISTS consumer
(
    id           serial    not null primary key,
    country_code varchar,
    phone_number varchar,
    first_name   varchar   not null,
    last_name    varchar,
    email        varchar   not null,
    created_at   timestamp not null,
    updated_at   timestamp not null,
    status       varchar   not null
);

create table IF NOT EXISTS consumer_location
(
    id           serial not null primary key,
    consumer_id  serial,
    location_alt varchar,
    location_lat varchar,
    country      varchar,
    city         varchar,
    region       varchar,
    street       varchar,
    home_number  varchar,
    floor        varchar,
    door         varchar,
    CONSTRAINT fk_consumer
        FOREIGN KEY (consumer_id)
            REFERENCES consumer (id)
);
