"use client"

import { useEffect, useMemo, useState } from "react"
import { UserRoundIcon } from "lucide-react"

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { LoadingState } from "@/components/ui/loading-state"
import { getStoredUserId } from "@/features/auth"
import { fetchUsers, type UserRecord } from "@/features/users"

export default function MemberProfilePage() {
  const memberId = getStoredUserId()
  const [profile, setProfile] = useState<UserRecord | null>(null)
  const [loading, setLoading] = useState(() => Boolean(memberId))
  const [error, setError] = useState<string | null>(
    memberId ? null : "User ID tidak ditemukan di token. Silakan login ulang.",
  )

  useEffect(() => {
    if (!memberId) {
      return
    }

    let active = true

    void (async () => {
      setLoading(true)
      setError(null)

      try {
        const response = await fetchUsers({ page: 1, limit: 200, search: memberId })
        const items = response.data ?? []
        const found = items.find((item) => item.id === memberId)

        if (!active) {
          return
        }

        setProfile(found ?? null)
        if (!found) {
          setError("Profile member tidak ditemukan.")
        }
      } catch (err) {
        if (active) {
          setError(err instanceof Error ? err.message : "Gagal memuat profile member")
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

  const rows = useMemo(
    () => [
      { label: "User ID", value: profile?.id ?? memberId ?? "-" },
      { label: "Nama", value: textOrDash(profile?.name) },
      { label: "Email", value: textOrDash(profile?.email) },
      { label: "Role", value: textOrDash(profile?.role) },
      { label: "Balance", value: typeof profile?.balance === "number" ? String(profile.balance) : "-" },
    ],
    [memberId, profile]
  )

  return (
    <Card className="rounded-3xl border-emerald-100 bg-white/95 shadow-lg dark:border-emerald-900/30 dark:bg-slate-900/95">
      <CardHeader className="space-y-3 p-8">
        <CardDescription>Member / User</CardDescription>
        <CardTitle className="text-2xl">Profile Saya</CardTitle>
        <p className="text-sm leading-6 text-slate-600 dark:text-slate-300">
          Informasi akun member yang sedang login.
        </p>
      </CardHeader>
      <CardContent className="space-y-4 px-8 pb-8">
        {loading ? <LoadingState rows={3} /> : null}
        {error ? <p className="text-sm text-red-600">{error}</p> : null}
        {!loading && !error ? (
          <div className="grid gap-4 md:grid-cols-2">
            {rows.map((row) => (
              <div key={row.label} className="rounded-2xl border border-emerald-100 bg-emerald-50/60 p-4 dark:border-emerald-900/30 dark:bg-slate-800/60">
                <div className="flex items-center justify-between gap-3">
                  <div>
                    <p className="text-sm text-slate-500 dark:text-slate-400">{row.label}</p>
                    <p className="mt-2 break-words text-sm font-medium text-slate-950 dark:text-white">{row.value}</p>
                  </div>
                  <UserRoundIcon className="size-5 text-emerald-600 dark:text-emerald-300" />
                </div>
              </div>
            ))}
          </div>
        ) : null}
      </CardContent>
    </Card>
  )
}

function textOrDash(value: string | null | undefined) {
  return value && value.trim() ? value : "-"
}
