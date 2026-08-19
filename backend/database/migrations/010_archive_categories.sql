-- +goose Up
ALTER TABLE categories
    ADD COLUMN archived_at TIMESTAMPTZ;

-- +goose Down
ALTER TABLE categories
    DROP COLUMN archived_at;