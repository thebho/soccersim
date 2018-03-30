-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE matches ADD COLUMN played boolean;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE matches DROP COLUMN played;
