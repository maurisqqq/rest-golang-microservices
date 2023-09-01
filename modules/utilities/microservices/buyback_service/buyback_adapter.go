package buybackservice

import (
	transaksiModels "jojonomic/modules/utilities/models/transaksi"
	hargaRepository "jojonomic/modules/utilities/repositories/harga"
	rekeningRepository "jojonomic/modules/utilities/repositories/rekening"
)

type Service interface {
	InputDataBuyback(input transaksiModels.TransaksiInput) error
}

type service struct {
	hargaRepository    hargaRepository.Repository
	rekeningRepository rekeningRepository.Repository
}

func NewService(hargaRepository hargaRepository.Repository, rekeningRepository rekeningRepository.Repository) *service {
	return &service{hargaRepository, rekeningRepository}
}
