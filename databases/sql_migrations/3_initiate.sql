-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE game (
    id VARCHAR(256) NOT NULL,
    name VARCHAR(256),
    description TEXT,
    category_id VARCHAR(256),
    PRIMARY KEY(id),
    CONSTRAINT fk_category
        FOREIGN KEY(category_id)
            REFERENCES category(id)
);

-- +migrate StatementEnd