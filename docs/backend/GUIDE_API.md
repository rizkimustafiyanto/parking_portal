# Backend API Guide

Dokumentasi ringkas endpoint backend `portal_digital`.

## Base URL

```text
http://localhost:8100/api
```

## Format Response

Sebagian besar endpoint menggunakan format respons berikut:

```json
{
  "status": true,
  "message": "success message",
  "data": {}
}
```

Jika endpoint paginasi, respons dapat memiliki `meta`.

```json
{
  "status": true,
  "message": "success message",
  "data": [],
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 100
  }
}
```

## Autentikasi

Endpoint yang dilindungi memakai header berikut:

```http
Authorization: Bearer <token>
```

Middleware juga menerima token mentah di header `Authorization`, tetapi format `Bearer` adalah yang disarankan.

## Login

### `POST /auth/login`

Masuk menggunakan email dan password.

Request body:

```json
{
  "email": "admin@example.com",
  "password": "password123"
}
```

Response sukses:

```json
{
  "status": true,
  "message": "login success",
  "data": {
    "token": "jwt-token"
  }
}
```

## Users

Semua endpoint berikut membutuhkan JWT.

### `GET /users`

Ambil daftar user.

Query opsional:

- `page`
- `limit`
- `sortBy`
- `order`
- `search`
- `role`

### `GET /users/:id`

Ambil detail user berdasarkan ID.

### `POST /users`

Hanya role `officer`.

Request body:

```json
{
  "name": "Budi",
  "email": "budi@example.com",
  "password": "password123",
  "role": "member"
}
```

### `PUT /users/:id`

Hanya role `officer`.

### `POST /users/:id/top-up-balance`

Hanya role `officer`.

Request body:

```json
{
  "amount": 100000
}
```

### `DELETE /users/:id`

Hanya role `officer`.

## Violations

Semua endpoint berikut membutuhkan JWT.

### `GET /violations`

Daftar pelanggaran.

Query opsional:

- `page`
- `limit`
- `sortBy`
- `order`
- `search`

### `GET /violations/:id`

Detail pelanggaran.

### `POST /violations`

Hanya role `officer`.

Request body:

```json
{
  "plate_number": "B 1234 CD",
  "violation_type_code": "SPEED",
  "location": "Jakarta Selatan",
  "occurred_at": "2026-06-24T10:00:00Z",
  "photo_url": "https://example.com/photo.jpg",
  "officer_id": "uuid-officer",
  "fine_rule_version_id": "uuid-rule-version"
}
```

### `PUT /violations/:id`

Hanya role `officer`.

### `DELETE /violations/:id`

Hanya role `officer`.

## Violation Types

Bagian ini didaftarkan pada grup `/violation-types`.

Semua endpoint membutuhkan JWT.

### `GET /violation-types`

### `GET /violation-types/:id`

### `POST /violation-types`

Hanya role `officer`.

Request body:

```json
{
  "code": "SPEED",
  "name": "Speeding",
  "base_amount": 500000,
  "created_by_id": "uuid-user"
}
```

### `PUT /violation-types/:id`

Hanya role `officer`.

### `DELETE /violation-types/:id`

Hanya role `officer`.

## Fine Rule Versions

Bagian ini didaftarkan pada grup `/fine-rule-versions`.

Semua endpoint membutuhkan JWT.

### `GET /fine-rule-versions`

### `GET /fine-rule-versions/:id`

### `POST /fine-rule-versions`

Hanya role `officer`.

Request body:

```json
{
  "version_number": 1,
  "is_active": true,
  "published_by": "uuid-user"
}
```

### `PUT /fine-rule-versions/:id`

Hanya role `officer`.

### `DELETE /fine-rule-versions/:id`

Hanya role `officer`.

## Fine Rule Details

Bagian ini didaftarkan pada grup `/fine-rule-details`.

Semua endpoint membutuhkan JWT.

### `GET /fine-rule-details`

### `GET /fine-rule-details/:id`

### `POST /fine-rule-details`

Hanya role `officer`.

Request body:

```json
{
  "rule_version_id": "uuid-rule-version",
  "rule_type": "base",
  "key": "amount",
  "value": "500000"
}
```

### `PUT /fine-rule-details/:id`

Hanya role `officer`.

### `DELETE /fine-rule-details/:id`

Hanya role `officer`.

## Invoices

Semua endpoint berikut membutuhkan JWT.

### `GET /invoice`

Daftar invoice.

Query opsional:

- `page`
- `limit`
- `sortBy`
- `order`
- `search`
- `status`

### `GET /invoice/:id`

Detail invoice.

### `POST /invoice`

Hanya role `officer`.

Request body:

```json
{
  "violation_id": "uuid-violation",
  "member_id": "uuid-member",
  "status": "unpaid"
}
```

### `PUT /invoice/:id`

Hanya role `officer`.

### `DELETE /invoice/:id`

Hanya role `officer`.

### `GET /invoices/:id/detail`

Mengambil detail invoice yang lebih lengkap.

### `GET /members/:id/transactions`

Riwayat transaksi member.

### `GET /members/:id/balance-history`

Riwayat perubahan saldo member.

## Payment Transactions

Semua endpoint berikut membutuhkan JWT.

### `GET /payment`

Daftar payment transaction.

Query opsional:

- `page`
- `limit`
- `sortBy`
- `order`
- `search`
- `status`
- `scenario`

### `GET /payment/:id`

### `POST /payment`

Role `member` dan `officer` sama-sama dapat memakai endpoint ini.  
Jika request datang dari member, backend hanya akan memproses invoice milik member tersebut.

Endpoint ini menggunakan payment provider mock internal.
Field `scenario` menentukan hasil simulasi:

- `SUCCESS` -> payment dianggap berhasil
- `FAILURE` -> payment dianggap gagal
- `TIMEOUT` -> payment timeout dan disimpan sebagai gagal

Request body:

```json
{
  "invoice_id": "uuid-invoice",
  "amount": 500000,
  "scenario": "SUCCESS",
  "paid_at": "2026-06-24T10:00:00Z"
}
```

### `PUT /payment/:id`

Hanya role `officer`.

### `DELETE /payment/:id`

Hanya role `officer`.

## Upload

### `POST /uploads`

Upload file.

Biasanya dipakai sebagai `multipart/form-data`.

## Public Endpoint

### `GET /health`

Endpoint cek status server.

Response sukses:

```json
{
  "status": true,
  "message": "server is running",
  "data": null
}
```

## Catatan

- Route yang tercantum di atas mengikuti definisi router backend saat ini.
- Jika ada perubahan pada handler atau route registration, dokumen ini perlu di-update agar tetap sinkron.
