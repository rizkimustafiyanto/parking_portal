"use client"

import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import { getStoredRole, getStoredToken } from "../store"
import type { AuthRole } from "../types"

export function useAuthCheck(expectedRole?: AuthRole) {
  const router = useRouter()
  const [isLoading] = useState(() => Boolean(getStoredToken()))

  useEffect(() => {
    const token = getStoredToken()
    const role = getStoredRole()

    if (!token) {
      return
    }

    if (!role) {
      router.replace("/login")
      return
    }

    if (expectedRole && role === expectedRole) {
      router.replace(role === "member" ? "/dashboard/member" : "/dashboard/officer")
      return
    }

    router.replace(role === "member" ? "/dashboard/member" : "/dashboard/officer")
  }, [expectedRole, router])

  return { isLoading }
}
