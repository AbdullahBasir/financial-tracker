-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1;
