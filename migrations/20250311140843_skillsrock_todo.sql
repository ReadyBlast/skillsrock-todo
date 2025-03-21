-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
    created_at TIMESTAMP DEFAULT now (),
    updated_at TIMESTAMP DEFAULT now ()
  );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
DROP TABLE tasks,
-- +goose StatementEnd