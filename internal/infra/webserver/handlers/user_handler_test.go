package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/go-chi/jwtauth"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"
	"github.com/mrpsousa/api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	expires, _ = strconv.Atoi((os.Getenv("JWT_EXPIRESIN")))
	secret     = []byte(os.Getenv("JWT_SECRET"))
)

func TestUserHandlerSuccess(t *testing.T) {

	user := map[string]interface{}{
		"name":     "Roger",
		"email":    "email@example.com",
		"password": "123456",
	}

	jsonData, err := json.Marshal(user)

	rr := httptest.NewRecorder()
	mockUserDB := &mocks.UserInterface{}
	mockUserDB.On("Create", mock.Anything).Return(nil)
	userHandler := handlers.NewUserHandler(mockUserDB, jwtauth.New("HS256", secret, nil), expires)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))

	assert.Nil(t, err)
	userHandler.UserCreate(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Result().StatusCode)
}

func TestUserHandlerFail(t *testing.T) {
	expectedErr := errors.New("specific_error")
	user := map[string]interface{}{
		"name":     "Roger",
		"email":    "email@example.com",
		"password": "123456",
	}

	jsonData, err := json.Marshal(user)

	rr := httptest.NewRecorder()
	mockUserDB := &mocks.UserInterface{}
	mockUserDB.On("Create", mock.Anything).Return(expectedErr)
	userHandler := handlers.NewUserHandler(mockUserDB, jwtauth.New("HS256", secret, nil), expires)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))

	assert.Nil(t, err)
	userHandler.UserCreate(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
}
