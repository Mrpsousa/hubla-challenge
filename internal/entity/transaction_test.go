package entity_test

import (
	"reflect"
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/pkg"
	"github.com/stretchr/testify/assert"
)

var (
	TypeIsEqual bool
)

func TestNewTransaction(t *testing.T) {
	tr, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	if reflect.TypeOf(tr) == reflect.TypeOf(&entity.Transaction{}) {
		TypeIsEqual = true
	}
	assert.NotEmpty(t, tr.ID)
	assert.True(t, TypeIsEqual)
	assert.Equal(t, "MARIA CANDIDA", tr.Seller)
	assert.Equal(t, false, tr.ForeignProduct)
	assert.Equal(t, 50000.0, tr.Value)
}

func TestNewTransactionWithForeignProduct(t *testing.T) {
	tr, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "FULL STACK DEVELOPER", "MARIA CANDIDA", 50000.0)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	if reflect.TypeOf(tr) == reflect.TypeOf(&entity.Transaction{}) {
		TypeIsEqual = true
	}
	assert.NotEmpty(t, tr.ID)
	assert.True(t, TypeIsEqual)
	assert.Equal(t, "MARIA CANDIDA", tr.Seller)
	assert.Equal(t, true, tr.ForeignProduct)
	assert.Equal(t, 50000.0, tr.Value)
}

func TestTransactionWhenTypeInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(0, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, pkg.ErrInvalidType.Error(), err.Error())
}

func TestTransactionWhenDateInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(1, "", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, pkg.ErrInvalidDate.Error(), err.Error())
}

func TestTransactionWhenProductInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "", "MARIA CANDIDA", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, pkg.ErrInvalidProduct.Error(), err.Error())
}

func TestTransactionWhenSellerInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, pkg.ErrInvalidSeller.Error(), err.Error())
}
func TestTransactionWhenValueInvalid(t *testing.T) {
	tr, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", -1)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, pkg.ErrInvalidValue.Error(), err.Error())
}

func TestTransactionWhenValueInvalidBiggerThenFour(t *testing.T) {
	tr, err := entity.NewTransaction(5, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Equal(t, pkg.ErrInvalidType.Error(), err.Error())
}
