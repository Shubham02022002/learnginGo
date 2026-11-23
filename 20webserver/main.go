package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	URL := "http://localhost:3000"
	requestBody := `{"name": "John Doe", "age": 30}`

	GetRequestHandler(URL)
	PostRequestHandler(URL, requestBody)
}

func GetRequestHandler(URL string) {
	response, err := http.Get(URL)
	checkNilError(err)
	defer response.Body.Close()

	println(response.Body)
	println(response.StatusCode)
	content, err := io.ReadAll(response.Body)
	checkNilError(err)
	println(string(content))
}

func PostRequestHandler(URL string, requestBody string) {
	response, err := http.Post(URL, "application/json", strings.NewReader(requestBody))
	checkNilError(err)
	defer response.Body.Close()

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
