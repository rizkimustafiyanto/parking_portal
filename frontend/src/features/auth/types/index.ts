export type LoginPayload = {
  email: string
  password: string
}

export type LoginResponse = {
  status: string
  message: string
  data: {
    token: string
  }
}

export * from "./role"
