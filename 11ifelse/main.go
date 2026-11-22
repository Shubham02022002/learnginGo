package main

import "fmt"

func main() {
	fmt.Println("In this module we'll learn about if else in Go")
	balance := 5000.0
	withdrawAmount := 2000.0

	if withdrawAmount <= 0 {
		fmt.Println("Invalid withdrawal amount")

	} else if withdrawAmount > balance {
		fmt.Println("Insufficient balance")

	} else {
		balance = balance - withdrawAmount

		if balance < 1000 {
			fmt.Println("Warning: Low account balance")
		} else {
			fmt.Println("Withdrawal successful")
		}

		fmt.Println("Remaining balance:", balance)
	}
}
