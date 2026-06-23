import axios from "axios"

import api from "@/lib/api/client"
import type {
  ApiEnvelope,
  CreateFineRuleDetailPayload,
  CreateFineRuleVersionPayload,
  CreateViolationPayload,
  CreateViolationTypePayload,
  FineRuleDetailRecord,
  FineRuleVersionRecord,
  UpdateFineRuleDetailPayload,
  UpdateFineRuleVersionPayload,
  UpdateViolationPayload,
  UpdateViolationTypePayload,
  ViolationRecord,
  ViolationTypeRecord,
} from "../types"

function getMessage(error: unknown, fallback: string) {
  if (axios.isAxiosError(error)) {
    return (error.response?.data as { message?: string } | undefined)?.message ?? error.message ?? fallback
  }
  if (error instanceof Error) return error.message
  return fallback
}

async function list<T>(path: string, params: Record<string, unknown> = {}) {
  const response = await api.get<ApiEnvelope<T[]>>(path, { params })
  return response.data
}

async function create(path: string, payload: unknown) {
  await api.post(path, payload)
}

async function update(path: string, payload: unknown) {
  await api.put(path, payload)
}

async function remove(path: string) {
  await api.delete(path)
}

export async function fetchViolations(params: Record<string, unknown> = {}) {
  try {
    return await list<ViolationRecord>("/api/violations", params)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil daftar violation"))
  }
}

export async function createViolation(payload: CreateViolationPayload) {
  try {
    await create("/api/violations", payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal membuat violation"))
  }
}

export async function updateViolation(id: string, payload: UpdateViolationPayload) {
  try {
    await update(`/api/violations/${id}`, payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal memperbarui violation"))
  }
}

export async function deleteViolation(id: string) {
  try {
    await remove(`/api/violations/${id}`)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal menghapus violation"))
  }
}

export async function fetchViolationTypes(params: Record<string, unknown> = {}) {
  try {
    return await list<ViolationTypeRecord>("/api/violation-types", params)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil violation type"))
  }
}

export async function createViolationType(payload: CreateViolationTypePayload) {
  try {
    await create("/api/violation-types", payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal membuat violation type"))
  }
}

export async function updateViolationType(id: string, payload: UpdateViolationTypePayload) {
  try {
    await update(`/api/violation-types/${id}`, payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal memperbarui violation type"))
  }
}

export async function deleteViolationType(id: string) {
  try {
    await remove(`/api/violation-types/${id}`)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal menghapus violation type"))
  }
}

export async function fetchFineRuleVersions(params: Record<string, unknown> = {}) {
  try {
    return await list<FineRuleVersionRecord>("/api/fine-rule-versions", params)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil fine rule version"))
  }
}

export async function createFineRuleVersion(payload: CreateFineRuleVersionPayload) {
  try {
    await create("/api/fine-rule-versions", payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal membuat fine rule version"))
  }
}

export async function updateFineRuleVersion(id: string, payload: UpdateFineRuleVersionPayload) {
  try {
    await update(`/api/fine-rule-versions/${id}`, payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal memperbarui fine rule version"))
  }
}

export async function deleteFineRuleVersion(id: string) {
  try {
    await remove(`/api/fine-rule-versions/${id}`)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal menghapus fine rule version"))
  }
}

export async function fetchFineRuleDetails(params: Record<string, unknown> = {}) {
  try {
    return await list<FineRuleDetailRecord>("/api/fine-rule-details", params)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil fine rule detail"))
  }
}

export async function createFineRuleDetail(payload: CreateFineRuleDetailPayload) {
  try {
    await create("/api/fine-rule-details", payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal membuat fine rule detail"))
  }
}

export async function updateFineRuleDetail(id: string, payload: UpdateFineRuleDetailPayload) {
  try {
    await update(`/api/fine-rule-details/${id}`, payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal memperbarui fine rule detail"))
  }
}

export async function deleteFineRuleDetail(id: string) {
  try {
    await remove(`/api/fine-rule-details/${id}`)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal menghapus fine rule detail"))
  }
}

