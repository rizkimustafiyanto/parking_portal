# Guide Fitur Frontend

Dokumen ini menjelaskan fitur utama yang tersedia di frontend `Portal Digital`.

## Gambaran umum

Frontend dibangun dengan pendekatan role-based portal. Ada dua jalur utama:

- `Officer / Admin` untuk area operasional
- `Member / User` untuk area personal

Setiap role punya layout, navigasi, dan ringkasan data yang berbeda.

## Fitur autentikasi

### 1. Role Gateway

Halaman awal menampilkan pilihan masuk berdasarkan role:

- Officer portal
- Member portal

Fungsi ini ada di `frontend/src/features/auth/components/RoleGateway.tsx`.

### 2. Login per role

Form login menggunakan:

- React Hook Form
- Zod validation
- request ke API via Axios

Role login ditentukan dari query `?role=officer` atau `?role=member`.

### 3. Proteksi halaman

Frontend memakai `RoleGuard` untuk memastikan:

- user sudah login
- token tersedia
- role sesuai dengan halaman yang dibuka

Jika tidak sesuai, user diarahkan ke halaman yang benar atau ke login.

### 4. Penyimpanan session

Token dan role disimpan di `localStorage`.

Token juga dipakai otomatis pada request API lewat interceptor Axios.

## Fitur dashboard member

Halaman member ada di `frontend/src/app/dashboard/member`.

Yang ditampilkan:

- total invoice
- jumlah invoice paid
- jumlah invoice pending
- status pembayaran terakhir
- jumlah balance history
- daftar invoice terbaru
- daftar histori saldo atau pembayaran

Member juga bisa melihat ringkasan data pribadinya melalui `getStoredUserId()` yang diambil dari payload token.

## Fitur dashboard officer

Halaman officer ada di `frontend/src/app/dashboard/officer`.

Yang ditampilkan:

- jumlah violations
- jumlah pending invoices
- jumlah paid invoices
- jumlah members
- jumlah officers
- jumlah payments
- jumlah successful payments
- jumlah failed payments

Dashboard officer mengambil data dari beberapa sumber sekaligus:

- users
- violations
- invoices
- payments

## Layout dashboard

Frontend memakai `DashboardShell` sebagai pembungkus utama dashboard.

Komponen ini menyediakan:

- sidebar navigasi
- topbar dengan tombol dashboard dan logout
- badge role
- aksen warna berbeda untuk setiap role

File utama:

- `frontend/src/components/layout/dashboard-shell.tsx`
- `frontend/src/components/layout/dashboard-topbar.tsx`

## Navigasi member

Menu member meliputi:

- Overview
- Profile
- Invoices
- Payments
- Support

## Navigasi officer

Menu officer meliputi:

- Overview
- Violations
- Invoices
- Payments
- Users
- Tasks
- Uploads

## Indikator realtime

Frontend juga punya komponen realtime berupa `DashboardPulseIndicator`.

Fungsinya untuk memberi status visual bahwa workspace sedang aktif atau terhubung.

## Komponen UI yang dipakai ulang

Beberapa komponen shared yang sering dipakai:

- `Button`
- `Card`
- `Input`
- `LoadingState`

Komponen-komponen ini membantu menjaga konsistensi tampilan antar halaman.

## Data yang diambil dari API

Frontend terhubung ke backend melalui base URL dari environment variable:

- `NEXT_PUBLIC_API_BASE_URL`

Data yang umum dipakai frontend:

- login dan validasi session
- daftar user
- daftar violation
- invoice
- payment
- riwayat balance member

## Ringkasan alur user

1. Buka halaman utama
2. Pilih role
3. Login sesuai role
4. Masuk ke dashboard yang sesuai
5. Lihat ringkasan data dan navigasi fitur
6. Logout jika selesai

