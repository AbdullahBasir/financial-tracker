-- name: GetTransactions :many
SELECT * FROM transactions 
WHERE account_id = $1
ORDER BY occurred_at DESC
LIMIT $2 OFFSET $3;