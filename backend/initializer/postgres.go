package initializer

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Driver wajib di-import di sini jika tidak di main
)

func NewPostgres(config DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
		config.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// --- CONFIGURATION (Connection Pool) ---
	// Sesuaikan angka ini dengan kebutuhan servermu
	db.SetMaxOpenConns(25)                 // Maksimal 25 koneksi terbuka
	db.SetMaxIdleConns(25)                 // Maksimal 25 koneksi idle
	db.SetConnMaxLifetime(5 * time.Minute) // Koneksi di-refresh tiap 5 menit
	db.SetConnMaxIdleTime(2 * time.Minute) // Koneksi idle dibuang setelah 2 menit

	// --- PING WITH TIMEOUT ---
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}