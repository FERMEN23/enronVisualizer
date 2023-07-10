package main

import (
	"log"
	"net/http"
	"os"

	"backend-go/controllers"
	"backend-go/middleware"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

var errGlobal = godotenv.Load()

var port = os.Getenv("DEFAULT_PORT")

func main() {

	router := chi.NewRouter()

	router.Use(middleware.CorsMiddleware)

	router.Get("/elementsId/{fromValue}", controllers.GetElementsFrom)
	router.Get("/elementsFilter/{email}/{last}", controllers.GetElementsIdAndFilter)
	router.Get("/elementsById/{id}", controllers.GetElementsByID)

	log.Printf("Server started on  %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
