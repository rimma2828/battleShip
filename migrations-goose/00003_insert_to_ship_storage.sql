-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO ship (length,count) VALUES
    (1, 4),
    (2, 3),
    (3, 2),
    (4, 1);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DELETE FROM ship;