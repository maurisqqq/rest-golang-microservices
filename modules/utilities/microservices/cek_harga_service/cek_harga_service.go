package cekhargaservice

import (
	models "jojonomic/modules/utilities/models/harga"
)

func (s *service) CekHarga() (models.Harga, error) {
	dataHarga, err := s.hargaRepository.CheckHarga()
	if err != nil {
		return models.Harga{}, err
	}
	return dataHarga, nil
}
