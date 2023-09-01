package ceksaldoservice

import (
	rekeningModels "jojonomic/modules/utilities/models/rekening"
)

func (s *service) CekSaldo(input rekeningModels.CekSaldoInput) (rekeningModels.Rekening, error) {
	infoSaldo, err := s.rekeningRepository.GetUserInfo(input.Norek)
	if err != nil {
		return rekeningModels.Rekening{}, err
	}
	return infoSaldo, nil
}
