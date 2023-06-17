package pkg

import (
	"errors"
	"strconv"

	"github.com/google/uuid"
)

var (
	ErrInvalidID      = errors.New("id_is_required_or_invalid")
	ErrInvalidType    = errors.New("type_is_required_or_invalid")
	ErrInvalidDate    = errors.New("date_is_required_or_invalid")
	ErrInvalidProduct = errors.New("product_is_required_or_invalid")
	ErrInvalidValue   = errors.New("value_is_required_or_invalid")
	ErrInvalidSeller  = errors.New("seller_is_required_or_invalid")
)

func StringToInt8(str string) (int8, error) {
	num, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, errors.New("Failed type convert")
	}
	return int8(num), nil
}

func StringToFloat64(str string) (float64, error) {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, errors.New("Failed type convert")
	}
	return num, nil
}

func FileNameGenerate(file string) string {
	return uuid.New().String() + "-" + file
}
