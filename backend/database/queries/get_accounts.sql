-- name: GetAccounts :many
SELECT * FROM accounts WHERE user_id = $1 
ORDER BY created_at DESC;