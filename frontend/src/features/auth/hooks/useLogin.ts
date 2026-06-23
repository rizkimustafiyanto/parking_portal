"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { useAuthStore } from "../store"
import { loginService } from "../services"
import type { AuthRole, LoginPayload } from "../types"

export function useLogin(role: AuthRole = "officer") {
  const router = useRouter()
  const setToken = useAuthStore((state) => state.setToken)
  const setRole = useAuthStore((state) => state.setRole)
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  async function login(payload: LoginPayload) {
    setIsLoading(true)
    setError(null)

    try {
      const token = await loginService(payload)
      setToken(token)
      setRole(role)
      router.push(role === "member" ? "/dashboard/member" : "/dashboard/officer")
    } catch (err) {
      const message = err instanceof Error ? err.message : "Gagal login"
      setError(message)
      setIsLoading(false)
    }
  }

  return {
    login,
    isLoading,
    error,
  }
}
