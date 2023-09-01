package rekening

import (
	rekeningRepository "jojonomic/modules/utilities/repositories/rekening"

	cekSaldoService "jojonomic/modules/utilities/microservices/cek_saldo_service"

	"gorm.io/gorm"
)

type RekeningHandler struct {
	cekSaldoService cekSaldoService.Service
}

func NewRekeningHandler(db *gorm.DB) *RekeningHandler {
	rekeningRepository := rekeningRepository.NewRepository(db)
	cekSaldoService := cekSaldoService.NewService(rekeningRepository)

	return &RekeningHandler{cekSaldoService}
}
