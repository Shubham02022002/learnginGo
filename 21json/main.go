package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
	Posts []string
}

func main() {
	// Encode Go struct into JSON
	EncodeJSON()

	// JSON array data
	data := []byte(`[
		{
			"Name": "Alice",
			"Age": 30,
			"Email": "alice@example.com",
			"Posts": ["Post 1", "Post 2"]
		},
		{
			"Name": "Bob",
			"Age": 25,
			"Email": "bob@example.com",
			"Posts": ["Post 3", "Post 4"]
		}
	]`)

	// Decode JSON into Go structs
	persons, err := DecodeJSON(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nDecoded JSON:")

	for _, person := range persons {
		fmt.Printf("\nName: %s\n", person.Name)
		fmt.Printf("Age: %d\n", person.Age)
		fmt.Printf("Email: %s\n", person.Email)
		fmt.Printf("Posts: %v\n", person.Posts)
	}
}

// Encode Go struct into JSON
func EncodeJSON() {
	persons := []Person{
		{
			Name:  "Alice",
			Age:   30,
			Email: "alice@example.com",
			Posts: []string{"Post 1", "Post 2"},
		},
		{
			Name:  "Bob",
			Age:   25,
			Email: "bob@example.com",
			Posts: []string{"Post 3", "Post 4"},
		},
	}

	finalJSON, err := json.MarshalIndent(persons, "", "  ")
	if err != nil {
		fmt.Println("Encoding error:", err)
		return
	}

	fmt.Println("Encoded JSON:")
	fmt.Println(string(finalJSON))
}

func DecodeJSON(data []byte) ([]Person, error) {
	var persons []Person

	err := json.Unmarshal(data, &persons)
	if err != nil {
		return nil, err
	}

	return persons, nil
}
