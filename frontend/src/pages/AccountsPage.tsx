import { useEffect, useState, type SyntheticEvent } from 'react'
import { Link } from 'react-router-dom'
import { accountService, type Account } from '../services/accounts'

export function AccountsPage() {
  const [accounts, setAccounts] = useState<Account[]>([])
  const [name, setName] = useState('')
  const [type, setType] = useState<'checking' | 'savings' | 'credit'>('credit')
  const [starting_balance, setBalance] = useState<number | ''>('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  async function loadAccounts() {
    try {
      setLoading(true)
      const data = await accountService.list()
      setAccounts(data)
    } catch (err) {
      setError((err as Error).message)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    async function load() {
      await loadAccounts()
    }
    load()
  }, [])

  async function handleCreate(e: SyntheticEvent<HTMLFormElement>) {
  e.preventDefault()
  if (!name.trim()) return
  try {
    const newAccount = await accountService.create({
      name,
      type,
      starting_balance: starting_balance === '' ? 0 : starting_balance,
    })
    setAccounts(prev => [...prev, newAccount])
    setName('')
    setType('credit')
    setBalance('')
  } catch (err) {
    setError((err as Error).message)
  }
}

  async function handleDelete(id: string) {
    try {
      await accountService.remove(id)
      setAccounts(prev => prev.filter(a => a.id !== id))
    } catch (err) {
      setError((err as Error).message)
    }
  }

  if (loading) return <p>Loading accounts...</p>

  return (
    <div>
      <h1>Accounts</h1>
      {error && <p className="error">{error}</p>}

      <form onSubmit={handleCreate}>
        <input
          value={name}
          onChange={e => setName(e.target.value)}
          placeholder="New account name"
        />

        <input
          type="number"
          step="0.01"
          placeholder="Starting Balance"
          value={starting_balance}
          onChange={e => setBalance(e.target.value === '' ? '' : Number(e.target.value))}
        />

        <select value={type} onChange={e => {
          const value = e.target.value as 'checking' | 'savings' | 'credit' 
          setType(value)
        }}>
          <option value="checking">Checking</option>
          <option value="savings">Savings</option>
          <option value="credit">Credit</option>
        </select>
        <button type="submit">Add Account</button>
      </form>

      <ul>
        {accounts.map(acc => (
          <li key={acc.id}>
            <Link to={`/accounts/${acc.id}`}>{acc.name}</Link> — {acc.type}
            <button className="btn-delete" onClick={() => handleDelete(acc.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  )
}