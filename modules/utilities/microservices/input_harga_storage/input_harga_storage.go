package inputhargastorage

import (
	"encoding/json"
	models "jojonomic/modules/utilities/models/harga"
)

func (s *service) StoreHarga(request []byte) error {
	var modelsHargaInput models.HargaInput

	err := json.Unmarshal(request, &modelsHargaInput)
	if err != nil {
		return err
	}

	harga := models.Harga{
		AdminID:      modelsHargaInput.AdminID,
		HargaTopup:   modelsHargaInput.HargaTopup,
		HargaBuyback: modelsHargaInput.HargaBuyback,
	}

	err = s.hargaRepository.UpdateHarga(harga)
	if err != nil {
		return err
	}
	return nil
}
