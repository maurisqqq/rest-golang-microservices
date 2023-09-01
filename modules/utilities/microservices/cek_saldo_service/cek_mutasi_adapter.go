package ceksaldoservice

import (
	rekeningModels "jojonomic/modules/utilities/models/rekening"
	rekeningRepository "jojonomic/modules/utilities/repositories/rekening"
)

type Service interface {
	CekSaldo(input rekeningModels.CekSaldoInput) (rekeningModels.Rekening, error)
}

type service struct {
	rekeningRepository rekeningRepository.Repository
}

func NewService(rekeningRepository rekeningRepository.Repository) *service {
	return &service{rekeningRepository}
}
