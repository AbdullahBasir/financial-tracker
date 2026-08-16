import { api } from '../lib/api'

export interface Account {
  id: string
  name: string
  created_at: string
  starting_balance: number
  type: 'checking' | 'savings' | 'credit'
  user_id: string
}

export const accountService = {
  list: (): Promise<Account[]> => api.get('/accounts'),
  get: (id: string): Promise<Account> => api.get(`/accounts/${id}`),
  create: (data: Pick<Account, 'name' | 'type' | 'starting_balance'>): Promise<Account> =>
    api.post('/accounts', data),
  update: (id: string, data: Partial<Account>): Promise<Account> => api.patch(`/accounts/${id}`, data),
  remove: (id: string): Promise<void> => api.del(`/accounts/${id}`),
}