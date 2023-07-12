package client

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// GetEnv get the value from key enviroment variable
func GetEnv(key string) string {
	var errGlobal = godotenv.Load()
	if errGlobal != nil {
		log.Printf("Error ocurred:  %s", errGlobal)
	}

	var value, exists = os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not found", key)
	}
	//ToDo: error handling when key doesnt exist
	return value
}

// EnvsZincSearch get all variables needed to zincSearch connection using function defined above
func EnvsZincSearch() (string, string, string) {
	admin := GetEnv("ADMIN")
	password := GetEnv("PASSWORD")
	indexName := GetEnv("INDEXNAME")
	pathZincSearch := GetEnv("PATHZINCSEARCH")
	path := pathZincSearch + indexName + "/_search"
	return admin, password, path
}

func ZincSearchRequest(query string) ([]byte, error) {
	//Get enviroment variables to ZincSearch connection
	admin, password, path := EnvsZincSearch()

	//create a new HTTP request object to ZincSearch, with POST method
	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(query))

	if err != nil {
		log.Println("Error ocurred:", err)
		return nil, err
	}

	req.SetBasicAuth(admin, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	//hacer la solicitud a ZincSearch
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error ocurred:", err)
		return nil, err
	}

	//free up resources and avoid possible memory leaks.
	defer resp.Body.Close()

	//log for ths status response
	log.Println(resp.StatusCode)

	//read the response body
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error ocurred:", err)
		return nil, err
	}

	return body, nil
}
