package controllers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"backend-go/client"

	"github.com/go-chi/chi"
)

// Get max_size from .env
var max_result, err = client.GetEnv("MAX_RESULT")

// getIndexs returns startIndex and endIndex as int to pagination
func getIndexs(startIndex string) (int, int, error) {
	intSirstIndex, err := strconv.Atoi(startIndex)
	if err != nil {
		log.Println("Error ocurred:", err)
		return -1, -1, err
	}
	intSirstIndex--

	intMaxResult, err := strconv.Atoi(max_result)
	if err != nil {
		log.Println("Error ocurred:", err)
		return -1, -1, err
	}
	intEndIndex := intSirstIndex + intMaxResult + 1

	return intSirstIndex, intEndIndex, err
}

// validateID validate that the input is a string that match with a pattern
func validateID(id string) bool {
	pattern := "^[A-Z0-9a-z]{11}$"
	match, _ := regexp.MatchString(pattern, id)

	return match
}

// GetEmailPagination return emails in pagination with ID parameter controlling results
func GetEmailPagination(w http.ResponseWriter, r *http.Request) {
	startIndex := chi.URLParam(r, "startIndex")

	intStartIndex, intEndIndex, err := getIndexs(startIndex)
	if err != nil {
		intStartIndex, intEndIndex = 0, 20
		log.Println("Error ocurred", err, "and set indexes by defaul start", intStartIndex, "end", intEndIndex)
	}

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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// GetEmailByID resturn an email by _id unique indentifier
func GetEmailByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	valid := validateID(id)
	if !valid {
		http.Error(w, "invalid ID input", http.StatusBadRequest)
		return
	}

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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// GetEnvMaxResult return max result value from .env file
func GetEnvMaxResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
