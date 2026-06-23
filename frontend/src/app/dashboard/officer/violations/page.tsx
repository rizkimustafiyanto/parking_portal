"use client"

import { useCallback, useEffect, useMemo, useState } from "react"
import {
  AlertTriangleIcon,
  FileTextIcon,
  RefreshCcwIcon,
  ShieldAlertIcon,
  ShieldIcon,
  Trash2Icon,
} from "lucide-react"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { LoadingState } from "@/components/ui/loading-state"
import { getStoredUserId } from "@/features/auth"
import {
  createFineRuleDetail,
  createFineRuleVersion,
  createViolation,
  createViolationType,
  deleteFineRuleDetail,
  deleteFineRuleVersion,
  deleteViolation,
  deleteViolationType,
  fetchFineRuleDetails,
  fetchFineRuleVersions,
  fetchViolationTypes,
  fetchViolations,
  type FineRuleDetailRecord,
  type FineRuleVersionRecord,
  type ViolationRecord,
  type ViolationTypeRecord,
} from "@/features/violation"

type SectionState<T> = {
  items: T[]
  loading: boolean
  error: string | null
}

const money = new Intl.NumberFormat("id-ID", {
  style: "currency",
  currency: "IDR",
  maximumFractionDigits: 0,
})

export default function OfficerViolationsPage() {
  const userId = getStoredUserId()
  const [violations, setViolations] = useState<SectionState<ViolationRecord>>({ items: [], loading: true, error: null })
  const [violationTypes, setViolationTypes] = useState<SectionState<ViolationTypeRecord>>({
    items: [],
    loading: true,
    error: null,
  })
  const [ruleVersions, setRuleVersions] = useState<SectionState<FineRuleVersionRecord>>({
    items: [],
    loading: true,
    error: null,
  })
  const [ruleDetails, setRuleDetails] = useState<SectionState<FineRuleDetailRecord>>({
    items: [],
    loading: true,
    error: null,
  })
  const [query, setQuery] = useState("")
  const [busy, setBusy] = useState<string | null>(null)
  const [formError, setFormError] = useState<string | null>(null)

  const [violationTypeForm, setViolationTypeForm] = useState({
    code: "",
    name: "",
    base_amount: "50000",
  })
  const [ruleVersionForm, setRuleVersionForm] = useState({
    version_number: "1",
    is_active: true,
  })
  const [ruleDetailForm, setRuleDetailForm] = useState({
    rule_version_id: "",
    rule_type: "penalty",
    key: "",
    value: "",
  })
  const [violationForm, setViolationForm] = useState({
    plate_number: "",
    violation_type_code: "",
    location: "",
    occurred_at: "",
    photo_url: "",
  })

  const loadAll = useCallback(async () => {
    setFormError(null)

    const [vt, rv, rd, v] = await Promise.allSettled([
      fetchViolationTypes({ page: 1, limit: 50, search: query }),
      fetchFineRuleVersions({ page: 1, limit: 50, search: query }),
      fetchFineRuleDetails({ page: 1, limit: 50, search: query }),
      fetchViolations({ page: 1, limit: 50, search: query }),
    ])

    setViolationTypes({
      loading: false,
      error: vt.status === "rejected" ? vt.reason.message : null,
      items: vt.status === "fulfilled" ? vt.value.data ?? [] : [],
    })
    setRuleVersions({
      loading: false,
      error: rv.status === "rejected" ? rv.reason.message : null,
      items: rv.status === "fulfilled" ? rv.value.data ?? [] : [],
    })
    setRuleDetails({
      loading: false,
      error: rd.status === "rejected" ? rd.reason.message : null,
      items: rd.status === "fulfilled" ? rd.value.data ?? [] : [],
    })
    setViolations({
      loading: false,
      error: v.status === "rejected" ? v.reason.message : null,
      items: v.status === "fulfilled" ? v.value.data ?? [] : [],
    })
  }, [query])

  useEffect(() => {
    let active = true

    void (async () => {
      const [vt, rv, rd, v] = await Promise.allSettled([
        fetchViolationTypes({ page: 1, limit: 50, search: query }),
        fetchFineRuleVersions({ page: 1, limit: 50, search: query }),
        fetchFineRuleDetails({ page: 1, limit: 50, search: query }),
        fetchViolations({ page: 1, limit: 50, search: query }),
      ])

      if (!active) {
        return
      }

      setViolationTypes({
        loading: false,
        error: vt.status === "rejected" ? vt.reason.message : null,
        items: vt.status === "fulfilled" ? vt.value.data ?? [] : [],
      })
      setRuleVersions({
        loading: false,
        error: rv.status === "rejected" ? rv.reason.message : null,
        items: rv.status === "fulfilled" ? rv.value.data ?? [] : [],
      })
      setRuleDetails({
        loading: false,
        error: rd.status === "rejected" ? rd.reason.message : null,
        items: rd.status === "fulfilled" ? rd.value.data ?? [] : [],
      })
      setViolations({
        loading: false,
        error: v.status === "rejected" ? v.reason.message : null,
        items: v.status === "fulfilled" ? v.value.data ?? [] : [],
      })
    })()

    return () => {
      active = false
    }
  }, [loadAll, query])

  const summary = useMemo(
    () => ({
      violations: violations.items.length,
      violationTypes: violationTypes.items.length,
      versions: ruleVersions.items.length,
      details: ruleDetails.items.length,
    }),
    [ruleDetails.items.length, ruleVersions.items.length, violationTypes.items.length, violations.items.length]
  )

  async function submitAndReload(action: string, task: () => Promise<void>) {
    setBusy(action)
    setFormError(null)
    try {
      await task()
      await loadAll()
    } catch (err) {
      setFormError(err instanceof Error ? err.message : "Terjadi kesalahan")
    } finally {
      setBusy(null)
    }
  }

  function requireUserId() {
    if (!userId) {
      throw new Error("User ID tidak ditemukan di token, silakan login ulang")
    }
    return userId
  }

  return (
    <div className="space-y-6">
      <Card className="overflow-hidden rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
        <div className="h-2 bg-gradient-to-r from-slate-950 via-slate-800 to-slate-700" />
        <CardHeader className="space-y-3 p-8">
          <CardDescription>Officer / Admin</CardDescription>
          <CardTitle className="text-2xl">Violation Flow</CardTitle>
          <p className="max-w-3xl text-sm leading-6 text-slate-600 dark:text-slate-300">
            Master data dipisah dari transaksi utama. `Violation Types`, `Fine Rule Versions`, dan `Fine Rule Details`
            berfungsi sebagai konfigurasi, sedangkan `Violations` adalah flow operasional.
          </p>
        </CardHeader>
        <CardContent className="grid gap-4 px-8 pb-8 md:grid-cols-4">
          <Stat label="Violations" value={summary.violations} icon={ShieldAlertIcon} />
          <Stat label="Violation Types" value={summary.violationTypes} icon={AlertTriangleIcon} />
          <Stat label="Fine Rule Versions" value={summary.versions} icon={FileTextIcon} />
          <Stat label="Fine Rule Details" value={summary.details} icon={ShieldIcon} />
        </CardContent>
      </Card>

      <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
        <CardContent className="flex flex-col gap-3 p-6 md:flex-row md:items-center md:justify-between">
          <div>
            <p className="text-sm font-medium text-slate-700 dark:text-slate-300">Global search</p>
            <p className="text-sm text-slate-500 dark:text-slate-400">Cari kode, nama, plate number, atau key detail.</p>
          </div>
          <Input
            value={query}
            onChange={(event) => setQuery(event.target.value)}
            className="md:max-w-sm"
            placeholder="Cari data violation"
          />
        </CardContent>
      </Card>

      {formError ? (
        <div className="rounded-2xl border border-red-200 bg-red-50 p-4 text-sm text-red-700 dark:border-red-900/40 dark:bg-red-950/40 dark:text-red-200">
          {formError}
        </div>
      ) : null}

      <div className="grid gap-6 xl:grid-cols-2">
        <MasterCard
          title="Violation Types"
          description="Master data untuk jenis pelanggaran."
          onRefresh={loadAll}
          loading={violationTypes.loading}
          error={violationTypes.error}
        >
          <form
            className="grid gap-3"
            onSubmit={(event) => {
              event.preventDefault()
              void submitAndReload("create-violation-type", async () => {
                await createViolationType({
                  ...violationTypeForm,
                  base_amount: Number(violationTypeForm.base_amount),
                  created_by_id: requireUserId(),
                })
                setViolationTypeForm({ code: "", name: "", base_amount: "50000" })
              })
            }}
          >
            <Input
              value={violationTypeForm.code}
              onChange={(event) => setViolationTypeForm((current) => ({ ...current, code: event.target.value }))}
              placeholder="Code"
              required
            />
            <Input
              value={violationTypeForm.name}
              onChange={(event) => setViolationTypeForm((current) => ({ ...current, name: event.target.value }))}
              placeholder="Nama"
              required
            />
            <Input
              value={violationTypeForm.base_amount}
              onChange={(event) => setViolationTypeForm((current) => ({ ...current, base_amount: event.target.value }))}
              type="number"
              min="0"
              placeholder="Base amount"
            />
            <Button type="submit" loading={busy === "create-violation-type"}>
              Simpan Type
            </Button>
          </form>

          <ListTable
            items={violationTypes.items}
            renderRow={(item) => (
              <>
                <td className="px-3 py-3 font-medium">{item.code}</td>
                <td className="px-3 py-3">{item.name}</td>
                <td className="px-3 py-3">{money.format(item.base_amount)}</td>
                <td className="px-3 py-3">
                  <Button type="button" variant="destructive" size="sm" loading={busy === item.id} onClick={() => void submitAndReload(item.id, () => deleteViolationType(item.id))}>
                    <Trash2Icon className="size-4" />
                  </Button>
                </td>
              </>
            )}
          />
        </MasterCard>

        <MasterCard
          title="Fine Rule Versions"
          description="Version master yang dipakai saat pelanggaran dibuat."
          onRefresh={loadAll}
          loading={ruleVersions.loading}
          error={ruleVersions.error}
        >
          <form
            className="grid gap-3"
            onSubmit={(event) => {
              event.preventDefault()
              void submitAndReload("create-rule-version", async () => {
                await createFineRuleVersion({
                  version_number: Number(ruleVersionForm.version_number),
                  is_active: ruleVersionForm.is_active,
                  published_by: requireUserId(),
                })
                setRuleVersionForm({ version_number: "1", is_active: true })
              })
            }}
          >
            <Input
              value={ruleVersionForm.version_number}
              onChange={(event) => setRuleVersionForm((current) => ({ ...current, version_number: event.target.value }))}
              type="number"
              min="1"
              placeholder="Version number"
            />
            <label className="flex items-center gap-2 text-sm text-slate-600 dark:text-slate-300">
              <input
                type="checkbox"
                checked={ruleVersionForm.is_active}
                onChange={(event) => setRuleVersionForm((current) => ({ ...current, is_active: event.target.checked }))}
              />
              Active
            </label>
            <Button type="submit" loading={busy === "create-rule-version"}>
              Simpan Version
            </Button>
          </form>
          <ListTable
            items={ruleVersions.items}
            renderRow={(item) => (
              <>
                <td className="px-3 py-3 font-medium">{item.version_number}</td>
                <td className="px-3 py-3">{item.is_active ? "Aktif" : "Nonaktif"}</td>
                <td className="px-3 py-3">{item.publish?.name}</td>
                <td className="px-3 py-3">
                  <Button type="button" variant="destructive" size="sm" loading={busy === item.id} onClick={() => void submitAndReload(item.id, () => deleteFineRuleVersion(item.id))}>
                    <Trash2Icon className="size-4" />
                  </Button>
                </td>
              </>
            )}
          />
        </MasterCard>

        <MasterCard
          title="Fine Rule Details"
          description="Detail key-value untuk tiap versi aturan."
          onRefresh={loadAll}
          loading={ruleDetails.loading}
          error={ruleDetails.error}
        >
          <form
            className="grid gap-3"
            onSubmit={(event) => {
              event.preventDefault()
              void submitAndReload("create-rule-detail", async () => {
                await createFineRuleDetail(ruleDetailForm)
                setRuleDetailForm({ rule_version_id: "", rule_type: "penalty", key: "", value: "" })
              })
            }}
          >
            <select
              value={ruleDetailForm.rule_version_id}
              onChange={(event) => setRuleDetailForm((current) => ({ ...current, rule_version_id: event.target.value }))}
              className="h-9 rounded-md border border-input bg-background px-3 text-sm"
              required
            >
              <option value="">Pilih version</option>
              {ruleVersions.items.map((version) => (
                <option key={version.id} value={version.id}>
                  {version.version_number}
                </option>
              ))}
            </select>
            <Input value={ruleDetailForm.rule_type} onChange={(event) => setRuleDetailForm((current) => ({ ...current, rule_type: event.target.value }))} placeholder="Rule type" />
            <Input value={ruleDetailForm.key} onChange={(event) => setRuleDetailForm((current) => ({ ...current, key: event.target.value }))} placeholder="Key" />
            <Input value={ruleDetailForm.value} onChange={(event) => setRuleDetailForm((current) => ({ ...current, value: event.target.value }))} placeholder="Value" />
            <Button type="submit" loading={busy === "create-rule-detail"}>
              Simpan Detail
            </Button>
          </form>
          <ListTable
            items={ruleDetails.items}
            renderRow={(item) => (
              <>
                <td className="px-3 py-3 font-medium">{item.rule_version_id}</td>
                <td className="px-3 py-3">{item.rule_type}</td>
                <td className="px-3 py-3">{item.key}</td>
                <td className="px-3 py-3">{item.value}</td>
                <td className="px-3 py-3">
                  <Button type="button" variant="destructive" size="sm" loading={busy === item.id} onClick={() => void submitAndReload(item.id, () => deleteFineRuleDetail(item.id))}>
                    <Trash2Icon className="size-4" />
                  </Button>
                </td>
              </>
            )}
          />
        </MasterCard>

        <MasterCard
          title="Violations"
          description="Transaksi pelanggaran yang memakai data master di atas."
          onRefresh={loadAll}
          loading={violations.loading}
          error={violations.error}
        >
          <form
            className="grid gap-3"
            onSubmit={(event) => {
              event.preventDefault()
              void submitAndReload("create-violation", async () => {
                await createViolation({
                  ...violationForm,
                  officer_id: requireUserId(),
                  occurred_at: new Date(violationForm.occurred_at).toISOString(),
                })
                setViolationForm({
                  plate_number: "",
                  violation_type_code: "",
                  location: "",
                  occurred_at: "",
                  photo_url: "",
                })
              })
            }}
          >
            <Input value={violationForm.plate_number} onChange={(event) => setViolationForm((current) => ({ ...current, plate_number: event.target.value }))} placeholder="Plate number" required />
            <select
              value={violationForm.violation_type_code}
              onChange={(event) => setViolationForm((current) => ({ ...current, violation_type_code: event.target.value }))}
              className="h-9 rounded-md border border-input bg-background px-3 text-sm"
              required
            >
              <option value="">Pilih violation type</option>
              {violationTypes.items.map((type) => (
                <option key={type.id} value={type.code}>
                  {type.code} - {type.name}
                </option>
              ))}
            </select>
            <Input value={violationForm.location} onChange={(event) => setViolationForm((current) => ({ ...current, location: event.target.value }))} placeholder="Location" required />
            <Input value={violationForm.occurred_at} onChange={(event) => setViolationForm((current) => ({ ...current, occurred_at: event.target.value }))} type="datetime-local" required />
            <Input value={violationForm.photo_url} onChange={(event) => setViolationForm((current) => ({ ...current, photo_url: event.target.value }))} placeholder="Photo URL" />
            <Button type="submit" loading={busy === "create-violation"}>
              Simpan Violation
            </Button>
          </form>
          <ListTable
            items={violations.items}
            renderRow={(item) => (
              <>
                <td className="px-3 py-3 font-medium">{item.plate_number}</td>
                <td className="px-3 py-3">{item.violation_type_code}</td>
                <td className="px-3 py-3">{item.location}</td>
                <td className="px-3 py-3">{item.officer?.name}</td>
                <td className="px-3 py-3">
                  <Button type="button" variant="destructive" size="sm" loading={busy === item.id} onClick={() => void submitAndReload(item.id, () => deleteViolation(item.id))}>
                    <Trash2Icon className="size-4" />
                  </Button>
                </td>
              </>
            )}
          />
        </MasterCard>
      </div>
    </div>
  )
}

function Stat({
  label,
  value,
  icon: Icon,
}: {
  label: string
  value: number
  icon: React.ComponentType<{ className?: string }>
}) {
  return (
    <div className="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
      <div className="flex items-center justify-between">
        <p className="text-sm text-slate-500 dark:text-slate-400">{label}</p>
        <Icon className="size-4 text-slate-400" />
      </div>
      <p className="mt-2 text-3xl font-semibold text-slate-950 dark:text-white">{value}</p>
    </div>
  )
}

function MasterCard({
  title,
  description,
  loading,
  error,
  onRefresh,
  children,
}: {
  title: string
  description: string
  loading: boolean
  error: string | null
  onRefresh: () => void
  children: React.ReactNode
}) {
  return (
    <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
      <CardHeader className="space-y-3 p-6">
        <div className="flex items-center justify-between gap-3">
          <div>
            <CardTitle className="text-xl">{title}</CardTitle>
            <CardDescription>{description}</CardDescription>
          </div>
          <Button type="button" variant="outline" size="sm" onClick={onRefresh}>
            <RefreshCcwIcon className="size-4" />
            Refresh
          </Button>
        </div>
      </CardHeader>
      <CardContent className="space-y-4 px-6 pb-6">
        {loading ? <LoadingState rows={3} /> : null}
        {error ? <p className="text-sm text-red-600">{error}</p> : null}
        {!loading ? children : null}
      </CardContent>
    </Card>
  )
}

function ListTable<T>({
  items,
  renderRow,
}: {
  items: T[]
  renderRow: (item: T) => React.ReactNode
}) {
  if (!items.length) {
    return <div className="rounded-2xl border border-dashed border-slate-200 p-4 text-sm text-slate-500 dark:border-slate-700">Belum ada data.</div>
  }

  return (
    <div className="overflow-x-auto rounded-2xl border border-slate-200 dark:border-slate-700">
      <table className="min-w-full text-left text-sm">
        <tbody className="divide-y divide-slate-200 bg-white dark:divide-slate-700 dark:bg-slate-900">
          <tr className="bg-slate-50 text-xs uppercase tracking-wide text-slate-500 dark:bg-slate-800/60">
            <td className="px-3 py-3">Data</td>
            <td className="px-3 py-3">Data</td>
            <td className="px-3 py-3">Data</td>
            <td className="px-3 py-3">Data</td>
            <td className="px-3 py-3">Aksi</td>
          </tr>
          {items.map((item, index) => (
            <tr key={index}>{renderRow(item)}</tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
