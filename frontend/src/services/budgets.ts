// src/services/budgets.ts
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

export interface BudgetSummaryResponse {
  month: string
  items: BudgetSummaryItem[]
  total_budget: number
  total_spent: number
  total_remaining: number
}

export const budgetService = {
  list: async (): Promise<Budget[]> => {
    const data = await api.get('/budgets')
    return data ?? []
  },
  create: (data: Pick<Budget, 'category_id' | 'monthly_limit' | 'month'>): Promise<Budget> =>
    api.post('/budgets', data),
  summary: async (month: string): Promise<BudgetSummaryItem[]> => {
    const result: BudgetSummaryResponse = await api.get(`/budgets/summary?month=${month}`)
    return result.items ?? []
  },
}