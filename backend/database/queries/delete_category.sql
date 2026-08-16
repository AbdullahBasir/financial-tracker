-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1 AND user_id = $2;