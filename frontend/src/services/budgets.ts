import { api } from '../lib/api'

export interface Budget {
  id: string
  user_id: string
  category_id: string
  monthly_limit: number
  month: string
  created_at: string
}

export interface BudgetSummaryItem {
  category_id: string
  category_name: string
  monthly_limit: number
  spent: number
}

export const budgetService = {
  list: (): Promise<Budget[]> => api.get('/budgets'),
  create: (data: Pick<Budget, 'category_id' | 'monthly_limit' | 'month'>): Promise<Budget> =>
    api.post('/budgets', data),
  summary: (month: string): Promise<BudgetSummaryItem[]> => api.get(`/budgets/summary?month=${month}`),
}