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
	BaseDB := database.NewTransaction(db)
	transactionHandler := handlers.NewTransactionHandler(BaseDB)
	listHanlder := handlers.NewListHandler(BaseDB)
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/middleware", handlers.MiddlewareHandler)
	router.HandleFunc("/list", handlers.GetAllHandler)
	router.HandleFunc("/upload", transactionHandler.PageUploadFile)
	router.HandleFunc("/producers", listHanlder.ListProductorsBalance)
	router.HandleFunc("/associates", listHanlder.ListAssociatesBalance)

	fmt.Println("Server running in: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
