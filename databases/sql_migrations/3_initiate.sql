-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE game (
    id SERIAL NOT NULL,
    name VARCHAR(256),
    description TEXT,
    category_id INT,
    PRIMARY KEY(id),
    CONSTRAINT fk_category
        FOREIGN KEY(category_id)
            REFERENCES category(id)
);

-- +migrate StatementEnd