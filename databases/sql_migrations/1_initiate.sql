-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id SERIAL NOT NULL,
    name VARCHAR(256),
    email VARCHAR(256) UNIQUE,
    password VARCHAR(256),
    PRIMARY KEY(id)
);

-- +migrate StatementEnd