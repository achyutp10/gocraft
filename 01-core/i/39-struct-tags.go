package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// FirstName string `json:"first_name" db:"firstName" xml:"first"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"-"`
	// Age       int    `json:"age,omitempty"`
	// EmailAddress string  `json:"email,omitempty"`
	// Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	// person := Person{
	// 	FirstName: "John",
	// 	Age:       0,
	// }
	// person := Person{
	// 	FirstName: "John",
	// 	Age:       30,
	// }
	// person := Person{
	// 	FirstName: "John",
	// 	LastName:  "",
	// 	Age:       30,
	// }

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling to struct:", err)
		return
	}
	fmt.Println(string(jsonData))
}
