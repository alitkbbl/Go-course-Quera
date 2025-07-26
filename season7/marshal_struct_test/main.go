package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	person := Person{
		Name:  "John",
		Age:   20,
		Email: "john@example.com",
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		panic("failed to marshal")
	}
	jsonString := string(jsonData)
	fmt.Println(jsonString)
}
