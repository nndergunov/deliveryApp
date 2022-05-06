create table IF NOT EXISTS courier
(
    id        serial    not null primary key,
    username  varchar   not null unique ,
    password  varchar   not null,
    firstName varchar   not null,
    lastName  varchar,
    email     varchar   not null,
    createdat timestamp not null,
    updatedat timestamp not null,
    phone     varchar,
    status    varchar   not null,
    available boolean   not null
);
