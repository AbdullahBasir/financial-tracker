-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;