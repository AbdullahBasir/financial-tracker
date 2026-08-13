-- name: GetCategories :many
SELECT * FROM categories WHERE user_id = $1 
ORDER BY created_at DESC;