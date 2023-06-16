package pkg

import (
	"strconv"
)

func StringToInt8(str string) (int8, error) {
	num, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(num), nil
}

func StringToFloat64(str string) (float64, error) {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
