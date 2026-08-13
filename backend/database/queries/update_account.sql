-- name: UpdateAccount :one
UPDATE accounts
SET name = $1, starting_balance = $2, type = $3
WHERE id = $4
RETURNING *;