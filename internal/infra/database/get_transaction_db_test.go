package database_test

import (
	"testing"

	"github.com/mrpsousa/api/internal/dto"
	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestGetProductorBalanceSuccess(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Transaction{})
	TransactionDB := database.NewTransaction(db)

	transaction, err := entity.NewTransaction(1, "2022-02-01T23:35:43-03:00", "DESENVOLVEDOR FULL STACK", "ELIANA NOGUEIRA", 155000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction)
	assert.Nil(t, err)

	transaction2, err := entity.NewTransaction(2, "2022-02-03T17:23:37-03:00", "DESENVOLVEDOR FULL STACK", "CARLOS BATISTA", 155000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction2)
	assert.Nil(t, err)

	transaction3, err := entity.NewTransaction(4, "2022-02-03T17:23:37-03:00", "DESENVOLVEDOR FULL STACK", "NOGUEIRA SOUSA", 50000)
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

	transaction, err := entity.NewTransaction(2, "2022-02-03T17:23:37-03:00", "DESENVOLVEDOR BACKEND", "CARLOS BATISTA", 255000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction)
	assert.Nil(t, err)

	transaction2, err := entity.NewTransaction(4, "2022-02-03T17:23:37-03:00", "DESENVOLVEDOR BACKEND", "KAYKE BALDER", 70000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction2)
	assert.Nil(t, err)

	transaction3, err := entity.NewTransaction(2, "2022-02-03T17:23:37-03:00", "DESENVOLVEDOR FULL STACK", "MYCHAEL KALYPSO", 155000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction3)
	assert.Nil(t, err)

	associate, err := TransactionDB.GetAssociateBalance()
	assert.Nil(t, err)
	assert.NotNil(t, associate)
	assert.Equal(t, "KAYKE BALDER", associate[0].Seller)
	assert.Equal(t, float64(700), associate[0].TValue)
}

func TestGetAssociateBalanceWhenEmpty(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Transaction{})
	TransactionDB := database.NewTransaction(db)
	producers, err := TransactionDB.GetProductorBalance()
	assert.Nil(t, err)
	assert.Empty(t, producers)

}

func TestGetProductorBalanceWhenEmpty(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Transaction{})
	TransactionDB := database.NewTransaction(db)
	associantes, err := TransactionDB.GetAssociateBalance()
	assert.Nil(t, err)
	assert.Empty(t, associantes)

}

func TestGetAssociateBalanceWhenFail(t *testing.T) {
	var associatesList []dto.DtoSellers

	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}

	TransactionDB := database.NewTransaction(db)
	associates, err := TransactionDB.GetAssociateBalance()
	assert.NotNil(t, err)
	assert.Empty(t, associates)
	assert.Equal(t, associatesList, associates)
	assert.Equal(t, "fail_to_query_associates", err.Error())
}

func TestGetProductorBalanceWhenFail(t *testing.T) {
	var producersList []dto.DtoSellers

	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}

	TransactionDB := database.NewTransaction(db)
	producers, err := TransactionDB.GetProductorBalance()
	assert.NotNil(t, err)
	assert.Empty(t, producers)
	assert.Equal(t, producersList, producers)
	assert.Equal(t, "fail_to_query_producers", err.Error())
}

func TestGetForeignCoursesSuccess(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Transaction{})
	TransactionDB := database.NewTransaction(db)

	transaction, err := entity.NewTransaction(2, "2022-02-03T17:23:37-03:00", "HOW TO BE AN ENGINEER", "CARLOS BATISTA", 255000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction)
	assert.Nil(t, err)

	transaction2, err := entity.NewTransaction(4, "2022-02-03T17:21:37-03:00", "MECHANIC'S GUIDE", "EMERSON HAYKEN", 80000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction2)
	assert.Nil(t, err)

	transaction3, err := entity.NewTransaction(4, "2022-02-03T17:18:37-03:00", "CURSO BÁSICO DE PINTOR", "EVANDRO BOILER", 70000)
	assert.Nil(t, err)
	err = TransactionDB.Create(transaction3)
	assert.Nil(t, err)

	courses, err := TransactionDB.GetForeignCourses()
	assert.Nil(t, err)
	assert.NotNil(t, courses)
	assert.Equal(t, "CARLOS BATISTA", courses[0].Seller)
	assert.Equal(t, "HOW TO BE AN ENGINEER", courses[0].Product)
	assert.Equal(t, 2, len(courses))
}

func TestGetForeignCoursesFail(t *testing.T) {
	var emptyList = make([]dto.DtoCourses, 0)
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Transaction{})
	TransactionDB := database.NewTransaction(db)

	courses, err := TransactionDB.GetForeignCourses()
	assert.Nil(t, err)
	assert.NotNil(t, courses)
	assert.Empty(t, courses)
	assert.Equal(t, emptyList, courses)
}
