-- name: GetTransactions :many
SELECT t.*
FROM transactions t
JOIN accounts a ON t.account_id = a.id
WHERE a.user_id = $1
  AND (sqlc.narg('account_id')::uuid IS NULL OR t.account_id = sqlc.narg('account_id'))
  AND (sqlc.narg('category_id')::uuid IS NULL OR t.category_id = sqlc.narg('category_id'))
  AND (sqlc.narg('from_date')::timestamptz IS NULL OR t.occurred_at >= sqlc.narg('from_date'))
  AND (sqlc.narg('to_date')::timestamptz IS NULL OR t.occurred_at <= sqlc.narg('to_date'))
ORDER BY t.occurred_at DESC
LIMIT $2 OFFSET $3;