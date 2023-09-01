package harga

import (
	"encoding/json"
	"errors"

	modelHarga "jojonomic/modules/utilities/models/harga"
	"jojonomic/pkg"
	"net/http"
)

func (h *HargaHandler) InputHarga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//valiidate request
	var hargaInput modelHarga.HargaInput
	err := json.NewDecoder(r.Body).Decode(&hargaInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}
	validatorErr := pkg.Validator(hargaInput)
	if validatorErr != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: errors.New(validatorErr)}))
		return
	}

	err = h.hargaService.InputHarga(hargaInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: false, Method: "updt"}))
	return
}

func (h *HargaHandler) CekHarga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := h.cekHargaService.CekHarga()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: false, Data: data}))
	return
}
