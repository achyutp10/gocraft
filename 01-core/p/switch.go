package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	var operator string
	var result any

	fmt.Println("Enter n1: ")
	fmt.Scanln(&n1)
	fmt.Println("Enter n2: ")
	fmt.Scanln(&n2)
	fmt.Println("Enter operator: ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		result = n1 + n2
		fmt.Println(result)
	case "-":
		result = n1 - n2
		fmt.Println(result)
	case "/":
		result = n1 / n2
		fmt.Println(result)
	case "*":
		result = n1 * n2
		fmt.Println(result)
	default:
		fmt.Println("Error")
	}

}
