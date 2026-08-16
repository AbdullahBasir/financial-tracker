// src/pages/AccountDetailPage.tsx
import { useEffect, useState } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { accountService } from '../services/accounts'
import type { Account } from '../services/accounts'
import { transactionService } from '../services/transactions'
import type { Transaction } from '../services/transactions'

export function AccountDetailPage() {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()

  const [account, setAccount] = useState<Account | null>(null)
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [name, setName] = useState('')
  const [starting_balance, setBalance] = useState<number | ''>('')
  const [editingBalance, setEditingBalance] = useState(false)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    if (!id) return
    loadData(id)
  }, [id])

  async function loadData(accountId: string) {
    try {
      setLoading(true)
      const [acc, txs] = await Promise.all([
        accountService.get(accountId),
        transactionService.list({ account_id: accountId }),
      ])
      setAccount(acc)
      setName(acc.name)
      setBalance(Number(acc.starting_balance))
      setTransactions(txs)
    } catch (err) {
      setError((err as Error).message)
    } finally {
      setLoading(false)
    }
  }

  async function handleRename() {
    if (!id || !name.trim()) return
    try {
      const updated = await accountService.update(id, { name })
      setAccount(updated)
    } catch (err) {
      setError((err as Error).message)
    }
  }

   async function handleUpdateBalance() {
    if (!id || starting_balance === '') return
    try {
      const updated = await accountService.update(id, { starting_balance })
      setAccount(updated)
      setEditingBalance(false)
    } catch (err) {
      setError((err as Error).message)
    }
  }

  function handleCancelBalance() {
    setEditingBalance(false)
    setBalance(Number(account?.starting_balance || 0))
  }

  async function handleDelete() {
    if (!id) return
    try {
      await accountService.remove(id)
      navigate('/accounts')
    } catch (err) {
      setError((err as Error).message)
    }
  }

  if (loading) return <p>Loading account...</p>
  if (!account) return <p>Account not found.</p>

  return (
    <div>
      <h1>{account.name}</h1>
      <p>Type: {account.type}</p>
      
      <div className="balance-section">
        {!editingBalance ? (
          <div className="balance-display">
            <p>
              Starting balance: <strong>${Number(account.starting_balance).toFixed(2)}</strong>
            </p>
            <button onClick={() => setEditingBalance(true)} className="btn-edit">
              Edit Balance
            </button>
          </div>
        ) : (
          <div className="balance-edit">
            <label htmlFor="balance-input">New Starting Balance:</label>
            <input
              id="balance-input"
              type="number"
              step="0.01"
              value={starting_balance}
              onChange={e => setBalance(e.target.value === '' ? '' : Number(e.target.value))}
              placeholder="0.00"
            />
            <button onClick={handleUpdateBalance}>Save Balance</button>
            <button onClick={handleCancelBalance} className="btn-cancel">
              Cancel
            </button>
          </div>
        )}
      </div>

      {error && <p className="error">{error}</p>}

      <div className="account-actions">
        <div className="name-section">
          <label htmlFor="name-input">Account Name:</label>
          <input
            id="name-input"
            value={name}
            onChange={e => setName(e.target.value)}
            placeholder="Account name"
          />
          <button onClick={handleRename}>Save Name</button>
        </div>
        <button onClick={handleDelete} className="btn-delete">
          Delete Account
        </button>
      </div>

      <h2>Transactions</h2>
      {transactions.length === 0 && <p className="empty-state">No transactions for this account.</p>}
      {transactions.length > 0 && (
        <ul>
          {transactions.map(tx => (
            <li key={tx.id}>
              <div>
                <strong>{tx.description}</strong>
                <br />
                <span className="transaction-date">{tx.occurred_at}</span>
              </div>
              <span className="transaction-amount">${Number(tx.amount).toFixed(2)}</span>
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}