import { api } from '../lib/api'

export const budgetService = {
  list: () => api.get('/budgets'),
  create: (data: unknown) => api.post('/budgets', data),
  getSummary: (month: string) => api.get(`/budgets/summary?month=${month}`),
}