import { api } from '../lib/api'

export interface Category {
  id: string
  user_id: string
  name: string
  type: 'income' | 'expense'
  created_at: string
}

export const categoryService = {
  list: (): Promise<Category[]> => api.get('/categories'),
  create: (data: Pick<Category, 'name' | 'type'>): Promise<Category> => api.post('/categories', data),
}