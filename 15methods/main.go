package main

import "fmt"

func main() {
	// methods are covered in this module
	user := User{
		Name:  "Alice",
		Age:   30,
		email: "alice@example.com",
	}
	println(user.Greet())
	println(user.GetEmail())
	var newEmail string
	println("Enter new email:")
	fmt.Scanln(&newEmail)
	user.SetEmail(newEmail)
	println(user.GetEmail())
}

type User struct {
	Name  string
	Age   int
	email string
}

func (u User) Greet() string {
	return "Hello, " + u.Name + "!"
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u User) GetEmail() string {
	return u.email
}
