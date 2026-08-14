import { api } from '../lib/api'

<<<<<<< HEAD
export interface Category {
  id: string
  user_id: string
  name: string
  type: 'income' | 'expenses'
  created_at: string
}

=======
>>>>>>> parent of ec8236c (Created the basic structure of the website with react, login and register working)
export const categoryService = {
  list: () => api.get('/categories'),
  create: (data: unknown) => api.post('/categories', data),
}