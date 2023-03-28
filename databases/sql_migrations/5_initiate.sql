-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE rating (
    id SERIAL NOT NULL,
    name VARCHAR(256),
    rate INT,
    users_id INT,
    game_id INT,
    PRIMARY KEY(id),
    CONSTRAINT fk_users
        FOREIGN KEY(users_id)
            REFERENCES users(id),
    CONSTRAINT fk_game
        FOREIGN KEY(game_id)
            REFERENCES game(id)
);

-- +migrate StatementEnd