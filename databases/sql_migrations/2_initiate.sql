-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id SERIAL NOT NULL,
    name VARCHAR(256),
    PRIMARY KEY(id)
);

-- +migrate StatementEnd