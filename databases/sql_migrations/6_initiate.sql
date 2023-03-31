-- +migrate Up
-- +migrate StatementBegin

ALTER TABLE rating DROP COLUMN name;

-- +migrate StatementEnd