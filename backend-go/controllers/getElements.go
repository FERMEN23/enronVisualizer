package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

var err = godotenv.Load()

var admin = os.Getenv("ADMIN")
var password = os.Getenv("PASSWORD")
var indexName = os.Getenv("INDEXNAME")
var pathZincSearch = os.Getenv("PATHZINCSEARCH")
var max_size = os.Getenv("MAX_RESULT")
var path = pathZincSearch + indexName + "/_search"

func GetElementsFrom(w http.ResponseWriter, r *http.Request) {
	fromValue := chi.URLParam(r, "fromValue")

	query := fmt.Sprintf(`{	
		"search_type": "alldocuments",
		"from": %s,
		"max_results": %s
	}`, fromValue, max_size)

	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(query))
	if err != nil {
		log.Println("Error ocurred:", err)
	}
	req.SetBasicAuth(admin, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error ocurred:", err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error ocurred:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func GetElementsByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query := fmt.Sprintf(`{	
		"search_type": "querystring",
		"query":
		{
			"term": "_id: %s"
		},
		"from": 0
	}`, id)

	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(query))
	if err != nil {
		log.Println("Error ocurred:", err)
	}
	req.SetBasicAuth(admin, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error ocurred:", err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error ocurred:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func GetElementsIdAndFilter(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	last := chi.URLParam(r, "last")

	query := fmt.Sprintf(`{	
		"search_type": "querystring",
		"query":
		{
			"term": "from.Address: *%s*"
		},
		"from": %s,
		"max_results": %s
	}`, email, last, max_size)

	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(query))
	if err != nil {
		log.Println("Error ocurred:", err)
	}
	req.SetBasicAuth(admin, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error ocurred:", err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error ocurred:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
