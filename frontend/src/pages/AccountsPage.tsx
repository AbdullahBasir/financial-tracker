import { useEffect, useState, type FormEvent } from 'react'
import { Link } from 'react-router-dom'
import { accountService, type Account } from '../services/accounts'

export function AccountsPage() {
  const [accounts, setAccounts] = useState<Account[]>([])
  const [name, setName] = useState('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    loadAccounts()
  }, [])

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

  async function handleCreate(e: FormEvent) {
    e.preventDefault()
    if (!name.trim()) return
    try {
      const newAccount = await accountService.create({
        name,
        type: 'checking',
        starting_balance: 0,
      })
      setAccounts(prev => [...prev, newAccount])
      setName('')
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
        <button type="submit">Add Account</button>
      </form>

      <ul>
        {accounts.map(acc => (
          <li key={acc.id}>
            <Link to={`/accounts/${acc.id}`}>{acc.name}</Link> — {acc.type}
            <button onClick={() => handleDelete(acc.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  )
}