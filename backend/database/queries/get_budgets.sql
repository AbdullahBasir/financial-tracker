-- name: GetBudgets :many
SELECT * FROM budgets
WHERE user_id = $1
ORDER BY month DESC;