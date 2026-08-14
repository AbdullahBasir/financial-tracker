import { api } from '../lib/api'

export interface Account {
  id: string
  user_id: string
  name: string
  type: 'checking' | 'savings' | 'credit'
  starting_balance: number
  created_at: string
}

export const accountService = {
  list: (): Promise<Account[]> => api.get('/accounts'),
  get: (id: string): Promise<Account> => api.get(`/accounts/${id}`),
  create: (data: Pick<Account, 'name' | 'type' | 'starting_balance'>): Promise<Account> =>
    api.post('/accounts', data),
  update: (id: string, data: Partial<Account>): Promise<Account> => api.patch(`/accounts/${id}`, data),
  remove: (id: string): Promise<void> => api.del(`/accounts/${id}`),
}