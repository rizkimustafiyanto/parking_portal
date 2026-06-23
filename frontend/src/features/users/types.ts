export type UserRole = "officer" | "member" | "admin" | string

export type UserRecord = {
  id: string
  name: string
  email: string
  role: UserRole
  balance: number
  created_at: string
}

export type UserListMeta = {
  page: number
  limit: number
  total: number
  total_pages: number
}

export type UserListResponse = {
  message?: string
  data: UserRecord[]
  meta?: UserListMeta
}

export type CreateUserPayload = {
  name: string
  email: string
  password: string
  role?: string
}

export type UpdateUserPayload = {
  name?: string
  email?: string
  password?: string
  role?: string
}

export type TopUpBalancePayload = {
  amount: number
}

