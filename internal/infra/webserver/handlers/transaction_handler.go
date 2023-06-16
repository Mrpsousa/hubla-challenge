package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/mrpsousa/api/internal/infra/database"
)

type TransactionHandler struct {
	TransactionDB database.TransactionInterface
}

func NewTransactionHandler(db database.TransactionInterface) *TransactionHandler {
	return &TransactionHandler{
		TransactionDB: db,
	}
}

func (t *TransactionHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
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
		dst, err := os.Create("./uploads/" + "uuid-" + header.Filename)
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
		err = t.TransactionDB.SaveFromFile("./uploads/sales.txt")
		if err != nil {
			log.Printf("ERRO DOIDO") //////////////////////////////////////////////////////////////////////////
		}
	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

}
