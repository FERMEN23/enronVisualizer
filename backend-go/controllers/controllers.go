package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"backend-go/client"

	"github.com/go-chi/chi"
)

// Get max_size from .env
var max_result = client.GetEnv("MAX_RESULT")

// getIndexs returns startIndex and endIndex as int to pagination
func getIndexs(startIndex string) (int, int) {
	intSirstIndex, err := strconv.Atoi(startIndex)
	if err != nil {
		log.Println("Error ocurred:", err)
	}
	intSirstIndex--

	intMaxResult, err := strconv.Atoi(max_result)
	if err != nil {
		log.Println("Error ocurred:", err)
	}
	intEndIndex := intSirstIndex + intMaxResult + 1

	return intSirstIndex, intEndIndex
}

// GetEmailPagination return emails in pagination with ID parameter controlling results
func GetEmailPagination(w http.ResponseWriter, r *http.Request) {
	startIndex := chi.URLParam(r, "startIndex")

	intStartIndex, intEndIndex := getIndexs(startIndex)

	query := fmt.Sprintf(`{	
		"search_type": "querystring",
		"sort_fields": ["ID"],
		"query":
		{
			"term": "+ID: >%d + ID: < %d"
		},
		"max_results": %s
	}`, intStartIndex, intEndIndex, max_result)

	body, err := client.ZincSearchRequest(query)
	if err != nil {
		log.Println("Error ocurred:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// GetEmailByID resturn an email by _id unique indentifier
func GetEmailByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	//ToDo- validate id input
	query := fmt.Sprintf(`{	
		"search_type": "querystring",
		"query":
		{
			"term": "_id: %s"
		},
		"from": 0
	}`, id)

	body, err := client.ZincSearchRequest(query)
	if err != nil {
		log.Println("Error ocurred:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// GetEmailsWithFilter return emails filtered parameter term
func GetEmailsWithFilter(w http.ResponseWriter, r *http.Request) {
	var (
		term       = chi.URLParam(r, "term")
		startIndex = chi.URLParam(r, "startIndex")
	)

	query := fmt.Sprintf(`{	
		"search_type": "querystring",
		"query":
		{
			"term": "%s"
		},
		"from": %s,
		"max_results": %s
	}`, term, startIndex, max_result)

	body, err := client.ZincSearchRequest(query)
	if err != nil {
		log.Println("Error ocurred:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// GetEnvMaxResult return max result value from .env file
func GetEnvMaxResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	log.Println("Value max_size from backend:", max_result)
	w.Write([]byte(max_result))
}

// GetAllEmails return all elements from zincSearch
func GetAllEmails(w http.ResponseWriter, r *http.Request) {
	query := `{
		"search_type": "alldocuments",
		"from": 0
	}`

	body, err := client.ZincSearchRequest(query)
	if err != nil {
		log.Println("Error ocurred:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
