"use client"

import Link from "next/link"
import { useEffect, useMemo, useState } from "react"
import { BadgeCheckIcon, FileTextIcon, HeartHandshakeIcon, UserRoundIcon } from "lucide-react"

import { buttonVariants } from "@/components/ui/button"
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { LoadingState } from "@/components/ui/loading-state"
import { cn } from "@/lib/utils"
import { getStoredUserId } from "@/features/auth"
import { fetchMemberBalanceHistory, fetchMemberInvoices, type InvoiceRecord } from "@/features/finance"

const stats = [
  { label: "Status aktif", value: "Aman", icon: BadgeCheckIcon },
  { label: "Riwayat terbaru", value: "5", icon: FileTextIcon },
  { label: "Poin layanan", value: "92", icon: HeartHandshakeIcon },
]

export default function MemberDashboardPage() {
  const memberId = getStoredUserId()
  const [invoices, setInvoices] = useState<InvoiceRecord[]>([])
  const [history, setHistory] = useState<InvoiceRecord[]>([])
  const [loading, setLoading] = useState(Boolean(memberId))
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (!memberId) {
      return
    }

    let active = true

    void (async () => {
      setLoading(true)
      setError(null)
      try {
        const [invoiceRes, historyRes] = await Promise.all([
          fetchMemberInvoices(memberId),
          fetchMemberBalanceHistory(memberId),
        ])

        if (!active) {
          return
        }

        setInvoices(invoiceRes.data ?? [])
        setHistory(historyRes.data ?? [])
      } catch (err) {
        if (active) {
          setError(err instanceof Error ? err.message : "Gagal memuat data member")
        }
      } finally {
        if (active) {
          setLoading(false)
        }
      }
    })()

    return () => {
      active = false
    }
  }, [memberId])

  const summary = useMemo(
    () => ({
      totalInvoices: invoices.length,
      paid: invoices.filter((item) => item.status === "PAID").length,
      pending: invoices.filter((item) => item.status === "PENDING").length,
    }),
    [invoices]
  )

  return (
    <div className="min-h-screen bg-[radial-gradient(circle_at_top,rgba(16,185,129,0.16),transparent_35%),linear-gradient(180deg,#f8fafc_0%,#ecfeff_100%)] px-4 py-8 dark:bg-[radial-gradient(circle_at_top,rgba(20,184,166,0.18),transparent_35%),linear-gradient(180deg,#020617_0%,#064e3b_100%)]">
      <div className="mx-auto flex w-full max-w-6xl flex-col gap-8">
        {!memberId ? (
          <Card className="rounded-3xl border-emerald-100 bg-white p-8 shadow-xl dark:border-emerald-900/30 dark:bg-slate-900">
            <CardHeader className="p-0">
              <CardDescription>Member / User</CardDescription>
              <CardTitle className="text-2xl">Login ulang diperlukan</CardTitle>
              <p className="text-sm leading-6 text-slate-600 dark:text-slate-300">
                User ID tidak ditemukan di token. Silakan login ulang untuk melihat data member kamu.
              </p>
            </CardHeader>
          </Card>
        ) : null}

        <div className="rounded-3xl border border-emerald-100 bg-white p-8 shadow-xl dark:border-emerald-900/30 dark:bg-slate-900">
          <div className="flex flex-wrap items-center justify-between gap-4">
            <div>
              <div className="mb-3 inline-flex rounded-full bg-emerald-50 px-3 py-1 text-xs font-semibold text-emerald-700 dark:bg-emerald-950/40 dark:text-emerald-300">
                Member / User
              </div>
              <h1 className="text-3xl font-semibold text-slate-950 dark:text-white">Selamat datang, Member</h1>
              <p className="mt-2 max-w-2xl text-sm leading-6 text-slate-600 dark:text-slate-300">
                Lihat ringkasan status, riwayat, dan informasi penting dalam satu tempat yang ringan.
              </p>
            </div>
            <Link
              href="/dashboard"
              className={cn(buttonVariants({ variant: "outline", size: "lg" }), "rounded-full px-5")}
            >
              Kembali ke pemilih dashboard
            </Link>
          </div>
        </div>

        <div className="grid gap-4 md:grid-cols-3">
          {stats.map((item) => {
            const Icon = item.icon
            return (
              <Card key={item.label} className="rounded-3xl border-emerald-100 bg-white/90 shadow-lg dark:border-emerald-900/30 dark:bg-slate-900/90">
                <CardHeader className="space-y-4 p-6">
                  <div className="flex items-center justify-between">
                    <div className="flex size-11 items-center justify-center rounded-2xl bg-emerald-50 text-emerald-700 dark:bg-emerald-950/40 dark:text-emerald-300">
                      <Icon className="size-5" />
                    </div>
                    <UserRoundIcon className="size-4 text-slate-400" />
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
          <Card className="rounded-3xl border-emerald-100 bg-white/90 shadow-lg dark:border-emerald-900/30 dark:bg-slate-900/90">
            <CardHeader className="p-6">
              <CardDescription>Invoices</CardDescription>
              <CardTitle className="mt-1 text-3xl">{summary.totalInvoices}</CardTitle>
            </CardHeader>
          </Card>
          <Card className="rounded-3xl border-emerald-100 bg-white/90 shadow-lg dark:border-emerald-900/30 dark:bg-slate-900/90">
            <CardHeader className="p-6">
              <CardDescription>Paid</CardDescription>
              <CardTitle className="mt-1 text-3xl">{summary.paid}</CardTitle>
            </CardHeader>
          </Card>
          <Card className="rounded-3xl border-emerald-100 bg-white/90 shadow-lg dark:border-emerald-900/30 dark:bg-slate-900/90">
            <CardHeader className="p-6">
              <CardDescription>Pending</CardDescription>
              <CardTitle className="mt-1 text-3xl">{summary.pending}</CardTitle>
            </CardHeader>
          </Card>
        </div>

        <div className="grid gap-6 xl:grid-cols-2">
          <Card className="rounded-3xl border-emerald-100 bg-white/95 shadow-lg dark:border-emerald-900/30 dark:bg-slate-900/95">
            <CardHeader className="p-6">
              <CardDescription>Member / User</CardDescription>
              <CardTitle className="text-2xl">Invoice Terbaru</CardTitle>
              <p className="text-sm leading-6 text-slate-600 dark:text-slate-300">
                Ringkasan tagihan yang terhubung ke pelanggaran kamu.
              </p>
            </CardHeader>
            <div className="px-6 pb-6">
              {loading ? <LoadingState rows={3} /> : null}
              {!loading && error ? <p className="text-sm text-red-600">{error}</p> : null}
              {!loading && !error && invoices.length === 0 ? (
                <div className="rounded-2xl border border-dashed border-emerald-100 p-4 text-sm text-slate-500 dark:border-emerald-900/30">
                  Belum ada invoice.
                </div>
              ) : null}
              {!loading && !error && invoices.length > 0 ? (
                <div className="space-y-3">
                  {invoices.slice(0, 5).map((item) => (
                    <div key={item.id} className="rounded-2xl border border-emerald-100 bg-emerald-50/60 p-4 dark:border-emerald-900/30 dark:bg-slate-800/60">
                      <div className="flex items-center justify-between gap-3">
                        <div>
                          <p className="font-medium text-slate-950 dark:text-white">{item.violation?.plate_number}</p>
                          <p className="text-sm text-slate-500 dark:text-slate-400">{item.violation?.location}</p>
                        </div>
                        <span className="rounded-full bg-white px-3 py-1 text-xs font-medium text-slate-600 dark:bg-slate-900 dark:text-slate-300">
                          {item.status}
                        </span>
                      </div>
                      <div className="mt-3 grid gap-2 text-sm text-slate-600 dark:text-slate-300 md:grid-cols-2">
                        <p>Amount: {item.amount}</p>
                        <p>Payment: {item.payment?.status ?? "-"}</p>
                      </div>
                    </div>
                  ))}
                </div>
              ) : null}
            </div>
          </Card>

          <Card className="rounded-3xl border-emerald-100 bg-white/95 shadow-lg dark:border-emerald-900/30 dark:bg-slate-900/95">
            <CardHeader className="p-6">
              <CardDescription>Member / User</CardDescription>
              <CardTitle className="text-2xl">Balance History</CardTitle>
              <p className="text-sm leading-6 text-slate-600 dark:text-slate-300">
                Riwayat perubahan saldo dan pembayaran dari akun kamu.
              </p>
            </CardHeader>
            <div className="px-6 pb-6">
              {loading ? <LoadingState rows={3} /> : null}
              {!loading && !error && history.length > 0 ? (
                <div className="space-y-3">
                  {history.slice(0, 5).map((item) => (
                    <div key={item.id} className="rounded-2xl border border-emerald-100 bg-white p-4 dark:border-emerald-900/30 dark:bg-slate-800/60">
                      <p className="font-medium text-slate-950 dark:text-white">{item.violation?.plate_number}</p>
                      <p className="text-sm text-slate-500 dark:text-slate-400">{item.status} - {item.payment?.status ?? "-"}</p>
                      <p className="mt-2 text-sm text-slate-600 dark:text-slate-300">Amount: {item.amount}</p>
                    </div>
                  ))}
                </div>
              ) : null}
            </div>
          </Card>
        </div>
      </div>
    </div>
  )
}
