"use client"

import { useCallback, useEffect, useMemo, useState } from "react"
import { PencilIcon, RefreshCcwIcon, Trash2Icon } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { LoadingState } from "@/components/ui/loading-state"
import { createPayment, deletePayment, fetchPayments, updatePayment, type PaymentRecord } from "@/features/finance"
import { fetchInvoices, type InvoiceRecord } from "@/features/finance"

export default function OfficerPaymentsPage() {
  const [items, setItems] = useState<PaymentRecord[]>([])
  const [invoices, setInvoices] = useState<InvoiceRecord[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [query, setQuery] = useState("")
  const [busy, setBusy] = useState<string | null>(null)
  const [mode, setMode] = useState<"create" | "edit">("create")
  const [editingId, setEditingId] = useState<string | null>(null)
  const [selectedPayment, setSelectedPayment] = useState<PaymentRecord | null>(null)
  const [form, setForm] = useState({
    invoice_id: "",
    internal_transaction_id: "",
    amount: "0",
    status: "SUCCESS",
    scenario: "SUCCESS",
    paid_at: "",
  })

  const loadAll = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const [paymentRes, invoiceRes] = await Promise.all([
        fetchPayments({ page: 1, limit: 50, search: query }),
        fetchInvoices({ page: 1, limit: 100 }),
      ])
      setItems(paymentRes.data ?? [])
      setInvoices(invoiceRes.data ?? [])
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal memuat payment")
    } finally {
      setLoading(false)
    }
  }, [query])

  useEffect(() => {
    let active = true

    void (async () => {
      setLoading(true)
      setError(null)
      try {
        const [paymentRes, invoiceRes] = await Promise.all([
          fetchPayments({ page: 1, limit: 50, search: query }),
          fetchInvoices({ page: 1, limit: 100 }),
        ])

        if (!active) {
          return
        }

        setItems(paymentRes.data ?? [])
        setInvoices(invoiceRes.data ?? [])
      } catch (err) {
        if (active) {
          setError(err instanceof Error ? err.message : "Gagal memuat payment")
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
  }, [loadAll, query])

  const summary = useMemo(
    () => ({
      total: items.length,
      success: items.filter((item) => item.status === "SUCCESS").length,
      failed: items.filter((item) => item.status === "FAILED").length,
    }),
    [items]
  )

  async function submit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setBusy("form")
    setError(null)
    try {
      const payload = {
        invoice_id: form.invoice_id,
        amount: Number(form.amount),
        scenario: form.scenario,
        paid_at: new Date(form.paid_at).toISOString(),
      }

      if (mode === "edit" && editingId) {
        await updatePayment(editingId, payload)
      } else {
        await createPayment(payload)
      }

      setMode("create")
      setEditingId(null)
      setForm({
        invoice_id: "",
        internal_transaction_id: "",
        amount: "0",
        status: "SUCCESS",
        scenario: "SUCCESS",
        paid_at: "",
      })
      await loadAll()
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal menyimpan payment")
    } finally {
      setBusy(null)
    }
  }

  function edit(item: PaymentRecord) {
    setMode("edit")
    setEditingId(item.id)
    setForm({
      invoice_id: item.invoice_id,
      internal_transaction_id: item.internal_transaction_id,
      amount: String(item.amount),
      status: item.status,
      scenario: item.scenario,
      paid_at: item.paid_at.slice(0, 19),
    })
  }

  return (
    <div className="space-y-6">
      <Card className="overflow-hidden rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
        <div className="h-2 bg-gradient-to-r from-slate-950 via-slate-800 to-slate-700" />
        <CardHeader className="space-y-3 p-8">
          <CardDescription>Officer / Admin</CardDescription>
          <CardTitle className="text-2xl">Payment Transactions</CardTitle>
          <p className="max-w-3xl text-sm leading-6 text-slate-600 dark:text-slate-300">
            Payment transaction sukses akan menandai invoice paid dan dipakai untuk alur pembayaran operasional.
          </p>
        </CardHeader>
        <CardContent className="grid gap-4 px-8 pb-8 md:grid-cols-3">
          <Stat label="Total" value={summary.total} />
          <Stat label="Success" value={summary.success} />
          <Stat label="Failed" value={summary.failed} />
        </CardContent>
      </Card>

      <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
        <CardContent className="flex flex-col gap-3 p-6 md:flex-row md:items-center md:justify-between">
          <div>
            <p className="text-sm font-medium text-slate-700 dark:text-slate-300">Global search</p>
            <p className="text-sm text-slate-500 dark:text-slate-400">Cari invoice id, amount, atau internal transaction id.</p>
          </div>
          <Input value={query} onChange={(e) => setQuery(e.target.value)} className="md:max-w-sm" placeholder="Cari payment" />
        </CardContent>
      </Card>

      <div className="grid gap-6 xl:grid-cols-[420px_1fr]">
        <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
          <CardHeader className="space-y-2 p-6">
            <CardTitle className="text-xl">{mode === "edit" ? "Edit Payment" : "Buat Payment"}</CardTitle>
            <CardDescription>Pilih invoice dan isi data pembayaran.</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4 px-6 pb-6">
            <form className="space-y-4" onSubmit={submit}>
              <select
                value={form.invoice_id}
                onChange={(e) => setForm((current) => ({ ...current, invoice_id: e.target.value }))}
                className="h-9 w-full rounded-md border border-input bg-background px-3 text-sm"
                required
              >
                <option value="">Pilih invoice</option>
                {invoices.map((item) => (
                  <option key={item.id} value={item.id}>
                    {item.id} - {item.member?.name} - {item.status}
                  </option>
                ))}
              </select>
              <Input value={form.internal_transaction_id} onChange={(e) => setForm((current) => ({ ...current, internal_transaction_id: e.target.value }))} placeholder="Internal transaction id" required />
              <Input value={form.amount} onChange={(e) => setForm((current) => ({ ...current, amount: e.target.value }))} type="number" min="0" placeholder="Amount" />
              <select value={form.status} onChange={(e) => setForm((current) => ({ ...current, status: e.target.value }))} className="h-9 w-full rounded-md border border-input bg-background px-3 text-sm">
                <option value="SUCCESS">SUCCESS</option>
                <option value="PENDING">PENDING</option>
                <option value="FAILED">FAILED</option>
              </select>
              <select value={form.scenario} onChange={(e) => setForm((current) => ({ ...current, scenario: e.target.value }))} className="h-9 w-full rounded-md border border-input bg-background px-3 text-sm">
                <option value="SUCCESS">SUCCESS</option>
                <option value="FAILURE">FAILURE</option>
                <option value="TIMEOUT">TIMEOUT</option>
              </select>
              <Input value={form.paid_at} onChange={(e) => setForm((current) => ({ ...current, paid_at: e.target.value }))} type="datetime-local" required />
              {error ? <p className="text-sm text-red-600">{error}</p> : null}
              <div className="flex flex-wrap gap-3">
                <Button type="submit" loading={busy === "form"}>
                  {mode === "edit" ? <PencilIcon className="size-4" /> : null}
                  {mode === "edit" ? "Simpan" : "Buat"}
                </Button>
                {mode === "edit" ? (
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => {
                      setMode("create")
                      setEditingId(null)
                    }}
                  >
                    Batal
                  </Button>
                ) : null}
              </div>
            </form>
          </CardContent>
        </Card>

        <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
          <CardHeader className="space-y-4 p-6">
            <div className="flex items-center justify-between gap-3">
              <div>
                <CardTitle className="text-xl">Daftar Payment</CardTitle>
                <CardDescription>Menunjukkan transaksi pembayaran invoice.</CardDescription>
              </div>
              <Button type="button" variant="outline" onClick={loadAll}>
                <RefreshCcwIcon className="size-4" />
                Refresh
              </Button>
            </div>
          </CardHeader>
          <CardContent className="space-y-4 px-6 pb-6">
            {loading ? <LoadingState rows={4} /> : null}
            {!loading && items.length === 0 ? (
              <div className="rounded-2xl border border-dashed border-slate-200 p-4 text-sm text-slate-500">Belum ada payment.</div>
            ) : null}
            {!loading && items.length > 0 ? (
              <div className="overflow-x-auto rounded-2xl border border-slate-200 dark:border-slate-700">
                <table className="min-w-full text-left text-sm">
                  <tbody className="divide-y divide-slate-200 bg-white dark:divide-slate-700 dark:bg-slate-900">
                    {items.map((item) => (
                      <tr key={item.id}>
                        <td className="px-3 py-3">
                          <div className="font-medium">{textOrDash(item.internal_transaction_id)}</div>
                          <div className="text-xs text-slate-500">
                            {textOrDash(item.status)} / {textOrDash(item.scenario)}
                          </div>
                        </td>
                        <td className="px-3 py-3">{textOrDash(item.invoice_id)}</td>
                        <td className="px-3 py-3">{item.amount ?? "-"}</td>
                        <td className="px-3 py-3">{textOrDash(item.paid_at)}</td>
                        <td className="px-3 py-3">
                          <div className="flex flex-wrap gap-2">
                            <Button type="button" variant="secondary" size="sm" onClick={() => setSelectedPayment(item)}>
                              Detail
                            </Button>
                            <Button type="button" variant="outline" size="sm" onClick={() => edit(item)}>
                              <PencilIcon className="size-4" />
                            </Button>
                            <Button type="button" variant="destructive" size="sm" loading={busy === item.id} onClick={() => void deletePayment(item.id).then(loadAll)}>
                              <Trash2Icon className="size-4" />
                            </Button>
                          </div>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            ) : null}
          </CardContent>
        </Card>
      </div>

      <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
        <CardHeader className="space-y-2 p-6">
          <CardTitle className="text-xl">Detail Payment</CardTitle>
          <CardDescription>Ringkasan transaksi terpilih dan kaitannya ke invoice.</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4 px-6 pb-6">
          {!selectedPayment ? (
            <div className="rounded-2xl border border-dashed border-slate-200 p-4 text-sm text-slate-500">
              Pilih payment dari tabel untuk melihat detail.
            </div>
          ) : (
            <div className="grid gap-4 md:grid-cols-2">
              <DetailBlock label="Payment ID" value={selectedPayment.id} />
              <DetailBlock label="Invoice ID" value={textOrDash(selectedPayment.invoice_id)} />
              <DetailBlock label="Internal transaction" value={textOrDash(selectedPayment.internal_transaction_id)} />
              <DetailBlock label="Amount" value={String(selectedPayment.amount ?? "-")} />
              <DetailBlock label="Status" value={textOrDash(selectedPayment.status)} />
              <DetailBlock label="Scenario" value={textOrDash(selectedPayment.scenario)} />
              <DetailBlock label="Paid at" value={textOrDash(selectedPayment.paid_at)} />
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  )
}

function Stat({ label, value }: { label: string; value: number }) {
  return (
    <div className="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
      <p className="text-sm text-slate-500 dark:text-slate-400">{label}</p>
      <p className="mt-2 text-3xl font-semibold">{value}</p>
    </div>
  )
}

function DetailBlock({ label, value }: { label: string; value: string }) {
  return (
    <div className="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
      <p className="text-sm text-slate-500 dark:text-slate-400">{label}</p>
      <p className="mt-2 break-words text-sm font-medium text-slate-950 dark:text-white">{value}</p>
    </div>
  )
}

function textOrDash(value: string | number | null | undefined) {
  if (typeof value === "number") {
    return Number.isFinite(value) ? String(value) : "-"
  }

  return value && value.trim() ? value : "-"
}
