package pkg

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testInt        = "10"
	testFloat      = "20.5"
	testError      = "42-34*2&3"
	stringName     = "test-name"
	portugueseText = "Isso é um texto em PT-BR"
	foreignText    = "This is a foreign text"
	typeIsEqual    bool
)

func TestStringToInt8Success(t *testing.T) {
	tr, err := StringToInt8(testInt)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, int8(10), tr)
}

func TestStringToInt8Fail(t *testing.T) {
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

func TestStringToFloat64Fail(t *testing.T) {
	tr, err := StringToFloat64(testError)
	assert.NotNil(t, err)
	assert.NotNil(t, tr)
	assert.Equal(t, float64(0), tr)
	assert.Equal(t, "Failed type convert", err.Error())
}

func TestFileNameGenerate(t *testing.T) {
	tr := FileNameGenerate(stringName)
	assert.NotNil(t, tr)
	if reflect.TypeOf(tr) == reflect.TypeOf(stringName) {
		typeIsEqual = true
	}
	assert.True(t, typeIsEqual)
}

func TestForeignProductValidateIsForeign(t *testing.T) {
	tr := ForeignProductValidate(foreignText)
	assert.NotNil(t, tr)
	assert.True(t, tr)
}

func TestForeignProductValidateNotIsForeign(t *testing.T) {
	tr := ForeignProductValidate(portugueseText)
	assert.NotNil(t, tr)
	assert.False(t, tr)
}
