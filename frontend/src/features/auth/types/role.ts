export type AuthRole = "officer" | "member"

export type AuthRoleConfig = {
  role: AuthRole
  title: string
  description: string
  badge: string
  accent: string
  buttonLabel: string
  helperText: string
  dashboardPath: string
}

export const AUTH_ROLE_CONFIG: Record<AuthRole, AuthRoleConfig> = {
  officer: {
    role: "officer",
    title: "Officer Portal",
    description: "Untuk admin, officer, dan petugas yang mengelola data, verifikasi, serta monitoring.",
    badge: "Admin / Officer",
    accent: "from-slate-950 via-slate-800 to-slate-700",
    buttonLabel: "Masuk sebagai Officer",
    helperText: "Akses penuh untuk operasional dan manajemen portal.",
    dashboardPath: "/dashboard/officer",
  },
  member: {
    role: "member",
    title: "Member Portal",
    description: "Untuk user atau member yang ingin melihat informasi pribadi, status, dan riwayat layanan.",
    badge: "User / Member",
    accent: "from-emerald-600 via-teal-600 to-cyan-600",
    buttonLabel: "Masuk sebagai Member",
    helperText: "Akses ringkas untuk kebutuhan personal dan pemantauan status.",
    dashboardPath: "/dashboard/member",
  },
}
