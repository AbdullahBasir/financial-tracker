-- name: DeleteBudget :exec
DELETE FROM budgets
WHERE id = $1 AND user_id = $2;