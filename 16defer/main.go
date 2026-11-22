package main

func main() {
	// defer is covered in this module
	defer println("Goodbye!")
	defer println("Goodbye 1!")
	defer println("Goodbye 2!")
	println("Hello, World!")
	defered()
}

// LIFO (Last In First Out) - the last deferred function is executed first

func defered() {
	for i := 0; i < 5; i++ {
		defer println(i)
	}
}
