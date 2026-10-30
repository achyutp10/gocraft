package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Command:", os.Args[0])
	for i, arg := range os.Args {
		fmt.Println("Argument1:", i, ":", arg)
	}

	// Define flags
	var name string
	var age int
	var male bool
	flag.StringVar(&name, "name", "John", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of user")
	flag.BoolVar(&male, "male", true, "Gender of the user")
	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Gender:", male)
}
