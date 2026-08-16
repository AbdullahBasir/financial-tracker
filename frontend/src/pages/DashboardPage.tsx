import { useEffect, useState } from 'react'
import { budgetService } from '../services/budgets'
import type { BudgetSummaryItem } from '../services/budgets'

function currentMonth() {
  const now = new Date()
  return `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}`
}

export function DashboardPage() {
  const [month, setMonth] = useState(currentMonth())
  const [summary, setSummary] = useState<BudgetSummaryItem[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    loadSummary(month)
  }, [month])

  async function loadSummary(m: string) {
    try {
      setLoading(true)
      const data = await budgetService.summary(m)
      console.log('summary data:', data)
      setSummary(data)
    } catch (err) {
      setError((err as Error).message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <h1>Dashboard</h1>

      <label htmlFor="month">Month</label>
      <input
        id="month"
        type="month"
        value={month}
        onChange={e => setMonth(e.target.value)}
      />

      {error && <p className="error">{error}</p>}
      {loading && <p>Loading summary...</p>}

      {!loading && summary.length === 0 && <p>No budgets set for this month.</p>}

      {!loading && summary.length > 0 && (
        <table>
          <thead>
            <tr>
              <th>Category</th>
              <th>Budget</th>
              <th>Spent</th>
              <th>Remaining</th>
            </tr>
          </thead>
          <tbody>
            {summary.map(item => (
              <tr key={item.category_id} className={item.is_over_budget ? 'balance-negative' : ''}>
                <td>{item.category_name}</td>
                <td>${Number(item.monthly_limit).toFixed(2)}</td>
                <td>${Number(item.total_spent).toFixed(2)}</td>
                <td>${Number(item.remaining).toFixed(2)}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  )
}