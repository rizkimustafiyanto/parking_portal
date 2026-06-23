import { BellIcon, HomeIcon, ReceiptTextIcon, UserRoundIcon, WalletCardsIcon } from "lucide-react"

import { DashboardShell } from "@/components/layout/dashboard-shell"
import { RoleGuard } from "@/features/auth"

const navItems = [
  { label: "Overview", href: "/dashboard/member", icon: HomeIcon },
  { label: "Profile", href: "/dashboard/member/profile", icon: UserRoundIcon },
  { label: "Invoices", href: "/dashboard/member/history", icon: ReceiptTextIcon },
  { label: "Payments", href: "/dashboard/member/notifications", icon: WalletCardsIcon },
  { label: "Support", href: "/dashboard/member/support", icon: BellIcon },
]

export default function MemberDashboardLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <RoleGuard role="member">
      <DashboardShell
        title="Member Workspace"
        subtitle="Area personal untuk user dan member melihat status, riwayat, dan info penting."
        badge="Member / User"
        accentClassName="from-emerald-600 via-teal-600 to-cyan-600"
        homeHref="/dashboard/member"
        navItems={navItems}
      >
        {children}
      </DashboardShell>
    </RoleGuard>
  )
}
