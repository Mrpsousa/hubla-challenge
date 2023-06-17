package database_test

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
	"github.com/mrpsousa/api/internal/infra/database"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dirPath = "./uploads"

func tierDown() {
	err := os.RemoveAll(dirPath)
	if err != nil {
		log.Println("Failed to remove tmp directory:", err)
		return
	}
}

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
	tierDown()
}

// func TestSaveTransaction(t *testing.T) {
// 	setup()
// 	db, err := returnDBInstance()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	db.AutoMigrate(&entity.Transaction{})
// 	transaction, err := entity.NewTransaction(1, "2022-02-19T05:33:07-03", "DOMINANDO INVESTIMENTOS", "MARIA CANDIDA", 50000.0)
// 	assert.NoError(t, err)
// 	TransactionDB := database.NewTransaction(db)
// 	err = TransactionDB.Create(transaction)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, transaction.ID)

// }

// func (p *Transaction) Save(line string) error {
// 	setup()
// 	tp, err := pkg.StringToInt8(line[:1])
// 	if err != nil {
// 		return err
// 	}
// 	value, err := pkg.StringToFloat64(line[56:66])
// 	if err != nil {
// 		return err
// 	}
// 	createdAt := line[1:26]
// 	product := line[26:56]
// 	seller := line[66:]

// 	transaction, err := entity.NewTransaction(tp, createdAt, product, seller, value)
// 	err = p.Create(transaction)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
