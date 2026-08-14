import { useEffect, useState } from 'react'
import type { SyntheticEvent } from 'react'
import { budgetService } from '../services/budgets'
import type { Budget } from '../services/budgets'
import { categoryService } from '../services/categories'
import type { Category } from '../services/categories'

function currentMonth() {
  const now = new Date()
  return `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}`
}

export function BudgetsPage() {
  const [budgets, setBudgets] = useState<Budget[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  const [form, setForm] = useState({
    categoryId: '',
    monthlyLimit: '',
    month: currentMonth(),
  })

  useEffect(() => {
    loadData()
  }, [])

  async function loadData() {
    try {
      setLoading(true)
      const [budgetData, categoryData] = await Promise.all([
        budgetService.list(),
        categoryService.list(),
      ])
      setBudgets(budgetData)
      setCategories(categoryData)
    } catch (err) {
      setError((err as Error).message)
    } finally {
      setLoading(false)
    }
  }

  async function handleCreate(e: SyntheticEvent<HTMLFormElement>) {
    e.preventDefault()
    try {
      const newBudget = await budgetService.create({
        category_id: form.categoryId,
        monthly_limit: parseFloat(form.monthlyLimit),
        month: form.month,
      })
      setBudgets(prev => [...prev, newBudget])
      setForm({ categoryId: '', monthlyLimit: '', month: currentMonth() })
    } catch (err) {
      setError((err as Error).message)
    }
  }

  function categoryName(id: string) {
    return categories.find(c => c.id === id)?.name ?? 'Unknown'
  }

  return (
    <div>
      <h1>Budgets</h1>
      {error && <p className="error">{error}</p>}

      <form onSubmit={handleCreate}>
        <select
          value={form.categoryId}
          onChange={e => setForm(f => ({ ...f, categoryId: e.target.value }))}
          required
        >
          <option value="" disabled>Category</option>
          {categories.map(c => (
            <option key={c.id} value={c.id}>{c.name}</option>
          ))}
        </select>

        <input
          type="number"
          step="0.01"
          placeholder="Monthly limit"
          value={form.monthlyLimit}
          onChange={e => setForm(f => ({ ...f, monthlyLimit: e.target.value }))}
          required
        />

        <input
          type="month"
          value={form.month}
          onChange={e => setForm(f => ({ ...f, month: e.target.value }))}
          required
        />

        <button type="submit">Add Budget</button>
      </form>

      {loading ? (
        <p>Loading budgets...</p>
      ) : (
        <table>
          <thead>
            <tr>
              <th>Category</th>
              <th>Month</th>
              <th>Limit</th>
            </tr>
          </thead>
          <tbody>
            {budgets.map(b => (
              <tr key={b.id}>
                <td>{categoryName(b.category_id)}</td>
                <td>{b.month}</td>
                <td>${Number(b.monthly_limit).toFixed(2)}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  )
}