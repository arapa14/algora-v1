package main

import (
	"backend/initializer"
	"backend/utils/shared"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Load Konfigurasi
	initializer.LoadConfig()
	cfg := initializer.AppConfig

	// Inisialisasi Database
	db, err := initializer.NewPostgres(cfg.DB)
	if err != nil {
		log.Fatalf("Gagal inisialisasi database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error saat menutup database: %v", err)
		}
	}()
	log.Println("✅ Database terkoneksi")

	// Routing
	mux := http.NewServeMux()

	// Handler Dummy untuk tes awal
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Welcome to Algora V1 API"}`))
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		shared.RespondSuccess(w, http.StatusOK, "Server is healthy", map[string]string{
			"version": "1.0.0",
			"env":     initializer.AppConfig.Server.ENV,
		})
	})

	// mux.HandleFunc("/register", authHandler.Register)

	// Inisialisasi Server
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Menjalankan Server di Goroutine (agar tidak blocking untuk graceful shutdown)
	go func() {
		log.Printf("🚀 Server running on port %s (%s mode)\n", cfg.Server.Port, cfg.Server.ENV)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Gagal menjalankan server: %v", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Memberikan waktu 5 detik untuk menyelesaikan request yang sedang berjalan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited gracefully")
}