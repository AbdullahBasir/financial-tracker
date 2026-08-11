-- name: GetMonthlySpending :many
SELECT 
    t.category_id,
    SUM(t.amount)::numeric AS total_spent
FROM transactions t
JOIN accounts a ON a.id = t.account_id
WHERE a.user_id = $1
AND TO_CHAR(t.occurred_at, 'YYYY-MM') = $2::text
GROUP BY t.category_id;