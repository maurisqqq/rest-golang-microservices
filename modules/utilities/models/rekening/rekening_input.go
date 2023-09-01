package rekening

type CekSaldoInput struct {
	Norek string `json:"norek" validate:"required"`
}
