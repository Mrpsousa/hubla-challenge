package handlers

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/database"
	"github.com/mrpsousa/api/pkg"
)

type TransactionHandler struct {
	TransactionDB database.TransactionInterface
}

func NewTransactionHandler(db database.TransactionInterface) *TransactionHandler {
	return &TransactionHandler{
		TransactionDB: db,
	}
}

// 3. Exibir a lista de todas as transações de produtos importadas
// 4. Exibir o saldo final do produtor
// 5. Exibir o saldo final de um afiliado
//valor das transações em centavos /

func (t *TransactionHandler) PageUploadFile(w http.ResponseWriter, r *http.Request) {
	// done := make(chan error)
	if r.Method == "GET" {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		// upload of 10 MB files.
		err := r.ParseMultipartForm(10 << 20) // limit your max input length!
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileName := pkg.FileNameGenerate(header.Filename)
		dst, err := os.Create("./uploads/" + fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "File uploaded successfully!")

		//saving data in background
		go func() {
			err := t.SaveFromFile(fmt.Sprintf("./uploads/%s", fileName))
			if err != nil {
				log.Printf(err.Error())
			}
		}()
	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

}

func (t *TransactionHandler) Save(line string) error {
	tp, err := pkg.StringToInt8(line[:1])
	if err != nil {
		return err
	}
	value, err := pkg.StringToFloat64(line[56:66])
	if err != nil {
		return err
	}
	createdAt := line[1:26]
	product := line[26:56]
	seller := line[66:]

	transaction, err := entity.NewTransaction(tp, createdAt, product, seller, value)
	err = t.TransactionDB.Create(transaction)
	if err != nil {
		return errors.New("Failed to create/save DB transaction")
	}
	return nil
}

func (t *TransactionHandler) SaveFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.New("Failed to open path")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return errors.New("Failed to read file")
		}
		line = bytes.TrimSuffix(line, []byte{'\n'})
		err = t.Save(string(line))
		if err != nil {
			return err
		}
	}

	return nil
}
