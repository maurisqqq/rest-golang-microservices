package harga

import (
	models "jojonomic/modules/utilities/models/harga"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateHarga(input models.Harga) error
	CheckHarga() (models.Harga, error)
	ListHarga(id string) (models.Harga, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
