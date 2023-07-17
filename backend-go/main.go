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
		port = "8000"
		log.Println("Error getting port enviroment variable then set in a default value;", port)
	}
	router := chi.NewRouter()

	router.Use(middleware.CorsMiddleware)

	router.Get("/v1/emailsPagination/{startIndex}", controllers.GetEmailPagination)
	router.Get("/v1/emailById/{id}", controllers.GetEmailByID)
	router.Get("/v1/emailsFilter/{term}/{startIndex}", controllers.GetEmailsWithFilter)
	router.Get("/v1/getMaxResultVariable", controllers.GetEnvMaxResult)
	router.Get("/v1/getallEmails", controllers.GetAllEmails)

	//start an HTTP server and listen for incoming requests on a specified address
	log.Printf("Server started on  %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
