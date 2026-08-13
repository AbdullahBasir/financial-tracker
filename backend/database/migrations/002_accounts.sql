-- +goose Up
CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    starting_balance DECIMAL(10, 2) NOT NULL DEFAULT 0,
    type TEXT NOT NULL CHECK (type IN ('checking', 'savings', 'credit')),
    user_id UUID NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE accounts;