package main

import "fmt"

func main() {

	// if condition {
	// 	// block of code
	// }

	// age := 25
	// if age >= 18 {
	// 	fmt.Println("You're an adult")
	// }

	var a int
	fmt.Println("Enter youre age: ")
	fmt.Scanln(&a)
	if a < 18 {
		fmt.Println("Youre kid")
	} else if a >= 18 && a <= 65 {
		fmt.Println("Youre adult")
	} else if a > 65 && a <= 100 {
		fmt.Println("Youre old")
	} else {
		fmt.Println("Youre dead")
	}
}
