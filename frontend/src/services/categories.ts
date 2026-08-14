import { api } from '../lib/api'

export const categoryService = {
  list: () => api.get('/categories'),
  create: (data: unknown) => api.post('/categories', data),
}