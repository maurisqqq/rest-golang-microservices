package transaksi

import (
	"encoding/json"
	"errors"
	"net/http"

	modelTransaksi "jojonomic/modules/utilities/models/transaksi"
	"jojonomic/pkg"
)

func (h *TransaksiHandler) Topup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//valiidate request
	var topupInput modelTransaksi.TransaksiInput
	err := json.NewDecoder(r.Body).Decode(&topupInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}
	validatorErr := pkg.Validator(topupInput)
	if validatorErr != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: errors.New(validatorErr)}))
		return
	}

	err = h.topupService.InputDataTopUp(topupInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: false, Method: "updt"}))
	return
}

func (h *TransaksiHandler) Buyback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//valiidate request
	var buybackInput modelTransaksi.TransaksiInput
	err := json.NewDecoder(r.Body).Decode(&buybackInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}
	validatorErr := pkg.Validator(buybackInput)
	if validatorErr != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: errors.New(validatorErr)}))
		return
	}

	err = h.buybackService.InputDataBuyback(buybackInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: false, Method: "updt"}))
	return
}

func (h *TransaksiHandler) CekMutasi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//valiidate request
	var cekMutasiInput modelTransaksi.CekMutasiInput
	err := json.NewDecoder(r.Body).Decode(&cekMutasiInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}
	validatorErr := pkg.Validator(cekMutasiInput)
	if validatorErr != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: errors.New(validatorErr)}))
		return
	}

	data, err := h.cekMutasiService.CekMutasi(cekMutasiInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: false, Data: data}))
	return
}
