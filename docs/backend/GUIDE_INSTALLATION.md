# Backend Installation Guide

Panduan ini menjelaskan cara menjalankan backend `portal_digital` secara lokal.

## Prasyarat

- Go 1.22+.
- PostgreSQL.
- RabbitMQ, jika fitur event processing ingin diaktifkan.
- File konfigurasi `.env` di folder `backend`.

## Struktur Folder Backend

- `backend/cmd/api` - entry point aplikasi.
- `backend/internal` - source utama backend.
- `backend/internal/database/seed` - kumpulan seed per domain agar mudah diperluas.
- `backend/migrations` - file migrasi database.
- `backend/uploads` - file hasil upload.

## Setup Awal

1. Masuk ke folder backend.

```powershell
cd backend
```

2. Salin file environment contoh.

```powershell
Copy-Item .env.example .env
```

3. Sesuaikan nilai di `.env`.

## Konfigurasi Environment

Berikut variabel yang digunakan backend:

- `APP_PORT` - port aplikasi.
- `DB_HOST` - host PostgreSQL.
- `DB_PORT` - port PostgreSQL.
- `DB_USER` - username database.
- `DB_PASSWORD` - password database.
- `DB_NAME` - nama database.
- `DB_SSLMODE` - mode SSL untuk PostgreSQL.
- `DB_TIMEZONE` - timezone database.
- `DB_MAX_OPEN_CONNS` - jumlah koneksi terbuka maksimal.
- `DB_MAX_IDLE_CONNS` - jumlah koneksi idle maksimal.
- `CORS_ORIGINS` - daftar origin yang diizinkan.
- `JWT_SECRET` - secret untuk pembuatan dan validasi token.
- `JWT_EXPIRE` - masa berlaku token.
- `RABBITMQ_ENABLED` - aktif/nonaktif RabbitMQ.
- `RABBITMQ_URL` - connection string RabbitMQ.
- `RABBITMQ_EXCHANGE` - nama exchange RabbitMQ.
- `AUTO_MIGRATE` - menjalankan migrasi otomatis saat startup.
- `SEED_DATABASE` - menjalankan seeding data awal.
- `SEED_RESET_DATA` - mengosongkan tabel seed dulu lalu isi ulang dari nol.
- `SEED_ADMIN_NAME` - nama admin awal.
- `SEED_ADMIN_EMAIL` - email admin awal.
- `SEED_ADMIN_PASSWORD` - password admin awal.
- `SEED_ADMIN_ROLE` - role admin awal.

## Menjalankan Database

Pastikan PostgreSQL sudah aktif dan database sesuai dengan konfigurasi `.env` sudah dibuat.

Jika memakai Docker, jalankan service database dan RabbitMQ sesuai file compose yang tersedia di root project atau folder infrastruktur.

## Install Dependency

```powershell
go mod tidy
```

## Menjalankan Migrasi dan Seed

Backend ini melakukan migrasi dan seed saat startup jika variabel berikut aktif:

- `AUTO_MIGRATE=true`
- `SEED_DATABASE=true`
- `SEED_RESET_DATA=true` hanya bila ingin hapus data seed lama dan membuat data baru dari nol

Jika ingin data admin awal dibuat otomatis, pastikan `SEED_DATABASE=true`.

## Seed Scalable

Seed backend sekarang disusun per domain supaya mudah ditambah tanpa membuat satu file besar.

Struktur utamanya:

- `backend/internal/database/seed/seed.go` - orchestrator seed utama.
- `backend/internal/database/seed/users.go` - seed user dan helper admin.
- `backend/internal/database/seed/violations.go` - seed rule, violation type, dan violation.
- `backend/internal/database/seed/finance.go` - seed invoice dan payment transaction.
- `backend/internal/database/seed/helpers.go` - helper umum seperti truncate/reset.
- `backend/internal/database/seed/time.go` - helper waktu seed agar konsisten.

Mode reset sangat berguna saat:

- ingin menghapus data demo yang lama
- ingin memvalidasi ulang relasi antar tabel
- ingin memastikan semua tabel seed terisi ulang dengan urutan yang sama

Contoh konfigurasi untuk fresh seed:

```env
SEED_DATABASE=true
SEED_RESET_DATA=true
```

## Menjalankan Backend

```powershell
go run ./cmd/api
```

Server akan berjalan sesuai `APP_PORT`, misalnya:

```text
http://localhost:8100
```

## Health Check

Setelah server hidup, cek endpoint berikut:

```http
GET /health
```

Respons sukses:

```json
{
  "status": true,
  "message": "server is running",
  "data": null
}
```

## Catatan Penting

- Backend menggunakan JWT untuk endpoint yang dilindungi.
- File upload disimpan dan disajikan dari folder `uploads`.
- Jika RabbitMQ tidak tersedia, aplikasi tetap bisa startup, tetapi fitur event publisher/worker akan dinonaktifkan saat koneksi gagal.
