package buybackservice

import (
	"encoding/json"
	"errors"
	"jojonomic/app/connections"
	transaksiModels "jojonomic/modules/utilities/models/transaksi"
	"jojonomic/pkg"
	"strconv"

	"gorm.io/gorm"
)

func (s *service) InputDataBuyback(input transaksiModels.TransaksiInput) error {
	//Validate request
	err := pkg.ValidateRequest(input.Gram)
	if err != nil {
		return err
	}

	//Get initial harga & validate request
	initialHarga, err := s.hargaRepository.CheckHarga()
	if err != nil {
		return err
	}
	harga, _ := strconv.ParseInt(input.Harga, 10, 64)
	if harga != initialHarga.HargaBuyback {
		return errors.New("Data harga tidak sama!")
	}

	//Validate norek
	_, err = s.rekeningRepository.GetUserInfo(input.Norek)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Norek tidak ditemukan!")
		}
		return err
	}

	//Check saldo
	userInfo, err := s.rekeningRepository.GetUserInfo(input.Norek)
	if err != nil {
		return err
	}
	gramActual, err := strconv.ParseFloat(input.Gram, 64)
	if err != nil {
		return err
	}
	if userInfo.Saldo < gramActual {
		return errors.New("Saldo anda tidak mencukupi")
	}

	//Publish and send to kafka
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}
	kafka, err := connections.PublisherConnection()
	if err != nil {
		return err
	}
	err = kafka.SendData("buyback", data)
	if err != nil {
		return err
	}
	return nil
}
