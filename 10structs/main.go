package main

import "fmt"

func main() {
	fmt.Println("In this module we'll cover structs.")
	// Go lang has no concepts of classes, inheritance, super or parent

	shubham := User{"Shubham", 99, "shubham@gmail.com", true, "userDetails"}
	fmt.Println("User:", shubham)
	fmt.Printf("User details:%+v", shubham) //use in case of structs 
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
	extras string // unexported
}
