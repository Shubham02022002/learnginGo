package main

import "fmt"

func main() {
	fmt.Println("Loops are covered in this module")
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for i := 0; i < len(days); i++ {
		if days[i] == "Wednesday" {
			fmt.Println("Found Wednesday, breaking the loop")
			break
		}
		fmt.Println(days[i])
	}

	for i := 0; i < len(days); i++ {
		if days[i] == "Wednesday" {
			fmt.Println("Found Wednesday, skipping this iteration")
			continue
		}
		fmt.Println(days[i])
	}
}
