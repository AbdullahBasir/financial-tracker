import { useEffect, useState } from 'react'
import type { SyntheticEvent } from 'react'
import { categoryService } from '../services/categories'
import type { Category } from '../services/categories'
import { ApiError } from '../lib/api'

export function CategoriesPage() {
  const [categories, setCategories] = useState<Category[]>([])
  const [name, setName] = useState('')
  const [type, setType] = useState<'income' | 'expenses'>('expenses')
  const [editingId, setEditingId] = useState<string | null>(null)
  const [editName, setEditName] = useState('')
  const [editType, setEditType] = useState<'income' | 'expenses'>('expenses')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

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

  useEffect(() => {
    async function load() {
      await loadCategories()
    }
    load()
  }, [])

  async function handleCreate(e: SyntheticEvent<HTMLFormElement>) {
    e.preventDefault()
    if (!name.trim()) return
    try {
      const newCategory = await categoryService.create({ name, type })
      setCategories(prev => [...prev, newCategory])
      setName('')
    } catch (err) {
      if (err instanceof ApiError && err.status === 409) {
        setError('A category with this name already exists. Category names must use the same type.')
      } else {
        setError((err as Error).message)
      }
    }
  }

  function handleEdit(category: Category) {
    setEditingId(category.id)
    setEditName(category.name)
    setEditType(category.type)
    setError('')
  }

  function handleCancelEdit() {
    setEditingId(null)
    setEditName('')
  }

  async function handleSaveEdit(id: string) {
    if (!editName.trim()) return
    try {
      const updatedCategory = await categoryService.update(id, {
        name: editName.trim(),
        type: editType,
      })
      setCategories(prev => prev.map(category => (
        category.id === id ? updatedCategory : category
      )))
      handleCancelEdit()
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
            <li key={c.id}>
              {editingId === c.id ? (
                <>
                  <input
                    value={editName}
                    onChange={e => setEditName(e.target.value)}
                    aria-label={`Name for ${c.name}`}
                  />
                  <select
                    value={editType}
                    onChange={e => setEditType(e.target.value as 'income' | 'expenses')}
                    aria-label={`Type for ${c.name}`}
                  >
                    <option value="expenses">Expenses</option>
                    <option value="income">Income</option>
                  </select>
                  <button type="button" onClick={() => handleSaveEdit(c.id)} disabled={!editName.trim()}>
                    Save
                  </button>
                  <button className="btn-cancel" onClick={handleCancelEdit}>
                    Cancel
                  </button>
                </>
              ) : (
                <>
                  {c.name} ({c.type})
                  <button type="button" onClick={() => handleEdit(c)}>
                    Edit
                  </button>
                </>
              )}
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}