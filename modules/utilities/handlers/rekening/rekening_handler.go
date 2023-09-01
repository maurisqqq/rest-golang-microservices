package rekening

import (
	"encoding/json"
	"errors"
	"net/http"

	modelRekening "jojonomic/modules/utilities/models/rekening"
	"jojonomic/pkg"
)

func (h *RekeningHandler) CekSaldo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//valiidate request
	var cekSaldoInput modelRekening.CekSaldoInput
	err := json.NewDecoder(r.Body).Decode(&cekSaldoInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}
	validatorErr := pkg.Validator(cekSaldoInput)
	if validatorErr != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: errors.New(validatorErr)}))
		return
	}

	data, err := h.cekSaldoService.CekSaldo(cekSaldoInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: true, Err: err}))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pkg.GeneralResponse(pkg.ResponseParam{Val: false, Data: data}))
	return
}
