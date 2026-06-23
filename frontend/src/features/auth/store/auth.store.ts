import { create } from "zustand"
import { TOKEN_STORAGE_KEY } from "@/lib/api/client"

type AuthState = {
  token: string | null
  role: string | null
  setToken: (token: string | null) => void
  setRole: (role: string | null) => void
  logout: () => void
}

const ROLE_STORAGE_KEY = "portal_auth_role"

export const useAuthStore = create<AuthState>((set) => ({
  token: null,
  role: null,
  setToken: (token) => {
    if (typeof window !== "undefined") {
      if (token) {
        localStorage.setItem(TOKEN_STORAGE_KEY, token)
      } else {
        localStorage.removeItem(TOKEN_STORAGE_KEY)
      }
    }

    set({ token })
  },
  setRole: (role) => {
    if (typeof window !== "undefined") {
      if (role) {
        localStorage.setItem(ROLE_STORAGE_KEY, role)
      } else {
        localStorage.removeItem(ROLE_STORAGE_KEY)
      }
    }

    set({ role })
  },
  logout: () => {
    if (typeof window !== "undefined") {
      localStorage.removeItem(TOKEN_STORAGE_KEY)
      localStorage.removeItem(ROLE_STORAGE_KEY)
    }

    set({ token: null, role: null })
  },
}))

export function getStoredToken(): string | null {
  if (typeof window === "undefined") {
    return null
  }

  return localStorage.getItem(TOKEN_STORAGE_KEY)
}

export function getStoredRole(): string | null {
  if (typeof window === "undefined") {
    return null
  }

  return localStorage.getItem(ROLE_STORAGE_KEY)
}

export function getStoredUserId(): string | null {
  if (typeof window === "undefined") {
    return null
  }

  const token = localStorage.getItem(TOKEN_STORAGE_KEY)
  if (!token) {
    return null
  }

  const payload = token.split(".")[1]
  if (!payload) {
    return null
  }

  try {
    const normalized = payload.replace(/-/g, "+").replace(/_/g, "/")
    const decoded = atob(normalized.padEnd(Math.ceil(normalized.length / 4) * 4, "="))
    const parsed = JSON.parse(decoded) as { user_id?: string }
    return parsed.user_id ?? null
  } catch {
    return null
  }
}
