package database_test

import (
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/database"
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

func TestCreateTransaction(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Transaction{})
	transaction, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
	assert.NoError(t, err)
	TransactionDB := database.NewTransaction(db)
	err = TransactionDB.Create(transaction)
	assert.NoError(t, err)
	assert.NotEmpty(t, transaction.ID)

}
