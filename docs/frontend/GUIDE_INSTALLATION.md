# Guide Instalasi Frontend

Dokumen ini menjelaskan cara menjalankan frontend `Portal Digital` secara lokal.

## Teknologi yang dipakai

- Next.js 16 dengan App Router
- React 19
- TypeScript
- Tailwind CSS 4
- Zustand untuk state auth
- Axios untuk request API
- React Hook Form + Zod untuk form login
- Sonner untuk notifikasi

## Prasyarat

- Node.js versi 18 ke atas
- npm
- Backend API sudah berjalan
- Akses ke file `.env` frontend

## Struktur singkat frontend

- Root aplikasi ada di `frontend/src/app`
- Komponen UI reusable ada di `frontend/src/components/ui`
- Fitur auth ada di `frontend/src/features/auth`
- Fitur data dashboard ada di `frontend/src/features/finance`, `frontend/src/features/users`, `frontend/src/features/violation`, dan `frontend/src/features/realtime`

## Langkah instalasi

1. Masuk ke folder frontend:

```bash
cd frontend
```

2. Install dependency:

```bash
npm install
```

3. Pastikan file environment tersedia.

Frontend membaca base URL API dari variabel berikut:

```bash
NEXT_PUBLIC_API_BASE_URL=http://localhost:8080
```

Jika variabel ini tidak diisi, frontend akan memakai default `http://localhost:8080`.

4. Jalankan mode development:

```bash
npm run dev
```

5. Buka aplikasi di browser:

```bash
http://localhost:3000
```

## Script yang tersedia

- `npm run dev` untuk development
- `npm run build` untuk build production
- `npm run start` untuk menjalankan hasil build
- `npm run lint` untuk pengecekan kode

## Alur akses saat aplikasi dibuka

- Halaman awal menampilkan pemilih peran melalui `RoleGateway`
- Pengguna memilih masuk sebagai `Officer` atau `Member`
- Login akan menyimpan token dan role ke `localStorage`
- Dashboard yang dibuka disesuaikan dengan role pengguna

## Catatan penting

- Pastikan backend aktif sebelum login, karena frontend mengambil data dashboard dari API
- Jika token hilang atau role tidak cocok, user akan diarahkan kembali ke halaman login
- Logout akan menghapus token dan role dari browser storage

