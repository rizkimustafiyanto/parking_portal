import { fetchInvoices, fetchMemberBalanceHistory, fetchMemberInvoices, fetchPayments } from "@/features/finance"
import { fetchUsers } from "@/features/users"
import { fetchViolations } from "@/features/violation"
import type { DashboardRole, DashboardSnapshot } from "../types"

async function fetchMemberSnapshot(memberId: string): Promise<DashboardSnapshot> {
  const [invoiceRes, historyRes] = await Promise.all([
    fetchMemberInvoices(memberId),
    fetchMemberBalanceHistory(memberId),
  ])

  const invoices = invoiceRes.data ?? []
  const history = historyRes.data ?? []
  const latestPayment = history.find((item) => item.payment?.status) ?? null

  return {
    role: "member",
    invoiceCount: invoices.length,
    paidInvoiceCount: invoices.filter((item) => item.status === "PAID").length,
    pendingInvoiceCount: invoices.filter((item) => item.status === "PENDING").length,
    paymentCount: history.length,
    successfulPaymentCount: history.filter((item) => item.payment?.status === "SUCCESS").length,
    failedPaymentCount: history.filter((item) => item.payment?.status === "FAILED").length,
    latestStatusLabel: latestPayment?.payment?.status ?? latestPayment?.status ?? null,
  }
}

async function fetchOfficerSnapshot(): Promise<DashboardSnapshot> {
  const [userRes, violationRes, invoiceRes, paymentRes] = await Promise.all([
    fetchUsers({ page: 1, limit: 200 }),
    fetchViolations({ page: 1, limit: 200 }),
    fetchInvoices({ page: 1, limit: 200 }),
    fetchPayments({ page: 1, limit: 200 }),
  ])

  const users = userRes.data ?? []
  const violations = violationRes.data ?? []
  const invoices = invoiceRes.data ?? []
  const payments = paymentRes.data ?? []

  return {
    role: "officer",
    invoiceCount: invoices.length,
    paidInvoiceCount: invoices.filter((item) => item.status === "PAID").length,
    pendingInvoiceCount: invoices.filter((item) => item.status === "PENDING").length,
    paymentCount: payments.length,
    successfulPaymentCount: payments.filter((item) => item.status === "SUCCESS").length,
    failedPaymentCount: payments.filter((item) => item.status === "FAILED").length,
    latestStatusLabel: `users:${users.length} violations:${violations.length}`,
  }
}

export async function fetchDashboardSnapshot(role: DashboardRole, memberId?: string | null) {
  if (role === "member") {
    if (!memberId) {
      throw new Error("User ID tidak ditemukan untuk memuat snapshot member")
    }

    return fetchMemberSnapshot(memberId)
  }

  return fetchOfficerSnapshot()
}

