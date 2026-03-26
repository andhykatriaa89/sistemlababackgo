package store

import (
	"log"
	"os"
	"sistem-laba/models"

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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	err = db.AutoMigrate(&models.Transaksi{})
	if err != nil {
		log.Fatalf("Gagal migrasi database: %v", err)
	}
	DB = db
	log.Println("Database terhubung & migrasi berhasil")
}

func CreateTransaksi(t models.Transaksi) (models.Transaksi, error) {
	result := DB.Create(&t)
	return t, result.Error
}

func GetAllTransaksi() []models.Transaksi {
	var transaksi []models.Transaksi
	DB.Order("created_at desc").Find(&transaksi)
	return transaksi
}

func DeleteTransaksi(id uint) error {
	result := DB.Delete(&models.Transaksi{}, id)
	return result.Error
}
