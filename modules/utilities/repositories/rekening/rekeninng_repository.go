package rekening

import (
	models "jojonomic/modules/utilities/models/rekening"
)

func (r *repository) UpdateRekeningSaldo(input models.Rekening) error {
	err := r.db.Where("norek = ?", input.Norek).Updates(&input).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUserInfo(norek string) (models.Rekening, error) {
	var rekening models.Rekening
	err := r.db.Where("norek = ?", norek).First(&rekening).Error
	if err != nil {
		return models.Rekening{}, err
	}
	return rekening, nil
}
