-- name: GetTotalIncomeAmount :one
SELECT 
    COALESCE(SUM(t.amount), 0)::numeric AS total_expenses
FROM transactions t
JOIN accounts a ON a.id = t.account_id
JOIN categories c ON c.id = t.category_id
WHERE a.user_id = $1 
AND a.id = $2
AND c.type = 'income';