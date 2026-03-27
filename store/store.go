package store

import (
	"log"
	"os"
	"sistem-laba/models"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Println("DATABASE_URL tidak ditemukan")
		dsn = "host=localhost user=postgres password=postgres dbname=sistem_laba port=5432 sslmode=disable"
	}

	// For Supabase pooler (PgBouncer), disable prepared statements
	// by using simple protocol to avoid "prepared statement already exists" errors
	if strings.Contains(dsn, "pooler.supabase.com") || strings.Contains(dsn, "supabase") {
		if !strings.Contains(dsn, "default_query_exec_mode") {
			if strings.Contains(dsn, "?") {
				dsn += "&default_query_exec_mode=simple_protocol"
			} else {
				dsn += "?default_query_exec_mode=simple_protocol"
			}
		}
		log.Println("Supabase pooler detected, using simple_protocol mode")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	err = db.AutoMigrate(&models.Transaksi{})
	if err != nil {
		log.Printf("Peringatan migrasi database: %v", err)
	}
	DB = db
	log.Println("Database terhubung & sinkronisasi selesai")
}

func CreateTransaksi(t models.Transaksi) (models.Transaksi, error) {
	result := DB.Create(&t)
	return t, result.Error
}

func GetAllTransaksi() []models.Transaksi {
	transaksi := make([]models.Transaksi, 0)
	result := DB.Order("created_at desc").Find(&transaksi)
	if result.Error != nil {
		log.Printf("Error fetching transaksi: %v", result.Error)
	}
	return transaksi
}

func DeleteTransaksi(id uint) error {
	result := DB.Delete(&models.Transaksi{}, id)
	return result.Error
}
