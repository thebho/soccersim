-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS teams (
  abv text PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE IF NOT EXISTS team_season (
  id SERIAL PRIMARY KEY,
  season text NOT NULL,
  team_id text references teams(abv),
  games_won int NOT NULL DEFAULT 0,
  games_lost int NOT NULL DEFAULT 0,
  games_drawn int NOT NULL DEFAULT 0,
  goals_scored int NOT NULL DEFAULT 0,
  goals_allowed int NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS matches (
  id SERIAL PRIMARY KEY,
  home_team text NOT NULL,
  away_team text NOT NULL,
  home_team_goals int,
  away_team_goals int,
  match_week int NOT NULL,
  season text NOT NULL,
  played boolean
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE matches;
DROP TABLE team_season;
DROP TABLE teams;
