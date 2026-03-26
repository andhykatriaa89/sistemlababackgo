package main

import (
	"log"
	"net/http"
	"os"
	"sistem-laba/handlers"
	"sistem-laba/middleware"
	"sistem-laba/store"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using OS environment variables")
	}

	store.InitDB()

	http.HandleFunc("/api/hitung-laba", middleware.CORS(handlers.HitungLaba))
	http.HandleFunc("/api/transaksi", middleware.CORS(handlers.GetTransaksi))
	http.HandleFunc("/api/transaksi/delete", middleware.CORS(handlers.DeleteTransaksi))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Sistem Laba API berjalan di http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
