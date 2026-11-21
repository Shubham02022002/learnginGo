package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	greet := "Hey, good to see you back"
	fmt.Println(greet)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Can you pick a number b/w 1 to 10")
	//comma ok or error ok

	input, _ := reader.ReadString('\n')
	fmt.Println("You will recive 1000$ for,", input) // type -> string
}
