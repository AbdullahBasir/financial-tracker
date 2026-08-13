-- name: UpdateTransaction :one
UPDATE transactions
SET amount = $1, occurred_at = $2, description = $3, account_id = $4, category_id = $5
WHERE id = $6
RETURNING *;