package main

import "fmt"

func main() {
	fmt.Println("In this module we'll cover arrays in go.")

	var userList [3]string

	userList[0] = "Shubham"
	userList[1] = "Rahul"
	userList[2] = "Akshay"
	fmt.Println("Users are,", userList)
	fmt.Println("Total number of users,", len(userList))

	var emails = [3]string{"shubham@gmail.com", "rahul@gmail.com", "akshay@gmail.com"}
	println("Emails: ", emails)

}
