import axios from "axios"

import api from "@/lib/api/client"
import type { UploadResponse } from "../types"

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

export async function uploadFile(file: File, type: string): Promise<UploadResponse> {
  try {
    const formData = new FormData()
    formData.append("file", file)
    formData.append("type", type)

    const response = await api.post<ApiEnvelope<UploadResponse>>("/api/uploads", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })

    return response.data.data
  } catch (error) {
    throw new Error(getMessage(error, "Gagal mengunggah file"))
  }
}

