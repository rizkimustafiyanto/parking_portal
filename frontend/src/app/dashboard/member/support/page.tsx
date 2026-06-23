"use client"

import Link from "next/link"
import { HelpCircleIcon, MessageCircleIcon, ReceiptTextIcon } from "lucide-react"

import { buttonVariants } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { cn } from "@/lib/utils"

export default function MemberSupportPage() {
  return (
    <Card className="rounded-3xl border-emerald-100 bg-white/95 shadow-lg dark:border-emerald-900/30 dark:bg-slate-900/95">
      <CardHeader className="space-y-3 p-8">
        <CardDescription>Member / User</CardDescription>
        <CardTitle className="text-2xl">Support & Help</CardTitle>
        <p className="text-sm leading-6 text-slate-600 dark:text-slate-300">
          Butuh bantuan soal invoice, payment, atau status pelanggaran? Mulai dari sini.
        </p>
      </CardHeader>
      <CardContent className="grid gap-4 px-8 pb-8 md:grid-cols-3">
        <SupportCard
          icon={ReceiptTextIcon}
          title="Cek Invoice"
          description="Buka history invoice untuk melihat detail tagihan terbaru."
          href="/dashboard/member/history"
        />
        <SupportCard
          icon={MessageCircleIcon}
          title="Payment Updates"
          description="Lihat notifikasi pembayaran dan status transaksi."
          href="/dashboard/member/notifications"
        />
        <SupportCard
          icon={HelpCircleIcon}
          title="Butuh bantuan"
          description="Jika ada data yang terasa tidak sesuai, hubungi officer yang bertugas."
          href="/dashboard/member"
        />
      </CardContent>
    </Card>
  )
}

function SupportCard({
  icon: Icon,
  title,
  description,
  href,
}: {
  icon: React.ComponentType<{ className?: string }>
  title: string
  description: string
  href: string
}) {
  return (
    <div className="rounded-2xl border border-emerald-100 bg-emerald-50/60 p-5 dark:border-emerald-900/30 dark:bg-slate-800/60">
      <div className="flex items-center gap-3">
        <div className="flex size-10 items-center justify-center rounded-2xl bg-white text-emerald-700 dark:bg-slate-900 dark:text-emerald-300">
          <Icon className="size-5" />
        </div>
        <div>
          <h3 className="font-semibold text-slate-950 dark:text-white">{title}</h3>
          <p className="text-sm text-slate-500 dark:text-slate-400">{description}</p>
        </div>
      </div>
      <Link href={href} className={cn(buttonVariants({ variant: "outline", size: "sm" }), "mt-4 rounded-full")}>
        Buka
      </Link>
    </div>
  )
}
