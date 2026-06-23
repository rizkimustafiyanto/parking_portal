"use client"

import { useEffect, useState } from "react"
import { WalletCardsIcon } from "lucide-react"

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { LoadingState } from "@/components/ui/loading-state"
import { getStoredUserId } from "@/features/auth"
import { fetchMemberBalanceHistory, type InvoiceRecord } from "@/features/finance"

export default function MemberNotificationsPage() {
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
        const response = await fetchMemberBalanceHistory(memberId)
        if (active) {
          setItems(response.data ?? [])
        }
      } catch (err) {
        if (active) {
          setError(err instanceof Error ? err.message : "Gagal memuat notifikasi")
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
        <CardTitle className="text-2xl">Payment Updates</CardTitle>
        <p className="text-sm leading-6 text-slate-600 dark:text-slate-300">
          Update pembayaran dan perubahan status invoice untuk akun kamu.
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
            Belum ada update pembayaran.
          </div>
        ) : null}
        {!loading && !error && items.length > 0 ? (
          <div className="space-y-3">
            {items.map((item) => (
              <div key={item.id} className="rounded-2xl border border-emerald-100 bg-white p-4 dark:border-emerald-900/30 dark:bg-slate-800/60">
                <div className="flex items-center justify-between gap-3">
                  <div>
                    <p className="flex items-center gap-2 font-medium text-slate-950 dark:text-white">
                      <WalletCardsIcon className="size-4 text-emerald-600" />
                      {item.violation?.plate_number}
                    </p>
                    <p className="text-sm text-slate-500 dark:text-slate-400">{item.violation?.location}</p>
                  </div>
                  <span className="rounded-full bg-emerald-50 px-3 py-1 text-xs font-medium text-emerald-700 dark:bg-emerald-950/40 dark:text-emerald-300">
                    {item.payment?.status ?? item.status}
                  </span>
                </div>
                <div className="mt-3 grid gap-2 text-sm text-slate-600 dark:text-slate-300 md:grid-cols-2">
                  <p>Amount: {item.amount}</p>
                  <p>Invoice status: {item.status}</p>
                  <p>Paid at: {item.payment?.paid_at ?? "-"}</p>
                  <p>Transaction: {item.payment?.internal_transaction_id ?? "-"}</p>
                </div>
              </div>
            ))}
          </div>
        ) : null}
      </CardContent>
    </Card>
  )
}
