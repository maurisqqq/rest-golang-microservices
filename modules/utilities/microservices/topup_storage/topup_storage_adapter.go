package topupstorage

import (
	hargaRepository "jojonomic/modules/utilities/repositories/harga"
	rekeningRepository "jojonomic/modules/utilities/repositories/rekening"
	transaksiRepository "jojonomic/modules/utilities/repositories/transaksi"
)

type Service interface {
	StoreTopupData(request []byte) error
}

type service struct {
	hargaRepository     hargaRepository.Repository
	rekeningRepository  rekeningRepository.Repository
	transaksiRepository transaksiRepository.Repository
}

func NewService(hargaRepository hargaRepository.Repository, rekeningRepository rekeningRepository.Repository, transaksiRepository transaksiRepository.Repository) *service {
	return &service{hargaRepository, rekeningRepository, transaksiRepository}
}
