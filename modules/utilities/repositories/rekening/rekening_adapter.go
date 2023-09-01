package rekening

import (
	models "jojonomic/modules/utilities/models/rekening"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateRekeningSaldo(input models.Rekening) error
	GetUserInfo(norek string) (models.Rekening, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
