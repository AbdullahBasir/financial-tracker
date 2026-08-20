-- name: CreateCategory :one
INSERT INTO categories (id, name, type, user_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;