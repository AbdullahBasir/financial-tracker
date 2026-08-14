import { api } from '../lib/api'

export const accountService = {
  list: () => api.get('/accounts'),
  create: (data: unknown) => api.post('/accounts', data),
  get: (id: string) => api.get(`/accounts/${id}`),
  update: (id: string, data: unknown) => api.patch(`/accounts/${id}`, data),
  remove: (id: string) => api.del(`/accounts/${id}`),
}