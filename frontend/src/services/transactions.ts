import { api } from '../lib/api'

export interface Transaction {
  id: string
  amount: number
  created_at: string
  occurred_at: string
  description: string
  account_id: string
  category_id: string
}

export interface TransactionPage {
  transaction?: Transaction[]
  page?: number
  page_size?: number
}

export interface TransactionFilters {
  account_id?: string
  category_id?: string
  from?: string
  to?: string
  page?: string
}

export const transactionService = {
  list: async (filters: TransactionFilters = {}): Promise<Transaction[]> => {
    const params = new URLSearchParams(filters as Record<string, string>)
    const result: TransactionPage = await api.get(`/transactions?${params}`)
    return result.transaction ?? result.transaction ?? []
  },
  get: (id: string): Promise<Transaction> => api.get(`/transactions/${id}`),
  create: (data: Pick<Transaction, 'account_id' | 'category_id' | 'amount' | 'description' | 'occurred_at'>): Promise<Transaction> =>
    api.post('/transactions', data),
  update: (id: string, data: Partial<Transaction>): Promise<Transaction> => api.patch(`/transactions/${id}`, data),
  remove: (id: string): Promise<void> => api.del(`/transactions/${id}`),
}