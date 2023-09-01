package transaksi

import (
	models "jojonomic/modules/utilities/models/transaksi"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	StoreTransaksi(input models.Transaksi) error
	MutasiTransaksi(id string, startDate, endDate time.Time) ([]models.Transaksi, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
