package models

import (
	"time"

	"gorm.io/gorm"
)

type Harga struct {
	HargaID      string         `gorm:"primaryKey" json:"-"`
	AdminID      string         `gorm:"size:255;not null" json:"-"`
	HargaTopup   int64          `gorm:"not null" json:"harga_topup"`
	HargaBuyback int64          `gorm:"not null" json:"harga_buyback"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
