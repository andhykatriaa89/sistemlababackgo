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
