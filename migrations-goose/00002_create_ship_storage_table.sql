-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE ship (
  id bigserial NOT NULL PRIMARY KEY,
  length int NOT NULL,
  count int NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE ship;