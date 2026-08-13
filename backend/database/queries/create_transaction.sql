-- name: CreateTransaction :one
INSERT INTO transactions (id, amount, occurred_at, description, account_id, category_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;