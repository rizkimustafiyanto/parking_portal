"use client"

import { useState, type FormEvent } from "react"
import { FileUpIcon, RefreshCcwIcon } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { LoadingState } from "@/components/ui/loading-state"
import { uploadFile, type UploadResponse } from "@/features/uploads"

export default function OfficerSettingsPage() {
  const [file, setFile] = useState<File | null>(null)
  const [type, setType] = useState("document")
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [result, setResult] = useState<UploadResponse | null>(null)

  async function handleUpload(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()

    if (!file) {
      setError("Pilih file terlebih dahulu")
      return
    }

    setIsLoading(true)
    setError(null)

    try {
      const response = await uploadFile(file, type)
      setResult(response)
    } catch (err) {
      setError(err instanceof Error ? err.message : "Gagal mengunggah file")
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="space-y-6">
      <Card className="overflow-hidden rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
        <div className="h-2 bg-gradient-to-r from-slate-950 via-slate-800 to-slate-700" />
        <CardHeader className="space-y-3 p-8">
          <CardDescription>Officer / Admin</CardDescription>
          <CardTitle className="text-2xl">Uploads</CardTitle>
          <p className="max-w-3xl text-sm leading-6 text-slate-600 dark:text-slate-300">
            Upload file ke backend lewat `POST /uploads` dengan multipart form-data.
          </p>
        </CardHeader>
      </Card>

      <div className="grid gap-6 xl:grid-cols-[420px_1fr]">
        <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
          <CardHeader className="space-y-2 p-6">
            <CardTitle className="text-xl">Upload File</CardTitle>
            <CardDescription>Pilih tipe file dan kirim ke endpoint upload backend.</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4 px-6 pb-6">
            <form className="space-y-4" onSubmit={handleUpload}>
              <div className="space-y-2">
                <label className="text-sm font-medium text-slate-700 dark:text-slate-300" htmlFor="upload-type">
                  Tipe
                </label>
                <select
                  id="upload-type"
                  value={type}
                  onChange={(event) => setType(event.target.value)}
                  className="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-3 focus-visible:ring-ring/50"
                >
                  <option value="document">Document</option>
                  <option value="profile">Profile</option>
                  <option value="vehicle">Vehicle</option>
                </select>
              </div>

              <div className="space-y-2">
                <label className="text-sm font-medium text-slate-700 dark:text-slate-300" htmlFor="file">
                  File
                </label>
                <Input
                  id="file"
                  type="file"
                  onChange={(event) => setFile(event.target.files?.[0] ?? null)}
                  accept="image/*,application/pdf"
                />
              </div>

              {error ? <p className="text-sm text-red-600">{error}</p> : null}

              <div className="flex flex-wrap gap-3">
                <Button type="submit" loading={isLoading}>
                  <FileUpIcon className="size-4" />
                  Upload
                </Button>
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => {
                    setFile(null)
                    setResult(null)
                    setError(null)
                  }}
                >
                  <RefreshCcwIcon className="size-4" />
                  Reset
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>

        <Card className="rounded-3xl border-slate-200 bg-white/95 shadow-lg dark:border-slate-700 dark:bg-slate-900/95">
          <CardHeader className="space-y-2 p-6">
            <CardTitle className="text-xl">Hasil Upload</CardTitle>
            <CardDescription>Informasi file terakhir yang berhasil diunggah.</CardDescription>
          </CardHeader>
          <CardContent className="px-6 pb-6">
            {isLoading ? <LoadingState rows={2} /> : null}

            {!isLoading && result ? (
              <div className="space-y-3 rounded-2xl border border-slate-200 bg-slate-50 p-5 text-sm dark:border-slate-700 dark:bg-slate-800/60">
                <div className="flex items-center justify-between gap-4">
                  <span className="text-slate-500 dark:text-slate-400">File name</span>
                  <span className="font-medium text-slate-950 dark:text-white">{result.file_name}</span>
                </div>
                <div className="flex items-center justify-between gap-4">
                  <span className="text-slate-500 dark:text-slate-400">Type</span>
                  <span className="font-medium text-slate-950 dark:text-white">{result.type}</span>
                </div>
                <div className="flex items-center justify-between gap-4">
                  <span className="text-slate-500 dark:text-slate-400">Mime</span>
                  <span className="font-medium text-slate-950 dark:text-white">{result.mime_type}</span>
                </div>
                <div className="flex items-center justify-between gap-4">
                  <span className="text-slate-500 dark:text-slate-400">Size</span>
                  <span className="font-medium text-slate-950 dark:text-white">{result.size} bytes</span>
                </div>
                <div className="break-all rounded-xl bg-white p-3 text-slate-700 ring-1 ring-slate-200 dark:bg-slate-900 dark:text-slate-300 dark:ring-slate-700">
                  {result.file_url}
                </div>
              </div>
            ) : null}

            {!isLoading && !result ? (
              <div className="rounded-2xl border border-dashed border-slate-200 bg-slate-50 p-6 text-sm text-slate-500 dark:border-slate-700 dark:bg-slate-800/50 dark:text-slate-300">
                Hasil upload akan tampil di sini setelah file berhasil dikirim.
              </div>
            ) : null}
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
