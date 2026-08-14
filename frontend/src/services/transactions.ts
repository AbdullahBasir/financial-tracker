import { api } from '../lib/api'

export interface Transaction {
  id: string
  account_id: string
  category_id: string
  amount: number
  description: string
  occurred_at: string
  created_at: string
}

export interface TransactionFilters {
  account_id?: string
  category_id?: string
  from?: string
  to?: string
  page?: string
}

export const transactionService = {
  list: (filters: TransactionFilters = {}): Promise<Transaction[]> => {
    const params = new URLSearchParams(filters as Record<string, string>)
    return api.get(`/transactions?${params}`)
  },
  get: (id: string): Promise<Transaction> => api.get(`/transactions/${id}`),
  create: (data: Pick<Transaction, 'account_id' | 'category_id' | 'amount' | 'description' | 'occurred_at'>): Promise<Transaction> =>
    api.post('/transactions', data),
  update: (id: string, data: Partial<Transaction>): Promise<Transaction> => api.patch(`/transactions/${id}`, data),
  remove: (id: string): Promise<void> => api.del(`/transactions/${id}`),
}