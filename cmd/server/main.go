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
// securityDefinitions.apikey ApiKeyAuth
// in header
// name Authorization

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
	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

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
