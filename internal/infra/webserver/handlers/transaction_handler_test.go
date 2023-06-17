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

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"

	"github.com/mrpsousa/api/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dirPath = "./uploads"

func returnDBInstance() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entity.Transaction{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func tierDown() {
	err := os.RemoveAll(dirPath)
	if err != nil {
		log.Println("Failed to remove tmp directory:", err)
		return
	}
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

func setupFail() *httptest.ResponseRecorder {
	db, err := returnDBInstance()
	if err != nil {
		log.Fatal(err)
	}
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

	req, err := http.NewRequest("PUT", "/", body)
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

func existFile() bool {
	dirPath := "./uploads"

	dir, err := os.Open(dirPath)
	if err != nil {
		log.Println("Failed to open dir")
	}
	defer dir.Close()

	files, err := dir.ReadDir(-1)
	if err != nil {
		log.Println("Failed to read dir")
	}
	return (len(files) > 0)
}
func TestUploadHandlerSuccess(t *testing.T) {
	rr := setup()

	expectedResponse := "File uploaded successfully!"
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expectedResponse, rr.Body.String())
	assert.Equal(t, existFile(), true)
	tierDown()
}

func TestUploadHandlerFail(t *testing.T) {
	rr := setupFail()

	expectedResponse := "Method not supported\n"
	assert.NotEqual(t, http.StatusOK, rr.Code)
	assert.Equal(t, expectedResponse, rr.Body.String())
	assert.Equal(t, existFile(), false)
	tierDown()
}

func TestSaveSuccess(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		log.Fatal(err)
	}
	transactionDB := database.NewTransaction(db)
	transactionHandler := handlers.NewTransactionHandler(transactionDB)
	line := "12021-12-03T11:46:02-03:00DOMINANDO INVESTIMENTOS       0000050000MARIA CANDIDA"

	err = transactionHandler.Save(line)
	assert.Nil(t, err)

	var transaction entity.Transaction
	db.First(&transaction)

	assert.Equal(t, int8(1), transaction.Type)
	assert.Equal(t, "MARIA CANDIDA", transaction.Seller)
	assert.Equal(t, "DOMINANDO INVESTIMENTOS       ", transaction.Product)
}

// func TestSaveFail(t *testing.T) {
// 	db, err := returnDBInstance()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	transactionDB := database.NewTransaction(db)
// 	transactionHandler := handlers.NewTransactionHandler(transactionDB)
// 	line := "12021-12-03T11:46:02-03:00DOMINANDO INVESTIMENTOS       0000050000MARIA CANDIDA"

// 	err = transactionHandler.Save(line)
// 	assert.Nil(t, err)

// 	var transaction entity.Transaction
// 	db.First(&transaction)

// 	assert.Equal(t, int8(1), transaction.Type)
// 	assert.Equal(t, "MARIA CANDIDA", transaction.Seller)
// 	assert.Equal(t, "DOMINANDO INVESTIMENTOS       ", transaction.Product)
// }
