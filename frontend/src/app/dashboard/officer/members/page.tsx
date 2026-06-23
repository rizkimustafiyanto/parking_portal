"use client"

import { useCallback, useEffect, useMemo, useState, type FormEvent } from "react"
import {
  BanknoteIcon,
  PencilIcon,
  RefreshCcwIcon,
  SearchIcon,
  Trash2Icon,
  UserPlusIcon,
} from "lucide-react"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { LoadingState } from "@/components/ui/loading-state"
import { createUser, deleteUser, fetchUsers, topUpBalance, updateUser, type UserRecord } from "@/features/users"

type FormState = {
  name: string
  email: string
  password: string
  role: string
}

const emptyForm: FormState = {
  name: "",
  email: "",
  password: "",
  role: "member",
}

const currency = new Intl.NumberFormat("id-ID", {
  style: "currency",
  currency: "IDR",
  maximumFractionDigits: 0,
})

export default function OfficerMembersPage() {
  const [users, setUsers] = useState<UserRecord[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [search, setSearch] = useState("")
  const [role, setRole] = useState("")
  const [form, setForm] = useState<FormState>(emptyForm)
  const [editingUser, setEditingUser] = useState<UserRecord | null>(null)
  const [topUpAmount, setTopUpAmount] = useState("50000")
  const [busyId, setBusyId] = useState<string | null>(null)
  const [submitMode, setSubmitMode] = useState<"create" | "edit">("create")

  const loadUsers = useCallback(async () => {
    setIsLoading(true)
    setError(null)

    try {
      const response = await fetchUsers({
        page: 1,
        limit: 50,
        search,
        role,
        sortBy: "created_at",
        order: "desc",
      })

      setUsers(response.data ?? [])
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal memuat user")
    } finally {
      setIsLoading(false)
    }
  }, [role, search])

  useEffect(() => {
    const timeout = window.setTimeout(() => {
      void loadUsers()
    }, 250)

    return () => window.clearTimeout(timeout)
  }, [loadUsers])

  const summary = useMemo(
    () => ({
      total: users.length,
      officers: users.filter((item) => item.role === "officer").length,
      members: users.filter((item) => item.role === "member").length,
    }),
    [users]
  )

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setBusyId("form")
    setError(null)

    try {
      if (submitMode === "edit" && editingUser) {
        await updateUser(editingUser.id, {
          name: form.name,
          email: form.email,
          password: form.password || undefined,
          role: form.role,
        })
      } else {
        await createUser(form)
      }

      setForm(emptyForm)
      setEditingUser(null)
      setSubmitMode("create")
      await loadUsers()
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal menyimpan data")
    } finally {
      setBusyId(null)
    }
  }

  function startEdit(user: UserRecord) {
    setEditingUser(user)
    setSubmitMode("edit")
    setForm({
      name: user.name,
      email: user.email,
      password: "",
      role: user.role || "member",
    })
  }

  function resetForm() {
    setEditingUser(null)
    setSubmitMode("create")
    setForm(emptyForm)
  }

  async function handleDelete(user: UserRecord) {
    const confirmed = window.confirm(`Hapus user ${user.name}?`)
    if (!confirmed) {
      return
    }

    setBusyId(user.id)
    setError(null)

    try {
      await deleteUser(user.id)
      await loadUsers()
      if (editingUser?.id === user.id) {
        resetForm()
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal menghapus user")
    } finally {
      setBusyId(null)
    }
  }

  async function handleTopUp(user: UserRecord) {
    const amount = Number(topUpAmount)
    if (!Number.isFinite(amount) || amount <= 0) {
      setError("Nominal top up harus lebih besar dari 0")
      return
    }

    setBusyId(user.id)
    setError(null)

    try {
      await topUpBalance(user.id, { amount })
      await loadUsers()
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal top up saldo")
    } finally {
      setBusyId(null)
    }
  }

  return (
    <div className="space-y-6">
      <Card className="overflow-hidden rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
        <div className="h-2 bg-gradient-to-r from-slate-950 via-slate-800 to-slate-700" />
        <CardHeader className="space-y-3 p-8">
          <CardDescription>Officer / Admin</CardDescription>
          <CardTitle className="text-2xl">User CRUD</CardTitle>
          <p className="max-w-3xl text-sm leading-6 text-slate-600 dark:text-slate-300">
            Kelola user, role, password, balance, dan aksi operasional lain dari satu layar.
          </p>
        </CardHeader>
        <CardContent className="grid gap-4 px-8 pb-8 md:grid-cols-3">
          <div className="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
            <p className="text-sm text-slate-500 dark:text-slate-400">Total user</p>
            <p className="mt-2 text-3xl font-semibold text-slate-950 dark:text-white">{summary.total}</p>
          </div>
          <div className="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
            <p className="text-sm text-slate-500 dark:text-slate-400">Officer</p>
            <p className="mt-2 text-3xl font-semibold text-slate-950 dark:text-white">{summary.officers}</p>
          </div>
          <div className="rounded-2xl border border-slate-200 bg-slate-50 p-4 dark:border-slate-700 dark:bg-slate-800/60">
            <p className="text-sm text-slate-500 dark:text-slate-400">Member</p>
            <p className="mt-2 text-3xl font-semibold text-slate-950 dark:text-white">{summary.members}</p>
          </div>
        </CardContent>
      </Card>

      <div className="grid gap-6 xl:grid-cols-[420px_1fr]">
        <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
          <CardHeader className="space-y-2 p-6">
            <CardTitle className="text-xl">{submitMode === "edit" ? "Edit User" : "Tambah User"}</CardTitle>
            <CardDescription>Isi data dasar user untuk backend `POST /users` dan `PUT /users/:id`.</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4 px-6 pb-6">
            <form className="space-y-4" onSubmit={handleSubmit}>
              <div className="space-y-2">
                <label className="text-sm font-medium text-slate-700 dark:text-slate-300" htmlFor="name">
                  Nama
                </label>
                <Input
                  id="name"
                  value={form.name}
                  onChange={(event) => setForm((current) => ({ ...current, name: event.target.value }))}
                  placeholder="Nama lengkap"
                  required
                />
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium text-slate-700 dark:text-slate-300" htmlFor="email">
                  Email
                </label>
                <Input
                  id="email"
                  type="email"
                  value={form.email}
                  onChange={(event) => setForm((current) => ({ ...current, email: event.target.value }))}
                  placeholder="user@example.com"
                  required
                />
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium text-slate-700 dark:text-slate-300" htmlFor="password">
                  Password {submitMode === "edit" ? "(opsional)" : ""}
                </label>
                <Input
                  id="password"
                  type="password"
                  value={form.password}
                  onChange={(event) => setForm((current) => ({ ...current, password: event.target.value }))}
                  placeholder={submitMode === "edit" ? "Kosongkan jika tidak diubah" : "Password baru"}
                  minLength={submitMode === "edit" ? undefined : 8}
                  required={submitMode === "create"}
                />
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium text-slate-700 dark:text-slate-300" htmlFor="role">
                  Role
                </label>
                <select
                  id="role"
                  value={form.role}
                  onChange={(event) => setForm((current) => ({ ...current, role: event.target.value }))}
                  className="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-3 focus-visible:ring-ring/50"
                >
                  <option value="member">Member</option>
                  <option value="officer">Officer</option>
                </select>
              </div>

              {error ? <p className="text-sm text-red-600">{error}</p> : null}

              <div className="flex flex-wrap gap-3">
                <Button type="submit" loading={busyId === "form"}>
                  {submitMode === "edit" ? <PencilIcon className="size-4" /> : <UserPlusIcon className="size-4" />}
                  {submitMode === "edit" ? "Simpan Perubahan" : "Buat User"}
                </Button>
                {submitMode === "edit" ? (
                  <Button type="button" variant="outline" onClick={resetForm}>
                    Batal
                  </Button>
                ) : null}
              </div>
            </form>

            <div className="rounded-2xl border border-dashed border-slate-200 bg-slate-50 p-4 text-sm text-slate-600 dark:border-slate-700 dark:bg-slate-800/50 dark:text-slate-300">
              Edit user dengan memilih tombol <strong>Edit</strong> di tabel kanan. Password hanya dikirim bila diisi.
            </div>
          </CardContent>
        </Card>

        <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
          <CardHeader className="space-y-4 p-6">
            <div className="flex flex-wrap items-center justify-between gap-3">
              <div>
                <CardTitle className="text-xl">Daftar User</CardTitle>
                <CardDescription>Filter, cari, edit, hapus, dan top up balance dari sini.</CardDescription>
              </div>
              <Button type="button" variant="outline" onClick={loadUsers}>
                <RefreshCcwIcon className="size-4" />
                Refresh
              </Button>
            </div>

            <div className="grid gap-3 md:grid-cols-[1fr_180px]">
              <div className="relative">
                <SearchIcon className="pointer-events-none absolute left-3 top-1/2 size-4 -translate-y-1/2 text-slate-400" />
                <Input
                  className="pl-9"
                  value={search}
                  onChange={(event) => setSearch(event.target.value)}
                  placeholder="Cari nama atau email"
                />
              </div>
              <select
                value={role}
                onChange={(event) => setRole(event.target.value)}
                className="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-3 focus-visible:ring-ring/50"
              >
                <option value="">Semua role</option>
                <option value="member">Member</option>
                <option value="officer">Officer</option>
              </select>
            </div>
          </CardHeader>

          <CardContent className="space-y-4 px-6 pb-6">
            {isLoading ? <LoadingState rows={4} /> : null}

            {!isLoading && users.length === 0 ? (
              <div className="rounded-2xl border border-dashed border-slate-200 bg-slate-50 p-6 text-sm text-slate-500 dark:border-slate-700 dark:bg-slate-800/50 dark:text-slate-300">
                Tidak ada user yang cocok dengan filter saat ini.
              </div>
            ) : null}

            {!isLoading && users.length > 0 ? (
              <div className="overflow-hidden rounded-2xl border border-slate-200 dark:border-slate-700">
                <div className="overflow-x-auto">
                  <table className="min-w-full divide-y divide-slate-200 text-left text-sm dark:divide-slate-700">
                    <thead className="bg-slate-50 text-slate-600 dark:bg-slate-800/60 dark:text-slate-300">
                      <tr>
                        <th className="px-4 py-3 font-medium">Nama</th>
                        <th className="px-4 py-3 font-medium">Email</th>
                        <th className="px-4 py-3 font-medium">Role</th>
                        <th className="px-4 py-3 font-medium">Balance</th>
                        <th className="px-4 py-3 font-medium">Aksi</th>
                      </tr>
                    </thead>
                    <tbody className="divide-y divide-slate-200 bg-white dark:divide-slate-700 dark:bg-slate-900">
                      {users.map((user) => (
                        <tr key={user.id} className="align-top">
                          <td className="px-4 py-4">
                            <div className="font-medium text-slate-950 dark:text-white">{user.name}</div>
                            <div className="text-xs text-slate-500 dark:text-slate-400">{user.id}</div>
                          </td>
                          <td className="px-4 py-4 text-slate-600 dark:text-slate-300">{user.email}</td>
                          <td className="px-4 py-4 capitalize text-slate-600 dark:text-slate-300">{user.role}</td>
                          <td className="px-4 py-4 text-slate-600 dark:text-slate-300">
                            {currency.format(user.balance ?? 0)}
                          </td>
                          <td className="px-4 py-4">
                            <div className="flex flex-wrap gap-2">
                              <Button type="button" variant="outline" size="sm" onClick={() => startEdit(user)}>
                                <PencilIcon className="size-4" />
                                Edit
                              </Button>
                              <Button
                                type="button"
                                variant="destructive"
                                size="sm"
                                loading={busyId === user.id}
                                onClick={() => handleDelete(user)}
                              >
                                <Trash2Icon className="size-4" />
                                Hapus
                              </Button>
                            </div>

                            <div className="mt-3 flex flex-wrap items-center gap-2">
                              <Input
                                value={topUpAmount}
                                onChange={(event) => setTopUpAmount(event.target.value)}
                                className="w-32"
                                type="number"
                                min="1"
                                step="1000"
                              />
                              <Button
                                type="button"
                                variant="secondary"
                                size="sm"
                                loading={busyId === user.id}
                                onClick={() => handleTopUp(user)}
                              >
                                <BanknoteIcon className="size-4" />
                                Top Up
                              </Button>
                            </div>
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            ) : null}
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
