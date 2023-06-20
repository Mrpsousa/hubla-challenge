package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"

	"github.com/go-chi/chi/v5"
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
	router := chi.NewRouter()
	// router := mux.NewRouter()

	router.Get("/", handlers.IndexHandler)
	router.Get("/middleware", handlers.MiddlewareHandler)
	router.Get("/list", handlers.GetAllHandler)
	router.Post("/upload", transactionHandler.PageUploadFile)
	router.Get("/producers", listHanlder.ListProductorsBalance)
	router.Get("/associates", listHanlder.ListAssociatesBalance)
	router.Get("/courses/foreign", listHanlder.ListForeignCourses)

	fmt.Println("Server running in: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
