package cekmutasiservice

import (
	"errors"
	transaksiModels "jojonomic/modules/utilities/models/transaksi"
	"time"
)

func (s *service) CekMutasi(input transaksiModels.CekMutasiInput) ([]transaksiModels.MutasiRes, error) {
	//Get user info
	userInfo, err := s.rekeningRepository.GetUserInfo(input.Norek)
	if err != nil {
		return nil, err
	}
	startDate := time.Unix(input.StartDate, 0)
	endDate := time.Unix(input.EndDate, 0)
	infoMutasi, err := s.transaksiRepository.MutasiTransaksi(userInfo.RekeningID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	if len(infoMutasi) == 0 {
		return nil, errors.New("record not found")
	}
	//Manage response
	var resMustasi []transaksiModels.MutasiRes
	for _, dataMutasi := range infoMutasi {
		res := transaksiModels.MutasiRes{
			Type:         dataMutasi.Type,
			Gram:         dataMutasi.Gram,
			HargaTopup:   dataMutasi.HargaTopup,
			HargaBuyback: dataMutasi.HargaBuyback,
			Date:         dataMutasi.CreatedAt.Unix(),
		}
		resMustasi = append(resMustasi, res)
	}
	return resMustasi, nil
}
