package handlers_test

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mrpsousa/api/internal/infra/webserver/handlers"

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

func setup() *httptest.ResponseRecorder {
	db, err := returnDBInstance()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "test.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, bytes.NewReader([]byte("This is a test file")))
	if err != nil {
		log.Fatal(err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()

	transactionDB := database.NewTransaction(db)
	transactionHandler := handlers.NewTransactionHandler(transactionDB)

	handler := http.HandlerFunc(transactionHandler.UploadHandler)

	handler.ServeHTTP(rr, req)
	return rr
}

func TestUploadHandler(t *testing.T) {
	rr := setup()

	expectedResponse := "File uploaded successfully!"
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expectedResponse, rr.Body.String())

	uploadPath := "./uploads/uuid-test.txt"
	_, err := os.Stat(uploadPath)
	if os.IsNotExist(err) {
		t.Errorf("uploaded file was not created: %v", err)
	} else {
		err = os.Remove(uploadPath)
		if err != nil {
			t.Errorf("error deleting uploaded file: %v", err)
		}
	}
}
