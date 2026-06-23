"use client"

import { useEffect, useState } from "react"
import { ReceiptTextIcon } from "lucide-react"

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { LoadingState } from "@/components/ui/loading-state"
import { getStoredUserId } from "@/features/auth"
import { fetchMemberInvoices, type InvoiceRecord } from "@/features/finance"

export default function MemberHistoryPage() {
  const memberId = getStoredUserId()
  const [items, setItems] = useState<InvoiceRecord[]>([])
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
        const response = await fetchMemberInvoices(memberId)
        if (active) {
          setItems(response.data ?? [])
        }
      } catch (err) {
        if (active) {
          setError(err instanceof Error ? err.message : "Gagal memuat riwayat invoice")
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

  return (
    <Card className="rounded-3xl border-emerald-100 bg-white/95 shadow-lg dark:border-emerald-900/30 dark:bg-slate-900/95">
      <CardHeader className="space-y-3 p-8">
        <CardDescription>Member / User</CardDescription>
        <CardTitle className="text-2xl">Invoice History</CardTitle>
        <p className="text-sm leading-6 text-slate-600 dark:text-slate-300">
          Daftar tagihan yang terhubung ke pelanggaran milik akun kamu.
        </p>
      </CardHeader>
      <CardContent className="space-y-4 px-8 pb-8">
        {!memberId ? (
          <div className="rounded-2xl border border-dashed border-emerald-100 p-4 text-sm text-slate-500 dark:border-emerald-900/30">
            User ID tidak ditemukan. Silakan login ulang.
          </div>
        ) : null}
        {loading ? <LoadingState rows={4} /> : null}
        {!loading && error ? <p className="text-sm text-red-600">{error}</p> : null}
        {!loading && !error && items.length === 0 ? (
          <div className="rounded-2xl border border-dashed border-emerald-100 p-4 text-sm text-slate-500 dark:border-emerald-900/30">
            Belum ada invoice.
          </div>
        ) : null}
        {!loading && !error && items.length > 0 ? (
          <div className="space-y-3">
            {items.map((item) => (
              <div key={item.id} className="rounded-2xl border border-emerald-100 bg-emerald-50/60 p-4 dark:border-emerald-900/30 dark:bg-slate-800/60">
                <div className="flex items-center justify-between gap-3">
                  <div>
                    <p className="flex items-center gap-2 font-medium text-slate-950 dark:text-white">
                      <ReceiptTextIcon className="size-4 text-emerald-600" />
                      {textOrDash(item.violation?.plate_number)}
                    </p>
                    <p className="text-sm text-slate-500 dark:text-slate-400">{textOrDash(item.violation?.location)}</p>
                  </div>
                  <span className="rounded-full bg-white px-3 py-1 text-xs font-medium text-slate-600 dark:bg-slate-900 dark:text-slate-300">
                    {item.status}
                  </span>
                </div>
                <div className="mt-3 grid gap-2 text-sm text-slate-600 dark:text-slate-300 md:grid-cols-2">
                  <p>Amount: {item.amount ?? "-"}</p>
                  <p>Payment: {textOrDash(item.payment?.status)}</p>
                  <p>Fine rule version: {item.violation?.fine_rule_version?.version_number ?? "-"}</p>
                  <p>Created at: {textOrDash(item.created_at)}</p>
                </div>
              </div>
            ))}
          </div>
        ) : null}
      </CardContent>
    </Card>
  )
}

function textOrDash(value: string | number | null | undefined) {
  if (typeof value === "number") {
    return Number.isFinite(value) ? String(value) : "-"
  }

  return value && value.trim() ? value : "-"
}
