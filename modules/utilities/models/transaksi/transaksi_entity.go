package transaksi

import (
	"time"

	"gorm.io/gorm"
)

//handling debit dan credit
type Transaksi struct {
	TransaskiID  string         `gorm:"primaryKey" json:"transaksi_id"`
	Type         string         `gorm:"size:255;not null" json:"type"`
	Gram         float64        `gorm:"not null" json:"gram"`
	HargaTopup   int64          `gorm:"not null" json:"harga_topup"`
	HargaBuyback int64          `gorm:"not null" json:"harga_buyback"`
	RekeningID   string         `gorm:"size:255;not null" json:"rekening_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type MutasiRes struct {
	Type         string  `gorm:"size:255;not null" json:"type"`
	Gram         float64 `gorm:"not null" json:"gram"`
	HargaTopup   int64   `gorm:"not null" json:"harga_topup"`
	HargaBuyback int64   `gorm:"not null" json:"harga_buyback"`
	Date         int64   `json:"date"`
}
