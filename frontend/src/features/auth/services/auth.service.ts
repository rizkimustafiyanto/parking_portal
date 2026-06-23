import { LoginPayload, LoginResponse } from "../types"
import api from "@/lib/api/client"
import axios from "axios"

export async function loginService(payload: LoginPayload): Promise<string> {
  try {
    const response = await api.post<LoginResponse>("/api/auth/login", payload)

    if (!response.data?.data?.token) {
      throw new Error("Token tidak diterima dari server")
    }

    return response.data.data.token
  } catch (error) {
    if (axios.isAxiosError(error)) {
      const message = error.response?.data?.message || error.message
      throw new Error(message || "Gagal login, periksa email dan password")
    }

    if (error instanceof Error) {
      throw error
    }

    throw new Error("Terjadi kesalahan saat login")
  }
}

