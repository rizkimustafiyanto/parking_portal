# About Frontend

Frontend `Portal Digital` dibangun sebagai portal berbasis role untuk dua jenis pengguna utama:

- `Officer / Admin`
- `Member / User`

Tujuan utamanya adalah menyediakan antarmuka yang rapi untuk login, navigasi dashboard, dan konsumsi data dari backend API.

## Teknologi yang dipakai

- `Next.js 16` sebagai framework utama
- `React 19` untuk UI component
- `TypeScript` untuk typing yang lebih aman
- `Tailwind CSS 4` untuk styling
- `Zustand` untuk state auth sederhana
- `Axios` untuk komunikasi ke backend
- `React Hook Form` untuk form login
- `Zod` untuk validasi input
- `sonner` untuk toast/notifikasi
- `lucide-react` untuk ikon

## Kenapa teknologi ini dipakai

- `Next.js` dipakai karena mendukung App Router, struktur halaman yang jelas, dan cocok untuk portal multi-role
- `React Hook Form` + `Zod` dipakai supaya form login ringan dan validasinya konsisten
- `Zustand` dipakai untuk menyimpan token dan role secara sederhana
- `Axios` dipakai agar request ke backend mudah diberi interceptor header `Authorization`
- `Tailwind CSS` dipakai supaya pengembangan UI cepat dan konsisten

## Arsitektur frontend

Frontend dibagi menjadi beberapa bagian utama:

- `frontend/src/app` untuk halaman dan layout berbasis route
- `frontend/src/components` untuk komponen reusable
- `frontend/src/features/auth` untuk login, guard, dan manajemen session
- `frontend/src/features/finance` untuk data invoice, payment, dan history
- `frontend/src/features/users` untuk data user
- `frontend/src/features/violation` untuk data pelanggaran
- `frontend/src/features/realtime` untuk indikator status realtime

## Alur kerja utama

1. User membuka halaman utama
2. User memilih role melalui `RoleGateway`
3. User login sesuai role
4. Token dan role disimpan di `localStorage`
5. `RoleGuard` memastikan halaman yang dibuka sesuai dengan role
6. Dashboard mengambil data dari backend
7. User dapat logout untuk menghapus session lokal

## Contoh fitur yang ada

### Role Gateway

Halaman awal menampilkan dua pintu masuk:

- Portal Officer
- Portal Member

Fitur ini memudahkan pengguna langsung masuk ke jalur yang sesuai tanpa perlu bingung memilih halaman login.

### Login per role

Login memanfaatkan query parameter seperti `?role=officer` atau `?role=member`.

Contohnya:

- Officer login akan diarahkan ke dashboard operasional
- Member login akan diarahkan ke dashboard personal

### Dashboard Officer

Dashboard officer dipakai untuk melihat data operasional seperti:

- jumlah user
- jumlah pelanggaran
- jumlah invoice
- jumlah payment

### Dashboard Member

Dashboard member dipakai untuk melihat data personal seperti:

- invoice milik akun sendiri
- status pembayaran
- riwayat saldo atau transaksi

### Topbar dan shell dashboard

`DashboardShell` dan `DashboardTopbar` dipakai agar tampilan dashboard konsisten, dengan sidebar navigasi, header workspace, tombol dashboard, dan logout.

## Integrasi ke backend

Frontend membaca base URL dari environment variable:

- `NEXT_PUBLIC_API_BASE_URL`

Kalau variabel ini tidak diisi, frontend akan memakai default `http://localhost:8080`.

Semua request API yang butuh autentikasi akan membawa token dari `localStorage` melalui interceptor Axios.

## Data flow singkat

- Frontend menampilkan form dan state UI
- Backend menyediakan endpoint login dan data resource
- Frontend mengambil data dashboard dan menampilkannya
- State login disimpan di browser agar navigasi tetap konsisten

