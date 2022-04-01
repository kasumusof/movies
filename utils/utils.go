package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kasumusof/movies/models"
	"github.com/kasumusof/movies/types"
)

var cacheHost string
var cachePassword string

func init() {
	godotenv.Load("../.env")
	cacheHost = os.Getenv("REDIS_HOST")
	if cacheHost == "" {
		cacheHost = "localhost:6379"
	}
	cachePassword = os.Getenv("REDIS_PASS")
	if cachePassword == "" {
		cachePassword = ""
	}
	fmt.Println(cacheHost, cachePassword)
}

// Fetch helper function to make requests
func Fetch(url string) (*http.Response, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func CheckAndCache(key string, expiration int, data interface{}) interface{} {
	cacher := NewCache(cacheHost, cachePassword, 0)
	cacherItem, err := cacher.Get("url")
	if err != nil {
		// fetch
		x, err := Fetch(key)
		if err != nil {
			log.Fatal(err)
		}
		json.NewDecoder(x.Body).Decode(&data)
		body, _ := json.Marshal(data)
		cacher.Set(key, string(body), 0)
	}
	json.Unmarshal([]byte(cacherItem), &data)
	return data
}

func MakeResponse(w *http.ResponseWriter, result interface{}, ok bool, status int) {
	var response types.Response
	response.OK = ok
	response.Result = result
	switch dum := result.(type) {
	case *models.Comments:
		response.Count = len(*dum)
	default:
	}
	(*w).WriteHeader(status)
	if err := json.NewEncoder(*w).Encode(response); err != nil {
		fmt.Println(err)
	}
}

func CheckMyCache(key string, expiration int, data interface{}) (interface{}, error) {
	// func CheckCache(key string) (string, error) {
	cacher := NewCache(cacheHost, cachePassword, 0)
	cacherItem, err := cacher.Get("url")
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(cacherItem), &data)
	return data, nil
}

func MakeMyCache(key string, expiration int, data interface{}) error {
	body, _ := json.Marshal(data)
	cacher := NewCache(cacheHost, cachePassword, 0)
	err := cacher.Set(key, body, 0)
	if err != nil {
		return err
	}
	return err
}
