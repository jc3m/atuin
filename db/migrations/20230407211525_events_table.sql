-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE SCHEMA tubul;
CREATE TABLE IF NOT EXISTS tubul.events (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  title text NOT NULL CONSTRAINT namechk CHECK (char_length(title) <= 255 AND char_length(title) > 3),
  event_date date NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS tubul.events;
DROP SCHEMA IF EXISTS tubul;
-- +goose StatementEnd
