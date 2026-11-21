package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	greet := "Hey, good to see you back"
	fmt.Println(greet)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Can you pick a number b/w 1 to 10")

	input, _ := reader.ReadString('\n')
	fmt.Println("You will recive 1000$ for,", input)

	potentialWin, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		potentialWin *= 10
		fmt.Println("Wining ammount can be a muiliplier of,", potentialWin)
	}
}
