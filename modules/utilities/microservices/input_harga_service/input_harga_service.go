package inputhargaservice

import (
	"encoding/json"
	"errors"
	"jojonomic/app/connections"
	models "jojonomic/modules/utilities/models/harga"

	"gorm.io/gorm"
)

func (s *service) InputHarga(input models.HargaInput) error {
	// Marshing input model
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}
	_, err = s.hargaRepository.ListHarga(input.AdminID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Admin id not found!")
		}
		return err
	}

	// Publish to kafka
	kafka, err := connections.PublisherConnection()
	if err != nil {
		return err
	}
	err = kafka.SendData("input_harga", data)
	if err != nil {
		return err
	}
	return nil
}
