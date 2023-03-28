-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id VARCHAR(256) NOT NULL,
    name VARCHAR(256),
    PRIMARY KEY(id)
);

-- +migrate StatementEnd