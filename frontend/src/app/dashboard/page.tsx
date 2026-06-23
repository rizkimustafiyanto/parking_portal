"use client"

import { useEffect } from "react"
import { useRouter } from "next/navigation"

import { LoadingState } from "@/components/ui/loading-state"
import { getStoredRole } from "@/features/auth"

export default function DashboardLandingPage() {
  const router = useRouter()

  useEffect(() => {
    const role = getStoredRole()

    if (role === "member") {
      router.replace("/dashboard/member")
      return
    }

    if (role === "officer") {
      router.replace("/dashboard/officer")
      return
    }

    router.replace("/login")
  }, [router])

  return (
    <div className="flex min-h-screen items-center justify-center bg-slate-50 px-4 dark:bg-slate-950">
      <div className="w-full max-w-md rounded-3xl border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-700 dark:bg-slate-900">
        <LoadingState rows={3} />
      </div>
    </div>
  )
}
