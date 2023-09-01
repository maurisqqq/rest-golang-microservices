package harga

import (
	hargaRepository "jojonomic/modules/utilities/repositories/harga"

	cekHargaService "jojonomic/modules/utilities/microservices/cek_harga_service"
	hargaService "jojonomic/modules/utilities/microservices/input_harga_service"

	"gorm.io/gorm"
)

type HargaHandler struct {
	hargaService    hargaService.Service
	cekHargaService cekHargaService.Service
}

func NewHargaHandler(db *gorm.DB) *HargaHandler {
	hargaRepository := hargaRepository.NewRepository(db)
	hargaService := hargaService.NewService(hargaRepository)
	cekHargaService := cekHargaService.NewService(hargaRepository)

	return &HargaHandler{hargaService, cekHargaService}
}
