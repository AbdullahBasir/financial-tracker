-- name: GetTransactions :many
SELECT t.* FROM transactions t
JOIN accounts a ON t.account_id = a.id
WHERE a.user_id = $1
  AND ($2::uuid IS NULL OR t.account_id = $2)
ORDER BY t.occurred_at DESC
LIMIT $3 OFFSET $4;