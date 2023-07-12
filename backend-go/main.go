package main

import (
	"log"
	"net/http"

	"backend-go/client"
	"backend-go/controllers"
	"backend-go/middleware"

	"github.com/go-chi/chi"
)

func main() {
	port, err := client.GetEnv("DEFAULT_PORT")
	if err != nil {
		log.Fatal("Error getting port enviroment variable ")
	}
	router := chi.NewRouter()

	router.Use(middleware.CorsMiddleware)

	router.Get("/emailsPagination/{startIndex}", controllers.GetEmailPagination)
	router.Get("/emailById/{id}", controllers.GetEmailByID)
	router.Get("/emailsFilter/{term}/{startIndex}", controllers.GetEmailsWithFilter)
	router.Get("/getMaxResultVariable", controllers.GetEnvMaxResult)
	router.Get("/getallEmails", controllers.GetAllEmails)

	//start an HTTP server and listen for incoming requests on a specified address
	log.Printf("Server started on  %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
