package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	//load enviroment variables from .env file
	var errGlobal = godotenv.Load()

	if errGlobal != nil {
		log.Printf("Error ocurred  %s", errGlobal)
	}

	//Get the value of the key enviroment variable if this exists
	var value, exists = os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not found", key)
	}

	return value
}

// Function to get all variables needed to zincSearch connection using function defined above
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

	//Set the basic authentication credentials on the request
	req.SetBasicAuth(admin, password)

	//Indicate that the request parameter is a json format
	req.Header.Set("Content-Type", "application/json")

	//Set the User-Agent
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
