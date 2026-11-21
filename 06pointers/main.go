package main

import "fmt"

func main() {

	fmt.Println("In moudle teaches us about pointers")
	// var ptr *int;
	// fmt.Println("Value of pointer ptr is,",ptr) //nil

	number := 24
	var ptrNumber = &number
	fmt.Println("Address of number in memory is,", ptrNumber)
	fmt.Println("Actual value is,", *ptrNumber)
	// Address of number in memory is, 0xbc1e2062130
	// Actual value is, 24

	*ptrNumber = *ptrNumber + 10
	fmt.Println("Updated value is,", *ptrNumber)

}
