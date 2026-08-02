-- +goose Up
CREATE TABLE budgets (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    monthly_limit DECIMAL(10, 2) NOT NULL,
    month VARCHAR(7) NOT NULL,
    user_id UUID NOT NULL,
    category_id UUID NOT NULL,
    FOREIGN KEY (user_id) 
    REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id),
    UNIQUE (user_id, category_id, month)
);

-- +goose Down
DROP TABLE budgets;