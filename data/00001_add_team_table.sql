-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS teams (
  abv text PRIMARY KEY,
  name text NOT NULL,
  games_won int NOT NULL DEFAULT 0,
  games_lost int NOT NULL DEFAULT 0,
  games_drawn int NOT NULL DEFAULT 0,
  goals_scored int NOT NULL DEFAULT 0,
  goals_allowed int NOT NULL DEFAULT 0
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE teams;
