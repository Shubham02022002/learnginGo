package main

import (
	"fmt"
)

const PORT = 3000

func main() {
	var username string = "Shubham"
	fmt.Println(username)
	fmt.Printf("Type of username: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of isLoggedIn: %T\n", isLoggedIn)

	var smallValue uint8 = 255 // use alias instead (int)
	fmt.Println(smallValue)
	fmt.Printf("Type of smallValue: %T\n", smallValue)

	var samllFloat float32 = 23.4233232 //only 5 values after the decimal
	fmt.Println(samllFloat)
	fmt.Printf("Type of samllFloat: %T\n", samllFloat)

	var smallFloat float64 = 4231.123231232 //gives more precision
	fmt.Println(smallFloat)
	fmt.Printf("Type of smallFloat: %T\n", smallFloat)

	var temp int
	fmt.Println(temp)
	fmt.Printf("Type of smallFloat: %T\n", temp)

	var tempString string
	fmt.Println(tempString) // empty
	fmt.Printf("Type of smallFloat: %T\n", tempString)

	// no var style (walrus operator)
	API_KEY := "sk-gmt02i320dkfsdfi23fdkfsi"
	fmt.Println(API_KEY)

}
