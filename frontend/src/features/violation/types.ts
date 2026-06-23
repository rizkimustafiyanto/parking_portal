export type PaginationMeta = {
  page: number
  limit: number
  total: number
  total_pages: number
}

export type ApiEnvelope<T> = {
  message?: string
  data: T
  meta?: PaginationMeta
}

export type ViolationOfficer = {
  id: string
  name: string
  email: string
}

export type FineRuleVersionThrow = {
  id: string
  version_number: number
  is_active: boolean
}

export type ViolationRecord = {
  id: string
  plate_number: string
  violation_type_code: string
  location: string
  occurred_at: string
  photo_url: string
  officer: ViolationOfficer
  fine_rule_version: FineRuleVersionThrow
  created_at: string
  updated_at: string
}

export type ViolationTypeRecord = {
  id: string
  code: string
  name: string
  base_amount: number
  created_by: ViolationOfficer
  created_at: string
  updated_at: string
}

export type FineRuleDetailRecord = {
  id: string
  rule_version_id: string
  rule_type: string
  key: string
  value: string
  created_at: string
  updated_at: string
}

export type FineRuleVersionRecord = {
  id: string
  version_number: number
  is_active: boolean
  publish: ViolationOfficer
  details: FineRuleDetailRecord[]
  created_at: string
  updated_at: string
}

export type CreateViolationPayload = {
  plate_number: string
  violation_type_code: string
  location: string
  occurred_at: string
  photo_url?: string
  officer_id: string
}

export type UpdateViolationPayload = Partial<CreateViolationPayload> & {
  fine_rule_version_id?: string
}

export type CreateViolationTypePayload = {
  code: string
  name: string
  base_amount: number
  created_by_id: string
}

export type UpdateViolationTypePayload = Partial<CreateViolationTypePayload>

export type CreateFineRuleVersionPayload = {
  version_number: number
  is_active?: boolean
  published_by: string
}

export type UpdateFineRuleVersionPayload = Partial<CreateFineRuleVersionPayload>

export type CreateFineRuleDetailPayload = {
  rule_version_id: string
  rule_type?: string
  key: string
  value: string
}

export type UpdateFineRuleDetailPayload = Partial<CreateFineRuleDetailPayload>

