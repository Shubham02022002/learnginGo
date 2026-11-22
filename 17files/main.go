package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	println("File handling in go")
	content := "this is some kinda gibrishsss"
	file, err := os.Create("17files/info.txt")
	
	checkNilError(err)

	length, err := io.WriteString(file, content)
	
	checkNilError(err)
	
	println("Content length:", length)
	defer file.Close()
	readFile("./17files/info.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	
	checkNilError(err)
	
	println("Reading file data:....\n")
	println(string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
