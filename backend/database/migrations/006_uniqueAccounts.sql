-- +goose Up
ALTER TABLE accounts
ADD CONSTRAINT unique_user_account_name UNIQUE(user_id, name);

-- +goose Down
ALTER TABLE accounts
DROP CONSTRAINT unique_user_account_name;