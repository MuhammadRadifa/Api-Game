-- +migrate Up
-- +migrate StatementBegin
CREATE TYPE roleEnum AS ENUM ('user', 'admin');

CREATE TABLE users (
    id VARCHAR(256) NOT NULL,
    name VARCHAR(256),
    email VARCHAR(256) UNIQUE,
    password VARCHAR(256),
    role roleEnum DEFAULT 'user',
    PRIMARY KEY(id)
);

-- +migrate StatementEnd