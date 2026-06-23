"use client"

import { useEffect } from "react"
import { useRouter } from "next/navigation"

import { LoadingState } from "@/components/ui/loading-state"
import { getStoredRole, getStoredToken } from "@/features/auth"
import type { AuthRole } from "@/features/auth"

type RoleGuardProps = {
  role: AuthRole
  children: React.ReactNode
}

function RoleGuard({ role, children }: RoleGuardProps) {
  const router = useRouter()

  useEffect(() => {
    const token = getStoredToken()
    const storedRole = getStoredRole()

    if (!token) {
      router.replace("/login")
      return
    }

    if (!storedRole) {
      router.replace("/login")
      return
    }

    if (storedRole !== role) {
      router.replace(storedRole === "member" ? "/dashboard/member" : "/dashboard/officer")
    }
  }, [role, router])

  const token = getStoredToken()
  const storedRole = getStoredRole()

  if (!token || !storedRole || storedRole !== role) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-slate-50 px-4 dark:bg-slate-950">
        <div className="w-full max-w-md rounded-3xl border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-700 dark:bg-slate-900">
          <LoadingState rows={3} />
        </div>
      </div>
    )
  }

  return children
}

export { RoleGuard }
