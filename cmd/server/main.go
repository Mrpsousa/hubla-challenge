package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrpsousa/api/configs"
	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/mrpsousa/api/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	config := configs.NewConfig()

	db, err := gorm.Open(sqlite.Open("hubla.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Transaction{}, &entity.User{})
	baseDB := database.NewTransaction(db)
	userDB := database.NewUser(db)
	transactionHandler := handlers.NewTransactionHandler(baseDB)
	listHanlder := handlers.NewListHandler(baseDB)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)

	router := chi.NewRouter()

	router.Get("/users/create", handlers.CreateUserHandler)
	router.Get("/users/login", handlers.UserLoginHandler)

	router.Route("/", func(r chi.Router) {
		// r.Use(jwtauth.Verifier(config.TokenAuth)) // get the token and inject it into the context
		// r.Use(jwtauth.Authenticator)              // validate of token
		r.Get("/", handlers.IndexHandler)
		r.Get("/middleware", handlers.MiddlewareHandler)
		r.Get("/list", handlers.GetAllHandler)
		r.Post("/upload", transactionHandler.PageUploadFile)
		r.Get("/producers", listHanlder.ListProductorsBalance)
		r.Get("/associates", listHanlder.ListAssociatesBalance)
		r.Get("/courses/foreign", listHanlder.ListForeignCourses)
	})

	router.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Post("/generate_token", userHandler.GetJWT)
	})

	fmt.Println("Server running in: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
