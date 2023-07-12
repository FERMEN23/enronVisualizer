package main

import (
	"log"
	"net/http"

	"backend-go/controllers"
	"backend-go/middleware"
	"backend-go/utils"

	"github.com/go-chi/chi"
)

func main() {
	//get the port from .env file
	port := utils.GetEnv("DEFAULT_PORT")

	//Create a new instance of a router
	router := chi.NewRouter()

	//add a CORS middleware to router
	router.Use(middleware.CorsMiddleware)

	//Define the routes that route will be manage
	router.Get("/elementsId/{fromValue}", controllers.GetElementsFrom)
	router.Get("/elementsFilter/{email}/{last}", controllers.GetElementsIdAndFilter)
	router.Get("/elementsById/{id}", controllers.GetElementsByID)
	router.Get("/getMaxSize", controllers.GetEnvMaxSize)
	router.Get("/allElements", controllers.GetAllElements)

	//start an HTTP server and listen for incoming requests on a specified address
	log.Printf("Server started on  %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
