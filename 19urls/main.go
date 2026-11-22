package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://www.linkedin.com/messaging/thread/2-OGE0NWYI3MzE5XzIwMA==/"

func main() {
	println("This section talks about url hangling.")
	result, err := url.Parse(URL)
	checkNilError(err)
	// println(result.Scheme)
	// println(result.Host)
	// println(result.Path)

	qparams := result.Query()
	fmt.Printf("Query params are: %v\n", qparams)

	newUrl := &url.URL{
		Scheme:   result.Scheme,
		Host:     result.Host,
		Path:     result.Path,
		RawQuery: qparams.Encode(),
	}

	fmt.Printf("New url is: %v\n", newUrl.String())

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
