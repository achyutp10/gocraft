package main

import "fmt"

func main() {
	studMarks := map[string]int{
		"Ram":  60,
		"Sita": 50,
		"Hari": 70,
	}

	for k, v := range studMarks {
		fmt.Println(k, " : ", v)
	}
	studMarks["Gita"] = 85
	studMarks["Ram"] = 65

	for k, v := range studMarks {
		fmt.Println(k, " : ", v)
	}

	ages := map[string]int{
		"Ram":  22,
		"Sita": 21,
		"Hari": 24,
	}
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	age, exists := ages[name]
	fmt.Println(age, exists)
	if exists {
		fmt.Println("Age: ", age)
	} else {
		fmt.Println("Person not found")
	}
}
