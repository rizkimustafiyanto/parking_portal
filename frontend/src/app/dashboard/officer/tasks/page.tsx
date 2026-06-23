import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"

export default function OfficerTasksPage() {
  return (
    <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
      <CardHeader className="p-8">
        <CardDescription>Officer / Admin</CardDescription>
        <CardTitle className="text-2xl">Tasks</CardTitle>
        <p className="text-sm leading-6 text-slate-600 dark:text-slate-300">
          Halaman ini siap diisi daftar tugas, antrian verifikasi, dan action center officer.
        </p>
      </CardHeader>
    </Card>
  )
}
