package main

import "fmt"

func main() {
	// panic(interface())

	// Example of a valid input
	process(10)

	// Example of a invalid input
	process(-1)

}

func process(input int) {
	defer fmt.Println("Deffered 1")
	defer fmt.Println("Deffered 2")
	if input < 0 {
		fmt.Println("Before panic")
		panic("Input must be a non negative value")

		//Not rechable
		// fmt.Println("After panic")
		// defer fmt.Println("Deffered 3")
	}
	fmt.Println("Processing input: ", input)
}
