export type UploadType = "profile" | "document" | "vehicle" | string

export type UploadResponse = {
  file_name: string
  file_path: string
  file_url: string
  type: string
  mime_type: string
  size: number
}

