import { api } from '../lib/api'

export const transactionService = {
  list: (params: Record<string, string> = {}) =>
    api.get(`/transactions?${new URLSearchParams(params)}`),
  create: (data: unknown) => api.post('/transactions', data),
  get: (id: string) => api.get(`/transactions/${id}`),
  update: (id: string, data: unknown) => api.patch(`/transactions/${id}`, data),
  remove: (id: string) => api.del(`/transactions/${id}`),
}