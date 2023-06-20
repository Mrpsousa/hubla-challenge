package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrpsousa/api/internal/dto"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"
	"github.com/mrpsousa/api/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	associates = []dto.DtoSellers{
		{Seller: "John Doe", TValue: 100},
		{Seller: "Jane Smith", TValue: 0},
		{Seller: "Mike Johnson", TValue: 50},
	}
	producers = []dto.DtoSellers{
		{Seller: "Maria Maia", TValue: 100},
		{Seller: "Kelly Smith", TValue: 0},
		{Seller: "Kaio Jullius", TValue: 0},
		{Seller: "Roger Santana", TValue: 50},
	}
	sellersResult []dto.DtoSellers
)

func TestListHandlerAssociatesSuccess(t *testing.T) {
	expectedList := []dto.DtoSellers{
		{Seller: "John Doe", TValue: 100},
		{Seller: "Mike Johnson", TValue: 50},
	}
	rr := httptest.NewRecorder()
	mockTransactionDB := &mocks.TransactionInterface{}
	mockTransactionDB.On("GetAssociateBalance").Return(associates, nil)
	listHandler := handlers.NewListHandler(mockTransactionDB)

	req, err := http.NewRequest("GET", "/associates", nil)
	assert.Nil(t, err)
	listHandler.ListAssociatesBalance(rr, req)
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	err = json.Unmarshal(rr.Body.Bytes(), &sellersResult)
	assert.Nil(t, err)
	assert.Equal(t, expectedList, sellersResult)
}

func TestListHandlerAssociatesFail(t *testing.T) {
	expectedErr := errors.New("specific_error")
	var expectedEmptyList []dto.DtoSellers

	rr := httptest.NewRecorder()
	mockTransactionDB := &mocks.TransactionInterface{}
	mockTransactionDB.On("GetAssociateBalance").Return(nil, expectedErr)
	listHandler := handlers.NewListHandler(mockTransactionDB)

	req, err := http.NewRequest("GET", "/associates", nil)
	assert.Nil(t, err)
	listHandler.ListAssociatesBalance(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	err = json.Unmarshal(rr.Body.Bytes(), &sellersResult)
	assert.Nil(t, err)
	assert.Equal(t, expectedEmptyList, sellersResult)
}

func TestListHandlerProducersSuccess(t *testing.T) {
	expectedList := []dto.DtoSellers{
		{Seller: "Maria Maia", TValue: 100},
		{Seller: "Roger Santana", TValue: 50},
	}

	rr := httptest.NewRecorder()
	mockTransactionDB := &mocks.TransactionInterface{}
	mockTransactionDB.On("GetProductorBalance").Return(producers, nil)
	listHandler := handlers.NewListHandler(mockTransactionDB)

	req, err := http.NewRequest("GET", "/producers", nil)
	assert.Nil(t, err)
	listHandler.ListProductorsBalance(rr, req)
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	err = json.Unmarshal(rr.Body.Bytes(), &sellersResult)
	assert.Nil(t, err)
	assert.Equal(t, expectedList, sellersResult)
}

func TestListHandlerProducersFail(t *testing.T) {
	expectedErr := errors.New("specific_error")
	var expectedEmptyList []dto.DtoSellers

	rr := httptest.NewRecorder()
	mockTransactionDB := &mocks.TransactionInterface{}
	mockTransactionDB.On("GetProductorBalance").Return(nil, expectedErr)
	listHandler := handlers.NewListHandler(mockTransactionDB)

	req, err := http.NewRequest("GET", "/producers", nil)
	assert.Nil(t, err)
	listHandler.ListProductorsBalance(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	err = json.Unmarshal(rr.Body.Bytes(), &sellersResult)
	assert.Nil(t, err)
	assert.Equal(t, expectedEmptyList, sellersResult)
}
