package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["ts"] = "typescript"
	languages["py"] = "python"
	fmt.Println("Languages in map are,", languages)
	delete(languages, "py") // will remove python

	//peek on loops here

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
