-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO teams (abv, name, season) VALUES
  ('ARS', 'Arsenal', 'DEFAULT'),
  ('TOT', 'Tottenham Hotspur', 'DEFAULT'),
  ('MAN', 'Manchester United', 'DEFAULT'),
  ('MNC', 'Manchester City', 'DEFAULT'),
  ('LEI', 'Leicester City', 'DEFAULT'),
  ('EVE', 'Everton', 'DEFAULT'),
  ('CHE', 'Chelsea', 'DEFAULT'),
  ('STK', 'Stoke City', 'DEFAULT'),
  ('LIV', 'Liverpool', 'DEFAULT'),
  ('BOU', 'AFC Bournemouth', 'DEFAULT'),
  ('BUR', 'Burnley', 'DEFAULT'),
  ('WAT', 'Watford', 'DEFAULT'),
  ('WHU', 'West Ham United', 'DEFAULT'),
  ('BHA', 'Brighton & Hove Albion', 'DEFAULT'),
  ('HUD', 'Huddersfield Town', 'DEFAULT'),
  ('NEW', 'Newcastle United', 'DEFAULT'),
  ('SOU', 'Southampton', 'DEFAULT'),
  ('SWA', 'Swansea City', 'DEFAULT'),
  ('WBA', 'West Bromwich Albion', 'DEFAULT'),
  ('CRY', 'Crystal Palace', 'DEFAULT');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
