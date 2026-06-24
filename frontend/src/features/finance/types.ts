export type InvoiceStatus = "PENDING" | "PAID" | "OVERDUE" | "CANCELED" | string
export type PaymentStatus = "PENDING" | "SUCCESS" | "FAILED" | string
export type PaymentScenario = "SUCCESS" | "FAILURE" | "TIMEOUT" | string

export type UserThrow2 = {
  id: string
  name: string
}

export type ViolationFineRuleDetail = {
  id: string
  rule_type: string
  key: string
  value: string
  rule_version_id: string
}

export type ViolationFineRuleVersion = {
  id: string
  version_number: number
  is_active: boolean
  details: ViolationFineRuleDetail[]
}

export type ViolationHistory = {
  id: string
  plate_number: string
  location: string
  occurred_at: string
  photo_url: string
  fine_rule_version: ViolationFineRuleVersion
}

export type PaymentTransactionThrow = {
  id: string
  invoice_id: string
  internal_transaction_id: string
  amount: number
  status: PaymentStatus
  scenario: PaymentScenario
  paid_at: string
}

export type InvoiceRecord = {
  id: string
  violation_id: string
  amount: number
  status: InvoiceStatus
  member: UserThrow2
  violation: ViolationHistory
  payment: PaymentTransactionThrow
  created_at: string
  updated_at: string
}

export type PaymentRecord = {
  id: string
  invoice_id: string
  internal_transaction_id: string
  amount: number
  status: PaymentStatus
  scenario: PaymentScenario
  paid_at: string
  created_at: string
  updated_at: string
}

export type CreateInvoicePayload = {
  violation_id: string
  member_id: string
  status: InvoiceStatus
}

export type UpdateInvoicePayload = {
  violation_id?: string
  member_id?: string
  status?: InvoiceStatus
}

export type CreatePaymentPayload = {
  invoice_id: string
  amount: number
  scenario: PaymentScenario
  paid_at: string
}

export type UpdatePaymentPayload = {
  internal_transaction_id?: string
  amount?: number
  status?: PaymentStatus
  scenario?: PaymentScenario
  paid_at?: string
}
