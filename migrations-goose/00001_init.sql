-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE users (
  id bigserial NOT NULL PRIMARY KEY,
  username text NOT NULL UNIQUE
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE users;