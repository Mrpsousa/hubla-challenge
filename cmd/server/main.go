package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"

	"github.com/mrpsousa/api/internal/infra/database"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func main() {

	// config := configs.NewConfig()

	db, err := gorm.Open(sqlite.Open("hubla.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Transaction{})
	transactionDB := database.NewTransaction(db)
	transactionHandler := handlers.NewTransactionHandler(transactionDB)
	router := mux.NewRouter()

	// Defina as rotas e os manipuladores
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", transactionHandler.UploadHandler)
	// Inicie o servidor HTTP
	fmt.Println("Servidor iniciado. Acesse http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
