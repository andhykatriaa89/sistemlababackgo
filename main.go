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

	// HPP Routes
	http.HandleFunc("/api/bahan-baku", middleware.CORS(handlers.GetBahanBaku))
	http.HandleFunc("/api/bahan-baku/save", middleware.CORS(handlers.SaveBahanBaku)) // Combine Create/Update
	http.HandleFunc("/api/bahan-baku/delete", middleware.CORS(handlers.DeleteBahanBaku))
	http.HandleFunc("/api/produk", middleware.CORS(handlers.GetProduk))
	http.HandleFunc("/api/produk/save", middleware.CORS(handlers.CreateProduk))
	http.HandleFunc("/api/produk/delete", middleware.CORS(handlers.DeleteProduk))

	// Fixed Cost Routes
	http.HandleFunc("/api/biaya-tetap", middleware.CORS(handlers.GetBiayaTetap))
	http.HandleFunc("/api/biaya-tetap/save", middleware.CORS(handlers.SaveBiayaTetap))
	http.HandleFunc("/api/biaya-tetap/delete", middleware.CORS(handlers.DeleteBiayaTetap))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Sistem Laba API berjalan di http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
