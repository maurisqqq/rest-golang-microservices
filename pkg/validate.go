package pkg

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
)

func Validator(dataBind interface{}) string {
	var validate = validator.New()

	if err := validate.Struct(dataBind); err != nil {

		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			dataValid := ValidateField(ve)
			return dataValid
		}
		return "Invalid input!"
	}
	return ""
}

func ValidateField(ve validator.ValidationErrors) string {
	out := make([]string, len(ve))
	for i, fe := range ve {
		out[i] = fe.Field() + " is a required field"
	}
	return out[0]
}

func ValidateRequest(gram string) error {
	//validate minimum data
	gramActual, err := strconv.ParseFloat(gram, 64)
	if err != nil {
		return err
	}
	if gramActual < 0.001 {
		return errors.New("Jumlah emas minamal 0.001")
	}

	//validate 3 digits behind comas
	val := strings.Split(gram, ".")
	valArr := strings.Split(val[1], "")
	if len(valArr) > 3 {
		return errors.New("Jumlah emas tidak valid")
	}

	//validate data kelipatan 0.001
	countData := math.Mod(gramActual, 0.001)
	if countData > 0 {
		return errors.New("Jumalh emas hanya boleh kelipatan 0.001")
	}
	return nil
}
