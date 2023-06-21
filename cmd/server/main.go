package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrpsousa/api/configs"
	_ "github.com/mrpsousa/api/docs"
	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/mrpsousa/api/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Hubla Challenge
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Marcos Rogerio
// @contact.url    http://www.example.com.br
// @contact.email  urameshi.uba@gmail.com

// @license.name   License Name
// @license.url    http://www.github.com

// @host      localhost:8000
// @BasePath  /
func main() {
	fs := http.FileServer(http.Dir("./static"))
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
	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	router.Get("/middleware", handlers.MiddlewareHandler)
	router.Get("/list", handlers.GetAllHandler)
	router.Handle("/static/*", http.StripPrefix("/static/", fs))
	router.Get("/", handlers.IndexHandler)
	router.Post("/upload", transactionHandler.PageUploadFile)
	router.Get("/ping", handlers.Healthz)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.UserCreate)
		r.Post("/generate_token", userHandler.GetJWT)
	})

	//authorization protected endpoints
	router.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/producers", listHanlder.ListProductorsBalance)
		r.Get("/associates", listHanlder.ListAssociatesBalance)
		r.Get("/courses/foreign", listHanlder.ListForeignCourses)
	})

	fmt.Println("Server running in: http://localhost:8000/users/login")
	log.Fatal(http.ListenAndServe(":8000", router))
}
