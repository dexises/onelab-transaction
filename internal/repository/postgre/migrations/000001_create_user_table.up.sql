CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    name varchar unique,
    balance int not null default 0,
);