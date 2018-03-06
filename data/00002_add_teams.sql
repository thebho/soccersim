-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO teams (abv, name) VALUES
  ('ARS', 'Arsenal'),
  ('TOT', 'Tottenham Hotspur'),
  ('MAN', 'Manchester United'),
  ('MNC', 'Manchester City'),
  ('LEI', 'Leicester City'),
  ('EVE', 'Everton'),
  ('CHE', 'Chelsea'),
  ('STK', 'Stoke City'),
  ('LIV', 'Liverpool'),
  ('BOU', 'AFC Bournemouth'),
  ('BUR', 'Burnley'),
  ('WAT', 'Watford'),
  ('WHU', 'West Ham United'),
  ('BHA', 'Brighton & Hove Albion'),
  ('HUD', 'Huddersfield Town'),
  ('NEW', 'Newcastle United'),
  ('SOU', 'Southampton'),
  ('SWA', 'Swansea City'),
  ('WBA', 'West Bromwich Albion'),
  ('CRY', 'Crystal Palace');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
