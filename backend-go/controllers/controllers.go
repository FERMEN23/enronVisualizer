package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"backend-go/utils"

	"github.com/go-chi/chi"
)

// Get max_size from .env
var max_size = utils.GetEnv("MAX_RESULT")

func GetElementsFrom(w http.ResponseWriter, r *http.Request) {

	//Retrieve the fromValue parameter
	fromValue := chi.URLParam(r, "fromValue")

	//query to get all documents from an index (fromValue), that allows pagination

	intFirstIndex, err := strconv.Atoi(fromValue)
	intFirstIndex--

	intMaxSize, err := strconv.Atoi(max_size)

	endIndex := intFirstIndex + intMaxSize + 1

	fmt.Print(fromValue, endIndex)
	query := fmt.Sprintf(`{	
		"search_type": "querystring",
		"sort_fields": ["ID"],
		"query":
		{
			"term": "+ID: >%d + ID: < %d"
		},
		"max_results": %s
	}`, intFirstIndex, endIndex, max_size)

	body, err := utils.ZincSearchRequest(query)
	if err != nil {
		log.Println("Error ocurred:", err)
	}

	//Set the "Content-Type" header in the response to be sent to the client.
	w.Header().Set("Content-Type", "application/json")

	//Write the response that will be sent to the client.
	w.Write(body)
}

func GetElementsByID(w http.ResponseWriter, r *http.Request) {

	//Get max_size from .env
	max_size := utils.GetEnv("MAX_RESULT")

	//Retrieve the id parameter
	id := chi.URLParam(r, "id")

	//query to get the elements match with the id parameter
	query := fmt.Sprintf(`{	
		"search_type": "querystring",
		"query":
		{
			"term": "_id: %s"
		},
		"from": 0,
		"max_results": %s
	}`, id, max_size)

	//make the htttp request to ZincSearch
	body, err := utils.ZincSearchRequest(query)
	if err != nil {
		log.Println("Error ocurred:", err)
	}

	//Set the "Content-Type" header in the response to be sent to the client.
	w.Header().Set("Content-Type", "application/json")

	//Write the response that will be sent to the client.
	w.Write(body)
}

func GetElementsIdAndFilter(w http.ResponseWriter, r *http.Request) {
	//Get max_size from .env
	max_size := utils.GetEnv("MAX_RESULT")

	var (
		email = chi.URLParam(r, "email")
		last  = chi.URLParam(r, "last")
	)

	query := fmt.Sprintf(`{	
		"search_type": "querystring",
		"query":
		{
			"term": "from.Address: *%s*"
		},
		"from": %s,
		"max_results": %s
	}`, email, last, max_size)

	//make the htttp request to ZincSearch
	body, err := utils.ZincSearchRequest(query)
	if err != nil {
		log.Println("Error ocurred:", err)
	}

	//Set the "Content-Type" header in the response to be sent to the client.
	w.Header().Set("Content-Type", "application/json")

	//Write the response that will be sent to the client.
	w.Write(body)
}

func GetEnvMaxSize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	log.Println("Value max_size from backend:", max_size)
	w.Write([]byte(max_size))
}

func GetAllElements(w http.ResponseWriter, r *http.Request) {

	//query to get all documents from an index (fromValue), that allows pagination
	query := `{
		"search_type": "alldocuments",
		"from": 0
	}`

	body, err := utils.ZincSearchRequest(query)
	if err != nil {
		log.Println("Error ocurred:", err)
	}

	//Set the "Content-Type" header in the response to be sent to the client.
	w.Header().Set("Content-Type", "application/json")

	//Write the response that will be sent to the client.
	w.Write(body)
}
