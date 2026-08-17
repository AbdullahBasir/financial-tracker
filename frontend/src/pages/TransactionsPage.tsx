// src/pages/TransactionsPage.tsx
import { useEffect, useState } from 'react'
import type { SyntheticEvent } from 'react'
import { transactionService } from '../services/transactions'
import type { Transaction } from '../services/transactions'
import { categoryService } from '../services/categories'
import type { Category } from '../services/categories'
import { accountService } from '../services/accounts'
import type { Account } from '../services/accounts'
import { formatDate } from '../lib/format'

interface TransactionFilters {
  account_id: string
  category_id: string
  from: string
  to: string
}

interface TransactionFormState {
  accountId: string
  categoryId: string
  amount: string
  description: string
  occurredAt: string
}

const EMPTY_FORM: TransactionFormState = {
  accountId: '',
  categoryId: '',
  amount: '',
  description: '',
  occurredAt: '',
}

function toCleanFilters(filters: TransactionFilters) {
  return Object.fromEntries(Object.entries(filters).filter(([, v]) => v !== ''))
}

export function TransactionsPage() {
  // Data
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const [accounts, setAccounts] = useState<Account[]>([])

  // Page state
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  // Filters
  const [filters, setFilters] = useState<TransactionFilters>({
    account_id: '',
    category_id: '',
    from: '',
    to: '',
  })

  // Add-transaction form
  const [form, setForm] = useState<TransactionFormState>(EMPTY_FORM)
  const [formError, setFormError] = useState('')

  useEffect(() => {
    loadStaticData()
  }, [])

  useEffect(() => {
    let cancelled = false

    async function load() {
      try {
        setLoading(true)
        const data = await transactionService.list(toCleanFilters(filters))
        if (!cancelled) setTransactions(data)
      } catch (err) {
        if (!cancelled) setError((err as Error).message)
      } finally {
        if (!cancelled) setLoading(false)
      }
    }

    load()
    return () => {
      cancelled = true
    }
  }, [filters])

  async function loadStaticData() {
    try {
      const [cats, accs] = await Promise.all([categoryService.list(), accountService.list()])
      setCategories(cats)
      setAccounts(accs)
    } catch (err) {
      setError((err as Error).message)
    }
  }

  async function handleCreate(e: SyntheticEvent<HTMLFormElement>) {
    e.preventDefault()
    setFormError('')

    const amount = parseFloat(form.amount)
    if (isNaN(amount) || amount < 0) {
      setFormError('Amount cannot be negative or invalid.')
      return
    }

    try {
      const newTx = await transactionService.create({
        account_id: form.accountId,
        category_id: form.categoryId,
        amount,
        description: form.description,
        occurred_at: new Date(form.occurredAt).toISOString(),
      })
      setTransactions(prev => [newTx, ...prev])
      setForm(EMPTY_FORM)
    } catch (err) {
      setError((err as Error).message)
    }
  }

  async function handleDelete(id: string) {
    try {
      await transactionService.remove(id)
      setTransactions(prev => prev.filter(t => t.id !== id))
    } catch (err) {
      setError((err as Error).message)
    }
  }

  function updateFormField<K extends keyof TransactionFormState>(field: K, value: TransactionFormState[K]) {
    setForm(f => ({ ...f, [field]: value }))
    if (formError) setFormError('')
  }

  return (
    <div>
      <h1>Transactions</h1>

      {error && <p className="error">{error}</p>}
      {formError && <p className="error form-error">{formError}</p>}

      <section>
        <h2>Filters</h2>
        <select
          value={filters.account_id}
          onChange={e => setFilters(f => ({ ...f, account_id: e.target.value }))}
        >
          <option value="">Accounts</option>
          {accounts.map(a => (
            <option key={a.id} value={a.id}>{a.name}</option>
          ))}
        </select>

        <select
          value={filters.category_id}
          onChange={e => setFilters(f => ({ ...f, category_id: e.target.value }))}
        >
          <option value="">Categories</option>
          {categories.map(c => (
            <option key={c.id} value={c.id}>{c.name}</option>
          ))}
        </select>

        <input
          type="date"
          value={filters.from}
          onChange={e => setFilters(f => ({ ...f, from: e.target.value }))}
        />
        <input
          type="date"
          value={filters.to}
          onChange={e => setFilters(f => ({ ...f, to: e.target.value }))}
        />
      </section>

      <section>
        <h2>Add Transaction</h2>
        <form onSubmit={handleCreate}>
          <select
            value={form.accountId}
            onChange={e => updateFormField('accountId', e.target.value)}
            required
          >
            <option value="" disabled>Account</option>
            {accounts.map(a => (
              <option key={a.id} value={a.id}>{a.name}</option>
            ))}
          </select>

          <select
            value={form.categoryId}
            onChange={e => updateFormField('categoryId', e.target.value)}
            required
          >
            <option value="" disabled>Category</option>
            {categories.map(c => (
              <option key={c.id} value={c.id}>{c.name}</option>
            ))}
          </select>

          <div>
            <input
              type="number"
              step="0.01"
              min="0.01"
              placeholder="Amount"
              value={form.amount}
              onChange={e => updateFormField('amount', e.target.value)}
              required
            />
          </div>

          <input
            type="text"
            placeholder="Description"
            value={form.description}
            onChange={e => updateFormField('description', e.target.value)}
          />

          <input
            type="date"
            value={form.occurredAt}
            onChange={e => updateFormField('occurredAt', e.target.value)}
            required
          />

          <button type="submit">Add Transaction</button>
        </form>
      </section>

      {loading && <p>Loading transactions...</p>}

      {!loading && transactions.length === 0 && (
        <p className="empty-state">No transactions yet — add one above to get started.</p>
      )}

      {!loading && transactions.length > 0 && (
        <table>
          <thead>
            <tr>
              <th>Date</th>
              <th>Description</th>
              <th>Amount</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
          {transactions.map(tx => {
            const amount = Number(tx.amount)
            return (
              <tr key={tx.id}>
                <td>{formatDate(tx.occurred_at)}</td>
                <td>{tx.description}</td>
                <td>${amount.toFixed(2)}</td>
                <td><button onClick={() => handleDelete(tx.id)}>Delete</button></td>
              </tr>
            )
          })}
        </tbody>
        </table>
      )}
    </div>
  )
}