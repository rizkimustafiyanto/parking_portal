import { ClipboardListIcon, FileTextIcon, LayoutDashboardIcon, ShieldAlertIcon, UploadIcon, UsersIcon } from "lucide-react"

import { DashboardShell } from "@/components/layout/dashboard-shell"
import { RoleGuard } from "@/features/auth"

const navItems = [
  { label: "Overview", href: "/dashboard/officer", icon: LayoutDashboardIcon },
  { label: "Violations", href: "/dashboard/officer/violations", icon: ShieldAlertIcon },
  { label: "Invoices", href: "/dashboard/officer/invoices", icon: FileTextIcon },
  { label: "Payments", href: "/dashboard/officer/payments", icon: FileTextIcon },
  { label: "Users", href: "/dashboard/officer/members", icon: UsersIcon },
  { label: "Tasks", href: "/dashboard/officer/tasks", icon: ClipboardListIcon },
  { label: "Uploads", href: "/dashboard/officer/settings", icon: UploadIcon },
]

export default function OfficerDashboardLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <RoleGuard role="officer">
      <DashboardShell
        title="Officer Workspace"
        subtitle="Area kerja untuk admin dan officer yang menangani operasional portal."
        badge="Officer / Admin"
        accentClassName="from-slate-950 via-slate-800 to-slate-700"
        homeHref="/dashboard/officer"
        navItems={navItems}
      >
        {children}
      </DashboardShell>
    </RoleGuard>
  )
}
