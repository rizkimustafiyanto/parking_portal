export type DashboardRole = "member" | "officer"

export type DashboardSnapshot = {
  role: DashboardRole
  invoiceCount: number
  paidInvoiceCount: number
  pendingInvoiceCount: number
  paymentCount: number
  successfulPaymentCount: number
  failedPaymentCount: number
  latestStatusLabel: string | null
}

