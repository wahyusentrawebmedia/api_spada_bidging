# API SPADA Bridging Documentation

## Overview
API SPADA Bridging adalah REST API berbasis Golang (Fiber) untuk integrasi dan sinkronisasi data akademik, dengan dukungan PostgreSQL dan JWT eksternal.

## Fitur Utama
- Clean architecture (handler, service, repository, model)
- Koneksi database PostgreSQL (GORM)
- Middleware JWT eksternal
- Migrasi otomatis (GORM AutoMigrate)
- Modular context untuk user
- Contoh endpoint CRUD (user, item, PostgresConfig)

## Struktur Folder
- `cmd/` : Entry point aplikasi
- `internal/handler/` : HTTP handler (controller)
- `internal/model/` : Model data (GORM)
- `internal/service/` : Bisnis logic/service
- `internal/repository/` : Data access layer
- `internal/database/` : Koneksi & migrasi database
- `internal/middleware/` : Middleware (JWT, dsb)
- `internal/context/` : Context aplikasi (user, dsb)
- `config/` : File konfigurasi

## Konfigurasi Environment
Pastikan variabel berikut diatur (bisa di .env/.envrc):
- `DB_USER`, `DB_PASS`, `DB_HOST`, `DB_PORT`, `DB_NAME` : Koneksi PostgreSQL
- `JWT_CHECK_URL` : Endpoint validasi JWT eksternal

## Menjalankan Aplikasi
1. Install dependency:
   ```sh
   go mod tidy
   ```
2. Build:
   ```sh
   go build -o api-spada-bridging ./cmd/
   ```
3. Jalankan:
   ```sh
   ./api-spada-bridging
   ```

## Migrasi Database
Migrasi otomatis dijalankan saat aplikasi start (lihat `internal/database/migrate.go`).

## Contoh Endpoint
- `POST /user` : Sinkronisasi user (body: UserSyncRequest)
- `POST /user/sync` : Sinkronisasi batch user
- `POST /user/reset` : Reset password user
- `GET /api/items` : List item
- `POST /api/items` : Tambah item
- `GET /api/items/:id` : Detail item
- `PUT /api/items/:id` : Update item
- `DELETE /api/items/:id` : Hapus item
- `GET /api/users` : List user
- `POST /api/users` : Tambah user
- `GET /api/users/:id` : Detail user
- `PUT /api/users/:id` : Update user
- `DELETE /api/users/:id` : Hapus user
- `GET /api/ping` : Health check

## Middleware JWT
Semua endpoint dapat diamankan dengan middleware JWT eksternal (`internal/middleware/jwt.go`).

## Model Penting
- `UserSyncRequest`, `UserSyncResponse` : Sinkronisasi user
- `PostgresConfig` : Konfigurasi koneksi database

## Kontribusi
- Gunakan branch feature/ untuk pengembangan baru
- Pull request ke main setelah review

## Lisensi
MIT
