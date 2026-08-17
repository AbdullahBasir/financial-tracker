# Finance Tracker 💰

A full-stack personal finance management application built with Go and React. Track accounts, transactions, budgets, and categories with an intuitive, modern interface.

## Overview

Finance Tracker is a portfolio-grade project that demonstrates full-stack development with production-ready patterns. Manage your finances efficiently with features for tracking accounts, categorizing transactions, and monitoring budgets across monthly periods.

## Tech Stack

### Backend
- **Language:** Go 1.26.2
- **Framework:** Standard library (`net/http`)
- **Database:** PostgreSQL
- **Database Access:** SQLC (type-safe SQL code generation)
- **Authentication:** JWT + Argon2 password hashing
- **Migrations:** Goose
- **Key Packages:**
  - `github.com/golang-jwt/jwt/v5` - JWT authentication
  - `github.com/alexedwards/argon2id` - Password hashing
  - `github.com/lib/pq` - PostgreSQL driver
  - `github.com/google/uuid` - UUID generation
  - `github.com/shopspring/decimal` - Precise decimal arithmetic for currency

### Frontend
- **Language:** TypeScript
- **Framework:** React 19
- **Build Tool:** Vite
- **Styling:** CSS with design system utilities
- **Package Manager:** npm

## Prerequisites

### System Requirements
- **Go 1.26.2** or later
- **Node.js 18+** and npm 9+
- **PostgreSQL 12+**
- **Git**

### Environment Setup

#### Backend
1. PostgreSQL must be running and accessible
2. Create a `.env` file in the `backend/` directory:
   ```env
   DATABASE_URL=postgres://username:password@localhost:5432/finance_tracker?sslmode=disable
   PORT=8080
   JWT_SECRET=your_jwt_secret_key_here
   ```

#### Frontend
1. Node modules will be installed automatically with `npm install`

## Installation

### 1. Install the Repository
```bash
go install github.com/AbdullahBasir/financial-tracker@latest
```
This compiles and places the binary in $GOPATH/bin (usually ~/go/bin). Make sure ~/go/bin is in your PATH.

or clone the repository

```bash
git clone https://github.com/AbdullahBasir/financial-tracker.git
cd finance-tracker
```

### 2. Backend Setup

```bash
cd backend

# Install Go dependencies
go mod download

# Generate SQLC code from SQL queries
sqlc generate

# Run database migrations
./migrate.sh

# Start the backend server
go run main.go
```

The backend API will be available at `http://localhost:8080`

### 3. Frontend Setup

```bash
cd frontend

# Install npm dependencies
npm install

# Start the development server
npm run dev
```

The frontend will be available at `http://localhost:5173` (Vite default)

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
│   ├── main.go                  # Entry point
│   ├── go.mod                   # Go module definition
│   ├── sqlc.yaml                # SQLC configuration
│   └── migrate.sh               # Migration runner script
│
├── frontend/
│   ├── src/
│   │   ├── components/          # React components
│   │   ├── pages/               # Page components
│   │   ├── services/            # API communication
│   │   ├── lib/                 # Utilities and helpers
│   │   ├── App.tsx              # Root component
│   │   ├── main.tsx             # Entry point
│   │   └── index.css            # Global styles & design system
│   ├── package.json             # NPM dependencies
│   └── tsconfig.json            # TypeScript configuration
│
└── README.md                    # This file
```

## Features

### Accounts
- Create and manage financial accounts (checking, savings, etc.)
- Track account balances and account type
- View account history and transactions

### Categories
- Organize transactions with custom categories
- Track different spending categories (groceries, utilities, etc.)
- Budget monitoring per category

### Transactions
- Record financial transactions with date and amount
- Link transactions to accounts and categories
- Filter transactions by date range, account, or category
- Transaction descriptions and detailed history

### Budgets
- Set monthly budgets per category
- Track spending against budgeted amounts
- View budget summary for any month
- Monitor remaining budget balance

### Authentication
- Secure user registration and login
- JWT-based session management
- Password hashing with Argon2

## Running the Application

### Development Mode

**Terminal 1 - Backend:**
```bash
cd backend
go run main.go
```

**Terminal 2 - Frontend:**
```bash
cd frontend
npm run dev
```

Then open `http://localhost:5173` in your browser.

### Production Build

**Backend:**
```bash
cd backend
go build -o finance-tracker
./finance-tracker
```

**Frontend:**
```bash
cd frontend
npm run build
npm run preview
```

## API Endpoints

### Authentication
- `POST /auth/register` - Create new user
- `POST /auth/login` - Login user

### Accounts
- `GET /accounts` - List user accounts
- `POST /accounts` - Create account
- `GET /accounts/{id}` - Get account details
- `PATCH /accounts/{id}` - Update account
- `DELETE /accounts/{id}` - Delete account

### Categories
- `GET /categories` - List categories
- `POST /categories` - Create category
- `DELETE /categories/{id}` - Delete category

### Transactions
- `GET /transactions` - List transactions (with filtering)
- `POST /transactions` - Create transaction
- `GET /transactions/{id}` - Get transaction details
- `PATCH /transactions/{id}` - Update transaction
- `DELETE /transactions/{id}` - Delete transaction

### Budgets
- `GET /budgets` - List budgets
- `POST /budgets` - Create budget
- `GET /budgets/summary` - Get budget summary for month

## Database Schema

The database includes tables for:
- `users` - User accounts
- `accounts` - Financial accounts
- `categories` - Transaction categories
- `transactions` - Financial transactions
- `budgets` - Monthly budget allocations

Migrations are managed with Goose. Run `./migrate.sh` to apply all pending migrations.

## Development Workflow

### Adding a New Endpoint

1. **Define the SQL query** in `backend/database/queries/*.sql`
2. **Generate Go code** with `sqlc generate`
3. **Create handler** in `backend/internal/handler/`
4. **Register route** in `backend/main.go`
5. **Create frontend page** in `frontend/src/pages/`

### Code Organization

- **Handlers** follow a numbered naming convention (`01_health.go`, `02_register.go`, etc.) for logical ordering
- **Database access** is exclusively through generated SQLC code
- **Middleware** handles cross-cutting concerns (auth, CORS, logging)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Future Enhancements

- Recurring transactions
- Budget alerts and notifications
- Transaction tagging and notes
- Export to CSV/PDF
- Multi-currency support
- Advanced reporting and analytics
- Mobile app (React Native)


