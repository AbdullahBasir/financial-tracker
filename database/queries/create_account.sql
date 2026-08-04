-- name: CreateAccount :one
INSERT INTO accounts (id, name, starting_balance, type, user_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;