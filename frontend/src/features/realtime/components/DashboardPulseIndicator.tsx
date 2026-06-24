"use client"

import { RefreshCwIcon } from "lucide-react"

import type { DashboardRole } from "../types"
import { useDashboardPulse } from "../hooks/useDashboardPulse"

type DashboardPulseIndicatorProps = {
  role: DashboardRole
}

function DashboardPulseIndicator({ role }: DashboardPulseIndicatorProps) {
  const { snapshot, isRefreshing } = useDashboardPulse(role)

  const label = `${snapshot?.invoiceCount ?? 0} invoices · ${snapshot?.paymentCount ?? 0} payments`

  return (
    <div className="flex flex-wrap items-center gap-2">
      <span className="rounded-full border border-slate-200 bg-white/80 px-3 py-1 text-xs font-medium text-slate-600 dark:border-slate-700 dark:bg-slate-900/80 dark:text-slate-300">
        {label}
      </span>
      <span className="inline-flex items-center rounded-full border border-slate-200 bg-white/80 px-3 py-1 text-xs font-medium text-slate-600 dark:border-slate-700 dark:bg-slate-900/80 dark:text-slate-300">
        <RefreshCwIcon className={`mr-2 size-3 ${isRefreshing ? "animate-spin" : ""}`} />
        Live sync
      </span>
    </div>
  )
}

export { DashboardPulseIndicator }

