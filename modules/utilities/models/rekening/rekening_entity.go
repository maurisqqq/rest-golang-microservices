package rekening

import (
	"time"

	"gorm.io/gorm"

	"jojonomic/modules/utilities/models/transaksi"
)

//handling update saldo saja
type Rekening struct {
	RekeningID string                `gorm:"primaryKey" json:"-"`
	Norek      string                `gorm:"size:255;not null" json:"norek"`
	Saldo      float64               `gorm:"not null" json:"saldo"`
	Transaksi  []transaksi.Transaksi `gorm:"foreignKey:RekeningID;refrences:RekeningID;constraint:OnUpdate:CASCADE,OnDelete:Cascade" json:"-"`
	CreatedAt  time.Time             `json:"-"`
	UpdatedAt  time.Time             `json:"-"`
	DeletedAt  gorm.DeletedAt        `gorm:"index" json:"-"`
}
