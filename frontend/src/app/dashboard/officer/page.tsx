"use client"

import Link from "next/link"
import { useEffect, useMemo, useState } from "react"
import { BarChart3Icon, CheckCircle2Icon, MessageSquareTextIcon, ShieldCheckIcon } from "lucide-react"

import { buttonVariants } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { LoadingState } from "@/components/ui/loading-state"
import { cn } from "@/lib/utils"
import { fetchUsers } from "@/features/users"
import { fetchViolations } from "@/features/violation"
import { fetchInvoices, fetchPayments } from "@/features/finance"

type OverviewState = {
  loading: boolean
  error: string | null
  members: number
  officers: number
  violations: number
  invoices: number
  pendingInvoices: number
  paidInvoices: number
  payments: number
  successfulPayments: number
  failedPayments: number
}

const initialState: OverviewState = {
  loading: true,
  error: null,
  members: 0,
  officers: 0,
  violations: 0,
  invoices: 0,
  pendingInvoices: 0,
  paidInvoices: 0,
  payments: 0,
  successfulPayments: 0,
  failedPayments: 0,
}

export default function OfficerDashboardPage() {
  const [overview, setOverview] = useState<OverviewState>(initialState)

  useEffect(() => {
    let active = true

    void (async () => {
      setOverview((current) => ({ ...current, loading: true, error: null }))

      try {
        const [userRes, violationRes, invoiceRes, paymentRes] = await Promise.all([
          fetchUsers({ page: 1, limit: 200 }),
          fetchViolations({ page: 1, limit: 200 }),
          fetchInvoices({ page: 1, limit: 200 }),
          fetchPayments({ page: 1, limit: 200 }),
        ])

        if (!active) {
          return
        }

        const users = userRes.data ?? []
        const violations = violationRes.data ?? []
        const invoices = invoiceRes.data ?? []
        const payments = paymentRes.data ?? []

        setOverview({
          loading: false,
          error: null,
          members: users.filter((item) => item.role === "member").length,
          officers: users.filter((item) => item.role === "officer").length,
          violations: violations.length,
          invoices: invoices.length,
          pendingInvoices: invoices.filter((item) => item.status === "PENDING").length,
          paidInvoices: invoices.filter((item) => item.status === "PAID").length,
          payments: payments.length,
          successfulPayments: payments.filter((item) => item.status === "SUCCESS").length,
          failedPayments: payments.filter((item) => item.status === "FAILED").length,
        })
      } catch (err) {
        if (active) {
          setOverview((current) => ({
            ...current,
            loading: false,
            error: err instanceof Error ? err.message : "Gagal memuat overview",
          }))
        }
      }
    })()

    return () => {
      active = false
    }
  }, [])

  const stats = useMemo(
    () => [
      { label: "Violations", value: overview.violations, icon: ShieldCheckIcon },
      { label: "Pending Invoices", value: overview.pendingInvoices, icon: MessageSquareTextIcon },
      { label: "Paid Invoices", value: overview.paidInvoices, icon: CheckCircle2Icon },
    ],
    [overview.paidInvoices, overview.pendingInvoices, overview.violations]
  )

  return (
    <div className="min-h-screen bg-[radial-gradient(circle_at_top,rgba(15,23,42,0.18),transparent_35%),linear-gradient(180deg,#f8fafc_0%,#e2e8f0_100%)] px-4 py-8 dark:bg-[radial-gradient(circle_at_top,rgba(30,41,59,0.3),transparent_35%),linear-gradient(180deg,#020617_0%,#0f172a_100%)]">
      <div className="mx-auto flex w-full max-w-6xl flex-col gap-8">
        <div className="rounded-3xl border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-700 dark:bg-slate-900">
          <div className="flex flex-wrap items-center justify-between gap-4">
            <div>
              <div className="mb-3 inline-flex rounded-full bg-slate-100 px-3 py-1 text-xs font-semibold text-slate-600 dark:bg-slate-800 dark:text-slate-300">
                Officer / Admin
              </div>
              <h1 className="text-3xl font-semibold text-slate-950 dark:text-white">Overview Operasional</h1>
              <p className="mt-2 max-w-2xl text-sm leading-6 text-slate-600 dark:text-slate-300">
                Ringkasan data nyata dari user, violation, invoice, dan payment di project ini.
              </p>
            </div>
            <Link href="/dashboard" className={cn(buttonVariants({ size: "lg" }), "rounded-full px-5")}>
              Kembali ke pemilih dashboard
            </Link>
          </div>
        </div>

        {overview.loading ? <LoadingState rows={4} /> : null}
        {overview.error ? (
          <div className="rounded-2xl border border-red-200 bg-red-50 p-4 text-sm text-red-700 dark:border-red-900/40 dark:bg-red-950/40 dark:text-red-200">
            {overview.error}
          </div>
        ) : null}

        {!overview.loading && !overview.error ? (
          <>
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

            <div className="grid gap-4 md:grid-cols-3">
              <Stat label="Members" value={overview.members} />
              <Stat label="Officers" value={overview.officers} />
              <Stat label="Payments" value={overview.payments} />
            </div>

            <div className="grid gap-4 md:grid-cols-3">
              <Stat label="Invoices" value={overview.invoices} />
              <Stat label="Successful Payments" value={overview.successfulPayments} />
              <Stat label="Failed Payments" value={overview.failedPayments} />
            </div>
          </>
        ) : null}
      </div>
    </div>
  )
}

function Stat({ label, value }: { label: string; value: number }) {
  return (
    <Card className="rounded-3xl border-slate-200 bg-white/90 shadow-lg dark:border-slate-700 dark:bg-slate-900/90">
      <CardContent className="p-6">
        <CardDescription>{label}</CardDescription>
        <CardTitle className="mt-1 text-3xl">{value}</CardTitle>
      </CardContent>
    </Card>
  )
}
