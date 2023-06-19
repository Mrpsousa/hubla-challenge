package database_test

import (
	"errors"
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/database"
	"github.com/mrpsousa/api/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetProductorBalanceSuccess(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Transaction{})
	TransactionDB := database.NewTransaction(db)

	transaction, err := entity.NewTransaction(1, "2022-02-01T23:35:43-03", "DESENVOLVEDOR FULL STACK", "ELIANA NOGUEIRA", 155000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction)
	assert.Nil(t, err)

	transaction2, err := entity.NewTransaction(2, "2022-02-03T17:23:37-03", "DESENVOLVEDOR FULL STACK", "CARLOS BATISTA", 155000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction2)
	assert.Nil(t, err)

	transaction3, err := entity.NewTransaction(4, "2022-02-03T17:23:37-03", "DESENVOLVEDOR FULL STACK", "CARLOS BATISTA", 50000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction3)
	assert.Nil(t, err)

	producers, err := TransactionDB.GetProductorBalance()
	assert.Nil(t, err)
	assert.NotNil(t, producers)
	assert.Equal(t, "ELIANA NOGUEIRA", producers[0].Seller)
	assert.Equal(t, float64(2600), producers[0].TValue)
}

func TestGetAssociateBalanceSuccess(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Transaction{})
	TransactionDB := database.NewTransaction(db)

	transaction, err := entity.NewTransaction(2, "2022-02-03T17:23:37-03", "DESENVOLVEDOR BACKEND", "CARLOS BATISTA", 255000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction)
	assert.Nil(t, err)

	transaction2, err := entity.NewTransaction(4, "2022-02-03T17:23:37-03", "DESENVOLVEDOR BACKEND", "CARLOS BATISTA", 70000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction2)
	assert.Nil(t, err)

	transaction3, err := entity.NewTransaction(2, "2022-02-03T17:23:37-03", "DESENVOLVEDOR FULL STACK", "CARLOS BATISTA", 155000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction3)
	assert.Nil(t, err)

	transaction4, err := entity.NewTransaction(4, "2022-02-03T17:23:37-03", "DESENVOLVEDOR FULL STACK", "CARLOS BATISTA", 50000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction4)
	assert.Nil(t, err)

	associate, err := TransactionDB.GetAssociateBalance()
	assert.Nil(t, err)
	assert.NotNil(t, associate)
	assert.Equal(t, "CARLOS BATISTA", associate[0].Seller)
	assert.Equal(t, float64(1200), associate[0].TValue)
}

func TestGetAssociateBalanceFail(t *testing.T) {
	// var tt entity.Transaction
	expectedError := errors.New("specific_error")
	// db, err := returnDBInstance()
	// if err != nil {
	// 	t.Error(err)
	// }

	TransactionDB := &mocks.TransactionInterface{}

	TransactionDB.On("GetAssociateBalance", nil).Return(expectedError)
	associantes, err := TransactionDB.GetAssociateBalance()
	assert.NotNil(t, err)
	assert.Nil(t, associantes)

	// assert.NotNil(t, err)
	// assert.Equal(t, "record not found", err.Error())
}
