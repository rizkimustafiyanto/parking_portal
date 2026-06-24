"use client"

import { Suspense } from "react"
import { useSearchParams } from "next/navigation"

import { LoginForm, type AuthRole } from "@/features/auth"

function LoginPageContent() {
  const searchParams = useSearchParams()
  const rawRole = searchParams.get("role")
  const role: AuthRole = rawRole === "member" ? "member" : "officer"

  return <LoginForm role={role} />
}

export default function LoginPage() {
  return (
    <Suspense fallback={null}>
      <LoginPageContent />
    </Suspense>
  )
}
