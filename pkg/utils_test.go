package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testInt   = "10"
	testFloat = "20.5"
	testError = "42-34*2&3"
)

func TestStringToInt8Success(t *testing.T) {
	tr, err := StringToInt8(testInt)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, int8(10), tr)
}

func TestStringToInt8(t *testing.T) {
	tr, err := StringToInt8(testError)
	assert.NotNil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, int8(0), tr)
	assert.Equal(t, "Failed type convert", err.Error())

}

func TestStringToFloat64Success(t *testing.T) {
	tr, err := StringToFloat64(testFloat)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, float64(20.5), tr)
}

func TestStringToFloat64(t *testing.T) {
	tr, err := StringToFloat64(testError)
	assert.NotNil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, float64(0), tr)
	assert.Equal(t, "Failed type convert", err.Error())
}

// func FileNameGenerate(file string) string {
// 	return uuid.New().String() + "-" + file
// }
