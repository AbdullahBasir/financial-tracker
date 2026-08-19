// src/pages/AccountDetailPage.tsx
import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { accountService } from '../services/accounts'
import type { Account } from '../services/accounts'
import { transactionService } from '../services/transactions'
import type { Transaction } from '../services/transactions'
import { categoryService } from '../services/categories'
import type { Category } from '../services/categories'
import { formatDate } from '../lib/format'

export function AccountDetailPage() {
  const { id } = useParams<{ id: string }>()

  const [account, setAccount] = useState<Account | null>(null)
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const [name, setName] = useState('')
  const [type, setType] = useState<'checking' | 'savings' | 'credit'>('credit')
  const [starting_balance, setBalance] = useState<number | ''>('')
  const [isEditing, setIsEditing] = useState(false)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    if (!id) return
    loadData(id)
  }, [id])

  async function loadData(accountId: string) {
    try {
      setLoading(true)
      const [acc, txResult] = await Promise.all([
        accountService.get(accountId),
        transactionService.list({ account_id: accountId }),
      ])
      const categoryResult = await categoryService.list()
      setAccount(acc)
      setName(acc.name)
      setType(acc.type)
      setBalance(Number(acc.starting_balance))
      setTransactions(txResult.transaction)
      setCategories(categoryResult)
    } catch (err) {
      setError((err as Error).message)
    } finally {
      setLoading(false)
    }
  }

  async function handleSaveChanges() {
    if (!id) return
    try {
      const updates: Partial<Account> = {}
      if (name.trim() && name !== account?.name) {
        updates.name = name
      }
      if (type !== account?.type) {
        updates.type = type
      }
      if (starting_balance !== '' && starting_balance !== Number(account?.starting_balance)) {
        updates.starting_balance = starting_balance
      }

      if (Object.keys(updates).length === 0) {
        setIsEditing(false)
        return
      }

      const updated = await accountService.update(id, updates)
      setAccount(updated)
      setIsEditing(false)
    } catch (err) {
      setError((err as Error).message)
    }
  }

  function handleCancel() {
    setIsEditing(false)
    setName(account?.name || '')
    setType(account?.type || 'credit')
    setBalance(Number(account?.starting_balance || 0))
  }

  if (loading) return <p>Loading account...</p>
  if (!account) return <p>Account not found.</p>

  const balance = Number(account.starting_balance)

  return (
    <div>
      {!isEditing && (
        <div className="account-header">
          <div>
            <h1>{account.name}</h1>
          </div>
          <button onClick={() => setIsEditing(true)} className="btn-edit">
            Edit Account
          </button>
        </div>
      )}

      {isEditing && (
        <div className="account-edit-mode">
          <h1>Edit Account</h1>
        </div>
      )}

      {error && <p className="error">{error}</p>}

      <div className="account-details">
        <div className="detail-section">
          <label htmlFor="name-input">Name:</label>
          {!isEditing ? (
            <p className="detail-value">{account.name}</p>
          ) : (
            <input
              id="name-input"
              value={name}
              onChange={e => setName(e.target.value)}
              placeholder="Account name"
            />
          )}
        </div>

        <div className="detail-section">
          <label htmlFor="type-input">Type:</label>
          {!isEditing ? (
            <p className="detail-value">{account.type}</p>
          ) : (
            <select
              id="type-input"
              value={type}
              onChange={e => setType(e.target.value as 'checking' | 'savings' | 'credit')}
            >
              <option value="checking">Checking</option>
              <option value="savings">Savings</option>
              <option value="credit">Credit</option>
            </select>
          )}
        </div>

        <div className="detail-section">
          <label htmlFor="balance-input">Balance:</label>
          {!isEditing ? (
            <p className={`detail-value${balance < 0 ? ' balance-negative' : ''}`}>
              ${balance.toFixed(2)}
            </p>
          ) : (
            <input
              id="balance-input"
              type="number"
              step="0.01"
              value={starting_balance}
              onChange={e => setBalance(e.target.value === '' ? '' : Number(e.target.value))}
              placeholder="0.00"
            />
          )}
        </div>

        {isEditing && (
          <div className="action-buttons">
            <button onClick={handleSaveChanges} className="btn-primary">
              Save Changes
            </button>
            <button onClick={handleCancel} className="btn-cancel">
              Cancel
            </button>
          </div>
        )}
      </div>

      <h2>Transactions</h2>
      {transactions.length === 0 && <p className="empty-state">No transactions for this account.</p>}
      {transactions.length > 0 && (
        <ul>
          {transactions.map(tx => {
            const categoryType = categories.find(category => category.id === tx.category_id)?.type
            return (
            <li key={tx.id}>
              <div>
                <strong>{tx.description}</strong>
                <br />
                <span className="transaction-date">{formatDate(tx.occurred_at)}</span>
              </div>
              <span className={`transaction-amount${categoryType ? ` transaction-amount-${categoryType}` : ''}`}>
                ${Number(tx.amount).toFixed(2)}
              </span>
            </li>
            )
          })}
        </ul>
      )}
    </div>
  )
}


