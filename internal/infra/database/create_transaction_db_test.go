package database_test

import (
	"errors"
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/database"
	"github.com/mrpsousa/api/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func returnDBInstance() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestCreateTransactionSuccess(t *testing.T) {
	var tt entity.Transaction
	db, err := returnDBInstance()

	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Transaction{})
	transaction, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03:00", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.Nil(t, err)
	TransactionDB := database.NewTransaction(db)
	err = TransactionDB.Create(transaction)

	assert.Nil(t, err)
	assert.NotNil(t, transaction.ID)
	err = db.First(&tt, "id = ?", transaction.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, transaction.ID, tt.ID)

}

func TestCreateTransactionFail(t *testing.T) {
	var tt entity.Transaction
	expectedError := errors.New("specific_error")

	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Transaction{})
	transaction, err := entity.NewTransaction(2, "2022-02-19T05:33:07-03:00", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.NoError(t, err)
	TransactionDB := &mocks.TransactionInterface{}
	TransactionDB.On("Create", transaction).Return(expectedError)
	err = TransactionDB.Create(transaction)

	assert.NotNil(t, err)
	err = db.First(&tt, "id = ?", transaction.ID).Error
	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error())
}
