package main

func main() {
	// Functions are covered in this module
	greet("Alice")
	result := add(5, 3)
	println("The result of addition is:", result)
	proResult := proAdder(1, 2, 3, 4, 5)
	println("The result of proAdder is:", proResult)
}

func greet(name string) string {
	return "Hello, " + name + "!"
}

func add(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}
