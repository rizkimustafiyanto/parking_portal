"use client"

import { useSearchParams } from "next/navigation"

import { LoginForm, type AuthRole } from "@/features/auth"

export default function LoginPage() {
  const searchParams = useSearchParams()
  const rawRole = searchParams.get("role")
  const role: AuthRole = rawRole === "member" ? "member" : "officer"

  return <LoginForm role={role} />
}
