import axios from "axios"

import api from "@/lib/api/client"
import type {
  CreateUserPayload,
  TopUpBalancePayload,
  UpdateUserPayload,
  UserListResponse,
  UserRecord,
} from "../types"

type ApiEnvelope<T> = {
  message?: string
  data: T
}

function getMessage(error: unknown, fallback: string) {
  if (axios.isAxiosError(error)) {
    return (error.response?.data as { message?: string } | undefined)?.message ?? error.message ?? fallback
  }

  if (error instanceof Error) {
    return error.message
  }

  return fallback
}

export async function fetchUsers(params: {
  page?: number
  limit?: number
  search?: string
  role?: string
  sortBy?: string
  order?: string
}): Promise<UserListResponse> {
  try {
    const response = await api.get<UserListResponse>("/api/users", {
      params: {
        page: params.page ?? 1,
        limit: params.limit ?? 10,
        search: params.search || undefined,
        role: params.role || undefined,
        sortBy: params.sortBy || "created_at",
        order: params.order || "desc",
      },
    })

    return response.data
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil daftar user"))
  }
}

export async function createUser(payload: CreateUserPayload): Promise<void> {
  try {
    await api.post<ApiEnvelope<null>>("/api/users", payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal membuat user"))
  }
}

export async function updateUser(id: string, payload: UpdateUserPayload): Promise<void> {
  try {
    await api.put<ApiEnvelope<null>>(`/api/users/${id}`, payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal memperbarui user"))
  }
}

export async function deleteUser(id: string): Promise<void> {
  try {
    await api.delete<ApiEnvelope<null>>(`/api/users/${id}`)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal menghapus user"))
  }
}

export async function topUpBalance(id: string, payload: TopUpBalancePayload): Promise<UserRecord> {
  try {
    const response = await api.post<ApiEnvelope<UserRecord>>(`/api/users/${id}/top-up-balance`, payload)

    return response.data.data
  } catch (error) {
    throw new Error(getMessage(error, "Gagal top up balance"))
  }
}

