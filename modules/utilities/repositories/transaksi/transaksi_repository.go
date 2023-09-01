package transaksi

import (
	models "jojonomic/modules/utilities/models/transaksi"
	"time"
)

func (r *repository) StoreTransaksi(input models.Transaksi) error {
	err := r.db.Create(&input).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) MutasiTransaksi(id string, startDate, endDate time.Time) ([]models.Transaksi, error) {
	var transaksi []models.Transaksi
	err := r.db.Where("rekening_id = ?", id).Where("created_at Between ? And ?", startDate, endDate).Order("created_at desc").Find(&transaksi).Error
	if err != nil {
		return nil, err
	}
	return transaksi, nil
}
