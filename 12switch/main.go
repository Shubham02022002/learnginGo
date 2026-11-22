package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in go lang.")
	rand.Seed(int64(time.Now().Nanosecond()))
	diceNumber := rand.Int31n(6) + 1
	fmt.Println("Dice number:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a one")
	case 2:
		fmt.Println("You rolled a two")
	case 3:
		fmt.Println("You rolled a three")
	case 4:
		fmt.Println("You rolled a four")
	case 5:
		fmt.Println("You rolled a five")
		fallthrough // will run the next case as well, even if the condition is not met
	case 6:
		fmt.Println("You rolled a six")

	default:
		fmt.Println("Invalid dice number")
	}
}
