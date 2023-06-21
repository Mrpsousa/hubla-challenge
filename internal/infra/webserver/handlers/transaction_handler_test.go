package handlers_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"
	"github.com/mrpsousa/api/pkg"

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

func returnHanlder() *handlers.TransactionHandler {
	db, err := returnDBInstance()
	if err != nil {
		log.Fatal(err)
	}
	transactionDB := database.NewTransaction(db)
	transactionHandler := handlers.NewTransactionHandler(transactionDB)
	return transactionHandler
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

func TestPageUploadFileGET(t *testing.T) {
	defer pkg.RemoveFolder(dirPath)
	handler := returnHanlder()

	req := httptest.NewRequest(http.MethodGet, "/upload", nil)
	rec := httptest.NewRecorder()

	handler.PageUploadFile(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.NotNil(t, body)

}

func TestPageUploadFilePOST(t *testing.T) {
	defer pkg.RemoveFolder(dirPath)
	handler := returnHanlder()
	tmpFile, err := ioutil.TempFile("", "testfile")
	assert.Nil(t, err)
	defer os.Remove(tmpFile.Name())

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", tmpFile.Name())
	assert.Nil(t, err)
	_, err = io.Copy(part, tmpFile)
	assert.Nil(t, err)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()

	handler.PageUploadFile(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusSeeOther, res.StatusCode)
}

func TestPageUploadFileUnsupportedMethod(t *testing.T) {
	defer pkg.RemoveFolder(dirPath)
	handler := returnHanlder()

	req := httptest.NewRequest(http.MethodPut, "/upload", nil)
	rec := httptest.NewRecorder()

	handler.PageUploadFile(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status MethodNotAllowed; got %v", res.StatusCode)
	}
}
