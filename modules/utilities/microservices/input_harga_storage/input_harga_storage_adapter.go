package inputhargastorage

import hargaRepository "jojonomic/modules/utilities/repositories/harga"

type Service interface {
	StoreHarga(request []byte) error
}

type service struct {
	hargaRepository hargaRepository.Repository
}

func NewService(hargaRepository hargaRepository.Repository) *service {
	return &service{hargaRepository}
}
