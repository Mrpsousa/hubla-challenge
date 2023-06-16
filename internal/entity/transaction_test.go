package entity_test

import (
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/stretchr/testify/assert"
)

var (
	ErrInvalidType    = "type_is_required_or_invalid"
	ErrInvalidDate    = "date_is_required_or_invalid"
	ErrInvalidProduct = "product_is_required_or_invalid"
	ErrInvalidValue   = "value_is_required_or_invalid"
	ErrInvalidSeller  = "seller_is_required_or_invalid"
)

func TestNewTransaction(t *testing.T) {
	tr, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	assert.NotEmpty(t, tr.ID)
	assert.Equal(t, "MARIA CANDIDA", tr.Seller)
	assert.Equal(t, 50000.0, tr.Value)
}

func TestTransactionWhenTypeInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(0, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, ErrInvalidType, err.Error())
}

func TestTransactionWhenDateInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(1, "", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, ErrInvalidDate, err.Error())
}

func TestTransactionWhenProductInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "", "MARIA CANDIDA", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, ErrInvalidProduct, err.Error())
}

func TestTransactionWhenSellerInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, ErrInvalidSeller, err.Error())
}
func TestTransactionWhenValueInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", -1)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, ErrInvalidValue, err.Error())
}

func TestTransactionWhenValueInvalidBiggerThenFour(t *testing.T) {
	tr, err := entity.NewTransaction(5, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, ErrInvalidType, err.Error())
}
