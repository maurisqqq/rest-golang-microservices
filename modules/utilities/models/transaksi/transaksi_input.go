package transaksi

type TransaksiInput struct {
	Gram  string `json:"gram" validate:"required"`
	Harga string `json:"harga" validate:"required"`
	Norek string `json:"norek" validate:"required"`
}

type CekMutasiInput struct {
	Norek     string `json:"norek" validate:"required"`
	StartDate int64  `json:"start_date" validate:"required"`
	EndDate   int64  `json:"end_date" validate:"required"`
}
