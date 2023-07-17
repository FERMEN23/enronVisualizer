package client

import (
	"errors"
	"fmt"
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
		return value, errors.New(fmt.Sprintf("%s enviroment variable doesn't exist", key))
	}

	return value, nil
}

// EnvsZincSearch get all variables needed to zincSearch connection using function defined above
func envsZincSearch() (string, string, string, error) {
	admin, err := GetEnv("ADMIN")
	if err != nil {
		return "", "", "", err
	}

	password, err := GetEnv("PASSWORD")
	if err != nil {
		return "", "", "", err
	}

	indexName, err := GetEnv("INDEXNAME")
	if err != nil {
		return "", "", "", err
	}

	pathZincSearch, err := GetEnv("PATHZINCSEARCH")
	if err != nil {
		return "", "", "", err
	}

	path := pathZincSearch + indexName + "/_search"

	return admin, password, path, nil
}

func ZincSearchRequest(query string) ([]byte, error) {
	admin, password, path, err := envsZincSearch()

	if err != nil {
		log.Println("Error ocurred: ", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(query))
	if err != nil {
		log.Println("Error ocurred: ", err)
		return nil, err
	}

	req.SetBasicAuth(admin, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error ocurred: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	log.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		log.Println("Error ocurred and get response", resp.StatusCode)
		return nil, errors.New("Something went wrong")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error ocurred: ", err)
		return nil, err
	}

	return body, nil
}
