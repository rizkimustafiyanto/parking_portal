import Link from "next/link"
import type { LucideIcon } from "lucide-react"

import { buttonVariants } from "@/components/ui/button"
import { Card } from "@/components/ui/card"
import { cn } from "@/lib/utils"
import { DashboardTopbar } from "./dashboard-topbar"

export type DashboardNavItem = {
  label: string
  href: string
  icon: LucideIcon
}

type DashboardShellProps = {
  title: string
  subtitle: string
  badge: string
  accentClassName: string
  homeHref: string
  navItems: DashboardNavItem[]
  children: React.ReactNode
}

function DashboardShell({
  title,
  subtitle,
  badge,
  accentClassName,
  homeHref,
  navItems,
  children,
}: DashboardShellProps) {
  return (
    <div className="min-h-screen bg-[radial-gradient(circle_at_top,rgba(148,163,184,0.14),transparent_35%),linear-gradient(180deg,#f8fafc_0%,#eef2ff_100%)] px-4 py-6 dark:bg-[radial-gradient(circle_at_top,rgba(15,23,42,0.2),transparent_35%),linear-gradient(180deg,#020617_0%,#0f172a_100%)]">
      <div className="mx-auto grid w-full max-w-7xl gap-6 lg:grid-cols-[280px_1fr]">
        <aside className="space-y-4">
          <Card className="overflow-hidden rounded-3xl border-slate-200 bg-white/95 p-0 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
            <div className={cn("h-2 bg-gradient-to-r", accentClassName)} />
            <div className="space-y-6 p-6">
              <div>
                <div className="mb-3 inline-flex rounded-full bg-slate-100 px-3 py-1 text-xs font-semibold text-slate-600 dark:bg-slate-800 dark:text-slate-300">
                  {badge}
                </div>
                <h1 className="text-xl font-semibold text-slate-950 dark:text-white">{title}</h1>
                <p className="mt-2 text-sm leading-6 text-slate-600 dark:text-slate-300">{subtitle}</p>
              </div>

              <nav className="space-y-2">
                {navItems.map((item) => {
                  const Icon = item.icon

                  return (
                    <Link
                      key={item.href}
                      href={item.href}
                      className={cn(
                        buttonVariants({ variant: "ghost" }),
                        "w-full justify-start rounded-2xl px-4 py-6 text-sm font-medium"
                      )}
                    >
                      <Icon className="size-4" />
                      {item.label}
                    </Link>
                  )
                })}
              </nav>
            </div>
          </Card>

        </aside>

        <main className="space-y-6">
          <DashboardTopbar homeHref={homeHref} title={title} />
          {children}
        </main>
      </div>
    </div>
  )
}

export { DashboardShell }
