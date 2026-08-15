-- +goose Up
ALTER TABLE transactions
    ALTER COLUMN category_id DROP NOT NULL;

ALTER TABLE transactions
    DROP CONSTRAINT IF EXISTS transactions_category_id_fkey,
    ADD CONSTRAINT transactions_category_id_fkey
        FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL;

ALTER TABLE budgets
    DROP CONSTRAINT IF EXISTS budgets_category_id_fkey,
    ADD CONSTRAINT budgets_category_id_fkey
        FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE transactions
    DROP CONSTRAINT IF EXISTS transactions_category_id_fkey,
    ADD CONSTRAINT transactions_category_id_fkey
        FOREIGN KEY (category_id) REFERENCES categories(id);

ALTER TABLE budgets
    DROP CONSTRAINT IF EXISTS budgets_category_id_fkey,
    ADD CONSTRAINT budgets_category_id_fkey
        FOREIGN KEY (category_id) REFERENCES categories(id);

