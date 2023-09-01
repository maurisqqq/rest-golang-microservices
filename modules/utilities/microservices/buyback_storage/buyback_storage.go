package buybackstorage

import (
	"encoding/json"
	rekeningModels "jojonomic/modules/utilities/models/rekening"
	transaksiModels "jojonomic/modules/utilities/models/transaksi"

	"strconv"

	"github.com/teris-io/shortid"
)

func (s *service) StoreBuybackData(request []byte) error {
	var transaksiInput transaksiModels.TransaksiInput

	err := json.Unmarshal(request, &transaksiInput)
	if err != nil {
		return err
	}
	id, err := shortid.Generate()
	if err != nil {
		return err
	}
	//Get initial harga
	initialHarga, err := s.hargaRepository.CheckHarga()
	if err != nil {
		return err
	}
	//Get Rekening id from norek
	userData, err := s.rekeningRepository.GetUserInfo(transaksiInput.Norek)
	if err != nil {
		return err
	}
	//Parse data to float
	gramActual, err := strconv.ParseFloat(transaksiInput.Gram, 64)
	if err != nil {
		return err
	}
	//Store transaksi
	transaksiTopup := transaksiModels.Transaksi{
		TransaskiID:  id,
		Type:         "buyback",
		Gram:         gramActual,
		HargaTopup:   initialHarga.HargaTopup,
		HargaBuyback: initialHarga.HargaBuyback,
		RekeningID:   userData.RekeningID,
	}
	err = s.transaksiRepository.StoreTransaksi(transaksiTopup)
	if err != nil {
		return err
	}
	//Update saldo
	newestSaldo := userData.Saldo - gramActual
	updateSald0 := rekeningModels.Rekening{
		Norek: transaksiInput.Norek,
		Saldo: newestSaldo,
	}
	err = s.rekeningRepository.UpdateRekeningSaldo(updateSald0)
	return nil
}
