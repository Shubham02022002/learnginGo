package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	response, err := http.Get(url)
	checkNilError(err)
	fmt.Printf("Response is of type: %T\n", response)
	dataBytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)
	println("Response Body:", string(dataBytes))
	response.Body.Close() // don't forget to close the connection

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
