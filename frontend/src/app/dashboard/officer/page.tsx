import Link from "next/link"
import { BarChart3Icon, CheckCircle2Icon, MessageSquareTextIcon, ShieldCheckIcon } from "lucide-react"

import { buttonVariants } from "@/components/ui/button"
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { cn } from "@/lib/utils"

const stats = [
  { label: "Tugas hari ini", value: "18", icon: CheckCircle2Icon },
  { label: "Laporan masuk", value: "24", icon: MessageSquareTextIcon },
  { label: "Aksi prioritas", value: "7", icon: ShieldCheckIcon },
]

export default function OfficerDashboardPage() {
  return (
    <div className="min-h-screen bg-[radial-gradient(circle_at_top,rgba(15,23,42,0.18),transparent_35%),linear-gradient(180deg,#f8fafc_0%,#e2e8f0_100%)] px-4 py-8 dark:bg-[radial-gradient(circle_at_top,rgba(30,41,59,0.3),transparent_35%),linear-gradient(180deg,#020617_0%,#0f172a_100%)]">
      <div className="mx-auto flex w-full max-w-6xl flex-col gap-8">
        <div className="rounded-3xl border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-700 dark:bg-slate-900">
          <div className="flex flex-wrap items-center justify-between gap-4">
            <div>
              <div className="mb-3 inline-flex rounded-full bg-slate-100 px-3 py-1 text-xs font-semibold text-slate-600 dark:bg-slate-800 dark:text-slate-300">
                Officer / Admin
              </div>
              <h1 className="text-3xl font-semibold text-slate-950 dark:text-white">Selamat bekerja, Officer</h1>
              <p className="mt-2 max-w-2xl text-sm leading-6 text-slate-600 dark:text-slate-300">
                Fokus pada operasi, verifikasi data, dan monitoring aktivitas portal.
              </p>
            </div>
            <Link
              href="/dashboard"
              className={cn(buttonVariants({ size: "lg" }), "rounded-full px-5")}
            >
              Kembali ke pemilih dashboard
            </Link>
          </div>
        </div>

        <div className="grid gap-4 md:grid-cols-3">
          {stats.map((item) => {
            const Icon = item.icon
            return (
              <Card key={item.label} className="rounded-3xl border-slate-200 bg-white/90 shadow-lg dark:border-slate-700 dark:bg-slate-900/90">
                <CardHeader className="space-y-4 p-6">
                  <div className="flex items-center justify-between">
                    <div className="flex size-11 items-center justify-center rounded-2xl bg-slate-100 text-slate-700 dark:bg-slate-800 dark:text-slate-200">
                      <Icon className="size-5" />
                    </div>
                    <BarChart3Icon className="size-4 text-slate-400" />
                  </div>
                  <div>
                    <CardDescription>{item.label}</CardDescription>
                    <CardTitle className="mt-1 text-3xl">{item.value}</CardTitle>
                  </div>
                </CardHeader>
              </Card>
            )
          })}
        </div>
      </div>
    </div>
  )
}
