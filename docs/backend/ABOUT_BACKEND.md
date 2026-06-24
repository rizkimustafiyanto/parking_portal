# About Backend

Backend `Portal Digital` bertugas sebagai penyedia API, pengolah data, autentikasi, validasi role, dan integrasi proses asinkron.

Backend ditulis dengan pendekatan modular sehingga tiap domain dipisahkan ke dalam module tersendiri.

## Teknologi yang dipakai

- `Go` sebagai bahasa utama
- `Gin` sebagai web framework HTTP
- `GORM` untuk akses database
- `PostgreSQL` sebagai database utama
- `Viper` untuk loading konfigurasi environment
- `JWT` untuk autentikasi token
- `CORS` middleware untuk akses dari frontend
- `RabbitMQ` untuk event asynchronous

## Kenapa teknologi ini dipakai

- `Gin` dipakai karena ringan dan cocok untuk REST API
- `GORM` dipakai agar akses database lebih terstruktur
- `PostgreSQL` dipakai sebagai database relasional utama
- `JWT` dipakai untuk login tanpa session server-side
- `RabbitMQ` dipakai untuk memisahkan proses yang tidak harus selesai dalam satu request

## Arsitektur backend

Struktur backend dibagi menjadi beberapa lapisan:

- `internal/config` untuk konfigurasi aplikasi
- `internal/database` untuk koneksi, migrasi, dan seed data
- `internal/routes` untuk registrasi route utama
- `internal/middleware` untuk auth, role, dan context
- `internal/modules/*` untuk domain bisnis
- `internal/messaging` untuk event bus dan worker
- `pkg/*` untuk helper umum seperti JWT, response, pagination, dan password

## Modul utama

### Auth

Modul auth menangani:

- login
- validasi kredensial
- pembuatan JWT

Contoh endpoint:

- `POST /api/auth/login`

### User

Modul user menangani data pengguna seperti:

- member
- officer
- pengelompokan role

### Violation

Modul violation menangani data pelanggaran, tipe pelanggaran, dan rule denda.

### Invoice

Modul invoice menangani data tagihan dan terhubung ke proses pembuatan invoice.

### Payment Transaction

Modul payment transaction menangani pencatatan transaksi pembayaran.

### Upload

Modul upload menangani penyimpanan file dan pembuatan metadata file.

## Autentikasi dan authorization

Backend memakai JWT.

Alurnya:

1. User login dengan email dan password
2. Backend memeriksa password
3. Backend membuat token JWT
4. Frontend menyimpan token
5. Request berikutnya mengirim token lewat header `Authorization`
6. Middleware auth memvalidasi token dan menyimpan `user_id` dan `role` ke context

Selain auth, backend juga memakai pengecekan role untuk membatasi akses route tertentu.

## Database

Backend menggunakan PostgreSQL sebagai source of truth untuk data utama.

Fitur database yang dipakai:

- koneksi database
- migrasi schema
- seed data awal
- query lewat repository layer

## RabbitMQ

RabbitMQ dipakai untuk memproses event secara asynchronous.

### Kenapa RabbitMQ dipakai

Tidak semua proses harus selesai saat user menekan tombol. Beberapa proses lebih aman dan efisien dijalankan lewat event queue.

Contoh manfaatnya:

- pembuatan invoice bisa diproses sebagai event
- pembayaran yang selesai bisa memicu event lanjutan
- proses notifikasi atau logging bisa dipisah dari request utama

### Contoh alur RabbitMQ di project ini

- saat invoice dibuat, backend dapat mem-publish event `invoice.created`
- saat payment selesai, backend dapat mem-publish event `payment.completed`
- worker RabbitMQ mendengarkan route tersebut
- worker mengeksekusi handler logging atau proses lanjutan

### Nilai tambah RabbitMQ

- request API tetap cepat
- proses berat dipindah ke background
- arsitektur lebih mudah dikembangkan
- event antar modul lebih rapi

## Konfigurasi penting

Backend membaca konfigurasi dari file `.env` dan default value.

Beberapa variabel penting:

- `APP_PORT`
- `DB_HOST`
- `DB_PORT`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `JWT_SECRET`
- `RABBITMQ_URL`
- `RABBITMQ_EXCHANGE`
- `RABBITMQ_ENABLED`
- `AUTO_MIGRATE`
- `SEED_DATABASE`

## Flow backend singkat

- server start
- config dimuat
- database di-ensure
- koneksi database dibuka
- migrasi dan seed dijalankan jika aktif
- route didaftarkan
- middleware dipasang
- RabbitMQ worker dijalankan jika aktif

