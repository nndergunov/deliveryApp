\c consumer_db

DO
$do$
    BEGIN
        IF EXISTS (
                SELECT FROM pg_catalog.pg_roles
                WHERE  rolname = 'consumer_db') THEN

            RAISE NOTICE 'Role "consumer_db" already exists. Skipping.';
        ELSE
            CREATE ROLE consumer_db LOGIN PASSWORD 'consumer_db_pass';
        END IF;
    END
$do$;

alter user consumer_db with superuser;

grant all privileges on database consumer_db to consumer_db;

create table IF NOT EXISTS consumer
(
    id         serial primary key,
    firstname  varchar        NOT NULL,
    lastname   varchar        NOT NULL DEFAULT '',
    email      varchar unique NOT NULL DEFAULT '',
    phone      varchar unique NOT NULL DEFAULT '',
    created_at timestamp      NOT NULL,
    updated_at timestamp      NOT NULL
);

create table IF NOT EXISTS location
(
    user_id     serial  NOT NULL,
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