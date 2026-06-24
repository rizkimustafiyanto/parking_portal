"use client"

import { useCallback, useEffect, useMemo, useState } from "react"
import { PencilIcon, RefreshCcwIcon, Trash2Icon } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { LoadingState } from "@/components/ui/loading-state"
import { fetchViolations, type ViolationRecord } from "@/features/violation"
import {
  createInvoice,
  deleteInvoice,
  fetchInvoiceDetail,
  fetchInvoices,
  updateInvoice,
  type InvoiceRecord,
} from "@/features/finance"
import { fetchUsers, type UserRecord } from "@/features/users"

type FormState = {
  violation_id: string
  member_id: string
  status: string
}

export default function OfficerInvoicesPage() {
  const [invoices, setInvoices] = useState<InvoiceRecord[]>([])
  const [violations, setViolations] = useState<ViolationRecord[]>([])
  const [members, setMembers] = useState<UserRecord[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [query, setQuery] = useState("")
  const [busy, setBusy] = useState<string | null>(null)
  const [mode, setMode] = useState<"create" | "edit">("create")
  const [editingId, setEditingId] = useState<string | null>(null)
  const [selectedInvoice, setSelectedInvoice] = useState<InvoiceRecord | null>(null)
  const [detailLoading, setDetailLoading] = useState(false)
  const [detailError, setDetailError] = useState<string | null>(null)
  const [form, setForm] = useState<FormState>({ violation_id: "", member_id: "", status: "PENDING" })

  const loadAll = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const [invoiceRes, violationRes, userRes] = await Promise.all([
        fetchInvoices({ page: 1, limit: 50, search: query }),
        fetchViolations({ page: 1, limit: 100 }),
        fetchUsers({ page: 1, limit: 100, role: "member" }),
      ])

      setInvoices(invoiceRes.data ?? [])
      setViolations(violationRes.data ?? [])
      setMembers(userRes.data ?? [])
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal memuat data invoice")
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
        const [invoiceRes, violationRes, userRes] = await Promise.all([
          fetchInvoices({ page: 1, limit: 50, search: query }),
          fetchViolations({ page: 1, limit: 100 }),
          fetchUsers({ page: 1, limit: 100, role: "member" }),
        ])

        if (!active) {
          return
        }

        setInvoices(invoiceRes.data ?? [])
        setViolations(violationRes.data ?? [])
        setMembers(userRes.data ?? [])
      } catch (err) {
        if (active) {
          setError(err instanceof Error ? err.message : "Gagal memuat data invoice")
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
      total: invoices.length,
      paid: invoices.filter((item) => item.status === "PAID").length,
      pending: invoices.filter((item) => item.status === "PENDING").length,
    }),
    [invoices]
  )

  async function submit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setBusy("form")
    setError(null)
    try {
      if (mode === "edit" && editingId) {
        await updateInvoice(editingId, {
          violation_id: form.violation_id,
          member_id: form.member_id,
          status: form.status,
        })
      } else {
        await createInvoice({
          violation_id: form.violation_id,
          member_id: form.member_id,
          status: form.status,
        })
      }

      setForm({ violation_id: "", member_id: "", status: "PENDING" })
      setMode("create")
      setEditingId(null)
      await loadAll()
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal menyimpan invoice")
    } finally {
      setBusy(null)
    }
  }

  function editInvoice(item: InvoiceRecord) {
    setMode("edit")
    setEditingId(item.id)
    setForm({
      violation_id: item.violation_id,
      member_id: item.member?.id ?? "",
      status: item.status,
    })
  }

  async function openDetail(invoiceId: string) {
    setDetailLoading(true)
    setDetailError(null)
    try {
      const detail = await fetchInvoiceDetail(invoiceId)
      setSelectedInvoice(detail)
    } catch (err) {
      setDetailError(err instanceof Error ? err.message : "Gagal memuat detail invoice")
    } finally {
      setDetailLoading(false)
    }
  }

  return (
    <div className="space-y-6">
      <Card className="overflow-hidden rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
        <div className="h-2 bg-gradient-to-r from-slate-950 via-slate-800 to-slate-700" />
        <CardHeader className="space-y-3 p-8">
          <CardDescription>Officer / Admin</CardDescription>
          <CardTitle className="text-2xl">Invoices</CardTitle>
          <p className="max-w-3xl text-sm leading-6 text-slate-600 dark:text-slate-300">
            Invoice dibuat dari violation, lalu jadi dasar payment transaction dan top up balance.
          </p>
        </CardHeader>
        <CardContent className="grid gap-4 px-8 pb-8 md:grid-cols-3">
          <div className="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
            <p className="text-sm text-slate-500 dark:text-slate-400">Total</p>
            <p className="mt-2 text-3xl font-semibold">{summary.total}</p>
          </div>
          <div className="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
            <p className="text-sm text-slate-500 dark:text-slate-400">Paid</p>
            <p className="mt-2 text-3xl font-semibold">{summary.paid}</p>
          </div>
          <div className="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
            <p className="text-sm text-slate-500 dark:text-slate-400">Pending</p>
            <p className="mt-2 text-3xl font-semibold">{summary.pending}</p>
          </div>
        </CardContent>
      </Card>

      <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
        <CardContent className="flex flex-col gap-3 p-6 md:flex-row md:items-center md:justify-between">
          <div>
            <p className="text-sm font-medium text-slate-700 dark:text-slate-300">Global search</p>
            <p className="text-sm text-slate-500 dark:text-slate-400">Cari invoice berdasarkan amount atau status.</p>
          </div>
          <Input value={query} onChange={(e) => setQuery(e.target.value)} className="md:max-w-sm" placeholder="Cari invoice" />
        </CardContent>
      </Card>

      <div className="grid gap-6 xl:grid-cols-[420px_1fr]">
        <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
          <CardHeader className="space-y-2 p-6">
            <CardTitle className="text-xl">{mode === "edit" ? "Edit Invoice" : "Buat Invoice"}</CardTitle>
            <CardDescription>Pilih violation dan member yang sesuai.</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4 px-6 pb-6">
            <form className="space-y-4" onSubmit={submit}>
              <select
                value={form.violation_id}
                onChange={(e) => setForm((current) => ({ ...current, violation_id: e.target.value }))}
                className="h-9 w-full rounded-md border border-input bg-background px-3 text-sm"
                required
              >
                <option value="">Pilih violation</option>
                {violations.map((item) => (
                  <option key={item.id} value={item.id}>
                    {item.plate_number} - {item.location}
                  </option>
                ))}
              </select>
              <select
                value={form.member_id}
                onChange={(e) => setForm((current) => ({ ...current, member_id: e.target.value }))}
                className="h-9 w-full rounded-md border border-input bg-background px-3 text-sm"
                required
              >
                <option value="">Pilih member</option>
                {members.map((item) => (
                  <option key={item.id} value={item.id}>
                    {item.name} - {item.email}
                  </option>
                ))}
              </select>
              <select
                value={form.status}
                onChange={(e) => setForm((current) => ({ ...current, status: e.target.value }))}
                className="h-9 w-full rounded-md border border-input bg-background px-3 text-sm"
              >
                <option value="PENDING">PENDING</option>
                <option value="PAID">PAID</option>
                <option value="OVERDUE">OVERDUE</option>
                <option value="CANCELED">CANCELED</option>
              </select>
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
                      setForm({ violation_id: "", member_id: "", status: "PENDING" })
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
                <CardTitle className="text-xl">Daftar Invoice</CardTitle>
                <CardDescription>Hubungan antara violation, member, dan payment.</CardDescription>
              </div>
              <Button type="button" variant="outline" onClick={loadAll}>
                <RefreshCcwIcon className="size-4" />
                Refresh
              </Button>
            </div>
          </CardHeader>
          <CardContent className="space-y-4 px-6 pb-6">
            {loading ? <LoadingState rows={4} /> : null}
            {!loading && invoices.length === 0 ? (
              <div className="rounded-2xl border border-dashed border-slate-200 p-4 text-sm text-slate-500">Belum ada invoice.</div>
            ) : null}
            {!loading && invoices.length > 0 ? (
              <div className="overflow-x-auto rounded-2xl border border-slate-200 dark:border-slate-700">
                <table className="min-w-full text-left text-sm">
                  <tbody className="divide-y divide-slate-200 bg-white dark:divide-slate-700 dark:bg-slate-900">
                    {invoices.map((item) => (
                        <tr key={item.id}>
                          <td className="px-3 py-3">
                          <div className="font-medium">{item.amount ?? "-"}</div>
                          <div className="text-xs text-slate-500">{textOrDash(item.status)}</div>
                          </td>
                        <td className="px-3 py-3">{textOrDash(item.member?.name)}</td>
                        <td className="px-3 py-3">{textOrDash(item.violation?.plate_number)}</td>
                        <td className="px-3 py-3">{textOrDash(item.payment?.status)}</td>
                        <td className="px-3 py-3">
                          <div className="flex flex-wrap gap-2">
                            <Button type="button" variant="secondary" size="sm" onClick={() => void openDetail(item.id)}>
                              Detail
                            </Button>
                            <Button type="button" variant="outline" size="sm" onClick={() => editInvoice(item)}>
                              <PencilIcon className="size-4" />
                            </Button>
                            <Button type="button" variant="destructive" size="sm" loading={busy === item.id} onClick={() => void deleteInvoice(item.id).then(loadAll)}>
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
          <CardTitle className="text-xl">Detail Invoice</CardTitle>
          <CardDescription>Relasi violation, member, payment, dan fine rule untuk invoice terpilih.</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4 px-6 pb-6">
          {detailLoading ? <LoadingState rows={3} /> : null}
          {detailError ? <p className="text-sm text-red-600">{detailError}</p> : null}
          {!detailLoading && !selectedInvoice ? (
            <div className="rounded-2xl border border-dashed border-slate-200 p-4 text-sm text-slate-500">
              Pilih invoice dari tabel untuk melihat detail lengkap.
            </div>
          ) : null}
          {selectedInvoice ? (
            <div className="grid gap-4 md:grid-cols-2">
              <DetailBlock label="Invoice ID" value={selectedInvoice.id} />
              <DetailBlock label="Status" value={textOrDash(selectedInvoice.status)} />
              <DetailBlock label="Member" value={selectedInvoice.member?.name ?? "-"} />
              <DetailBlock label="Plate number" value={selectedInvoice.violation?.plate_number ?? "-"} />
              <DetailBlock label="Violation location" value={selectedInvoice.violation?.location ?? "-"} />
              <DetailBlock label="Payment status" value={textOrDash(selectedInvoice.payment?.status)} />
              <DetailBlock label="Payment scenario" value={textOrDash(selectedInvoice.payment?.scenario)} />
              <DetailBlock label="Fine rule version" value={String(selectedInvoice.violation?.fine_rule_version?.version_number ?? "-")} />
              <div className="md:col-span-2 rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
                <p className="text-sm text-slate-500 dark:text-slate-400">Fine rule details</p>
                <div className="mt-3 space-y-2">
                  {selectedInvoice.violation?.fine_rule_version?.details?.length ? (
                    selectedInvoice.violation.fine_rule_version.details.map((detail) => (
                      <div key={detail.id} className="rounded-xl bg-white p-3 text-sm ring-1 ring-slate-200 dark:bg-slate-900 dark:ring-slate-700">
                        <div className="font-medium">{detail.key}</div>
                        <div className="text-slate-500 dark:text-slate-400">{detail.value}</div>
                      </div>
                    ))
                  ) : (
                    <p className="text-sm text-slate-500">Tidak ada detail rule.</p>
                  )}
                </div>
              </div>
            </div>
          ) : null}
        </CardContent>
      </Card>
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
