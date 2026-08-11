-- name: GetBudgetsForMonth :many
SELECT 
    b.category_id,
    c.name AS category_name,
    b.monthly_limit
FROM budgets b
JOIN categories c ON c.id = b.category_id
WHERE b.user_id = $1
AND TO_CHAR(b.month, 'YYYY-MM') = $2;