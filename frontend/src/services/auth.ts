// src/services/auth.ts
import { api } from '../lib/api'

export interface User {
  id: string
  email: string
  created_at: string
}

export interface LoginResponse {
  token: string
  user?: User
}

export const authService = {
  register: (email: string, password: string): Promise<User> =>
    api.post('/auth/register', { email, password }),

  login: async (email: string, password: string): Promise<LoginResponse> => {
    const res: LoginResponse = await api.post('/auth/login', { email, password })
    localStorage.setItem('token', res.token)
    return res
  },

  logout: (): void => {
    localStorage.removeItem('token')
  },
}