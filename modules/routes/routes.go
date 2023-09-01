package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"jojonomic/app/config"
	hargaHandler "jojonomic/modules/utilities/handlers/harga"
	rekeningHandler "jojonomic/modules/utilities/handlers/rekening"
	transaksiHandler "jojonomic/modules/utilities/handlers/transaksi"
)

func SetupRoutes(db *gorm.DB) {
	// Set Handler
	hargaHandler := hargaHandler.NewHargaHandler(db)
	rekeningHandler := rekeningHandler.NewRekeningHandler(db)
	transaksiHandler := transaksiHandler.NewTransaksiHandler(db)
	// Set Routes
	router := mux.NewRouter()
	appRouter := router.PathPrefix("/api").Subrouter()
	appRouter.HandleFunc("/input-harga", hargaHandler.InputHarga).Methods("POST")
	appRouter.HandleFunc("/check-harga", hargaHandler.CekHarga).Methods("GET")
	appRouter.HandleFunc("/topup", transaksiHandler.Topup).Methods("POST")
	appRouter.HandleFunc("/saldo", rekeningHandler.CekSaldo).Methods("POST")
	appRouter.HandleFunc("/mutasi", transaksiHandler.CekMutasi).Methods("POST")
	appRouter.HandleFunc("/buyback", transaksiHandler.Buyback).Methods("POST")

	http.ListenAndServe(config.Env("APP_PORT"), router)
}
