-- name: CreateCategory :one
INSERT INTO categories (id, name, type, user_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
ON CONFLICT (user_id, name) DO UPDATE
SET archived_at = NULL
WHERE categories.archived_at IS NOT NULL
    AND categories.type = EXCLUDED.type
RETURNING *;