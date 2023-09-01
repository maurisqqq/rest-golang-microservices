package inputhargaservice

import (
	models "jojonomic/modules/utilities/models/harga"
	hargaRepository "jojonomic/modules/utilities/repositories/harga"
)

type Service interface {
	InputHarga(input models.HargaInput) error
}

type service struct {
	hargaRepository hargaRepository.Repository
}

func NewService(hargaRepository hargaRepository.Repository) *service {
	return &service{hargaRepository}
}
