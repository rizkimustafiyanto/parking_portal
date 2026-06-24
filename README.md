# Portal Digital

Portal Digital adalah aplikasi web full-stack untuk pengelolaan data pengguna, pelanggaran, invoice, dan transaksi pembayaran. Repository ini dibagi menjadi dua bagian utama:

- `backend`: API Go
- `frontend`: aplikasi Next.js

Dokumentasi rinci untuk masing-masing bagian sudah tersedia di:

- [Dokumentasi Backend](./docs/backend/ABOUT_BACKEND.md)
- [Panduan Install Backend](./docs/backend/GUIDE_INSTALLATION.md)
- [Panduan API Backend](./docs/backend/GUIDE_API.md)
- [Dokumentasi Frontend](./docs/frontend/ABOUT_FRONTEND.md)
- [Panduan Install Frontend](./docs/frontend/GUIDE_INSTALLATION.md)
- [Panduan Fitur Frontend](./docs/frontend/GUIDE_FEATURES.md)

## Ringkasan Teknologi

- Frontend: Next.js, React, TypeScript, Tailwind CSS
- Backend: Go, Gin, GORM, Viper
- Database: PostgreSQL
- Messaging: RabbitMQ

## Struktur Proyek

```text
portal_digital/
├─ backend/
├─ frontend/
├─ infrastructure/
├─ docs/
└─ README.md
```

## Prasyarat

- Node.js 18+ untuk frontend
- Go 1.22+ untuk backend
- PostgreSQL
- RabbitMQ, jika fitur event processing ingin diaktifkan

## Instalasi Cepat

### 1. Backend

Masuk ke folder backend:

```powershell
cd backend
```

Salin file environment contoh:

```powershell
Copy-Item .env.example .env
```

Sesuaikan isi `.env` dengan environment lokal Anda.

Install dependency Go:

```powershell
go mod tidy
```

Jalankan backend:

```powershell
go run ./cmd/api
```

### 2. Frontend

Masuk ke folder frontend:

```powershell
cd frontend
```

Install dependency:

```powershell
npm install
```

Jalankan frontend:

```powershell
npm run dev
```

## Konfigurasi Backend

Backend membaca konfigurasi dari file `.env` di folder `backend`.

Variabel yang umum dipakai:

- `APP_PORT`
- `DB_HOST`
- `DB_PORT`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `DB_SSLMODE`
- `DB_MAX_OPEN_CONNS`
- `DB_MAX_IDLE_CONNS`
- `DB_TIMEZONE`
- `CORS_ORIGINS`
- `JWT_SECRET`
- `JWT_EXPIRE`
- `RABBITMQ_ENABLED`
- `RABBITMQ_URL`
- `RABBITMQ_EXCHANGE`
- `AUTO_MIGRATE`
- `SEED_DATABASE`
- `SEED_RESET_DATA`
- `SEED_ADMIN_NAME`
- `SEED_ADMIN_EMAIL`
- `SEED_ADMIN_PASSWORD`
- `SEED_ADMIN_ROLE`

Contoh nilai awal tersedia di [backend/.env.example](./backend/.env.example) dan file aktif yang sedang dipakai ada di [backend/.env](./backend/.env).

## Database dan Seed

Backend dapat menjalankan migrasi dan seed otomatis saat startup jika:

- `AUTO_MIGRATE=true`
- `SEED_DATABASE=true`

Jika ingin mengulang data seed dari awal, gunakan:

- `SEED_RESET_DATA=true`

## Docker

Repository ini juga menyediakan file Docker Compose di [infrastructure/docker-compose.yml](./infrastructure/docker-compose.yml).

Gunakan file tersebut jika ingin menjalankan service pendukung seperti PostgreSQL dan RabbitMQ lewat Docker.

## Fitur Utama

- Autentikasi pengguna
- Authorization berbasis role
- Manajemen user
- Manajemen violation dan fine rule
- Invoice
- Payment transaction
- Upload file
- Event processing dengan RabbitMQ

## Endpoint Dasar

- `GET /health` untuk health check
- Semua endpoint API berada di prefix `/api`

## Catatan

- File upload disimpan di folder `backend/uploads`
- Jika RabbitMQ tidak tersedia, backend tetap bisa berjalan dan akan menonaktifkan komponen messaging yang gagal konek
- Untuk detail teknis dan panduan fitur, silakan baca dokumen di folder `docs/`

## Lisensi

Belum ditentukan.
