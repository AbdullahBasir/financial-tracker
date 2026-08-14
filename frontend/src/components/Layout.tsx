import { NavLink, Outlet, useNavigate } from 'react-router-dom'
import { authService } from '../services/auth'

export function Layout() {
  const navigate = useNavigate()

  function handleLogout() {
    authService.logout()
    navigate('/login')
  }

  return (
    <div className="app-shell">
      <nav>
        <NavLink to="/">Dashboard</NavLink>
        <NavLink to="/accounts">Accounts</NavLink>
        <NavLink to="/transactions">Transactions</NavLink>
        <NavLink to="/categories">Categories</NavLink>
        <NavLink to="/budgets">Budgets</NavLink>
        <button onClick={handleLogout}>Logout</button>
      </nav>

      <main>
        {/* This is where the matched child route renders */}
        <Outlet />
      </main>
    </div>
  )
}