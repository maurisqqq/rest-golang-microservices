package topupservice

import (
	"encoding/json"
	"errors"
	"jojonomic/app/connections"
	transaksiModels "jojonomic/modules/utilities/models/transaksi"
	"jojonomic/pkg"
	"strconv"

	"gorm.io/gorm"
)

func (s *service) InputDataTopUp(input transaksiModels.TransaksiInput) error {
	//Validate request
	err := pkg.ValidateRequest(input.Gram)
	if err != nil {
		return err
	}

	//Get initial harga
	initialHarga, err := s.hargaRepository.CheckHarga()
	if err != nil {
		return err
	}

	harga, _ := strconv.ParseInt(input.Harga, 10, 64)
	if harga != initialHarga.HargaTopup {
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

	//Publish and send to kafka
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	kafka, err := connections.PublisherConnection()
	if err != nil {
		return err
	}

	err = kafka.SendData("topup", data)
	if err != nil {
		return err
	}
	return nil
}
