package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpsousa/api/internal/infra/webserver/handlers"
)

func main() {

	router := mux.NewRouter()

	// Defina as rotas e os manipuladores
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)
	// Inicie o servidor HTTP
	fmt.Println("Servidor iniciado. Acesse http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
