-- name: UpdateCategory :one
UPDATE categories
SET name = $1, type = $2
WHERE id = $3 AND user_id = $4
RETURNING *;