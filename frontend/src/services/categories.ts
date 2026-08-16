import { api } from '../lib/api'

export interface Category {
  id: string
  name: string
  created_at: string
  type: 'income' | 'expenses'
  user_id: string
}

export const categoryService = {
  list: (): Promise<Category[]> => api.get('/categories'),
  create: (data: Pick<Category, 'name' | 'type'>): Promise<Category> => api.post('/categories', data),
}