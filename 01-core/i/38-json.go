package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName    string  `json:"first_name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{
		FirstName:    "John",
		Age:          30,
		EmailAddress: "sample.@example.com",
	}

	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{
		FirstName:    "Jane",
		Age:          32,
		EmailAddress: "jane@example.com",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	jsonData, err = json.Marshal(person1)
	if err != nil {
		fmt.Println("Error Marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	jsonData1 := `{
		"full_name": "Jenny Doe", 
		"emp_id": "009", 
		"age":30, 
		"address":{
			"city": "San Jose",
			"state": "CA"
	}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData1), &employeeFromJson)
	if err != nil {
		fmt.Println("Error Unmarshalling JSON:", err)
	}
	fmt.Println(employeeFromJson)
	fmt.Println("Age increased by 5:", employeeFromJson.Age+5)
	fmt.Println("Jenny City:", employeeFromJson.Address.City)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NY"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		fmt.Println("Error Marshalling JSON:", err)
	}
	fmt.Println("JSON LIST:", string(jsonList))

	// Handeling unknown JSON strictures
	jsonData2 := `{
		"name": "John", 
		"age":30, 
		"address":{
			"city": "New York",
			"state": "NY"
	}}`

	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		fmt.Println("Error UnMarshalling JSON:", err)
	}
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])
	fmt.Println("Decoded/Unmarshalled JSON:", data)

}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
