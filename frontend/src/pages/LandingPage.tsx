// src/pages/LandingPage.tsx
import { Link } from 'react-router-dom'

export function LandingPage() {
  return (
    <div className="landing-page">
      <header className="landing-hero">
        <h1>Take control of your money</h1>
        <p>
          Track accounts, categorize spending, and stay on budget —
          all in one place.
        </p>
        <div className="landing-actions">
          <Link to="/register" className="btn-primary">Get Started</Link>
          <Link to="/login" className="btn-secondary">Log In</Link>
        </div>
      </header>

      <section className="landing-features">
        <div className="feature-card">
          <h3>Accounts</h3>
          <p>Keep checking, savings, and credit accounts in sync.</p>
        </div>
        <div className="feature-card">
          <h3>Transactions</h3>
          <p>Log spending and income, filter by date or category.</p>
        </div>
        <div className="feature-card">
          <h3>Budgets</h3>
          <p>Set monthly limits per category and see spend vs. budget at a glance.</p>
        </div>
      </section>
    </div>
  )
}