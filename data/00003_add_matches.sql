-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS matches (
  id SERIAL PRIMARY KEY,
  home_team text NOT NULL,
  away_team text NOT NULL,
  home_team_goals int,
  away_team_goals int,
  match_week int NOT NULL,
  season text NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE matches;
