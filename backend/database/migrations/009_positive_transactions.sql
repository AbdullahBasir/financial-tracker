-- +goose Up
ALTER TABLE transactions
ADD CONSTRAINT check_amount_positive
CHECK (amount >= 0);

-- +goose Down
ALTER TABLE transactions
DROP CONSTRAINT check_amount_positive;