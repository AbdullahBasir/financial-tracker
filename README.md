# Finance Tracker 💰

A full-stack personal finance management application built with Go and React. This application allows users to efficiently manage their personal finances by tracking multiple financial accounts, categorizing spending, and monitoring budgets. Users can create accounts (checking, savings, investment, etc.), record transactions with automatic categorization, set monthly budgets by category, and view spending summaries. 

Built with production-ready patterns including JWT authentication, type-safe database access via SQLC, and a modern React frontend with TypeScript.

## Tech Stack

**Backend:** Go 1.26.2, PostgreSQL, SQLC, JWT + Argon2  
**Frontend:** TypeScript, React 19, Vite

## Installation

**Prerequisites:** Go 1.26.2+, Node.js 18+, PostgreSQL 12+

Clone the repository:
```bash
git clone https://github.com/AbdullahBasir/financial-tracker.git
cd finance-tracker
```

Create `backend/.env`:
```env
DATABASE_URL=postgres://username:password@localhost:5432/finance_tracker?sslmode=disable
PORT=8080
JWT_SECRET=your_jwt_secret_key_here
```

## Quick Start

**Backend:**
```bash
cd backend
go mod download
sqlc generate
./migrate.sh
go run main.go
```
Backend runs on `http://localhost:8080`

**Frontend:**
```bash
cd frontend
npm install
npm run dev
```
Frontend runs on `http://localhost:5173`

## Features

- User authentication (JWT + Argon2)
- Account management (checking, savings, etc.)
- Transaction tracking with categories
- Monthly budget monitoring
- Account balance tracking

## Project Structure

```
finance-tracker/
├── backend/
│   ├── database/
│   │   ├── migrations/          # SQL migration files (goose)
│   │   ├── queries/             # SQL query definitions for SQLC
│   │   └── sqlc/                # Generated Go code from SQL
│   ├── internal/
│   │   ├── auth/                # Authentication logic (JWT, hashing)
│   │   ├── handler/             # HTTP endpoint handlers
│   │   └── middleware/          # HTTP middleware (auth, CORS, logging)
│   ├── main.go
│   ├── go.mod
│   ├── sqlc.yaml
│   └── migrate.sh
│
├── frontend/
│   ├── src/
│   │   ├── components/          # React components
│   │   ├── pages/               # Page components
│   │   ├── services/            # API communication
│   │   ├── lib/                 # Utilities and helpers
│   │   ├── App.tsx
│   │   ├── main.tsx
│   │   └── index.css
│   ├── package.json
│   └── tsconfig.json
│
└── README.md
```

## Dependencies

### Backend
- `github.com/golang-jwt/jwt/v5` - JWT authentication
- `github.com/alexedwards/argon2id` - Password hashing
- `github.com/lib/pq` - PostgreSQL driver
- `github.com/google/uuid` - UUID generation
- `github.com/shopspring/decimal` - Precise decimal arithmetic for currency
- `github.com/joho/godotenv` - Environment variable loading

### Frontend
- React 19
- React Router DOM 7
- TypeScript
- Vite

## API Endpoints

- `GET /health` - Health check
- `POST /auth/register` - User registration
- `POST /auth/login` - User login
- `POST /accounts` - Create account
- `GET /accounts` - List user accounts
- `GET /accounts/{id}` - Get account details
- `PATCH /accounts/{id}` - Update account
- `DELETE /accounts/{id}` - Delete account
- `POST /categories` - Create category
- `GET /categories` - List categories
- `PATCH /categories/{id}` - Update category
- `GET /transactions` - List transactions
- `POST /transactions` - Create transaction
- `GET /transactions/{id}` - Get transaction details
- `PATCH /transactions/{id}` - Update transaction
- `DELETE /transactions/{id}` - Delete transaction
- `POST /budgets` - Create budget
- `GET /budgets` - List budgets
- `GET /budgets/summary` - Get budget summary for month
- `DELETE /budgets/{id}` - Delete budget