"use client"

import Link from "next/link"
import { LayoutDashboardIcon, LogOutIcon } from "lucide-react"

import { buttonVariants } from "@/components/ui/button"
import { useAuthStore } from "@/features/auth"
import { cn } from "@/lib/utils"

type DashboardTopbarProps = {
  homeHref: string
  title: string
}

function DashboardTopbar({ homeHref, title }: DashboardTopbarProps) {
  const logout = useAuthStore((state) => state.logout)

  function handleLogout() {
    logout()
    window.location.href = "/login"
  }

  return (
    <div className="flex flex-wrap items-center justify-between gap-3 rounded-3xl border border-slate-200 bg-white/90 px-5 py-4 shadow-lg dark:border-slate-700 dark:bg-slate-900/90">
      <div className="flex items-center gap-3">
        <div className="flex size-11 items-center justify-center rounded-2xl bg-slate-100 text-slate-700 dark:bg-slate-800 dark:text-slate-200">
          <LayoutDashboardIcon className="size-5" />
        </div>
        <div>
          <p className="text-sm font-medium text-slate-500 dark:text-slate-400">Active workspace</p>
          <h2 className="text-lg font-semibold text-slate-950 dark:text-white">{title}</h2>
        </div>
      </div>

      <div className="flex flex-wrap items-center gap-2">
        <Link href={homeHref} className={cn(buttonVariants({ variant: "outline" }), "rounded-full px-4")}>
          Dashboard
        </Link>
        <button
          type="button"
          onClick={handleLogout}
          className={cn(buttonVariants({ variant: "outline" }), "rounded-full px-4")}
        >
          <LogOutIcon className="size-4" />
          Logout
        </button>
      </div>
    </div>
  )
}

export { DashboardTopbar }
