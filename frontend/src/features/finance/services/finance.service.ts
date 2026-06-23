import axios from "axios"

import api from "@/lib/api/client"
import type {
  CreateInvoicePayload,
  CreatePaymentPayload,
  InvoiceRecord,
  PaymentRecord,
  UpdateInvoicePayload,
  UpdatePaymentPayload,
} from "../types"

type Envelope<T> = {
  message?: string
  data: T
}

function getMessage(error: unknown, fallback: string) {
  if (axios.isAxiosError(error)) {
    return (error.response?.data as { message?: string } | undefined)?.message ?? error.message ?? fallback
  }
  if (error instanceof Error) return error.message
  return fallback
}

export async function fetchInvoices(params: Record<string, unknown> = {}) {
  try {
    const response = await api.get<Envelope<InvoiceRecord[]>>("/api/invoice", { params })
    return response.data
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil invoice"))
  }
}

export async function fetchInvoiceDetail(id: string) {
  try {
    const response = await api.get<Envelope<InvoiceRecord>>(`/api/invoices/${id}/detail`)
    return response.data.data
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil detail invoice"))
  }
}

export async function createInvoice(payload: CreateInvoicePayload) {
  try {
    await api.post("/api/invoice", payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal membuat invoice"))
  }
}

export async function updateInvoice(id: string, payload: UpdateInvoicePayload) {
  try {
    await api.put(`/api/invoice/${id}`, payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal memperbarui invoice"))
  }
}

export async function deleteInvoice(id: string) {
  try {
    await api.delete(`/api/invoice/${id}`)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal menghapus invoice"))
  }
}

export async function fetchPayments(params: Record<string, unknown> = {}) {
  try {
    const response = await api.get<Envelope<PaymentRecord[]>>("/api/payment", { params })
    return response.data
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil payment transaction"))
  }
}

export async function createPayment(payload: CreatePaymentPayload) {
  try {
    await api.post("/api/payment", payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal membuat payment transaction"))
  }
}

export async function updatePayment(id: string, payload: UpdatePaymentPayload) {
  try {
    await api.put(`/api/payment/${id}`, payload)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal memperbarui payment transaction"))
  }
}

export async function deletePayment(id: string) {
  try {
    await api.delete(`/api/payment/${id}`)
  } catch (error) {
    throw new Error(getMessage(error, "Gagal menghapus payment transaction"))
  }
}

export async function fetchMemberInvoices(memberId: string) {
  try {
    const response = await api.get<Envelope<InvoiceRecord[]>>(`/api/members/${memberId}/transactions`)
    return response.data
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil invoice member"))
  }
}

export async function fetchMemberBalanceHistory(memberId: string) {
  try {
    const response = await api.get<Envelope<InvoiceRecord[]>>(`/api/members/${memberId}/balance-history`)
    return response.data
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengambil history saldo member"))
  }
}
