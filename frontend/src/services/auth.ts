import { api } from '../lib/api'

export const authService = {
  register: (email: string, password: string) =>
    api.post('/auth/register', { email, password }),
  
  login: async (email: string, password: string) => {
    const res = await api.post('/auth/login', { email, password })
    localStorage.setItem('token', res.token)
    return res
  },
  logout: () => localStorage.removeItem('token'),
}