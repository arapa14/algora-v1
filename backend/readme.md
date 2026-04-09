# Algora V1 - Backend

Project backend Algora menggunakan **Vanilla Go** (tanpa framework besar) dengan struktur **Clean Architecture** (DDD-Lite). Project ini dilengkapi dengan *live reload* menggunakan Air dan koneksi database PostgreSQL.

## 📋 Prasyarat
Sebelum memulai, pastikan kamu sudah menginstal:
* [Go](https://go.dev/doc/install) (Versi 1.20+)
* [PostgreSQL](https://www.postgresql.org/download/)
* [Air](https://github.com/air-verse/air) (Untuk Live Reload)

## 🛠️ Persiapan Database
1. Buat database baru di PostgreSQL dengan nama `algora_v1_db`.
2. Jalankan query berikut untuk membuat tabel `users`:
   ```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       username VARCHAR(50) UNIQUE NOT NULL,
       email VARCHAR(100) UNIQUE NOT NULL,
       password TEXT NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
🚀 Cara Menjalankan di Lokal
1. Clone & Install Dependency
Bash
git clone <url-repository-kamu>
cd backend
go mod tidy
2. Konfigurasi Environment
Salin file .env.example menjadi .env dan sesuaikan kredensial databasemu:

Bash
cp .env.example .env
Isi .env dengan:

Cuplikan kode
APP_PORT=3000
APP_ENV=development

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=root
DB_NAME=algora_v1_db
DB_SSLMODE=disable

JWT_SECRET=rahasia_banget
JWT_EXPIRY_HOUR=24
3. Jalankan Aplikasi
Untuk pengembangan (dengan Live Reload):

Bash
air
Untuk menjalankan tanpa Air:

Bash
go run cmd/main.go
📂 Struktur Folder
cmd/: Entry point aplikasi.

initializer/: Inisialisasi konfigurasi, database, dan logging.

internal/: Logika bisnis yang dibagi per domain (Auth, Post, dll).

utils/: Helper fungsi yang bersifat umum.

Dikembangkan dengan ❤️ oleh Alg.