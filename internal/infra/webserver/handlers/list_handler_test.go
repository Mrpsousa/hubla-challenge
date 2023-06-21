package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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
	sellersResult    []dto.DtoSellers
	foreignerCourses = []dto.DtoCourses{
		{
			Type:      1,
			CreatedAt: "2023-06-15",
			Product:   "Course A",
			Value:     9.99,
			Seller:    "Seller A",
		},
		{
			Type:      2,
			CreatedAt: "2023-06-16",
			Product:   "Course B",
			Value:     19.99,
			Seller:    "Seller B",
		},
	}
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
	mockTransactionDB.AssertCalled(t, "GetAssociateBalance")
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
	mockTransactionDB.AssertCalled(t, "GetAssociateBalance")
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
	mockTransactionDB.AssertCalled(t, "GetProductorBalance")
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
	mockTransactionDB.AssertCalled(t, "GetProductorBalance")
	err = json.Unmarshal(rr.Body.Bytes(), &sellersResult)
	assert.Nil(t, err)
	assert.Equal(t, expectedEmptyList, sellersResult)
}

func TestListForeignCourses_Success(t *testing.T) {
	mockTransactionDB := &mocks.TransactionInterface{}
	mockTransactionDB.On("GetForeignCourses").Return(foreignerCourses, nil)
	expectedBody := `[{"type":1,"created_at":"2023-06-15","product":"Course A","value":9.99,"seller":"Seller A"},{"type":2,"created_at":"2023-06-16","product":"Course B","value":19.99,"seller":"Seller B"}]`

	listHandler := handlers.NewListHandler(mockTransactionDB)
	req := httptest.NewRequest("GET", "/foreign-courses", nil)
	rec := httptest.NewRecorder()
	listHandler.ListForeignCourses(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockTransactionDB.AssertCalled(t, "GetForeignCourses")

	actualBody := rec.Body.String()
	assert.Equal(t, expectedBody, strings.TrimSpace(actualBody))
}

func TestListForeignCoursesEmpty(t *testing.T) {
	mockTransactionDB := &mocks.TransactionInterface{}
	courses := make([]dto.DtoCourses, 0)
	mockTransactionDB.On("GetForeignCourses").Return(courses, nil)

	listHandler := handlers.NewListHandler(mockTransactionDB)
	req := httptest.NewRequest("GET", "/foreign-courses", nil)
	rec := httptest.NewRecorder()
	listHandler.ListForeignCourses(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockTransactionDB.AssertCalled(t, "GetForeignCourses")

	actualBody := rec.Body.String()
	assert.Equal(t, "[]\n", actualBody)
}

func TestListForeignCoursesFailure(t *testing.T) {
	expectedError := errors.New("database error")
	mockTransactionDB := &mocks.TransactionInterface{}
	mockTransactionDB.On("GetForeignCourses").Return(nil, expectedError)

	listHandler := handlers.NewListHandler(mockTransactionDB)
	req := httptest.NewRequest("GET", "/foreign-courses", nil)
	rec := httptest.NewRecorder()
	listHandler.ListForeignCourses(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
