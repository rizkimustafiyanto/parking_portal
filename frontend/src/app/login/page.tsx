import { LoginForm, type AuthRole } from "@/features/auth"

type LoginPageProps = {
  searchParams?: {
    role?: string
  }
}

export default function LoginPage({ searchParams }: LoginPageProps) {
  const role: AuthRole = searchParams?.role === "member" ? "member" : "officer"

  return <LoginForm role={role} />
}
