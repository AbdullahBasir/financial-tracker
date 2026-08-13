-- name: CreateBudget :one
INSERT INTO budgets (id, monthly_limit, month, user_id, category_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;