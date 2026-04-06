package models

import (
	"time"

	"github.com/lib/pq"
)

type Transaksi struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Pendapatan float64        `json:"pendapatan"`
	Modal      float64        `json:"modal"`
	LabaBersih float64        `json:"laba_bersih"`
	Margin     float64        `json:"margin"`
	Status     string         `json:"status"`
	Catatan    string         `json:"catatan"`
	Items      pq.StringArray `json:"items" gorm:"type:text[]"`
	CreatedAt  time.Time      `json:"created_at"`
}
type BahanBaku struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Nama           string    `json:"nama"`
	Satuan         string    `json:"satuan"`
	HargaPerSatuan float64   `json:"harga_per_satuan"`
	CreatedAt      time.Time `json:"created_at"`
}

type ResepItem struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ProdukID    uint      `json:"produk_id"`
	BahanBakuID uint      `json:"bahan_baku_id"`
	BahanBaku   BahanBaku `json:"bahan_baku" gorm:"foreignKey:BahanBakuID"`
	Jumlah      float64   `json:"jumlah"`
}

type Produk struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Nama      string      `json:"nama"`
	HPP       float64     `json:"hpp"`
	Resep     []ResepItem `json:"resep" gorm:"foreignKey:ProdukID"`
	CreatedAt time.Time   `json:"created_at"`
}
