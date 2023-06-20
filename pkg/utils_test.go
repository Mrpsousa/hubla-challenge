package pkg

import (
	"os"
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

func TestRemoveFolder(t *testing.T) {
	tmpDir := "./tmp"
	err := os.Mkdir(tmpDir, 0755)
	assert.Nil(t, err)
	defer os.Remove(tmpDir)

	RemoveFolder(tmpDir)

	_, err = os.Stat(tmpDir)
	if os.IsExist(err) {
		t.Error("failed to remove tmp directory")
	}
}

func TestParseIDSuccess(t *testing.T) {
	validID := NewID()
	validIDString := validID.String()

	parsedID, err := ParseID(validIDString)
	assert.Nil(t, err)

	assert.Equal(t, parsedID, validID, parsedID)
}

func TestParseIDError(t *testing.T) {
	invalidIDString := "invalid_id"

	_, err := ParseID(invalidIDString)

	assert.NotNil(t, err)
}
