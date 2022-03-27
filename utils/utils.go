package utils

import (
	"net/http"
)

// Fetch helper function to make requests
func Fetch(url string) (*http.Response, error) {

	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return req, nil
}
