package cekmutasiservice

import (
	transaksiModels "jojonomic/modules/utilities/models/transaksi"
	rekeningRepository "jojonomic/modules/utilities/repositories/rekening"
	transaksiRepository "jojonomic/modules/utilities/repositories/transaksi"
)

type Service interface {
	CekMutasi(input transaksiModels.CekMutasiInput) ([]transaksiModels.MutasiRes, error)
}

type service struct {
	rekeningRepository  rekeningRepository.Repository
	transaksiRepository transaksiRepository.Repository
}

func NewService(rekeningRepository rekeningRepository.Repository, transaksiRepository transaksiRepository.Repository) *service {
	return &service{rekeningRepository, transaksiRepository}
}
