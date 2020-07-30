-- +goose Up
-- SQL in this section is executed when the migration is applied.

ALTER TABLE fight_act ADD COLUMN ship_id INT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

ALTER TABLE fight_act DROP COLUMN ship_id;