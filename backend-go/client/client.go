package client

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// GetEnv get the value from key enviroment variable
func GetEnv(key string) (string, error) {
	var errGlobal = godotenv.Load()
	if errGlobal != nil {
		return "", errGlobal
	}

	var value, exists = os.LookupEnv(key)
	if !exists {
		return value, errors.New("enviroment variable doesn't exist")
	}

	return value, nil
}

// EnvsZincSearch get all variables needed to zincSearch connection using function defined above
func EnvsZincSearch() (string, string, string, error) {
	admin, err := GetEnv("ADMIN")
	password, err := GetEnv("PASSWORD")
	indexName, err := GetEnv("INDEXNAME")
	pathZincSearch, err := GetEnv("PATHZINCSEARCH")
	path := pathZincSearch + indexName + "/_search"
	return admin, password, path, err
}

func ZincSearchRequest(query string) ([]byte, error) {
	admin, password, path, err := EnvsZincSearch()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(query))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(admin, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
