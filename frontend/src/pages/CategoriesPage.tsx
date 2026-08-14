import { useEffect, useState } from 'react'
import type { FormEvent } from 'react'
import { categoryService } from '../services/categories'
import type { Category } from '../services/categories'

export function CategoriesPage() {
  const [categories, setCategories] = useState<Category[]>([])
  const [name, setName] = useState('')
  const [type, setType] = useState<'income' | 'expenses'>('expenses')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    loadCategories()
  }, [])

  async function loadCategories() {
    try {
      setLoading(true)
      const data = await categoryService.list()
      setCategories(data)
    } catch (err) {
      setError((err as Error).message)
    } finally {
      setLoading(false)
    }
  }

  async function handleCreate(e: FormEvent) {
    e.preventDefault()
    if (!name.trim()) return
    try {
      const newCategory = await categoryService.create({ name, type })
      setCategories(prev => [...prev, newCategory])
      setName('')
    } catch (err) {
      setError((err as Error).message)
    }
  }

  return (
    <div>
      <h1>Categories</h1>
      {error && <p className="error">{error}</p>}

      <form onSubmit={handleCreate}>
        <input
          value={name}
          onChange={e => setName(e.target.value)}
          placeholder="Category name"
          required
        />
        <select value={type} onChange={e => setType(e.target.value as 'income' | 'expenses')}>
          <option value="expenses">Expenses</option>
          <option value="income">Income</option>
        </select>
        <button type="submit">Add Category</button>
      </form>

      {loading ? (
        <p>Loading categories...</p>
      ) : (
        <ul>
          {categories.map(c => (
            <li key={c.id}>{c.name} ({c.type})</li>
          ))}
        </ul>
      )}
    </div>
  )
}