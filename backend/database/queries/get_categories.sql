-- name: GetCategories :many
SELECT * FROM categories WHERE user_id = $1 AND archived_at IS NULL
ORDER BY created_at DESC;