package models

type HargaInput struct {
	AdminID      string `json:"admin_id" validate:"required"`
	HargaTopup   int64  `json:"harga_topup" validate:"required"`
	HargaBuyback int64  `json:"harga_buyback" validate:"required"`
}
