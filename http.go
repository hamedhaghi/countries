package countries

import (
	"io"
	"net/http"
)

const (
	API_URL = "https://restcountries.com/v3.1"
)

func get(path string) (content string, err error) {
	res, err := http.Get(API_URL + path)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	return string(body), err
}
