"use client"

import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { LoadingState } from "@/components/ui/loading-state"
import {
  AUTH_ROLE_CONFIG,
  loginSchema,
  type AuthRole,
  type LoginFormValues,
  useAuthCheck,
  useLogin,
} from "@/features/auth"
import Link from "next/link"
import { ArrowLeftIcon } from "lucide-react"

type LoginFormProps = {
  role?: AuthRole
}

function LoginForm({ role = "officer" }: LoginFormProps) {
  const roleConfig = AUTH_ROLE_CONFIG[role]
  const { isLoading: isCheckingAuth } = useAuthCheck(role)
  const { login, isLoading, error } = useLogin(role)

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  })

  if (isCheckingAuth) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-slate-50 px-4 dark:bg-slate-950">
        <div className="w-full max-w-md space-y-4 rounded-3xl border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-700 dark:bg-slate-900">
          <LoadingState rows={3} />
        </div>
      </div>
    )
  }

  async function onSubmit(values: LoginFormValues) {
    await login(values)
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-slate-50 px-4 py-10 dark:bg-slate-950">
      <Card className="w-full max-w-md overflow-hidden rounded-3xl border border-slate-200 bg-white shadow-xl dark:border-slate-700 dark:bg-slate-900">
        <div className={`h-2 bg-gradient-to-r ${roleConfig.accent}`} />
        <CardHeader className="space-y-3 pb-4 pt-8 text-center">
          <div className="mx-auto inline-flex rounded-full bg-slate-100 px-3 py-1 text-xs font-medium text-slate-600 dark:bg-slate-800 dark:text-slate-300">
            {roleConfig.badge}
          </div>
          <CardTitle className="text-2xl">{roleConfig.title}</CardTitle>
          <CardDescription className="text-sm text-slate-500 dark:text-slate-400">
            {roleConfig.description}
          </CardDescription>
          <p className="text-sm text-slate-500 dark:text-slate-400">{roleConfig.helperText}</p>
        </CardHeader>

        <CardContent>
          <div className="mb-5">
            <Link
              href="/"
              className="inline-flex items-center gap-2 text-sm font-medium text-slate-600 transition-colors hover:text-slate-950 dark:text-slate-300 dark:hover:text-white"
            >
              <ArrowLeftIcon className="size-4" />
              Back to Portal
            </Link>
          </div>
          <form className="space-y-5" onSubmit={handleSubmit(onSubmit)}>
            <div>
              <label className="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300" htmlFor="email">
                Email
              </label>
              <Input id="email" type="email" placeholder="admin@example.com" {...register("email")} />
              {errors.email ? (
                <p className="mt-2 text-sm text-destructive">{errors.email.message}</p>
              ) : null}
            </div>

            <div>
              <label className="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300" htmlFor="password">
                Password
              </label>
              <Input id="password" type="password" placeholder="Password Anda" {...register("password")} />
              {errors.password ? (
                <p className="mt-2 text-sm text-destructive">{errors.password.message}</p>
              ) : null}
            </div>

            {error ? <p className="text-sm text-destructive">{error}</p> : null}

            <Button className="w-full" type="submit" loading={isLoading}>
              {roleConfig.buttonLabel}
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}

export { LoginForm }
