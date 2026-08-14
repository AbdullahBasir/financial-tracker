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
      <p>Starting balance: ${account.starting_balance.toFixed(2)}</p>

      {error && <p className="error">{error}</p>}

      <div>
        <input value={name} onChange={e => setName(e.target.value)} />
        <button onClick={handleRename}>Save Name</button>
        <button onClick={handleDelete}>Delete Account</button>
      </div>

      <h2>Transactions</h2>
      {transactions.length === 0 && <p>No transactions for this account.</p>}
      <ul>
        {transactions.map(tx => (
          <li key={tx.id}>
            {tx.occurred_at} — {tx.description} — ${tx.amount.toFixed(2)}
          </li>
        ))}
      </ul>
    </div>
  )
}