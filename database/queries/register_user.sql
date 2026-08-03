-- name: RegisterUser :one
INSERT INTO users (id, email, password_hash)
VALUES (
    gen_random_uuid(),
    $1,
    $2
)
RETURNING *;