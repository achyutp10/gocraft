package main

import "fmt"

func main() {
	num := 424849
	fmt.Printf("%05d\n", num)

	// message := "Hello World! This is New Beginning"
	message := "Hello"
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	message1 := "Hello \nWorld"
	message2 := `Hello \nWorld`
	fmt.Printf(message1)
	fmt.Printf(message2)
	// `\d` "\\d"

	sqlQuery := `SELECT * FROM users WHERE age > 30`
}
