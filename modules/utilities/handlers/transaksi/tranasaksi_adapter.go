package transaksi

import (
	hargaRepository "jojonomic/modules/utilities/repositories/harga"
	rekeningRepository "jojonomic/modules/utilities/repositories/rekening"
	transaksiRepository "jojonomic/modules/utilities/repositories/transaksi"

	buybackService "jojonomic/modules/utilities/microservices/buyback_service"
	cekMutasiService "jojonomic/modules/utilities/microservices/cek_mutasi_service"
	topupService "jojonomic/modules/utilities/microservices/topup_service"

	"gorm.io/gorm"
)

type TransaksiHandler struct {
	topupService     topupService.Service
	buybackService   buybackService.Service
	cekMutasiService cekMutasiService.Service
}

func NewTransaksiHandler(db *gorm.DB) *TransaksiHandler {
	rekeningRepository := rekeningRepository.NewRepository(db)
	transaksiRepository := transaksiRepository.NewRepository(db)
	hargaRepository := hargaRepository.NewRepository(db)
	topupService := topupService.NewService(hargaRepository, rekeningRepository)
	buybackService := buybackService.NewService(hargaRepository, rekeningRepository)
	cekMutasiService := cekMutasiService.NewService(rekeningRepository, transaksiRepository)

	return &TransaksiHandler{topupService, buybackService, cekMutasiService}
}
