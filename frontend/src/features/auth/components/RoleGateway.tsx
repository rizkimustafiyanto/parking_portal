"use client"

import Link from "next/link"
import { ArrowRightIcon, ShieldCheckIcon, UsersIcon } from "lucide-react"

import { buttonVariants } from "@/components/ui/button"
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { cn } from "@/lib/utils"
import { AUTH_ROLE_CONFIG } from "@/features/auth"

function RoleGateway() {
  const roles = [AUTH_ROLE_CONFIG.officer, AUTH_ROLE_CONFIG.member]

  return (
    <div className="min-h-screen bg-[radial-gradient(circle_at_top,rgba(148,163,184,0.15),transparent_35%),linear-gradient(180deg,#f8fafc_0%,#eef2ff_100%)] px-4 py-8 dark:bg-[radial-gradient(circle_at_top,rgba(148,163,184,0.12),transparent_35%),linear-gradient(180deg,#020617_0%,#0f172a_100%)]">
      <div className="mx-auto flex min-h-screen w-full max-w-6xl flex-col justify-center gap-10">
        <div className="mx-auto max-w-3xl text-center">
          <div className="mb-4 inline-flex rounded-full border border-slate-200 bg-white px-4 py-1 text-sm font-medium text-slate-600 shadow-sm dark:border-slate-700 dark:bg-slate-900 dark:text-slate-300">
            Portal Digital
          </div>
          <h1 className="text-4xl font-semibold tracking-tight text-slate-950 dark:text-white md:text-6xl">
            Pilih pintu masuk sesuai peran Anda
          </h1>
          <p className="mx-auto mt-4 max-w-2xl text-base leading-7 text-slate-600 dark:text-slate-300 md:text-lg">
            Admin dan officer masuk ke area operasional. Member masuk ke ruang personal untuk melihat status dan informasi penting.
          </p>
        </div>

        <div className="grid gap-5 md:grid-cols-2">
          {roles.map((role) => (
            <Card
              key={role.role}
              className="group overflow-hidden rounded-3xl border-slate-200 bg-white/90 p-0 shadow-lg backdrop-blur transition-transform duration-300 hover:-translate-y-1 dark:border-slate-700 dark:bg-slate-900/90"
            >
              <div className={`h-2 bg-linear-to-r ${role.accent}`} />
              <CardHeader className="space-y-4 p-8">
                <div className="flex items-center justify-between">
                  <div className="flex size-12 items-center justify-center rounded-2xl bg-slate-100 text-slate-700 dark:bg-slate-800 dark:text-slate-200">
                    {role.role === "officer" ? <ShieldCheckIcon className="size-5" /> : <UsersIcon className="size-5" />}
                  </div>
                  <span className="rounded-full bg-slate-100 px-3 py-1 text-xs font-medium text-slate-600 dark:bg-slate-800 dark:text-slate-300">
                    {role.badge}
                  </span>
                </div>
                <div className="space-y-2">
                  <CardTitle className="text-2xl">{role.title}</CardTitle>
                  <CardDescription className="text-sm leading-6 text-slate-600 dark:text-slate-300">
                    {role.description}
                  </CardDescription>
                </div>
                <div className="flex flex-wrap gap-3 pt-2">
                  <Link
                    href={`/login?role=${role.role}`}
                    className={cn(buttonVariants({ size: "lg" }), "rounded-full px-5")}
                  >
                    Mulai
                    <ArrowRightIcon className="size-4" />
                  </Link>
                  <Link
                    href={role.dashboardPath}
                    className={cn(buttonVariants({ variant: "outline", size: "lg" }), "rounded-full px-5")}
                  >
                    Lihat dashboard
                  </Link>
                </div>
              </CardHeader>
            </Card>
          ))}
        </div>
      </div>
    </div>
  )
}

export { RoleGateway }
