-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE fight_act (
  user_id int NOT NULL,
  x_coord int NOT NULL,
  y_coord int NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE fight_act;