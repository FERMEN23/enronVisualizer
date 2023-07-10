package main

import (
	"log"
	"net/http"

	"backend-go/controllers"
	"backend-go/middleware"

	"github.com/go-chi/chi"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.CorsMiddleware)
	router.Get("/elementsId/{fromValue}", controllers.GetElementsFrom)
	router.Get("/elementsFilter/{email}/{last}", controllers.GetElementsIdAndFilter)
	router.Get("/elementsById/{id}", controllers.GetElementsByID)

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
