package harga

import (
	models "jojonomic/modules/utilities/models/harga"
)

func (r *repository) UpdateHarga(input models.Harga) error {
	err := r.db.Where("admin_id = ?", input.AdminID).Updates(&input).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CheckHarga() (models.Harga, error) {
	var harga models.Harga
	err := r.db.Order("created_at desc").First(&harga).Error
	if err != nil {
		return models.Harga{}, err
	}
	return harga, nil
}

func (r *repository) ListHarga(id string) (models.Harga, error) {
	var harga models.Harga
	err := r.db.Where("admin_id = ?", id).First(&harga).Error
	if err != nil {
		return models.Harga{}, err
	}
	return harga, nil
}
