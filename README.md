# REST API Boilerplate with Golang & Fiber

This is a simple REST API boilerplate using Golang and Fiber framework.

## Features
- Clean architecture
- CRUD example
- Config file

## Usage
1. Install dependencies: `go mod tidy`
2. Run the server: `go run cmd/main.go`

## Folder Structure
## Database Migration

Untuk melakukan migrasi database, jalankan perintah berikut di terminal:

```fish
go run internal/database/migrate.go
```

Pastikan konfigurasi database sudah benar di `config/config.yaml` sebelum menjalankan migrasi.
