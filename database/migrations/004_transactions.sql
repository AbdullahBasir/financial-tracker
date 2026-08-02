-- +goose Up
CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    occurred_at TIMESTAMP NOT NULL,
    description TEXT,
    account_id UUID NOT NULL,
    category_id UUID NOT NULL,
    FOREIGN KEY (account_id)
    REFERENCES accounts(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- +goose Down
DROP TABLE transactions;